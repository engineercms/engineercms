package models

import (
	"fmt"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/jinzhu/gorm"
	"time"
)

//购物车表
type Cart struct {
	// gorm.Model
	Id      int64     `json:"id"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now_add;type(datetime)"`
	// DeletedAt *time.Time
	UserId    int64
	ProductId int64
	Status    int `orm:"default(0)"`
}

func init() {
	orm.RegisterModel(new(Cart))
}

func CreateCart(productid, userid int64) (id int64, err error) {
	o := orm.NewOrm()
	//查询数据库中有无打卡
	var cart Cart
	//判断是否有重名
	err = o.QueryTable("cart").Filter("ProductId", productid).Filter("UserId", userid).Filter("Status", 0).One(&cart)
	if err == orm.ErrNoRows {
		// 没有找到记录
		cart := &Cart{
			ProductId: productid,
			UserId:    userid,
			Created:   time.Now(),
			Updated:   time.Now(),
		}
		id, err = o.Insert(cart)
		if err != nil {
			return 0, err
		}
		return id, nil
	} else if err == nil { //如果存在记录，或多于一条，则进行更新
		// cart2 := &Cart{Id: cart.Id}
		fmt.Printf("cart2.Updated")
		cart.Updated = time.Now()
		_, err = o.Update(&cart, "Updated")
		if err != nil {
			return cart.Id, err
			fmt.Printf("heh")
		}
		return cart.Id, nil
	}
	fmt.Printf("he")
	return id, nil
}

//购物车表
type UserCart struct {
	// gorm.Model
	Id int64 `json:"id"`
	// Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated"`
	// DeletedAt *time.Time
	UserId       int64  `json:"userid"`
	ProductId    int64  `json:"productid"`
	TopProjectId int64  `json:"topprojectid"`
	Status       int    `json:"status"`
	UserNickname string `json:"usernickname"`
	ProductTitle string `json:"producttitle"`
	ProjectTitle string `json:"projecttitle"`
}

//查询某个用户借阅记录
func GetUserCart(uid int64, limit, offset, status int, searchText string, isadmin bool) (usercarts []UserCart, err error) {
	//获取DB Where("product.title LIKE ?", "%searchText%").不对
	db := GetDB()
	if isadmin {
		// 必须要写权select，坑爹啊
		err = db.Order("cart.updated desc").Table("cart").Select("cart.id,cart.user_id,cart.product_id,cart.status,cart.updated,user.nickname as user_nickname, product.title as product_title, project.title as project_title").Where("cart.status=?", status).
			Joins("left JOIN user on user.id = cart.user_id").
			Joins("left join product on product.id = cart.product_id").
			Joins("left join project on project.id = product.project_id").
			Limit(limit).Offset(offset).Scan(&usercarts).Error
	} else if searchText != "" {
		err = db.Order("cart.updated desc").Table("cart").Select("cart.id,cart.user_id,cart.product_id,cart.status,cart.updated,user.nickname as user_nickname, product.title as product_title, project.title as project_title").Where("user_id=? AND cart.status=?", uid, status).
			Joins("left JOIN user on user.id = cart.user_id").
			Joins("left join product on product.id = cart.product_id").
			Joins("left join project on project.id = product.project_id").
			Limit(limit).Offset(offset).Scan(&usercarts).Error
	} else {
		err = db.Order("cart.updated desc").Table("cart").Select("cart.id,cart.user_id,cart.product_id,cart.status,cart.updated,user.nickname as user_nickname, product.title as product_title, product.top_project_id as top_project_id,project.title as project_title").Where("user_id=? AND cart.status=?", uid, status).
			Joins("left JOIN user on user.id = cart.user_id").
			Joins("left join product on product.id = cart.product_id").
			Joins("left join project on project.id = product.project_id").
			Limit(limit).Offset(offset).Scan(&usercarts).Error
	}
	return usercarts, err
	// 多连接及参数
	// db.Joins("JOIN pays ON pays.user_id = users.id", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("user_id = ?", uid).Find(&pays)
}

//查询某个用户借阅记录总数
func GetUserCartCount(uid int64, status int, searchText string, isadmin bool) (count int64, err error) {
	//获取DB
	db := GetDB()
	if isadmin {
		err = db.Table("cart").Where("cart.status=?", status).
			Count(&count).Error
	} else if searchText != "" {
		err = db.Table("cart").Where("user_id=? AND cart.status=?", uid, status).
			Count(&count).Error
	} else {
		err = db.Table("cart").Where("user_id=? AND cart.status=?", uid, status).
			Count(&count).Error
	}
	return count, err
}

//查询一个cart
func GetUserCartbyId(id int64) (cart Cart, err error) {
	//获取DB
	db := GetDB()
	err = db.Where("id = ?", id).Find(&cart).Error
	return cart, err
}

// 删除
func DeleteUserCart(id int64) error {
	//获取DB
	db := GetDB()
	err := db.Where("id = ?", id).Delete(Cart{}).Error
	return err
}

// func CreateCart(cart Cart) (Cart, error) {
// 	db := GetDB()
// 	err := db.Where("product_i_d = ? AND user_i_d = ?", cart.ProductID, cart.UserID).First(&cart).Error
// 	if err != nil {
// 		return cart, err
// 	}
// 	cart.CreatedAt = time.Now()
// 	cart.UpdatedAt = time.Now()
// 	db = db.Table("Cart").Create(&cart)
// 	return cart, db.Error
// }

// 声明project_pics数据表结构
type ProjectPic struct {
	Id        int    `json:"id"`
	ProjectId int    `json:"project_id"`
	Url       string `json:"url"`
	Type      string `json:"type"`
}

//gorm批量插入——参考用，没有意义
func AddProjectPics(data []string, project_pic_type string, project_id int) bool {
	sql := "INSERT INTO `project_pics` (`project_id`,`url`,`type`) VALUES "
	// 循环data数组,组合sql语句
	for key, value := range data {
		if len(data)-1 == key {
			//最后一条数据 以分号结尾
			sql += fmt.Sprintf("(%d,'%s','%s');", project_id, value, project_pic_type)
		} else {
			sql += fmt.Sprintf("(%d,'%s','%s'),", project_id, value, project_pic_type)
		}
	}
	_db.Exec(sql)
	return true
}

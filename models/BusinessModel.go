package models

import (
	// "fmt"
	// "github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"time"
)

//出差登记信息表
type Business struct {
	// gorm.Model
	ID           uint      `json:"id" gorm:"primary_key"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	UserID       int64 `gorm:"column:user_id;foreignkey:UserId;"` // One-To-One (属于 - 本表的BillingAddressID作外键
	Location     string
	Lat          float64
	Lng          float64
	StartDate    time.Time
	EndDate      time.Time
	Projecttitle string
	Drivername   string
	Subsidy      int
	Carfare      int
	Hotelfee     int
	BusinessUser []BusinessUser `gorm:"foreignkey:UserId"` //这个外键难道不是错的么？应该是UserID?没错，因为column:user_id
	ArticleID    int64          `gorm:"column:article_id;foreignkey:ArticleId;"`
	Article      Article        `gorm:"foreignkey:ArticleId"`
	Worktime     float64
	Overtime     int
}

//出差人员表
type BusinessUser struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	UserID     int64 `gorm:"column:user_id;foreignkey:UserId;"` // 外键 (属于), tag `index`是为该列创建索引
	BusinessID uint  `json:"amount" gorm:"column:amount"`
	User       User  `gorm:"foreignkey:UserId"`
}

// 出差每天打卡，还是另外一个功能吧，不要放一起
type BusinessCheckin struct {
	gorm.Model
	Location   string `json:"F_Location"`
	Lat        float64
	Lng        float64
	CheckTime  time.Time `gorm:"autoCreateTime"`
	SelectDate time.Time `gorm:"type:date"`
}

//用户充值记录
// type Recharge struct {
// 	gorm.Model
// 	UserID int64 `gorm:"column:user_id;foreignkey:UserId;"` // 外键 (属于), tag `index`是为该列创建索引
// 	Amount int   `gorm:"column:amount"`
// 	User   User  `gorm:"foreignkey:UserId"`
// }

func init() {
	_db.CreateTable(&Business{}, &BusinessUser{})
}

func CreateBusiness(business Business) (id uint, err error) {
	db := GetDB()
	// projectuser := ProjectUser{ProjectId: pid, UserId: uid}
	result := db.Create(&business) // 通过数据的指针来创建
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
	// result.RowsAffected // 返回插入记录的条数
	return business.ID, result.Error
}

// 添加用户~出差关联表格
func CreateUserBusiness(businessuser BusinessUser) (id uint, err error) {
	db := GetDB()
	// projectuser := ProjectUser{ProjectId: pid, UserId: uid}
	result := db.Create(&businessuser) // 通过数据的指针来创建
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
	// result.RowsAffected // 返回插入记录的条数
	return businessuser.ID, result.Error
}

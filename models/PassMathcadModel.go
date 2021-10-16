package models

import (
	// "fmt"
	// "github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	// "time"
	"errors"
)

//用户-模板表
type UserTemple struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	ClassID    int64  `gorm:"column:class_id;foreignkey:ProjectId;"`
	UserID     int64  `gorm:"column:user_id;foreignkey:UserId;"` // 外键 (属于), tag `index`是为该列创建索引
	TempTitle  string `json:"temptitle" gorm:"column:temp_title"`
	TempTitleB string `json:"temptitleb" gorm:"column:temp_title_b"` //简要名称，无日期，无版本号，无扩展名
	TempPath   string `json:"temppath" gorm:"column:temp_path"`
	// TempUrl   string `json:"tempurl" gorm:"column:temp_url"`
	Status      bool   `json:"status"`
	Version     string `json:"version"`
	User        User   `gorm:"foreignkey:UserId"`
	MathArticle MathArticle
}

//模板-输入表
type TempleInputs struct {
	gorm.Model
	UserTempleID      uint
	InputAlias        string            `json:"inputalias" gorm:"column:input_alias"`
	InputValue        string            `json:"inputvalue" gorm:"column:input_value"` //默认值
	ResultType        string            `json:"resulttype" gorm:"column:result_type"`
	Units             string            `json:"units" gorm:"column:units"`
	Comment           string            `json:"comment"`
	RealMin           string            `json:"realmin"`
	RealMax           string            `json:"realmax"`
	HistoryInputValue HistoryInputValue //`gorm:"foreignkey:ID"` //带条件的预加载,条件是UserHistoryID
	SelectValue       []Select2         `json:"selectvalue"`
	TextArealValue    TextAreal         `json:"textarealvalue"`
	// RealValue    float64 `json:"realvalue" gorm:"column:real_value"`
}

type Select2 struct {
	Id    string `json:"id"`
	Text  string `json:"text"`
	Value string `json:"value"`
}

type TextAreal struct {
	Value string `json:"value"`
}

// type Select struct {
// 	Value []string
// }

//模板-输出表
type TempleOutputs struct {
	gorm.Model
	UserTempleID       uint
	OutputAlias        string             `json:"outputalias" gorm:"column:output_alias"`
	OutputValue        string             `json:"outputvalue" gorm:"column:output_value"`
	ResultType         string             `json:"resulttype" gorm:"column:result_type"`
	Units              string             `json:"units" gorm:"column:units"`
	Comment            string             `json:"comment"`
	HistoryOutputValue HistoryOutputValue //`gorm:"foreignkey:ID"`
	// RealValue    float64 `json:"realvalue" gorm:"column:real_value"`
}

//用户-历史计算记录表
type UserHistory struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	UserID       int64 `json:"userid" gorm:"column:user_id;foreignkey:UserId;"` // 外键 (属于), tag `index`是为该列创建索引
	UserTempleID uint  `json:"tempid" gorm:"column:user_temple_id"`
	User         User  `gorm:"foreignkey:UserId"`
	UserTemple   UserTemple
	PdfUrl       string `json:"pdfurl"`
	// TempleInputs  TempleInputs //根据usertempleid查询到templeinputs
	// TempleOutputs TempleOutputs
	// HistoryInputValue  []HistoryInputValue  `gorm:"foreignkey:ID"`
	// HistoryOutputValue []HistoryOutputValue `gorm:"foreignkey:ID"`
}

//历史计算记录-输入参数记录表
type HistoryInputValue struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	UserHistoryID  uint   `json:"user_history_id" foreignkey:UserHistoryID;" gorm:"column:user_history_id"`
	TempleInputsID uint   `json:"templeinputsid" gorm:"column:temple_inputs_id"`
	InputValue     string `json:"inputvalue" gorm:"column:input_value"`
	// TempleInputs   TempleInputs//invalid recursive type
}

//历史计算记录-输出参数记录表
type HistoryOutputValue struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	UserHistoryID   uint   `foreignkey:UserHistoryID;" gorm:"column:user_history_id"` // 外键 (属于), tag `index`是为该列创建索引
	TempleOutputsID uint   `json:"templeoutputsid" gorm:"column:temple_outputs_id"`
	OutputValue     string `json:"outputvalue" gorm:"column:output_value"`
	// TempleOutputs   TempleOutputs//invalid recursive type
}

// 用户自定义计算书封面
type UserCalFace struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	UserID    int64 `gorm:"column:user_id;foreignkey:UserId;"` // 外键 (属于), tag `index`是为该列创建索引
	CalFaceID uint  `json:"calfaceid" gorm:"column:cal_face_id"`
}

type CalFace struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	Tittle  string `gorm:"column:title;"` // 外键 (属于), tag `index`是为该列创建索引
	Content string `json:"content" gorm:"column:content"`
}

// 用户添加计算书说明（前说明和后分析）explain
type UserCalExplain struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	UserID       int64 `gorm:"column:user_id;foreignkey:UserId;"` // 外键 (属于), tag `index`是为该列创建索引
	CalExplainID uint  `json:"calexplainid" gorm:"column:cal_explain_id"`
}

type CalExplain struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	Tittle      string `gorm:"column:title"` // 外键 (属于), tag `index`是为该列创建索引
	Content     string `json:"content" gorm:"column:content"`
	FrontOrBack bool   `json:"frontorback" gorm:"column:front_or_back"`
}

// mathcad的文章
type MathArticle struct {
	gorm.Model
	UserTempleID uint
	Title        string `json:"title" gorm:"column:title;size:20"`
	Subtext      string `json:"subtext" gorm:"column:subtext;size:20"`
	Content      string `json:"html" gorm:"column:content;size:5000"`
	// UserTemple   UserTemple
}

func init() {
	_db.CreateTable(&UserTemple{})
	_db.CreateTable(&TempleInputs{})
	_db.CreateTable(&TempleOutputs{})

	_db.CreateTable(&UserHistory{})
	_db.CreateTable(&HistoryInputValue{})
	_db.CreateTable(&HistoryOutputValue{})
	_db.CreateTable(&MathArticle{})

	// _db.CreateTable(&AdminIpsegment{})
}

// 用户模板写入数据库
func AddTemple(classid, userid int64, templetitle, templetitleb, templepath, version string) (id uint, err error) {
	db := GetDB()
	//查询数据库中有无打卡
	// var businesscheckin BusinessCheckin
	usertemple := &UserTemple{
		ClassID:    classid,
		UserID:     userid,
		TempTitle:  templetitle,
		TempTitleB: templetitleb,
		TempPath:   templepath,
		// TempUrl:   templeurl,
		Version: version,
		Status:  true,
	}
	//判断是否有重名
	err = db.Where("user_id = ? AND temp_title = ?", userid, templetitle).FirstOrCreate(&usertemple).Error
	// err = o.QueryTable("business_checkin").Filter("ActivityId", ActivityId).Filter("UserId", UserId).Filter("SelectDate", SelectDate).One(&check1, "Id")
	// if err == orm.ErrNoRows {
	// 没有找到记录
	return usertemple.ID, err
}

// 修改输入参数的信息
func UpdateTemple(templeid uint, fieldname, value string) (err error) {
	//获取DB
	db := GetDB()
	// 条件更新
	var new_value bool
	switch value {
	case "true":
		new_value = true
	case "false":
		new_value = false
	}
	// 	case "comment":
	result := db.Model(&UserTemple{}).Where("id = ?", templeid).Update(fieldname, new_value)
	return result.Error
}

// 解析模板的输入参数写入数据库
func AddInputAlias(usertempleid uint, inputvalue, inputalias, resulttype, units string) (id uint, err error) {
	//获取DB
	db := GetDB()
	//保证id正确
	var templeinputs TempleInputs
	result := db.Where("user_temple_id = ? AND input_alias=?", usertempleid, inputalias).First(&templeinputs)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result = db.Create(&TempleInputs{UserTempleID: usertempleid, InputAlias: inputalias, InputValue: inputvalue, ResultType: resulttype, Units: units}) // 通过数据的指针来创建
		return templeinputs.ID, result.Error
	} else {
		return 0, result.Error
	}
	// result := db.Create(&TempleInputs{UserTempleID: usertempleid, InputAlias: inputalias}) // 通过数据的指针来创建
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
	// result.RowsAffected // 返回插入记录的条数
	// return templeinputs.ID, result.Error

	// if err = tx.Create(&Pay{UserID: uid, User2ID: product.Uid, ArticleID: articleid, Amount: newamount}).Error; err != nil {
	// 	return err
	// }
}

// 修改输入参数的备注信息
func UpdateTempleInputs(templeinputid uint, fieldname, value string) (err error) {
	//获取DB
	db := GetDB()
	// 条件更新
	switch fieldname {
	case "realmin":
		fieldname = "real_min"
	case "realmax":
		fieldname = "real_max"
	}
	// 	case "comment":
	result := db.Model(&TempleInputs{}).Where("id = ?", templeinputid).Update(fieldname, value)
	return result.Error
}

// 修改输出参数的备注信息
func UpdateTempleOutputs(templeoutputid uint, comment string) (err error) {
	//获取DB
	db := GetDB()
	// 条件更新
	result := db.Model(&TempleOutputs{}).Where("id = ?", templeoutputid).Update("comment", comment)
	return result.Error
}

// 解析模板的输出参数写入数据库
func AddOutputAlias(usertempleid uint, outputvalue, outputalias, resulttype, units string) (id uint, err error) {
	//获取DB
	db := GetDB()
	//保证id正确
	var templeoutputs TempleOutputs
	result := db.Where("user_temple_id = ? AND output_alias=?", usertempleid, outputalias).First(&templeoutputs)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result = db.Create(&TempleOutputs{UserTempleID: usertempleid, OutputAlias: outputalias, OutputValue: outputvalue, ResultType: resulttype, Units: units}) // 通过数据的指针来创建
		return templeoutputs.ID, result.Error
	} else {
		return 0, result.Error
	}
}

// 查出所有模板
func GetMathTemples(classid int64, limit, offset int) (usertemples []UserTemple, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("BusinessUsers.NickNames")——嵌套预加载！！
	db := GetDB()
	err = db.Order("updated_at desc").
		Preload("User").
		Preload("MathArticle").
		// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
		// Preload("User.Nickname").
		Where("class_id = ?", classid).
		Limit(limit).Offset(offset).
		Find(&usertemples).Error
	return usertemples, err
}

// 查出所有模板
func GetMathTempleCount(classid int64) (count int64, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("BusinessUsers.NickNames")——嵌套预加载！！
	db := GetDB()
	err = db.Model(&UserTemple{}).
		// Preload("User").
		// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
		// Preload("User.Nickname").
		Where("class_id = ?", classid).
		Count(&count).Error
	return count, err
}

// 查出某个模板
func GetMathTemple(templeid uint) (usertemple UserTemple, err error) {
	db := GetDB()
	err = db.
		// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
		Preload("User").
		Where("id = ?", templeid).
		Find(&usertemple).Error
	return usertemple, err
}

// 查出模板输入参数
func GetTempleInputs(templeid uint) (templeinputs []TempleInputs, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("BusinessUsers.NickNames")——嵌套预加载！！
	db := GetDB()
	err = db.
		// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
		// Preload("User.Nickname").
		Where("user_temple_id = ?", templeid).
		Find(&templeinputs).Error
	return templeinputs, err
}

// 查出模板输出参数
func GetTempleOutputs(templeid uint) (templeoutputs []TempleOutputs, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("BusinessUsers.NickNames")——嵌套预加载！！
	db := GetDB()
	err = db.
		// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
		// Preload("User.Nickname").
		Where("user_temple_id = ?", templeid).
		Find(&templeoutputs).Error
	return templeoutputs, err
}

// 用户计算历史写入
func CreateUserHistory(userid int64, usertempleid uint, pdfurl string) (id uint, err error) {
	//获取DB
	db := GetDB()
	//保证id正确
	userhistory := UserHistory{UserID: userid, UserTempleID: usertempleid, PdfUrl: pdfurl}
	result := db.Create(&userhistory) // 通过数据的指针来创建
	return userhistory.ID, result.Error
}

// 用户计输入算数据写入, templeinputsid
func AddHistoryInputValue(userhistoryid uint, templeinputs []TempleInputs) (id uint, err error) {
	db := GetDB()
	var historyinputvalue HistoryInputValue
	// result := map[string]interface{}{}
	var result *gorm.DB
	for _, v := range templeinputs {
		// for i := 0; i < len(templeinputs); i++ {
		result = db.Where("user_history_id = ? AND temple_inputs_id = ? ", userhistoryid, v.ID).First(&historyinputvalue)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			historyinputvalue = HistoryInputValue{UserHistoryID: userhistoryid, TempleInputsID: v.ID, InputValue: v.InputValue}
			result = db.Create(&historyinputvalue)
			// 开始这里有return，导致只算一次就返回了，百思不得其解！！不应该这里放return
		}
	}
	return historyinputvalue.ID, result.Error
}

// 用户计算输出数据写入, templeoutputsid
func AddHistoryOutputValue(userhistoryid uint, templeoutputs []TempleOutputs) (id uint, err error) {
	db := GetDB()
	var historyoutputvalue HistoryOutputValue
	// result := map[string]interface{}{}
	var result *gorm.DB
	for _, v := range templeoutputs {
		// for i := 0; i < len(templeinputs); i++ {
		result = db.Where("user_history_id = ? AND temple_outputs_id = ? ", userhistoryid, v.ID).First(&historyoutputvalue)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			historyoutputvalue = HistoryOutputValue{UserHistoryID: userhistoryid, TempleOutputsID: v.ID, OutputValue: v.OutputValue}
			result = db.Create(&historyoutputvalue)
			// 开始这里有return，导致只算一次就返回了，百思不得其解！！不应该这里放return
		}
	}
	return historyoutputvalue.ID, result.Error
}

// 查询用户计算历史列表
func GetUserHistory(userid int64, usertempleid uint, limit, offset int) (userhistory []UserHistory, err error) {
	db := GetDB()
	err = db.Order("user_history.updated_at desc").
		Preload("UserTemple", "id = ?", usertempleid).
		Preload("User", "id = ?", userid).
		Where("user_id = ? AND user_temple_id=?", userid, usertempleid).
		Limit(limit).Offset(offset).
		Find(&userhistory).Error
	return userhistory, err
}

func GetUserHistoryCount(userid int64, usertempleid uint) (count int64, err error) {
	db := GetDB()
	err = db.Model(&UserHistory{}).
		Preload("UserTemple", "id = ?", usertempleid).
		Preload("User", "id = ?", userid).
		Where("user_id = ? AND user_temple_id=?", userid, usertempleid).
		Count(&count).Error
	return count, err
}

// 查询历史计算数值_作废！！
func GetUserHistoryValue(userhistoryid uint) (userhistory UserHistory, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("BusinessUsers.NickNames")——嵌套预加载！！
	db := GetDB()
	err = db.
		Preload("User").
		Preload("TempleInputs").
		Where("id = ?", userhistoryid).
		Joins("left join history_input_value AS t1 on t1.user_history_id = userhistoryid").
		Joins("left join history_input_value AS t2 ON t2.temple_inputs_id = temple_inputs.id").
		Joins("left join history_output_value AS t3 on t3.user_history_id = userhistoryid").
		Joins("left join history_output_value AS t4 ON t4.temple_outputs_id = temple_outputs.id").
		Find(&userhistory).Error
	return userhistory, err
}

// 根据historyid查询历史计算
func GetHistory(historyid uint) (userhistory UserHistory, err error) {
	db := GetDB()
	err = db.
		Where("id = ?", historyid).
		Find(&userhistory).Error
	return userhistory, err
}

// 查出模板历史输入参数
func GetUserHistoryInputs(usertempleid, userhistoryid uint) (templeinputs []TempleInputs, err error) {
	db := GetDB()
	err = db.
		Preload("HistoryInputValue", "user_history_id = ?", userhistoryid).
		Where("user_temple_id = ?", usertempleid).
		// Joins("left join user_history AS t1 on t1.UserTempleID = HistoryInputValue.").
		// Joins("left join history_input_value AS t2 ON t2.temple_inputs_id = temple_inputs.id").
		Find(&templeinputs).Error
	return templeinputs, err
}

// 查出模板历史输出参数
func GetUserHistoryOutputs(usertempleid, userhistoryid uint) (templeoutputs []TempleOutputs, err error) {
	db := GetDB()
	err = db.
		Preload("HistoryOutputValue", "user_history_id = ?", userhistoryid).
		Where("user_temple_id = ?", usertempleid).
		// Joins("left join user_history AS t1 on t1.UserTempleID = HistoryInputValue.").
		// Joins("left join history_input_value AS t2 ON t2.temple_inputs_id = temple_inputs.id").
		Find(&templeoutputs).Error
	return templeoutputs, err
}

// 删除模板
func DeleteTemple(templeid uint) error {
	db := GetDB()
	result := db.Where("id = ?", templeid).Delete(&UserTemple{})
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
	// result.RowsAffected // 返回插入记录的条数
	return result.Error
}

//添加文章作为用户模板的附件
func AddMathArticle(usertempleid uint, title, subtext, content string) (id uint, err1, err2 error) {
	db := GetDB()
	var result2 *gorm.DB
	//保证id正确
	var matharticle MathArticle
	result1 := db.Where("user_temple_id = ?", usertempleid).First(&matharticle)
	if errors.Is(result1.Error, gorm.ErrRecordNotFound) {
		result2 = db.Create(&MathArticle{UserTempleID: usertempleid, Title: title, Subtext: subtext, Content: content}) // 通过数据的指针来创建
		return matharticle.ID, result1.Error, result2.Error                                                             //这样无法返回ID，参见CreateUserHistory
	} else {
		return matharticle.ID, result1.Error, nil
	}
}

// 根据matharticle取得文章
func GetMathArticle(matharticleid uint) (matharticle MathArticle, err error) {
	db := GetDB()
	err = db.
		Where("id = ?", matharticleid).
		Find(&matharticle).Error
	return matharticle, err
}

//编辑文章
func UpdateMathArticle(matharticleid uint, title, subtext, content string) (err error) {
	//获取DB
	db := GetDB()
	// 条件更新
	result := db.Model(MathArticle{}).Where("id = ?", matharticleid).Updates(MathArticle{Title: title, Subtext: subtext, Content: content})
	// db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
	return result.Error
}

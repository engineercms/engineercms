// Business Trip
package controllers

import (
	"encoding/json"
	"github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"
	"strconv"
	// "strings"
	"time"
)

type BusinessController struct {
	beego.Controller
}

type wxuser struct {
	Name string `json:"name"`
}

// @Title post business by userid
// @Description post business by userid
// @Param id query string true "The id of project"
// @Param title query string true "The title of article"
// @Param content query string true "The content of article"
// @Success 200 {object} models.CreateBusiness
// @Failure 400 Invalid page supplied
// @Failure 404 pas not found
// @router /addbusiness/:id [post]
//用户id打赏一个文章id
func (c *BusinessController) AddBusiness() {
	// var user models.User
	// var err error
	// openID := c.GetSession("openID")
	// if openID != nil {
	// 	user, err = models.GetUserByOpenID(openID.(string))
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// } else {
	// 	c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
	// 	c.ServeJSON()
	// 	return
	// 	// user.Id = 9
	// }
	_, _, uid, _, isLogin := checkprodRole(c.Ctx)
	if !isLogin {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		c.Data["json"] = "未登陆"
		c.ServeJSON()
		// return
	}
	// user, err := models.GetUserByUsername(uname)
	// if err != nil {
	// 	beego.Error(err)
	// }

	// pid := c.Ctx.Input.Param(":id")

	location := c.Input().Get("location")
	lat := c.Input().Get("lat")
	latfloat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		beego.Error(err)
	}
	lng := c.Input().Get("lng")
	lngfloat, err := strconv.ParseFloat(lng, 64)
	if err != nil {
		beego.Error(err)
	}
	startDate := c.Input().Get("startDate")
	endDate := c.Input().Get("endDate")
	projecttitle := c.Input().Get("projecttitle")
	drivername := c.Input().Get("drivername")
	subsidy := c.Input().Get("subsidy")
	subsidyint, err := strconv.Atoi(subsidy)
	if err != nil {
		beego.Error(err)
	}
	carfare := c.Input().Get("carfare")
	carfareint, err := strconv.Atoi(carfare)
	if err != nil {
		beego.Error(err)
	}
	hotelfee := c.Input().Get("hotelfee")
	hotelfeeint, err := strconv.Atoi(hotelfee)
	if err != nil {
		beego.Error(err)
	}
	users := c.Input().Get("users")
	beego.Info(users)
	m := []wxuser{}
	err = json.Unmarshal([]byte(users), &m)
	if err != nil {
		beego.Error(err)
	}
	beego.Info(m)

	// 2020/11/15 17:08:02.488 [I] [business_trip.go:60] [{"index":1,"name":"6","showta
	// g":true},{"index":2,"name":"61","showtag":true}]
	//id转成64为
	// pidNum, err := strconv.ParseInt(pid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//根据pid查出项目id
	// proj, err := models.GetProj(pidNum)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// var topprojectid int64
	// if proj.ParentIdPath != "" {
	// 	parentidpath := strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
	// 	parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
	// 	patharray := strings.Split(parentidpath1, "-")
	// 	topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// } else {
	// 	topprojectid = proj.Id
	// }
	// code := time.Now().Format("2006-01-02 15:04")
	// code = strings.Replace(code, "-", "", -1)
	// code = strings.Replace(code, " ", "", -1)
	// code = strings.Replace(code, ":", "", -1)
	//根据项目id添加出差
	var business models.Business
	// 添加business
	const lll = "2006-01-02"
	starttime, _ := time.Parse(lll, startDate)
	endtime, _ := time.Parse(lll, endDate)
	// d:=a.Sub(b)
	// fmt.Println(d.Hours()/24)
	// worktime := endtime.Sub(starttime)

	// var business models.Business
	business = models.Business{
		UserID: uid,
		// ArticleID:    aid,
		StartDate:    starttime,
		EndDate:      endtime,
		Location:     location,
		Lat:          latfloat,
		Lng:          lngfloat,
		Projecttitle: projecttitle,
		Drivername:   drivername,
		Subsidy:      subsidyint,
		Carfare:      carfareint,
		Hotelfee:     hotelfeeint,
		// Worktime:  worktime.Hours(),
		// Overtime:  0,
	}
	Id, err := models.CreateBusiness(business)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "写入数据错误", "id": 1}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"info": "SUCCESS"}
		c.ServeJSON()
	}

	// 循环添加出差同行人员与出差的关联表
	for _, v := range m {
		beego.Info(v.Name)
		tripuser := models.GetUserByNickname(v.Name)
		beego.Info(tripuser.Id)
		var businessuser models.BusinessUser
		businessuser.UserID = tripuser.Id
		businessuser.BusinessID = Id
		_, err = models.CreateUserBusiness(businessuser)
	}
	// 如果文章时打开的，则添加文章并关联到出差
	aid, err := models.AddArticle("title", "content", int64(Id))
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERR", "id": aid}
		c.ServeJSON()
	} else {
		// c.Data["json"] = id
		c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "id": aid}
		c.ServeJSON()
	}
}

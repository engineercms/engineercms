// Business Trip
package controllers

import (
	"github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"time"
)

type BusinessController struct {
	beego.Controller
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
func (c *PayController) AddBusiness() {
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
	uname, _, uid, _, isLogin := checkprodRole(c.Ctx)
	if !isLogin {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		c.Data["json"] = "未登陆"
		c.ServeJSON()
		return
	}
	user, err := models.GetUserByUsername(uname)
	if err != nil {
		beego.Error(err)
	}

	// 添加文章
	pid := c.Ctx.Input.Param(":id")

	title := c.Input().Get("title")
	content := c.Input().Get("content")
	//id转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据pid查出项目id
	proj, err := models.GetProj(pidNum)
	if err != nil {
		beego.Error(err)
	}
	var topprojectid int64
	if proj.ParentIdPath != "" {
		parentidpath := strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else {
		topprojectid = proj.Id
	}
	code := time.Now().Format("2006-01-02 15:04")
	code = strings.Replace(code, "-", "", -1)
	code = strings.Replace(code, " ", "", -1)
	code = strings.Replace(code, ":", "", -1)
	//根据项目id添加成果code, title, label, principal, content string, projectid int64
	Id, err := models.AddProduct(code, title, "wx", user.Nickname, user.Id, pidNum, topprojectid)
	if err != nil {
		beego.Error(err)
	}
	//将文章添加到成果id下
	aid, err := models.AddArticle(title, content, Id)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERR", "id": aid}
		c.ServeJSON()
	} else {
		// c.Data["json"] = id
		c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "id": aid}
		c.ServeJSON()
	}

	// 添加business
	fromdate := c.Input().Get("content")
	enddate := c.Input().Get("content")
	const lll = "2006-01-02"
	starttime, _ := time.Parse(lll, fromdate)
	endtime, _ := time.Parse(lll, enddate)
	// d:=a.Sub(b)
	// fmt.Println(d.Hours()/24)
	worktime := endtime.Sub(starttime)

	// var business models.Business
	business := models.Business{
		UserID:    uid,
		ArticleID: aid,
		FromDate:  starttime,
		EndDate:   endtime,
		Worktime:  worktime.Hours(),
		Overtime:  0,
	}

	_, err = models.CreateBusiness(business)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "写入数据错误", "id": 1}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"info": "SUCCESS"}
		c.ServeJSON()
	}
}

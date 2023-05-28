package service

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/amatsuzero/ginchat/models"
	"github.com/amatsuzero/ginchat/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Accept json
// @Produce json
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// GetUserList
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Accept json
// @Produce json
// @Success 200 {string} json{"code", "message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	usr := models.UserBasic{}
	usr.Name = c.Query("name")
	pwd := c.Query("password")
	repwd := c.Query("repassword")

	if repwd != pwd {
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	usr.Salt = fmt.Sprintf("%06d", rand.Int31())
	usr.Password = utils.MakePassword(pwd, usr.Salt)
	empty := models.UserBasic{}
	if models.FindUserByName(usr.Name) != empty {
		c.JSON(-1, gin.H{
			"message": "用户名已经注册!",
		})
		return
	}

	models.CreateUser(usr)
	c.JSON(200, gin.H{
		"message": "新增用户成功!",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "用户id"
// @Accept json
// @Produce json
// @Success 200 {string} json{"code", "message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	usr := models.UserBasic{}
	ID, _ := strconv.Atoi(c.Query("id"))
	usr.ID = uint(ID)
	models.DeleteUser(usr)
	c.JSON(200, gin.H{
		"message": "删除成功！",
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "用户 ID"
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @param email formData string false "邮箱"
// @param phone formData string false "电话"
// @Accept multipart/form-data
// @Produce json
// @Success 200 {string} json{"code", "message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	usr := models.UserBasic{}
	ID, _ := strconv.Atoi(c.PostForm("id"))
	usr.ID = uint(ID)
	usr.Name = c.PostForm("name")
	usr.Password = c.PostForm("password")
	usr.Email = c.PostForm("email")
	usr.Phone = c.PostForm("phone")

	if ret, err := govalidator.ValidateStruct(usr); !ret {
		c.JSON(-1, gin.H{
			"message": err.Error(),
		})
		return
	}

	models.UpdateUser(usr)
	c.JSON(200, gin.H{
		"message": "更新成功！",
	})
}

// FindUserByNameAndPassword
// @Summary 按用户名和密码查找
// @Tags 用户模块
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @Accept multipart/form-data
// @Produce json
// @Success 200 {string} json{"code", "message"}
// @Router /user/findUserByNameAndPassword [post]
func FindUserByNameAndPassword(c *gin.Context) {
	name := c.PostForm("name")
	pwd := c.PostForm("password")

	model := models.FindUserByName(name)
	empty := models.UserBasic{}
	if model == empty {
		c.JSON(-1, gin.H{
			"message": "该用户不存在",
		})
		return
	}

	ret := utils.ValidPassword(pwd, model.Salt, model.Password)
	if !ret {
		c.JSON(-1, gin.H{
			"message": "密码不正确",
		})
		return
	}

	pwd = utils.MakePassword(pwd, model.Salt)
	data := models.FindUserByNameAndPassword(name, pwd)
	if !ret || data != model {
		c.JSON(-1, gin.H{
			"message": "密码不正确",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": data,
	})
}

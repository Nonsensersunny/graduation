package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb/uuid"
	"graduation/internal/config"
	"graduation/internal/log"
	"graduation/internal/middleware"
	"graduation/internal/utils"
	"graduation/pkg/modules/model"
	"graduation/pkg/service"
	"net/http"
	"strconv"
	"time"
)

type App struct {
	serverConfig *config.ServerConf
	userService *service.UserService
	contentService *service.ContentService
}

func ErrorHelper(err error, statusCode int) gin.H {
	return gin.H{
		"error": err,
		"code": statusCode,
		"data": "",
	}
}

type M map[string]interface{}

func SetData(key string, val interface{}) M {
	return M{
		key: val,
	}
}

func SuccessResp() M {
	return SetData("status", "success")
}

func RespHelper(m M) gin.H {
	return gin.H{
		"code": utils.SUCCESS,
		"data": m,
		"error": "",
	}
}

func Init() (app App) {
	log.Info("Initializing APP")
	var sc config.ServerConf
	serverConfig := sc.GetConfig()
	app = App{
		serverConfig: serverConfig,
	}

	config.InitDB(*serverConfig)
	app.userService = service.NewUserService(config.GetMySQLClient())
	app.contentService = service.NewContentService(config.GetMySQLClient())
	config.InitScheme()
	return
}

func (app *App) UserRegister(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	log.Info(user)
	err := app.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorHelper(err, utils.REGISTER_FAIL))
		return
	}
	c.JSON(http.StatusOK, RespHelper(SuccessResp()))
}

func (app *App) GetRankedUsers(c *gin.Context) {
	users, err := app.userService.GetRankedUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorHelper(err, utils.FETCH_CONTENT_FAIL))
		return
	}
	c.JSON(http.StatusOK, RespHelper(SetData("users", users)))
}

func (app *App) UserLogin(c *gin.Context)  {
	var user model.User
	c.BindJSON(&user)
	err := app.userService.ValidateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorHelper(err, utils.LOGIN_FAIL))
		return
	} else {
		info, err := app.userService.GetUserProfileByName(user.Username)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorHelper(err, utils.LOGIN_FAIL))
			return
		}
		sessionId := uuid.FromTime(time.Now()).String()
		config.RedisClient.Set(user.Username + "-session_id", sessionId, time.Hour*12)
		c.SetCookie("username", user.Username, 3600, "/", "", false, false)
		c.SetCookie("session_id", sessionId, 3600, "/", "", false, false)
		c.JSON(http.StatusOK, RespHelper(SetData("data", info)))
	}
}

func (app *App) CheckLoginStatus(c *gin.Context) {
	username, err := c.Cookie("username")
	if err != nil {
		c.String(http.StatusUnauthorized, "verification failed")
		return
	}
	user, err := app.userService.GetUserProfileByName(username)
	if err != nil {
		c.String(http.StatusUnauthorized, "verification failed")
		return
	}
	c.JSON(http.StatusOK, RespHelper(SetData("data", user)))
}

func (app *App) CreateContent(c *gin.Context) {
	var content model.Content
	c.BindJSON(&content)
	err := app.contentService.ValidateContent(content)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorHelper(err, utils.INVALID_CONTENT))
		return
	} else {
		err = app.contentService.CreateContent(content)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorHelper(err, utils.CREATE_CONTENT_FAIL))
			return
		}
		err = app.userService.UpdateUserGrades(content.Author, content.Category)
		if err != nil {
			log.Warningf("update user grade found error:%v", err)
			return
		}
	}
	c.JSON(http.StatusOK, RespHelper(SuccessResp()))
}

func (app *App) GetContentById(c *gin.Context) {
	pathId := c.Param("id")
	id, err := strconv.Atoi(pathId)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorHelper(err, utils.INVALID_PATH_PARAMETER))
		return
	}
	content, err := app.contentService.GetContentById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorHelper(err, utils.FETCH_CONTENT_FAIL))
		return
	}
	err = app.contentService.ContentVisited(id)
	if err != nil {
		log.Warningf("update content views found error:%v", err)
	}
	c.JSON(http.StatusOK, RespHelper(SetData("data", content)))
}

func (app *App) GetRankedContent(c *gin.Context) {
	content, err := app.contentService.GetRankedContent()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorHelper(err, utils.FETCH_CONTENT_FAIL))
		return
	}
	c.JSON(http.StatusOK, RespHelper(SetData("data", content)))
}

func (app *App) GetTopContent(c *gin.Context) {
	content, err := app.contentService.GetTopContent()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorHelper(err, utils.FETCH_CONTENT_FAIL))
		return
	}
	c.JSON(http.StatusOK, RespHelper(SetData("data", content)))
}

func (app *App) GetContentByCat(c *gin.Context) {
	cat := c.Param("cat")
	content, err := app.contentService.GetContentByCat(cat)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorHelper(err, utils.FETCH_CONTENT_FAIL))
		return
	}
	c.JSON(http.StatusOK, RespHelper(SetData("data", content)))
}

func (app *App) GetUserProfileById(c *gin.Context) {
	pathId := c.Param("id")
	id, err := strconv.Atoi(pathId)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorHelper(err, utils.INVALID_PATH_PARAMETER))
		return
	}
	user, err := app.userService.GetUserProfileById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorHelper(err, utils.FETCH_PROFILE_FAIL))
		return
	}
	c.JSON(http.StatusOK, RespHelper(SetData("data", user)))
}

func (app *App) GetUserProfileByName(c *gin.Context) {
	username := c.Param("username")
	user, err := app.userService.GetUserIdByName(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorHelper(err, utils.FETCH_PROFILE_FAIL))
		return
	}
	c.JSON(http.StatusOK, RespHelper(SetData("data", user)))
}

func (app *App) UpdateUserProfile(c *gin.Context) {
	var user model.User
	log.Info(user)
	err := c.BindJSON(&user)
	log.Info(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorHelper(err, utils.UPDATE_PROFILE_FAIL))
		return
	}
	log.Info(user)
	err = app.userService.UpdateUserProfile(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorHelper(err, utils.UPDATE_PROFILE_FAIL))
		return
	}
	newUser, err := app.userService.GetUserProfileById(user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorHelper(err, utils.UPDATE_PROFILE_FAIL))
		return
	}
	c.JSON(http.StatusOK, RespHelper(SetData("data", newUser)))
}

func main() {
	app := Init()
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = app.serverConfig.Http.AllowOrigin
	fmt.Print(app.serverConfig.Http.AllowOrigin)
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	guestRouter := r.Group("/g")
	{
		guestRouter.POST("/register", app.UserRegister)
		guestRouter.POST("/login", app.UserLogin)
		guestRouter.GET("/content/:id", app.GetContentById)
		guestRouter.GET("/contents", app.GetRankedContent)
		guestRouter.GET("/contents/:cat", app.GetContentByCat)
		guestRouter.GET("/profile/id/:id", app.GetUserProfileById)
		guestRouter.GET("/profile/username/:username", app.GetUserProfileByName)
	}

	clientRouter := r.Group("/u")
	clientRouter.Use(middleware.Auth())
	{
		clientRouter.POST("/test", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "still",
			})
		})
		clientRouter.POST("/content", app.CreateContent)
		clientRouter.GET("/status", app.CheckLoginStatus)
		clientRouter.GET("/contents", app.GetTopContent)
		clientRouter.GET("/rank", app.GetRankedUsers)
		clientRouter.POST("/profile/update", app.UpdateUserProfile)
	}

	r.Run(":" + strconv.Itoa(app.serverConfig.Http.Port))
}
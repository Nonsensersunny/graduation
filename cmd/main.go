package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb/uuid"
	"graduation/internal/config"
	"graduation/internal/log"
	"graduation/internal/middleware"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
}

func (app *App) UserLogin(c *gin.Context)  {
	var user model.User
	c.BindJSON(&user)
	err := app.userService.ValidateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		sessionId := uuid.FromTime(time.Now()).String()
		config.RedisClient.Set(user.Username + "-session_id", sessionId, time.Hour)
		c.SetCookie("username", user.Username, 3600, "/", "", false, false)
		c.SetCookie("session_id", sessionId, 3600, "/", "", false, false)
		c.String(http.StatusOK, "success")
	}
}

func (app *App) CreateContent(c *gin.Context) {
	var content model.Content
	c.BindJSON(&content)
	err := app.contentService.ValidateContent(content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err = app.contentService.CreateContent(content)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (app *App) GetContentById(c *gin.Context) {
	pathId := c.Query("id")
	id, err := strconv.Atoi(pathId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	content, err := app.contentService.GetContentById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"content": content,
	})
}

func main() {
	app := Init()
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins=[]string{app.serverConfig.Http.AllowOrigin}
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	guestRouter := r.Group("/g")
	{
		guestRouter.POST("/register", app.UserRegister)
		guestRouter.POST("/login", app.UserLogin)
	}

	clientRouter := r.Group("/u")
	clientRouter.Use(middleware.Auth())
	{
		clientRouter.POST("/test", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "still",
			})
		})
		clientRouter.GET("/content/:id", app.GetContentById)
		clientRouter.POST("/content", app.CreateContent)
	}

	r.Run(":" + strconv.Itoa(app.serverConfig.Http.Port))
	fmt.Println("main")
}
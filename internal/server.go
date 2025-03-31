package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nrednav/cuid2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

func (s *Server) Run() error {
	var err error
	s.db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := s.db.AutoMigrate(&User{}, &Location{}); err != nil {
		return err
	}

	router := gin.Default()

	router.StaticFS("/", http.Dir("../ui"))

	api := router.Group("api")
	api.POST("user", s.CreateUser)
	api.POST("location", s.CreateLocation)

	return router.Run()
}

func (s *Server) CreateUser(c *gin.Context) {
	var user User
	if err := c.Bind(&user); err != nil {
		c.Status(http.StatusBadRequest)
	}

	userid, err := c.Cookie("userid")
	if err != nil && cuid2.IsCuid(userid) {
		user.ID = userid
	} else {
		user.ID = cuid2.Generate()
	}

	// TODO: make sure data is not overwritten if user already exists
	if s.db.Save(&user).Error != nil {
		c.Status(http.StatusBadRequest)
	}

	c.SetCookie("userid", user.ID, 0, "/", "localhost", false, true)
}

func (s *Server) CreateLocation(c *gin.Context) {
	var location Location

	if err := c.BindJSON(&location); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	userid, err := c.Cookie("userid")
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	location.UserID = userid
	location.ID = cuid2.Generate()

	if err := s.db.Save(&location).Error; err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

}

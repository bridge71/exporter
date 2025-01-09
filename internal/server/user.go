package server

import (
	"exporter/internal/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *Server) CreateRootHandler(c *gin.Context) {
	user := &models.User{}
	user.UserID = 1
	user.UserName = "root"
	user.Email = os.Getenv("rootEmail")
	user.Password = os.Getenv("rootPasswd")
	user.Priority = 2
	var err error
	user.Password, err = EncryptPassword(user.Password)

	isUser := &models.User{}
	s.db.FindUserByID(isUser, 1)
	if isUser.Email != "" {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "there is a first root",
		})
		return
	}

	err = s.db.Save(user)

	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to create root",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "create root successfully",
		})
	}
}

func (s *Server) FindUserHandler(c *gin.Context) {
	User := []models.User{}

	s.db.Find(&User)
	all := len(User)
	for i := 0; i < all; i++ {
		User[i].Password = ""
	}
	c.JSON(http.StatusOK, models.Message{
		Users: User,
	})
}

func (s *Server) SaveUserHandler(c *gin.Context) {
	user := &models.User{}
	user.Email = c.PostForm("Email")
	user.UserName = c.PostForm("UserName")
	user.EmplID = s.Str2Uint(c.PostForm("EmplID"))
	if user.Email == "" || user.UserName == "" || user.EmplID == 0 {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "不可为空",
		})
		return
	}
	if user.Password == "" {
		user.Password = os.Getenv("defaultPasswd")
	}
	isLong, message := CheckStringLen(*user, false)
	log.Printf("sss %s", user.Password)
	if isLong {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: message,
		})
		return
	}
	user.Password, _ = EncryptPassword(user.Password)

	err := s.db.Save(user)

	user.Password = ""
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "email may be illegal",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "register successfully",
			User:       *user,
		})
	}
}

func CheckStringLen(user models.User, flag bool) (bool, string) {
	if len(user.Password) > 36 || len(user.Password) == 0 {
		return true, "password length is illegal"
	}
	if len(user.Email) > 36 || len(user.Email) == 0 {
		return true, "email length is illegal"
	}
	if flag {
		return false, ""
	}
	if len(user.UserName) > 57 || len(user.UserName) == 0 {
		return true, "userName length is illegal"
	}
	return false, ""
}

func (s *Server) LoginHandler(c *gin.Context) {
	log.Printf("sss\n")
	user := &models.User{}
	err := c.ShouldBind(user)
	log.Printf("%v ss\n", *user)
	if err != nil {
		log.Printf("error in bind of user\n")
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of user",
		})
		return
	}
	isLong, message := CheckStringLen(*user, true)
	if isLong {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: message,
		})
		return
	}

	auth := &models.User{}
	s.db.FindUserByEmail(auth, user.Email)
	err = bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to login",
		})
	}
	auth.Password = ""
	c.JSON(http.StatusOK, models.Message{
		RetMessage: "login successfully",
		User:       *auth,
	})
}

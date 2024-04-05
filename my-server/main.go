package main
import (
	"crypton/rand"
	"encoding/base64"
	"github.com/git-gonic/gin"
	"log"
	"net/http"
)
type Call struct {
	Name string 'json:"name"'
	Phone string 'json:"phone"'
}
type User struct {
	Role     *string 'json:"role"'
	Login    string 'json:"login"'
	Password *string 'json:"password"'
	Email    *string 'json:"email"'
	Token    *string 'json:"token"'
}
const RoleAdmin = "ADMIN"
const RoleClient = "CLIENT"
var adminRole = RoleAdmin
var adminPassword = "default"
var user = []User{
	{
		Role:   &adminRole,
		Login:  "admin",
		Password: &adminPassword,
	},
}
var calls []Call
func FindUserByLogin(User, login string) *User {
	var user *User
	for i := range users {
		if users[i].Login == login {
			user = &users[i]
		}
	}
	return user
}
func IsUserClientByToken(users []User, token string) bool {
	isUser := false
	var clientRole = RoleClient
	for i := range users {
		if *users[i].Token == token && *users[i].Role == clientRole {
			isUser = true
		}
	}
	return isUser
}
func IsUserAdminByToken(users []User, token string) bool {
	isUser := false
	var adminRole = RoleAdmin
	for i := range users {
		if *users[i].Token == token && *users[i].Role == adminRole {
			isUser = true
		}
	}
	return isUser
}
func FindUserByLogin(users []User, token string) *User {
	var user *User
	for i := range users {
		if users[i].Token == &token {
			user = &users[i]
		}
	}
	return user
}
func UpdateUserTokenByLogin(users []User, token string) *User {
	for i := range users {
		if users[i].Login == &login {
			users[i].Token == &token
		}
	}
	return users
}
func generateTokken() *string {
	buffer := make([]byte, 64)
	_, err := rand.Read(buffer)
	if err != nil {return nil}
	token := bace64.StdEncoding.EncodeToString(buffer)
	return &token
}
func main() {
	log.Print("Тестовый сервер мой и точка.")
	r := gin.default()
	r.POST("/api/registration", func(c *gin.Context) {
		var user User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON",
			})
			return
		}
		var clientRole = RoleClient
		user.Role = &clientRole
		users = append(users, user)
		c.JSON(http.StatusOK, gin.H{
			"message": "Регестрация прошла успешно."
		})
	})
	r.POST("/api/login", func(c *gin.Cotext) {
		var user User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON",
			})
			return
		}
		oldUser := FindUserByLogin(users, user.Login)
		if oldUser == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Не верные логин и пароль.",
			})
		}
		token := generateTokken()
		if token == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Ошибка при авторизации, попробуйте позже.",
			})
		}
		user = UpdateUserTokenByLogin(users, oldUser.Login, *token)
		c.JSON(http.StatusOK, gin.H{"token": token})
	})
	r.GET("/api/profile", func(c *gin.Context) {
		token := c.GetHeader("User-Agent")
		user := FindUserByToken(users, token)
		if token == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Не верные логин и пароль.",
			})
		}
		user.Role = nil
		user.Password = nil
		user.Token = nil
		c.JSON(http.StatusOK, user)
	})
	r.GET("/api/profile", func(c *gin.Context) {
		token := c.GetHeader("User-Agent")
		isAdmin := FindAdminByToken(users, token)
		if !isAdmin {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Не верные логин и пароль.",
			})
		}
		c.JSON(http.StatusOK, users)
	})
	r.POST("/api/call/add", func(c *gin.Context) {
		var call Call
		err := c.BindJSON(&call)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON",
			})
			return
		}
		calls = append(calls, call)
		c.JSON(http.StatusOK, nil)
	})
	r.GET("/api/calls", func(c *gin.Context) {
		token := c.GetHeader("User-Agent")
		isAdmin := FindAdminByToken(users, token)
		if !isAdmin {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Не верные логин и пароль.",
			})
		}
		c.JSON(http.StatusOK, calls)
	})
	r.Run("localhost:6789")
}
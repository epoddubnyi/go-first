package users

import (
	"fmt"
	"log"
	"my-first-go-project/internal/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email"`
}

// Get godoc
// @Summary Получить пользователей или одного пользователя
// @Description Если userId указан, возвращает одного пользователя, иначе — всех пользователей
// @Tags users
// @Accept json
// @Produce json
// @Param userId query int false "ID пользователя"
// @Success 200 {object} users.User
// @Router /users/get [get]
func Get(c *gin.Context) {
	userIDStr := c.Query("userId")

	if userIDStr != "" {
		id, err := strconv.Atoi(userIDStr) // convert string to
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userId"})
			return
		}

		var user User
		err = db.Db.QueryRow("SELECT id, name, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		c.JSON(http.StatusOK, user)
		return
	}

	users, err := db.Db.Query("SELECT id, name, email from users")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer users.Close()

	allData := []User{}

	for users.Next() {
		var id int
		var name string
		var email string
		if err := users.Scan(&id, &name, &email); err != nil {
			log.Fatal(err)
		}
		allData = append(allData, User{id, name, email})
	}

	c.JSON(http.StatusCreated, allData)
}

// Save godoc
// @Summary Создать пользователя
// @Description Создает нового пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param name query string false "Имя пользователя"
// @Param email query string false "Эмейл пользователя"
// @Success 200 {object} users.User
// @Router /users/create [post]
func Create(c *gin.Context) {
	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email"
	err := db.Db.QueryRow(query, newUser.Name, newUser.Email).Scan(&newUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

// Update godoc
// @Summary Обновить пользователя
// @Description Обновляет поля пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param name query string false "Новое имя пользователя"
// @Param email query string false "Новый эмейл пользователя"
// @Param userId query string false "ID пользователя"
// @Success 200 {object} users.User
// @Router /users/update [put]
func Update(c *gin.Context) {
	var req User
	userId := c.Query("userId")

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newName := req.Name
	newEmail := req.Email

	fmt.Println(userId, newName, newEmail)

	if userId != "" && newName != "" {
		userIdInt, err := strconv.Atoi(userId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updatedUser User
		err = db.Db.QueryRow("UPDATE users SET name=$1, email=$3 WHERE id=$2 RETURNING id, name, email", newName, userIdInt, newEmail).Scan(&updatedUser.ID, &updatedUser.Name, &updatedUser.Email)

		fmt.Println(updatedUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(updatedUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, updatedUser)
	}
}

package handler

import (
	todo "ToDoApp"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
	"strconv"
)

var store = sessions.NewCookieStore([]byte("your-secret-key"))

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type signInInput struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.Bind(&input); err != nil {
		newJSONResponse(c, http.StatusBadRequest, "error", err.Error())
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newJSONResponse(c, http.StatusInternalServerError, "error", err.Error())
		//c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Set the JWT token in a cookie
	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)

	// Извлекаем пользователя из базы данных
	name, id, err := h.services.Authorization.GetIdName(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Сохраняем имя пользователя в сессии
	session, _ := store.Get(c.Request, "session-name")
	session.Values["username"] = name
	session.Values["id"] = strconv.Itoa(id)
	session.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, statusResponse{"ok"})

}

func (h *Handler) signOut(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Вы успешно вышли из системы"})
}

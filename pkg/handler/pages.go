package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) mainPage(c *gin.Context) {
	// Получаем имя пользователя из сессии
	session, _ := store.Get(c.Request, "session-name")
	name := session.Values["username"].(string)
	id := session.Values["id"].(string)

	data := map[string]interface{}{
		"name": name,
		"id":   id,
	}
	c.HTML(http.StatusOK, "main.html", data)
}

func (h *Handler) usersPage(c *gin.Context) {
	usersStatuses, err := h.services.Status.GetUsersStatuses()
	if err != nil {
		newJSONResponse(c, http.StatusInternalServerError, "error", err.Error())
		return
	}
	c.HTML(http.StatusOK, "users.html", gin.H{"usersStatuses": usersStatuses})

}

func (h *Handler) adminPage(c *gin.Context) {

	c.HTML(http.StatusOK, "console.html", gin.H{})
}

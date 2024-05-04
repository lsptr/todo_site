package handler

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) mainPage(c *gin.Context) {
	// Получаем имя пользователя из сессии
	session, _ := store.Get(c.Request, "session-name")
	name := session.Values["username"].(string)

	tmpl, err := template.ParseFiles("templates/main.html")
	if err != nil {
		log.Println(err)
		return
	}

	data := map[string]interface{}{
		"name": name,
	}

	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		log.Println(err)
		return
	}
}

func (h *Handler) usersPage(c *gin.Context) {
	usersStatuses, err := h.services.Status.GetUsersStatuses()
	if err != nil {
		newJSONResponse(c, http.StatusInternalServerError, "error", err.Error())
		return
	}

	c.HTML(http.StatusOK, "users.html", gin.H{"usersStatuses": usersStatuses})

}

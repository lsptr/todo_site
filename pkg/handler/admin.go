package handler

import (
	todo "ToDoApp"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) adminDelete(c *gin.Context) {
	// Извлечение userId из запроса
	var usersStatuses todo.UserStatusPage
	if err := c.BindJSON(&usersStatuses); err != nil {
		newJSONResponse(c, http.StatusBadRequest, "error", err.Error())
		return
	}

	fmt.Println(usersStatuses.Id, usersStatuses.Name, usersStatuses.Role)
	// Проверка роли пользователя
	if usersStatuses.Role == "" {
		fmt.Println(usersStatuses.Id, "NO ROLE")
		// Удаление пользователя
		if err := h.services.Admin.DeleteUser(usersStatuses.Id); err != nil {
			fmt.Println("err!")
			newJSONResponse(c, http.StatusInternalServerError, "error", err.Error())
			return
		}
		fmt.Println(usersStatuses.Id, "DELETED!")
		c.JSON(http.StatusOK, gin.H{})
	} else {
		newJSONResponse(c, http.StatusBadRequest, "error", "Пользователь не может быть удален, так как у него есть роль")
	}
}

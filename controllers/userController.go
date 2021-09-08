package controllers

import (
	"aysf/day6r1/lib/database"
	"aysf/day6r1/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return c.String(http.StatusBadRequest, "error request")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "status ok",
		"user":    users,
	})
}

func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := database.GetUser(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "query success",
		"user":    user,
	})
}

func CreateUserController(c echo.Context) error {
	userInput := new(models.User)
	if err := c.Bind(userInput); err != nil {
		return err
	}
	user, err := database.CreateUser(userInput)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})

}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	userInput := new(models.User)
	if err := c.Bind(userInput); err != nil {
		return err
	}
	user, err := database.UpdateUser(id, userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to update data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully updated",
		"user":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := database.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "failed to delete ",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "successfully delete data user " + c.Param("id"),
	})
}

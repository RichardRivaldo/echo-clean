package controllers

import (
	"echo-clean/src/configs"
	"echo-clean/src/dtos"
	"echo-clean/src/models"
	"echo-clean/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllMembersHandler(c echo.Context) error {
	result, err := services.GetAllMembers()
	if err != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed retrieving all members!",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := dtos.Response{
		Status:  http.StatusOK,
		Message: "Successfully retrieved all members!",
		Data:    &echo.Map{"data": result},
	}

	return c.JSON(http.StatusOK, response)
}

func CreateNewMemberHandler(c echo.Context) error {
	member := models.Member{}
	if c.Bind(&member) != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed creating new member!",
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if validationErr := configs.Validator.Struct(&member); validationErr != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed creating new member!",
			Data: &echo.Map{
				"data": validationErr.Error(),
			},
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	data, result := services.CreateNewMember(member)
	if result.Error != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed creating new member!",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := dtos.Response{
		Status:  http.StatusCreated,
		Message: "Successfully created new member!",
		Data:    &echo.Map{"data": data},
	}

	return c.JSON(http.StatusCreated, response)
}

func UpdateMemberHandler(c echo.Context) error {
	member := models.Member{}
	if c.Bind(&member) != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed updating new member!",
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	id := c.Param("member_id")
	result := services.UpdateMember(id, member)
	if result.Error != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed updating new member!",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.NoContent(http.StatusNoContent)
}

func DeleteMember(c echo.Context) error {
	id := c.Param("member_id")
	result := services.DeleteMember(id)
	if result.Error != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed deleting new member!",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.NoContent(http.StatusNoContent)
}

func AddMemberRoutes(e *echo.Group) {
	root := "/members"

	e.GET(root, GetAllMembersHandler)
	e.POST(root, CreateNewMemberHandler)
	e.PUT(root+"/:member_id", UpdateMemberHandler)
	e.DELETE(root+"/:member_id", DeleteMember)
}

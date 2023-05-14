package controllers

import (
	"echo-clean/src/dtos"
	"echo-clean/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProductReviewHandler(c echo.Context) error {
	id := c.Param("product_id")
	data, result := services.GetProductReview(id)
	if result.Error != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed retrieving product reviews!",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := dtos.Response{
		Status:  http.StatusOK,
		Message: "Successfully retrieved product reviews!",
		Data:    &echo.Map{"data": data},
	}

	return c.JSON(http.StatusOK, response)
}

func AddReviewRoutes(e *echo.Group) {
	root := "/reviews"

	e.GET(root+"/:product_id", GetProductReviewHandler)
}

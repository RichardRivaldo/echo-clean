package controllers

import (
	"echo-clean/src/configs"
	"echo-clean/src/dtos"
	"echo-clean/src/models"
	"echo-clean/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LikeHandler(c echo.Context) error {
	likeReview := models.LikeReview{}
	if c.Bind(&likeReview) != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed to like the review!",
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if validationErr := configs.Validator.Struct(&likeReview); validationErr != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed to like the review!",
			Data: &echo.Map{
				"data": validationErr.Error(),
			},
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	data, result := services.Like(likeReview)
	if result.Error != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to like the review!",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := dtos.Response{
		Status:  http.StatusCreated,
		Message: "Successfully liked the review!",
		Data:    &echo.Map{"data": data},
	}

	return c.JSON(http.StatusCreated, response)

}

func UnlikeHandler(c echo.Context) error {
	likeReview := models.LikeReview{}
	if c.Bind(&likeReview) != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed to unlike the review!",
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if validationErr := configs.Validator.Struct(&likeReview); validationErr != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed to unlike the review!",
			Data: &echo.Map{
				"data": validationErr.Error(),
			},
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	result := services.Unlike(likeReview)
	if result.Error != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to unlike the review!",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.NoContent(http.StatusNoContent)
}

func AddLikeReviewRoutes(e *echo.Group) {
	root := "/like_reviews"

	e.POST(root+"/like", LikeHandler)
	e.POST(root+"/unlike", UnlikeHandler)
}

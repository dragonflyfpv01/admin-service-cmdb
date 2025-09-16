package handler

import (
	"net/http"

	"sllpklls/admin-service/model"
	"sllpklls/admin-service/repository"
	"sllpklls/admin-service/security"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type InfraComponentHandler struct {
	InfraComponentRepo repository.InfraComponentRepo
}

func (i *InfraComponentHandler) GetInfraComponents(c echo.Context) error {
	// Lấy JWT claims từ context
	claims, err := security.GetClaimsFromContext(c)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Unauthorized: " + err.Error(),
			Data:       nil,
		})
	}

	// Kiểm tra role admin
	if !security.IsAdmin(claims) {
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    "Forbidden: Only admin can access this resource",
			Data:       nil,
		})
	}

	// Lấy danh sách infra components
	components, err := i.InfraComponentRepo.GetAllInfraComponents(c.Request().Context())
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get infra components: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Get infra components successfully",
		Data:       components,
	})
}

package handler

import (
	"net/http"
	"time"

	"sllpklls/admin-service/model"
	"sllpklls/admin-service/model/req"
	"sllpklls/admin-service/repository"
	"sllpklls/admin-service/security"

	validator "github.com/go-playground/validator/v10"
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

	// Parse pagination parameters từ query string
	var pagination model.PaginationRequest
	if err := c.Bind(&pagination); err != nil {
		log.Error("Error parsing pagination parameters:", err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid pagination parameters: " + err.Error(),
			Data:       nil,
		})
	}

	// Validate và set default values
	pagination.Validate()

	// Lấy danh sách infra components với phân trang
	components, totalCount, err := i.InfraComponentRepo.GetInfraComponentsPaginated(c.Request().Context(), pagination)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get infra components: " + err.Error(),
			Data:       nil,
		})
	}

	// Tạo response với thông tin phân trang
	paginationResponse := model.BuildPaginationResponse(pagination, totalCount)
	paginatedData := model.PaginatedResponse{
		Data:       components,
		Pagination: paginationResponse,
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Get infra components successfully",
		Data:       paginatedData,
	})
}

// GetAllInfraComponents lấy tất cả danh sách infra components không phân trang
func (i *InfraComponentHandler) GetAllInfraComponents(c echo.Context) error {
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

	// Lấy tất cả danh sách infra components
	components, err := i.InfraComponentRepo.GetAllInfraComponents(c.Request().Context())
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get all infra components: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Get all infra components successfully",
		Data:       components,
	})
}

// GetPendingInfraComponents lấy các infra components có status "Đang chờ"
func (i *InfraComponentHandler) GetPendingInfraComponents(c echo.Context) error {
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

	// Lấy danh sách infra components có status "Đang chờ"
	components, err := i.InfraComponentRepo.GetInfraComponentsByStatus(c.Request().Context(), "Đang chờ")
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get pending infra components: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Get pending infra components successfully",
		Data:       components,
	})
}

// UpdateInfraComponentStatus cập nhật status của infra component theo ID và hostname
func (i *InfraComponentHandler) UpdateInfraComponentStatus(c echo.Context) error {
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

	// Parse request body
	var req req.ReqUpdateStatus
	if err := c.Bind(&req); err != nil {
		log.Error("Error parsing request:", err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request format: " + err.Error(),
			Data:       nil,
		})
	}

	// Validate request
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Error("Validation error:", err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Validation error: " + err.Error(),
			Data:       nil,
		})
	}

	// Cập nhật status
	err = i.InfraComponentRepo.UpdateInfraComponentStatus(c.Request().Context(), req.ID, req.Hostname, req.NewStatus)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update infra component status: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Update infra component status successfully",
		Data: map[string]interface{}{
			"id":         req.ID,
			"hostname":   req.Hostname,
			"new_status": req.NewStatus,
		},
	})
}

// UpdateInfraComponent cập nhật thông tin của infra component theo ID
func (i *InfraComponentHandler) UpdateInfraComponent(c echo.Context) error {
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

	// Parse request body
	var req req.ReqUpdateInfraComponent
	if err := c.Bind(&req); err != nil {
		log.Error("Error parsing request:", err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request format: " + err.Error(),
			Data:       nil,
		})
	}

	// Validate request
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Error("Validation error:", err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Validation error: " + err.Error(),
			Data:       nil,
		})
	}

	// Tạo model InfraComponent từ request
	// created_at sẽ được set về thời gian hiện tại (thời gian sửa)
	component := model.InfraComponent{
		Hostname:        req.Hostname,
		DNS:             req.DNS,
		Description:     req.Description,
		PublicInternet:  req.PublicInternet,
		Class:           req.Class,
		IPAddress:       req.IPAddress,
		Subnet:          req.Subnet,
		Site:            req.Site,
		ITComponentType: req.ITComponentType,
		RequestType:     req.RequestType,
		AppID:           req.AppID,
		VLAN:            req.VLAN,
		AppName:         req.AppName,
		AppOwner:        req.AppOwner,
		Level:           req.Level,
		CIOwners:        req.CIOwners,
		IMCM:            req.IMCM,
		CreatedAt:       time.Now().Format("2006-01-02 15:04:05"), // Thời gian sửa
	}

	// Cập nhật thông tin component
	err = i.InfraComponentRepo.UpdateInfraComponent(c.Request().Context(), req.ID, component)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update infra component: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Update infra component successfully",
		Data: map[string]interface{}{
			"id":         req.ID,
			"hostname":   req.Hostname,
			"updated_at": component.CreatedAt,
		},
	})
}

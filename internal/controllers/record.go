package controllers

import (
	"errors"
	"git.epam.com/ryan_wang/go-web-service/internal/dto"
	customErrors "git.epam.com/ryan_wang/go-web-service/internal/errors"
	"git.epam.com/ryan_wang/go-web-service/internal/models"
	"git.epam.com/ryan_wang/go-web-service/internal/services"
	"git.epam.com/ryan_wang/go-web-service/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RecordController struct {
	service services.RecordService
}

func NewRecordController(service services.RecordService) *RecordController {
	return &RecordController{service: service}
}

func errorResponse(c echo.Context, status int, err error) error {
	return c.JSON(status, utils.ErrorMessage{err.Error()})
}

func successResponse(c echo.Context, status int, data any) error {
	return c.JSON(status, data)
}

// Create godoc
// @Summary      Create a record
// @Tags         records
// @Accept       json
// @Produce      json
// @Param        record body dto.CreateRecordRequest true "Create Record Request"
// @Success      201  {object}  models.Record
// @Failure      400  {object}  utils.ErrorMessage
// @Failure      500  {object}  utils.ErrorMessage
// @Router       /records [post]
func (ctrl *RecordController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var req dto.CreateRecordRequest
	var record *models.Record
	var err error

	if err = c.Bind(&req); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	if err = c.Validate(&req); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	if record, err = ctrl.service.Create(ctx, req.Load()); err != nil {
		return errorResponse(c, http.StatusInternalServerError, err)
	}

	return successResponse(c, http.StatusCreated, record)
}

// Update godoc
// @Summary      Update a record
// @Tags         records
// @Accept       json
// @Produce      json
// @Param        id path int true "Update Record ID"
// @Param        record body dto.UpdateRecordRequest true "Update Record Request"
// @Success      200  {object}  models.Record
// @Failure      400  {object}  utils.ErrorMessage
// @Failure      404  {object}  utils.ErrorMessage
// @Failure      500  {object}  utils.ErrorMessage
// @Router       /records/{id} [put]
func (ctrl *RecordController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	var req dto.UpdateRecordRequest
	var record *models.Record
	var err error

	if err = c.Bind(&req); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	if err = c.Validate(&req); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	if record, err = ctrl.service.Update(ctx, req.Load()); err != nil {
		var errNotFound *customErrors.RecordNotFoundError
		if errors.As(err, &errNotFound) {
			return errorResponse(c, http.StatusNotFound, errNotFound)
		}
		return errorResponse(c, http.StatusInternalServerError, err)
	}

	return successResponse(c, http.StatusOK, record)
}

// Delete godoc
// @Summary      Delete a record
// @Tags         records
// @Param        id path int true "Delete Record ID"
// @Success      204
// @Failure      400  {object}  utils.ErrorMessage
// @Failure      404  {object}  utils.ErrorMessage
// @Failure      500  {object}  utils.ErrorMessage
// @Router       /records/{id} [delete]
func (ctrl *RecordController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	var req dto.DeleteRecordRequest
	var err error

	if err = c.Bind(&req); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	if err = ctrl.service.Delete(ctx, req.ID); err != nil {
		var errNotFound *customErrors.RecordNotFoundError
		if errors.As(err, &errNotFound) {
			return errorResponse(c, http.StatusNotFound, errNotFound)
		}
		return errorResponse(c, http.StatusInternalServerError, err)
	}

	return successResponse(c, http.StatusNoContent, nil)
}

// Get godoc
// @Summary      Get a record by ID
// @Tags         records
// @Param        id path int true "Get Record ID"
// @Success      200  {object}  models.Record
// @Failure      400  {object}  utils.ErrorMessage
// @Failure      404  {object}  utils.ErrorMessage
// @Failure      500  {object}  utils.ErrorMessage
// @Router       /records/{id} [get]
func (ctrl *RecordController) Get(c echo.Context) error {
	ctx := c.Request().Context()
	var req dto.GetRecordRequest
	var record *models.Record
	var err error

	if err = c.Bind(&req); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	if record, err = ctrl.service.Get(ctx, req.ID); err != nil {
		var errNotFound *customErrors.RecordNotFoundError
		if errors.As(err, &errNotFound) {
			return errorResponse(c, http.StatusNotFound, errNotFound)
		}
		return errorResponse(c, http.StatusInternalServerError, err)
	}

	return successResponse(c, http.StatusOK, record)
}

// Query godoc
// @Summary      Get records list with pagination and search
// @Description  Returns a page of records with specified page number and size.<br>
// @Description  Currently only supports search by displayName.
// @Tags         records
// @Param        displayName query string false "Search by displayName"
// @Param        pageNum     query int    false "Page number (default is 0)"
// @Param        pageSize    query int    false "Page size (default is 10)"
// @Success      200  {object}  []models.Record
// @Failure      400  {object}  utils.ErrorMessage
// @Failure      500  {object}  utils.ErrorMessage
// @Router       /records [get]
func (ctrl *RecordController) Query(c echo.Context) error {
	ctx := c.Request().Context()
	var req dto.QueryRecordRequest
	var records []*models.Record
	var err error

	if err = c.Bind(&req); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	if records, err = ctrl.service.Query(ctx, req.DisplayName, req.PageNum, req.PageSize); err != nil {
		return errorResponse(c, http.StatusInternalServerError, err)
	}

	return successResponse(c, http.StatusOK, records)
}

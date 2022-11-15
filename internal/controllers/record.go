package controllers

import (
	"errors"
	"git.epam.com/ryan_wang/crud-demo/internal/dto"
	customErrors "git.epam.com/ryan_wang/crud-demo/internal/errors"
	"git.epam.com/ryan_wang/crud-demo/internal/models"
	"git.epam.com/ryan_wang/crud-demo/internal/services"
	"git.epam.com/ryan_wang/crud-demo/internal/utils"
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
	return c.JSON(status, utils.ResponseBody{
		Success: false,
		Errors:  err.Error(),
	})
}

func successResponse(c echo.Context, status int, data any) error {
	return c.JSON(status, utils.ResponseBody{
		Success: true,
		Data:    data,
	})
}

// Create godoc
// @Summary      Create a record
// @Tags         records
// @Accept       json
// @Produce      json
// @Param        record   body   object  true  "Request Body"
// @Success      201  {object}  utils.ResponseBody{success=bool,data=models.Record}
// @Failure      400  {object}  utils.ResponseBody{success=bool,errors=string}
// @Failure      500  {object}  utils.ResponseBody{success=bool,errors=string}
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

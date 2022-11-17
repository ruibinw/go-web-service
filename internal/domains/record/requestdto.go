package record

import (
	"git.epam.com/ryan_wang/go-web-service/internal/models"
)

type CreateRecordRequest struct {
	Url         string `json:"url"          validate:"required" example:"/url"`
	DisplayName string `json:"display_name" validate:"required" example:"name"`
	Description string `json:"description"  example:"This is description"`
}

type UpdateRecordRequest struct {
	ID          int64  `json:"id,omitempty" param:"id" validate:"required" swaggerignore:"true"`
	Url         string `json:"url"          validate:"required" example:"/url"`
	DisplayName string `json:"display_name" validate:"required" example:"name"`
	Description string `json:"description"  example:"This is a description"`
}

type DeleteRecordRequest struct {
	ID int64 `param:"id" validate:"required"`
}

type GetRecordRequest struct {
	ID int64 `param:"id" validate:"required"`
}

type QueryRecordRequest struct {
	DisplayName string `query:"displayName"`
	PageNum     int    `query:"pageNum"`
	PageSize    int    `query:"pageSize"`
}

func (req *CreateRecordRequest) Load() *models.Record {
	return &models.Record{
		Url:         req.Url,
		DisplayName: req.DisplayName,
		Description: req.Description,
	}
}

func (req *UpdateRecordRequest) Load() *models.Record {
	return &models.Record{
		ID:          req.ID,
		Url:         req.Url,
		DisplayName: req.DisplayName,
		Description: req.Description,
	}
}

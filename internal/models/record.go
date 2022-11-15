package models

import "time"

type Record struct {
	// ID of a record
	ID int64 `json:"id"             example:"1"`
	// Url of a record
	Url string `json:"url"            example:"/record/learning-go"`
	// DisplayName of a record
	DisplayName string `json:"display_name"   example:"Learning Go"`
	// Description of a record
	Description string `json:"description"    example:"A book to learn Go"`
	// CreatedTime of a record
	CreatedTime time.Time `json:"created_time"   example:"2022-10-29T18:31:22.378373+08:00" format:"date-time"`
	// UpdatedTime of a record
	UpdatedTime time.Time `json:"updated_time"   example:"2022-10-29T18:31:22.378373+08:00" format:"date-time"`
}

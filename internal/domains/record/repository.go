package record

import (
	"context"
	customErrors "git.epam.com/ryan_wang/go-web-service/internal/errors"
	"git.epam.com/ryan_wang/go-web-service/internal/models"
	"gorm.io/gorm"
)

const (
	FirstPage       = 0
	DefaultPageSize = 10
)

// Repository interface defines how to operate Records in database
type Repository interface {
	// Create inserts a new record
	Create(ctx context.Context, record *models.Record) (*models.Record, error)
	// Update updates the record with specified ID (ID is in record parameter)
	Update(ctx context.Context, record *models.Record) (*models.Record, error)
	// Delete removes the record with specified ID
	Delete(ctx context.Context, id int64) error
	// Get returns the album with specified ID
	Get(ctx context.Context, id int64) (*models.Record, error)
	// Query returns the list of records with specified page number and size and query condition
	// Currently only supports display_name as a query condition
	Query(ctx context.Context, displayName string, pageNum int, pageSize int) ([]*models.Record, error)
}

// repositoryImpl using gorm to interact with the database
type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}

func (repo *repositoryImpl) Create(ctx context.Context, record *models.Record) (*models.Record, error) {
	//ID is auto inclement
	res := repo.db.WithContext(ctx).Create(record)
	return record, res.Error
}

func (repo *repositoryImpl) Update(ctx context.Context, record *models.Record) (*models.Record, error) {
	res := repo.db.WithContext(ctx).Save(record)
	return record, res.Error
}

func (repo *repositoryImpl) Delete(ctx context.Context, id int64) error {
	res := repo.db.WithContext(ctx).Delete(&models.Record{}, id)
	return res.Error
}

func (repo *repositoryImpl) Get(ctx context.Context, id int64) (*models.Record, error) {
	record := &models.Record{ID: id}
	res := repo.db.WithContext(ctx).Find(record)
	if res.RowsAffected == 0 {
		return nil, customErrors.NewRecordNotFoundError(id)
	}
	return record, nil
}

func (repo *repositoryImpl) Query(ctx context.Context, displayName string, pageNum int, pageSize int) ([]*models.Record, error) {
	var records []*models.Record
	res := repo.db.WithContext(ctx).
		Scopes(
			condition(displayName),
			pagination(pageNum, pageSize)).
		Find(&records)

	return records, res.Error
}

// to set where condition in sql if keyword not empty
func condition(keyword string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(keyword) > 0 {
			db.Where("display_name like ?", "%"+keyword+"%")
		}
		return db
	}
}

func pagination(pageNum int, pageSize int) func(*gorm.DB) *gorm.DB {
	if pageNum < 0 {
		pageNum = FirstPage
	}
	if pageSize == 0 {
		pageSize = DefaultPageSize
	}
	if pageSize < 1 {
		pageSize = 1
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pageNum * pageSize).Limit(pageSize)
	}
}

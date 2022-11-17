package record

import (
	"context"
	"git.epam.com/ryan_wang/go-web-service/internal/models"
	"time"
)

type Service interface {
	Create(ctx context.Context, record *models.Record) (*models.Record, error)
	Update(ctx context.Context, record *models.Record) (*models.Record, error)
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.Record, error)
	Query(ctx context.Context, displayName string, pageNum int, pageSize int) ([]*models.Record, error)
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{repo: repo}
}

func (srv *serviceImpl) Create(ctx context.Context, record *models.Record) (*models.Record, error) {
	now := time.Now()
	record.CreatedTime = now
	record.UpdatedTime = now
	return srv.repo.Create(ctx, record)
}

func (srv *serviceImpl) Update(ctx context.Context, record *models.Record) (*models.Record, error) {
	//check if the record exists
	original, err := srv.repo.Get(ctx, record.ID)
	if err != nil {
		return nil, err //RecordNotFoundError(id)
	}
	record.CreatedTime = original.CreatedTime
	record.UpdatedTime = time.Now()
	return srv.repo.Update(ctx, record)
}

func (srv *serviceImpl) Delete(ctx context.Context, id int64) error {
	//check if the record exists
	if _, err := srv.repo.Get(ctx, id); err != nil {
		return err //RecordNotFoundError(id)
	}
	return srv.repo.Delete(ctx, id)
}

func (srv *serviceImpl) Get(ctx context.Context, id int64) (*models.Record, error) {
	return srv.repo.Get(ctx, id)
}

func (srv *serviceImpl) Query(ctx context.Context, displayName string, pageNum int, pageSize int) ([]*models.Record, error) {
	return srv.repo.Query(ctx, displayName, pageNum, pageSize)
}

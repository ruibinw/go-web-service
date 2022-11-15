package services

import (
	"context"
	customErrors "git.epam.com/ryan_wang/go-web-service/internal/errors"
	"git.epam.com/ryan_wang/go-web-service/internal/models"
	"git.epam.com/ryan_wang/go-web-service/internal/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func newRecord(name, url, desc string) *models.Record {
	return &models.Record{
		DisplayName: name,
		Url:         url,
		Description: desc,
	}
}

func TestRecordService_Create(t *testing.T) {
	mockRepo := mocks.NewMockRecordRepository(gomock.NewController(t))
	testService := NewRecordService(mockRepo)
	ctx := context.Background()

	toCreate := newRecord("name1", "url1", "description1")
	mockRepo.EXPECT().Create(ctx, toCreate).Return(toCreate, nil)

	created, err := testService.Create(ctx, toCreate)
	assert.NoError(t, err)
	assert.NotEmpty(t, created.CreatedTime)
	assert.Equal(t, created.CreatedTime, created.UpdatedTime)
}

func TestRecordService_Update(t *testing.T) {
	mockRepo := mocks.NewMockRecordRepository(gomock.NewController(t))
	testService := NewRecordService(mockRepo)
	ctx := context.Background()

	toUpdate := newRecord("name1", "url1", "description1")
	createdTime := time.Now().Add(time.Duration(-1) * time.Hour)
	toUpdate.ID = 1
	toUpdate.CreatedTime = createdTime
	toUpdate.UpdatedTime = createdTime

	mockRepo.EXPECT().Get(ctx, toUpdate.ID).Return(toUpdate, nil)
	mockRepo.EXPECT().Update(ctx, toUpdate).Return(toUpdate, nil)

	updated, err := testService.Update(ctx, toUpdate)
	assert.NoError(t, err)
	assert.True(t, updated.UpdatedTime.After(updated.CreatedTime))
}

func TestRecordService_Get(t *testing.T) {
	mockRepo := mocks.NewMockRecordRepository(gomock.NewController(t))
	testService := NewRecordService(mockRepo)
	ctx := context.Background()

	toGet := newRecord("name1", "url1", "description1")
	toGet.ID = 1

	//to get existing record
	mockRepo.EXPECT().Get(ctx, toGet.ID).Return(toGet, nil)
	got, err := testService.Get(ctx, toGet.ID)
	assert.NoError(t, err)
	assert.Equal(t, toGet, got)

	//to get not existing record
	expectedErr := customErrors.NewRecordNotFoundError(toGet.ID)
	mockRepo.EXPECT().Get(ctx, toGet.ID).Return(nil, expectedErr)
	got, err = testService.Get(ctx, toGet.ID)
	assert.Equal(t, expectedErr, err)
	assert.Empty(t, got)
}

func TestRecordServiceImpl_QueryByDisplayName(t *testing.T) {
	mockRepo := mocks.NewMockRecordRepository(gomock.NewController(t))
	testService := NewRecordService(mockRepo)
	ctx := context.Background()

	toGet := []*models.Record{newRecord("name1", "url1", "description1")}
	mockRepo.EXPECT().Query(ctx, "name1", 0, 10).Return(toGet, nil)
	result, err := testService.Query(ctx, "name1", 0, 10)
	assert.NoError(t, err)
	assert.Equal(t, toGet, result)
}

func TestRecordServiceImpl_QueryWithPagination(t *testing.T) {
	mockRepo := mocks.NewMockRecordRepository(gomock.NewController(t))
	testService := NewRecordService(mockRepo)
	ctx := context.Background()

	toGet := []*models.Record{
		newRecord("name1", "/url1", "description1"),
		newRecord("name2_go", "/url2", "description2"),
		newRecord("name3", "/url3", "description3"),
		newRecord("name4", "/url4", "description4"),
		newRecord("name5_golang", "/url5", "description5"),
	}
	mockRepo.EXPECT().Query(ctx, "", 0, 5).Return(toGet, nil)
	result, err := testService.Query(ctx, "", 0, 5)
	assert.NoError(t, err)
	assert.Equal(t, toGet, result)
}

func TestRecordServiceImpl_Delete(t *testing.T) {
	mockRepo := mocks.NewMockRecordRepository(gomock.NewController(t))
	testService := NewRecordService(mockRepo)
	ctx := context.Background()

	toDelete := &models.Record{ID: 1}

	//to delete existing record
	mockRepo.EXPECT().Get(ctx, toDelete.ID).Return(toDelete, nil)
	mockRepo.EXPECT().Delete(ctx, toDelete.ID).Return(nil)
	err := testService.Delete(ctx, toDelete.ID)
	assert.NoError(t, err)

	//to delete not existing record
	expectedErr := customErrors.NewRecordNotFoundError(toDelete.ID)
	mockRepo.EXPECT().Get(ctx, toDelete.ID).Return(nil, expectedErr)
	err = testService.Delete(ctx, toDelete.ID)
	assert.Equal(t, expectedErr, err)
}

package record

import (
	"context"
	"gorm.io/driver/sqlite"
	"os"
	"testing"
	"time"

	"git.epam.com/ryan_wang/go-web-service/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var db *gorm.DB

// Set up before testing
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	//open database connection
	db, _ = gorm.Open(sqlite.Open("file::memory:?cache=shared"))
	//create Record table
	db.AutoMigrate(&models.Record{})
	//init data
	db.Create([]*models.Record{
		newRecordForRepoTest("name1", "/url1", "description1"),
		newRecordForRepoTest("name2_go", "/url2", "description2"),
		newRecordForRepoTest("name3", "/url3", "description3"),
		newRecordForRepoTest("name4", "/url4", "description4"),
		newRecordForRepoTest("name5_golang", "/url5", "description5"),
	})
}

func TestRepository_Create(t *testing.T) {
	// start tx and rollback at the end in order to prevent effect on other unit tests
	tx := db.Begin()
	defer tx.Rollback()
	testRepo := NewRepository(tx)

	toCreate := newRecordForRepoTest("name", "url", "description")
	created, err := testRepo.Create(context.Background(), toCreate)
	assert.NoError(t, err)
	assert.NotEmpty(t, created.ID)
}

func TestRepository_Update(t *testing.T) {
	// start tx and rollback at the end in order to prevent effect on other unit tests
	tx := db.Begin()
	defer tx.Rollback()
	testRepo := NewRepository(tx)

	// get the record with id=1 for testing update operation
	toUpdate := &models.Record{ID: 1}
	tx.Find(toUpdate)
	toUpdate.DisplayName = "new_name"
	toUpdate.Url = "/new_url"
	toUpdate.Description = "new_description"

	updated, err := testRepo.Update(context.Background(), toUpdate)
	assert.NoError(t, err)

	// check if record in db is updated
	actual := &models.Record{ID: 1}
	tx.Find(actual)
	assert.Equal(t, updated, actual)
}

func TestRepository_Delete(t *testing.T) {
	// start tx and rollback at the end in order to prevent effect on other unit tests
	tx := db.Begin()
	defer tx.Rollback()
	testRepo := NewRepository(tx)

	err := testRepo.Delete(context.Background(), 1)
	assert.NoError(t, err)

	//check if the record is deleted from db
	result := tx.Find(&models.Record{ID: 1})
	assert.Equal(t, int64(0), result.RowsAffected)
}

func TestRepository_Get(t *testing.T) {
	testRepo := NewRepository(db)

	//Get existing record
	record, _ := testRepo.Get(context.Background(), 1)
	assert.NotEmpty(t, record)

	//Get not existing record
	record, _ = testRepo.Get(context.Background(), 100)
	assert.Empty(t, record)
}

func TestRepository_QueryByDisplayName(t *testing.T) {
	testRepo := NewRepository(db)

	records, err := testRepo.Query(context.Background(), "go", 0, 10)
	assert.NoError(t, err)
	assert.Len(t, records, 2)
	assert.Contains(t, records[0].DisplayName, "go")
	assert.Contains(t, records[1].DisplayName, "go")
}

func TestRepository_QueryWithPagination(t *testing.T) {
	testRepo := NewRepository(db)

	records, err := testRepo.Query(context.Background(), "", 0, 2)
	assert.NoError(t, err)
	assert.Len(t, records, 2)

	records, err = testRepo.Query(context.Background(), "", 0, 10)
	assert.NoError(t, err)
	assert.Len(t, records, 5)

	records, err = testRepo.Query(context.Background(), "", 1, 2)
	assert.NoError(t, err)
	assert.Len(t, records, 2)

	records, err = testRepo.Query(context.Background(), "", 1, 10)
	assert.NoError(t, err)
	assert.Len(t, records, 0)
}

func newRecordForRepoTest(name, url, desc string) *models.Record {
	now := time.Now()
	return &models.Record{
		DisplayName: name,
		Url:         url,
		Description: desc,
		CreatedTime: now,
		UpdatedTime: now,
	}
}

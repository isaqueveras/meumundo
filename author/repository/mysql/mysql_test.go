package mysql_test

import (
	"context"
	"testing"
	"time"

	repository "nossobr/author/repository/mysql"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "updated_at", "created_at"}).
		AddRow(1, "Iman Tumorang", time.Now(), time.Now())

	prep := mock.ExpectPrepare("SELECT id, name, created_at, updated_at FROM author WHERE id=\\?")
	userID := int64(1)
	prep.ExpectQuery().WithArgs(userID).WillReturnRows(rows)

	a := repository.NewMysqlAuthorRepository(db)

	anArticle, err := a.GetByID(context.TODO(), userID)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}

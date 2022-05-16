package model_test

import (
	"go_simple_api/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFindByIdMock(t *testing.T) {
	//MockDBを初期化する
	mockDB, mock, err := model.InitDBMock()

	if err != nil {
		t.Fatal(err)
	}

	//UserModelの初期化する
	userModel := model.NewUserModel(mockDB)

	//期待されるカラム
	expectedColumns := []string{
		"id",
		"first_name",
		"last_name",
	}

	//期待されるuser
	expect := model.User{
		ID:        1,
		FirstName: "Michael",
		LastName:  "Jordan",
	}

	//期待されるSQL
	query := "SELECT * FROM `users` WHERE `users`.`id` = ?"

	//DBの振る舞いを定義する
	mock.ExpectQuery(query).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(expectedColumns).AddRow(expect.ID, expect.FirstName, expect.LastName))

	s, _ := userModel.FindById(1)

	//実際のuserと期待されるuserを比較する
	assert.Equal(t, expect, s)
}

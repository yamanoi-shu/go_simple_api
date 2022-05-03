package model_test

import (
	"go_simple_api/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFindUserByIdMock(t *testing.T) {
	mockDB, mock, err := model.GetDBMock()

	if err != nil {
		t.Fatal(err)
	}

	userModel := model.NewUserModel(mockDB)

	expectedColumns := []string{
		"id",
		"first_name",
		"last_name",
	}

	expect := model.User{
		ID:        1,
		FirstName: "a",
		LastName:  "b",
	}

	query := "SELECT * FROM `users` WHERE `users`.`id` = ?"

	mock.ExpectQuery(query).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(expectedColumns).AddRow(1, "a", "b"))

	s, _ := userModel.FindUserById(1)

	assert.Equal(t, expect, s)
}

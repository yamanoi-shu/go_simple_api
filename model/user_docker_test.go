package model_test

import (
	"go_simple_api/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUserByIdDocker(t *testing.T) {
	pool, err := model.NewPool()

	if err != nil {
		t.Fatal(err)
	}

	//コンテナの起動、DB接続
	dbContainer, err := model.NewDBContainer(pool)

	if err != nil {
		pool.Purge(dbContainer.Resource)
		t.Fatal(err)
	}

	userModel := model.NewUserModel(dbContainer.DB)

	expect := model.User{
		ID:        1,
		FirstName: "Michael",
		LastName:  "Jordan",
	}

	user, err := userModel.FindById(1)

	if err != nil {
		pool.Purge(dbContainer.Resource)
		t.Fatal(err)
	}

	pool.Purge(dbContainer.Resource)
	assert.Equal(t, expect, user)
}

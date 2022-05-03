package model_test

import (
	"go_simple_api/model"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
)

func TestFindUserByIdDocker(t *testing.T) {
	pool, err := model.NewPool()

	if err != nil {
		t.Fatal(err)
	}

	dbContainer, err := model.NewDBContainer(pool)

	dbContainer.Resource.Exec([]string{"ls testdata"}, dockertest.ExecOptions{StdOut: os.Stdout})

	dbContainer.Resource.Exec([]string{"mysql", "-uroot", "-psecret", "test_db", "-e\"$(cat /testdata/insert_data.sql)\""}, dockertest.ExecOptions{StdOut: os.Stdout})

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

	user, err := userModel.FindUserById(1)

	if err != nil {
		pool.Purge(dbContainer.Resource)
		t.Fatal(err)
	}

	pool.Purge(dbContainer.Resource)
	assert.Equal(t, expect, user)
}

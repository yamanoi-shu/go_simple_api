package model

import (
	"fmt"
	"os"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		return nil, mock, err
	}

	db, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn:                      mockDB,
			SkipInitializeWithVersion: true,
		}),
		&gorm.Config{},
	)

	return db, mock, err
}

type DBContainer struct {
	Resource *dockertest.Resource
	DB       *gorm.DB
}

func NewPool() (*dockertest.Pool, error) {
	return dockertest.NewPool("")
}

func NewDBContainer(pool *dockertest.Pool) (*DBContainer, error) {
	workDir, _ := os.Getwd()

	// コンテナ起動
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8.0",
		Env:        []string{"MYSQL_ROOT_PASSWORD=secret"},
		Mounts: []string{
			workDir + "/testdata/create_tables.sql:/docker-entrypoint-initdb.d/create_tables.sql",
			workDir + "/testdata:/testdata",
		},
	})

	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("root:secret@(localhost:%s)/test_db?parseTime=true", resource.GetPort("3306/tcp"))

	var db *gorm.DB

	// exponential backoff でコンテナの起動を待つ
	if err := pool.Retry(func() error {
		var err error

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}

		sqlDB, err := db.DB()
		if err != nil {
			return err
		}

		return sqlDB.Ping()
	}); err != nil {
		pool.Purge(resource)
		return nil, err
	}

	return &DBContainer{Resource: resource, DB: db}, nil
}

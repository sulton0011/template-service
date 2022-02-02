package postgres

import (
	"home_work/task-service/config"
	"home_work/task-service/pkg/db"
	"home_work/task-service/pkg/logger"
	"log"
	"os"
	"testing"
)

var pgRepo *TaskRepo

func TestMain(m *testing.M) {
	cfg := config.Load()
	conn, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("test sqlx connetc error", logger.Error(err))
	}
	pgRepo = NewTaskRepo(conn)
	os.Exit(m.Run())
}

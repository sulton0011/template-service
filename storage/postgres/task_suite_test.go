package postgres

import (
	"home_work/task-service/config"
	pb "home_work/task-service/genproto/task"
	"home_work/task-service/pkg/db"
	"home_work/task-service/storage/repo"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TaskRepoTestSuite struct {
	suite.Suite
	CleanUpFunc func()
	Repository  repo.TaskStorageI
}

func (suite *TaskRepoTestSuite) SetupSuite() {
	pgPool, cleanup := db.ConnectDbForSuite(config.Load())
	suite.Repository = NewTaskRepo(pgPool)
	suite.CleanUpFunc = cleanup
}

// testing CRUD
func (suite *TaskRepoTestSuite) TestTaskCRUD() {
	task := pb.Task{
		Id:       "fe66a4bb-4c17-40a1-a5a2-ae8b3eb0cb62",
		Assignee: "Komron",
		Title:    "CRUD",
		Summary:  "Task crud",
		Deadline: "2022-01-31 17:00",
		Status:   "done",
	}
	_, err := suite.Repository.Delete(&pb.IdReq{Id: task.Id})
	suite.Nil(err)

	task1, err := suite.Repository.Create(&task)
	suite.Nil(err)
	suite.NotNil(task1)

	getTask, err := suite.Repository.Get(task.Id)
	suite.Nil(err)
	suite.NotNil(getTask, "task must not be nil")
	suite.Equal(task.Assignee, getTask.Assignee, "asignees must match")

	task.Title = "New title"
	updatedTask, err := suite.Repository.Update(&task)
	suite.Nil(err)
	suite.Equal(updatedTask.Title, task.Title, "titles must match")

	listTasks, err := suite.Repository.List(&pb.ListReq{
		Page:  1,
		Limit: 10,
	})
	suite.Nil(err)
	suite.NotNil(listTasks)

	listOverdueTasks, err := suite.Repository.ListOverdue(&pb.ListOverReq{
		Page:  1,
		Limit: 10,
		Time:  "2022-01-10 17:00",
	})
	suite.Nil(err)
	suite.NotNil(listOverdueTasks)
	// _, err = suite.Repository.Delete(&pb.IdReq{Id: task.Id})
	// suite.Nil(err)
	// suite.TearDownSuite()
}

func (suite *TaskRepoTestSuite) TearDownSuite() {
	suite.CleanUpFunc()
}

func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepoTestSuite))
}

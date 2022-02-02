package postgres

import (
	pb "home_work/task-service/genproto/task"
	"reflect"
	"testing"
)

func TaskTestRepoCreat(t *testing.T) {
	tasks := []struct {
		name string
		input pb.Task
		output pb.Task
	}{
		{
			name: "Create",
			input: pb.Task{
				Id:       "39bf6bdf-2e62-4687-bbc6-2bc1abd626e5",
				Assignee: "Assignee",
				Title:    "Title",
				Summary:  "Summary",
				Deadline: "2022-04-04",
				Status:   "Aktive",
			},
			output: pb.Task{
				Id:       "39bf6bdf-2e62-4687-bbc6-2bc1abd626e5",
				Assignee: "Assignee",
				Title:    "Title",
				Summary:  "Summary",
				Deadline: "2022-04-04T00:00:00Z",
				Status:   "Aktive",
			},
		},
	}

	for _, value := range tasks {
		t.Run(value.name, func(t *testing.T){
			got, err := pgRepo.Create(&value.input)
			if err != nil {
				t.Fatalf("Failed create task err:%v", err)
			}
			got.CreatedAt = ""
			if !reflect.DeepEqual(value.output, got) {
				t.Fatalf("%s: expected:%v, got:%v", value.name, value.output, got)
			}
		})
		pgRepo.Delete(&pb.IdReq{Id: value.input.Id})
	}
}
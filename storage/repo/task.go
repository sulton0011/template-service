package repo

import (
	pb "home_work/task-service/genproto/task"
)

//TaskStorageI ...
type TaskStorageI interface {
	Create(*pb.Task) (*pb.Task, error)
	Get(string) (*pb.Task, error)
	List(*pb.ListReq) (*pb.ListResp, error)
	Update(*pb.Task) (*pb.Task, error)
	Delete(*pb.IdReq) (*pb.EmptyResp, error)
	ListOverdue(*pb.ListOverReq) (*pb.ListOverResp, error)
}

package repo

import (
	pb "github.com/sulton0011/home_work/microservice/genproto/users"
)

//UserStorageI ...
type UserStorageI interface {
	Create(*pb.User) (*pb.User, error)
	Get(int64) (*pb.User, error)
	List(*pb.ListReq) (*pb.ListResp, error)
	Update(*pb.User) (*pb.User, error)
	Delete(*pb.IdReq) (*pb.EmptyResp, error)
}

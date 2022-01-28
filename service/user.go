package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	pb "github.com/sulton0011/home_work/microservice/genproto/users"
	l "github.com/sulton0011/home_work/microservice/pkg/logger"
	"github.com/sulton0011/home_work/microservice/storage"
)


//UserService ...
type UserService struct {
    storage storage.IStorage
    logger  l.Logger
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
    return &UserService{
        storage: storage.NewStoragePg(db),
        logger:  log,
    }
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) { 
    return nil, nil
}

func (s *UserService) Get(ctx context.Context, req *pb.IdReq) (*pb.User, error) {    
    return nil, nil
}

func (s *UserService) List(ctx context.Context, req *pb.ListReq) (*pb.ListResp, error) {
    return nil, nil
}

func (s *UserService) Update(ctx context.Context, req *pb.User) (*pb.User, error) {
    return nil, nil
}

func (s *UserService) Delete(ctx context.Context, req *pb.IdReq) (*pb.EmptyResp, error) {
    return nil, nil
}
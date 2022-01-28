package postgres

import (
    "github.com/jmoiron/sqlx"
    pb "github.com/sulton0011/home_work/microservice/genproto/users"
)

type userRepo struct {
    db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
    return &userRepo{db: db}
}


func (r *userRepo) Create(user *pb.User) (*pb.User, error) {
    return nil, nil
}


func (r *userRepo) Get(id int64) (*pb.User, error) {
    return nil, nil
}


func (r *userRepo) List(req *pb.ListReq) (*pb.ListResp, error) {
    return nil, nil
}


func (r *userRepo) Update(user *pb.User) (*pb.User, error) {
    return nil, nil
}

func (r *userRepo) Delete(id *pb.IdReq) (*pb.EmptyResp, error) {
    return nil, nil
}
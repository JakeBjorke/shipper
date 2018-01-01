package main

import (
	"context"

	pb "github.com/jakebjorke/shipper/user-service/proto/user"
)

type service struct {
	repo         Repository
	tokenService Authable
}

//Get is used to get a user
func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}

	res.User = user
	return nil
}

//GetAll gets all of the current system users.
func (srv *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}

	res.Users = users
	return nil
}

//Auth is used to authenticate a user.
func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	_, err := srv.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}

	//ignoring for the most part.
	res.Token = "testingabc"
	return nil
}

//Create is used to create a new user
func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	if err := srv.repo.Create(req); err != nil {
		return err
	}

	res.User = req
	return nil
}

//ValidateToken validates a token
func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	//ignoring for now...
	return nil
}

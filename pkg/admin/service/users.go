package service

import (
	"context"
	"fmt"

	adminpb "github.com/Shakezidin/pkg/admin/pb"
	pb "github.com/Shakezidin/pkg/admin/userclient/userpb"
)

func (a *AdminService) ViewUsersSVC(p *adminpb.AdminView) (*adminpb.AdminUsers, error) {

	var ctx = context.Background()
	responce, err := a.userClient.UserViewUsers(ctx, &pb.UserView{
		Id:     p.Id,
		Status: p.Status,
		Page:   p.Page,
	})
	if err != nil {
		fmt.Println("heey")
		return nil, err
	}
	var users []*adminpb.AdminUser
	for _, user := range responce.Users {
		users = append(users, &adminpb.AdminUser{
			Id:       user.Id,
			Name:     user.Name,
			Email:    user.Email,
			Phone:    user.Phone,
			Password: user.Password,
			Role:     user.Role,
		})
	}
	return &adminpb.AdminUsers{
		Users: users,
	}, nil
}

func (a *AdminService) ViewUserSVC(p *adminpb.AdminView) (*adminpb.AdminUser, error) {
	var ctx = context.Background()
	responce, err := a.userClient.UserViewUser(ctx, &pb.UserView{
		Id:     p.Id,
		Status: p.Status,
		Page:   p.Page,
	})
	if err != nil {
		return nil, err
	}

	return &adminpb.AdminUser{
		Id:    responce.Id,
		Name:  responce.Name,
		Email: responce.Email,
		Phone: responce.Phone,
		Role:  responce.Role,
	}, nil
}

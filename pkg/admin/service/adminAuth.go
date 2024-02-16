package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	clientpb "github.com/Shakezidin/pkg/admin/client/pb"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
	inter "github.com/Shakezidin/pkg/admin/repository/interface"
	interr "github.com/Shakezidin/pkg/admin/service/interface"
	pb "github.com/Shakezidin/pkg/admin/userclient/userpb"
	"github.com/Shakezidin/utils"
)

type AdminService struct {
	Repo       inter.RepoInterface
	CodClient  clientpb.CoordinatorClient
	userClient pb.UserClient
}

func (a *AdminService) LoginService(p *adminpb.AdminLogin) (*adminpb.AdminResponse, error) {
	admin, err := a.Repo.FetchAdmin(p.Email)
	if err != nil {
		return nil, err
	}
	if !utils.CheckPasswordMatch([]byte(admin.Password), p.Password) {
		fmt.Println("password error")
		return nil, errors.New("incorrect password")
	}

	token, err := utils.GenerateToken(admin.Email, p.Role)
	if err != nil {
		log.Print("error while generating jwt")
		return nil, errors.New("error while generating jwt")
	}
	adminResp := &adminpb.AdminResponse{
		Status:  "Success",
		Message: token,
	}

	return adminResp, nil
}

func (a *AdminService) ViewDashboard(p *adminpb.AdminView) (*adminpb.AdminDashboard, error) {
	ctx := context.Background()
	result, err := a.CodClient.ViewDashboard(ctx, &clientpb.View{})
	if err != nil {
		return nil, err
	}

	rslt, errrr := a.userClient.UsersCount(ctx, &pb.UserView{})
	if errrr != nil {
		fmt.Println(errrr)
	}
	admin, err := a.Repo.FetchAdmin(p.Status)
	if err != nil {
		return nil, errors.New("error while fetching admin")
	}
	return &adminpb.AdminDashboard{
		Wallet:           int64(admin.Wallet),
		Today:            result.Today,
		Monthly:          result.Monthly,
		CoordinatorCount: result.CoordinatorCount,
		UsersCount:       rslt.Usercount,
	}, nil
}

func NewAdminService(repos inter.RepoInterface, codClient clientpb.CoordinatorClient, userclient pb.UserClient) interr.ServiceInterface {
	return &AdminService{
		Repo:       repos,
		CodClient:  codClient,
		userClient: userclient,
	}
}

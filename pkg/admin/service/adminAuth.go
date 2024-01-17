package service

import (
	"errors"
	"fmt"
	"log"

	clientpb "github.com/Shakezidin/pkg/admin/client/pb"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
	inter "github.com/Shakezidin/pkg/admin/repository/interface"
	interr "github.com/Shakezidin/pkg/admin/service/interface"
	"github.com/Shakezidin/utils"
	jwt "github.com/Shakezidin/utils"
)

type AdminService struct {
	Repo      inter.RepoInterface
	codClient clientpb.CoordinatorClient
}

func (a *AdminService) LoginService(p *adminpb.AdminLogin) (*adminpb.AdminResponce, error) {
	admin, err := a.Repo.FetchAdmin(p.Email)
	if err != nil {
		return nil, err
	}
	if !utils.CheckPasswordMatch([]byte(admin.Password), p.Password) {
		fmt.Println("password error")
		return nil, errors.New("passwor incorrect")
	}

	token, err := jwt.GenerateToken(admin.Email, p.Role)
	if err != nil {
		log.Print("Generate jwt error")
		return nil, err
	}
	adminn := &adminpb.AdminResponce{
		Status:  "Success",
		Message: token,
	}

	return adminn, nil
}

func NewAdminService(repos inter.RepoInterface) interr.ServiceInterface {
	return &AdminService{
		Repo: repos,
	}
}

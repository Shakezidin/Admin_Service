package handler

import (
	"context"

	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminHandler) AdminLoginRequest(ctx context.Context, p *adminpb.AdminLogin) (*adminpb.AdminResponce, error) {
	admin, err := a.svc.LoginService(p)
	if err != nil {
		return &adminpb.AdminResponce{
			Status:  "False",
			Message: "error while admin login",
		}, err
	}
	return admin, nil
}

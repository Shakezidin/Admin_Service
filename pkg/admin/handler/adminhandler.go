package handler

import (
	"context"

	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminHandler) AdminLoginRequest(ctx context.Context, p *adminpb.AdminLogin) (*adminpb.AdminResponce, error) {
	admin, err := a.svc.LoginService(p)
	if err != nil {
		return &adminpb.AdminResponce{
			Status:  admin.Status,
			Message: admin.Message,
		}, err
	}
	return admin, nil
}

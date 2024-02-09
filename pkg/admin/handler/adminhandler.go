package handler

import (
	"context"

	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminHandler) AdminLoginRequest(ctx context.Context, p *adminpb.AdminLogin) (*adminpb.AdminResponse, error) {
	admin, err := a.svc.LoginService(p)
	if err != nil {
		return &adminpb.AdminResponse{
			Status:  admin.Status,
			Message: admin.Message,
		}, err
	}
	return admin, nil
}

func (a *AdminHandler) AdminViewDashBord(ctx context.Context, p *adminpb.AdminView) (*adminpb.AdminDashboard, error) {
	dashbord, err := a.svc.ViewDashboard(p)
	if err != nil {
		return nil, err
	}
	return dashbord, nil
}

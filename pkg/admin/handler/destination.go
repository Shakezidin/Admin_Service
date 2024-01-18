package handler

import (
	"context"

	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminHandler) AdminViewDestination(ctx context.Context, p *adminpb.AdminView) (*adminpb.AdminDestination, error) {
	destination, err := a.svc.ViewDestination(p)
	if err != nil {
		return &adminpb.AdminDestination{}, err
	}
	return destination, nil
}

func (a *AdminHandler) AdminViewActivity(ctx context.Context, p *adminpb.AdminView) (*adminpb.AdminActivity, error) {
	destination, err := a.svc.ViewActivity(p)
	if err != nil {
		return &adminpb.AdminActivity{}, err
	}
	return destination, nil
}

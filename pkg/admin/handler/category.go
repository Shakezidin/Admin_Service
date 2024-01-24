package handler

import (
	"context"

	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminHandler) AdminAddCategory(ctx context.Context, p *adminpb.AdminCategory) (*adminpb.AdminResponce, error) {
	admin, err := a.svc.AddCategory(p)
	if err != nil {
		return &adminpb.AdminResponce{
			Status:  admin.Status,
			Message: admin.Message,
		}, err
	}
	return admin, nil
}

func (a *AdminHandler) AdminViewCategories(ctx context.Context, p *adminpb.AdminView) (*adminpb.AdminCatagories, error) {
	catagories, err := a.svc.ViewCategories(p)
	if err != nil {
		return nil, err
	}
	return catagories, nil
}

package handler

import (
	"context"

	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminHandler) AdminAddCategory(ctx context.Context, p *adminpb.AdminCategory) (*adminpb.AdminResponse, error) {
	admin, err := a.svc.AddCategory(p)
	if err != nil {
	return nil,err
	}
	return admin, nil
}

func (a *AdminHandler) AdminViewCategories(ctx context.Context, p *adminpb.AdminView) (*adminpb.AdminCategories, error) {
	catagories, err := a.svc.ViewCategories(p)
	if err != nil {
		return nil, err
	}
	return catagories, nil
}

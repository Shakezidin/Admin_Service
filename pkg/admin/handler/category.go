package handler

import (
	"context"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminHandler) AddCategory(ctx context.Context,p *adminpb.AdminCategory)(*adminpb.AdminResponce,error){
	admin, err := a.svc.AddCategory(p)
	if err != nil {
		return &adminpb.AdminResponce{
			Status:  "False",
			Message: "error while admin login",
		}, err
	}
	return admin, nil
}
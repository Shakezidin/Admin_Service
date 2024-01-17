package service

import (
	"context"

	cpb "github.com/Shakezidin/pkg/admin/client/pb"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminService) AddCategory(p *adminpb.AdminCategory) (*adminpb.AdminResponce, error) {
	var ctx = context.Background()
	var catogory = cpb.Category{
		CategoryName: p.Category,
	}
	resp, err := a.codClient.AddCatagory(ctx, &catogory)
	if err != nil {
		return &adminpb.AdminResponce{
			Status:  "fail",
			Message: "error while adding catagory",
		}, err
	}

	return &adminpb.AdminResponce{
		Status:  "success",
		Message: resp.Message,
	}, nil
}

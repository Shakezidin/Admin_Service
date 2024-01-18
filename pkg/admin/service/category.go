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

func (a *AdminService) ViewCategories(p *adminpb.AdminView) (*adminpb.AdminCatagories, error) {
	var ctx = context.Background()
	resp, err := a.codClient.ViewCatagories(ctx, &cpb.View{})
	if err != nil {
		return &adminpb.AdminCatagories{}, err
	}
	var catagory adminpb.AdminCategory
	var catagories []*adminpb.AdminCategory
	for _, ctgry := range resp.Catagories {
		catagory.Categoryid = ctgry.CategoryId
		catagory.Category = ctgry.CategoryName
		catagories = append(catagories, &catagory)
	}
	return &adminpb.AdminCatagories{
		Catagory: catagories,
	}, nil
}

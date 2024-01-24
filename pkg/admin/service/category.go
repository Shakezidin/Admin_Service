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
			Status:  resp.Status,
			Message: resp.Message,
		}, err
	}

	return &adminpb.AdminResponce{
		Status:  resp.Status,
		Message: resp.Message,
	}, nil
}

func (a *AdminService) ViewCategories(p *adminpb.AdminView) (*adminpb.AdminCatagories, error) {
	var ctx = context.Background()
	resp, err := a.codClient.ViewCatagories(ctx, &cpb.View{
		Page: p.Page,
	})
	if err != nil {
		return &adminpb.AdminCatagories{}, err
	}
	var catagories []*adminpb.AdminCategory
	for _, ctgry := range resp.Catagories {
		var catagory adminpb.AdminCategory
		catagory.Categoryid = ctgry.CatagoryId
		catagory.Category = ctgry.CategoryName
		catagories = append(catagories, &catagory)
	}
	return &adminpb.AdminCatagories{
		Catagory: catagories,
	}, nil
}

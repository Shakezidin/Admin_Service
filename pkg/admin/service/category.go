package service

import (
	"context"

	cpb "github.com/Shakezidin/pkg/admin/client/pb"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminService) AddCategory(p *adminpb.AdminCategory) (*adminpb.AdminResponse, error) {
	var ctx = context.Background()
	category := &cpb.Category{
		CategoryName: p.Category,
	}

	resp, err := a.CodClient.AddCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	return &adminpb.AdminResponse{
		Status:  resp.Status,
		Message: resp.Message,
	}, nil
}

func (a *AdminService) ViewCategories(p *adminpb.AdminView) (*adminpb.AdminCategories, error) {
	var ctx = context.Background()
	resp, err := a.CodClient.ViewCatagories(ctx, &cpb.View{
		Page: p.Page,
	})
	if err != nil {
		return nil, err
	}

	var categories []*adminpb.AdminCategory
	for _, ctgry := range resp.Catagories {
		category := &adminpb.AdminCategory{
			Categoryid: ctgry.CatagoryId,
			Category:   ctgry.CategoryName,
		}
		categories = append(categories, category)
	}

	return &adminpb.AdminCategories{
		Catagory: categories,
	}, nil
}

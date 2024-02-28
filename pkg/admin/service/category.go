package service

import (
	"context"

	cpb "github.com/Shakezidin/pkg/admin/client/pb"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminService) AddCategory(p *adminpb.AdminCategory) (*adminpb.AdminResponse, error) {
	var ctx = context.Background()
	category := &cpb.Category{
		Category_Name: p.Category,
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
	for _, ctgry := range resp.Categories {
		category := &adminpb.AdminCategory{
			Category_ID: ctgry.Category_ID,
			Category:    ctgry.Category_Name,
		}
		categories = append(categories, category)
	}

	return &adminpb.AdminCategories{
		Categories: categories,
	}, nil
}

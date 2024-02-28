package service

import (
	"context"

	clientpb "github.com/Shakezidin/pkg/admin/client/pb"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminService) ViewCoordinators(p *adminpb.AdminView) (*adminpb.AdminUsers, error) {
	var ctx = context.Background()
	response, err := a.CodClient.ViewCoordinators(ctx, &clientpb.View{
		Status: p.Status,
		Page:   p.Page,
	})
	if err != nil {
		return nil, err
	}

	var coordinators []*adminpb.AdminUser
	for _, coordinator := range response.Users {
		coordinators = append(coordinators, &adminpb.AdminUser{
			ID:       coordinator.ID,
			Name:     coordinator.Name,
			Email:    coordinator.Email,
			Phone:    coordinator.Phone,
			Password: coordinator.Password,
			Role:     coordinator.Role,
		})
	}

	return &adminpb.AdminUsers{
		Users: coordinators,
	}, nil
}

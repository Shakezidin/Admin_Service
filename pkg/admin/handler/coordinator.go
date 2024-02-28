package handler

import (
	"context"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)
func (a *AdminHandler)AdminViewCoordinators(ctx context.Context,p *adminpb.AdminView)(*adminpb.AdminUsers,error){
	coordinators, err := a.svc.ViewCoordinators(p)
	if err != nil {
		return nil, err
	}
	return coordinators, nil
}
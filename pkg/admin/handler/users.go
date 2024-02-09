package handler

import (
	"context"

	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminHandler) AdminViewUsers(ctx context.Context, p *adminpb.AdminView) (*adminpb.AdminUsers, error) {
	rslt, err := a.svc.ViewUsersSVC(p)
	if err != nil {
		return nil, err
	}
	return rslt, nil
}

func (a *AdminHandler) AdminViewUser(ctx context.Context, p *adminpb.AdminView) (*adminpb.AdminUser, error) {
	rslt, err := a.svc.ViewUserSVC(p)
	if err != nil {
		return nil, err
	}
	return rslt, nil
}

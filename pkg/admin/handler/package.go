package handler

import (
	"context"

	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminHandler) AdminViewPackages(ctx context.Context, p *adminpb.AdminView) (*adminpb.AdminPackages, error) {
	rslt, err := a.svc.ViewPackagesSVC(p)
	if err != nil {
		return &adminpb.AdminPackages{}, err
	}
	return rslt, nil
}

func (a *AdminHandler) AdminViewpackage(ctx context.Context, p *adminpb.AdminView) (*adminpb.AdminPackage, error) {
	rslt, err := a.svc.ViewPackageSVC(p)
	if err != nil {
		return &adminpb.AdminPackage{}, err
	}
	return rslt, nil
}

func (a *AdminHandler)AdminPacakgeStatus(ctx context.Context, p *adminpb.AdminView) (*adminpb.AdminResponse, error) {
	rslt, err := a.svc.PackageStatusSVC(p)
	if err != nil {
		return &adminpb.AdminResponse{}, err
	}
	return rslt, nil
}
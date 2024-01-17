package interfaces

import (
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

type ServiceInterface interface {
	LoginService(p *adminpb.AdminLogin) (*adminpb.AdminResponce, error)
	AddCategory(p *adminpb.AdminCategory) (*adminpb.AdminResponce, error)
	ViewPackagesSVC() (*adminpb.AdminPackages, error)
	ViewPackageSVC(p *adminpb.AdminView) (*adminpb.AdminPackage, error)
	PackageStatusSVC(p *adminpb.AdminView) (*adminpb.AdminResponce, error)
}

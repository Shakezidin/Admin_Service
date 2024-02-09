package interfaces

import (
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

type ServiceInterface interface {
	LoginService(p *adminpb.AdminLogin) (*adminpb.AdminResponse, error)
	AddCategory(p *adminpb.AdminCategory) (*adminpb.AdminResponse, error)
	ViewPackagesSVC(p *adminpb.AdminView) (*adminpb.AdminPackages, error)
	ViewPackageSVC(p *adminpb.AdminView) (*adminpb.AdminPackage, error) 
	PackageStatusSVC(p *adminpb.AdminView) (*adminpb.AdminResponse, error)
	ViewCategories(p *adminpb.AdminView) (*adminpb.AdminCategories, error)
	ViewDestination(p *adminpb.AdminView) (*adminpb.AdminDestination, error)
	ViewActivity(p *adminpb.AdminView) (*adminpb.AdminActivity, error) 
	AddWalletSVC(p *adminpb.AdminAddWallet)(*adminpb.AdminResponse,error)
	ReduseWalletSVC(p *adminpb.AdminAddWallet) (*adminpb.AdminResponse, error)
	ViewCoordinators(p *adminpb.AdminView) (*adminpb.AdminUsers, error)
	ViewBookings(p *adminpb.AdminView) (*adminpb.AdminHistories, error)
	ViewBooking(p *adminpb.AdminView) (*adminpb.AdminHistory, error)
	ViewDashboard(p *adminpb.AdminView) (*adminpb.AdminDashboard, error)
	SearchBooking(p *adminpb.AdminBookingSearchCriteria)(*adminpb.AdminHistories,error)
	ViewUsersSVC(p *adminpb.AdminView)(*adminpb.AdminUsers,error)
	ViewUserSVC(p *adminpb.AdminView) (*adminpb.AdminUser, error) 
}

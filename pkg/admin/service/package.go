package service

import (
	"context"
	"errors"

	cpb "github.com/Shakezidin/pkg/admin/client/pb"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminService) ViewPackagesSVC(p *adminpb.AdminView) (*adminpb.AdminPackages, error) {
	var ctx = context.Background()

	result, err := a.codClient.AvailablePackages(ctx, &cpb.View{
		Status: p.Status,
	})
	if err != nil {
		return &adminpb.AdminPackages{
			Packages: nil,
		}, errors.New("error while fetching all packages ")
	}
	var pkgs []*adminpb.AdminPackage
	for _, pakg := range result.Packages {
		var pkg adminpb.AdminPackage
		pkg.PackageId = pakg.PackageId
		pkg.Destination = pakg.Destination
		pkg.DestinationCount = int64(pakg.DestinationCount)
		pkg.Enddate = pakg.Enddatetime
		pkg.Endlocation = pakg.Endlocation
		pkg.Image = pakg.Image
		pkg.Packagename = pakg.Packagename
		pkg.Price = int64(pakg.Price)
		pkg.Startdate = pakg.Startdatetime
		pkg.Startlocation = pakg.Startlocation
		pkg.Description = pakg.Description
		pkg.CoorinatorId = pakg.CoorinatorId
		pkgs = append(pkgs, &pkg)
	}

	return &adminpb.AdminPackages{
		Packages: pkgs,
	}, nil
}

func (a *AdminService) ViewPackageSVC(p *adminpb.AdminView) (*adminpb.AdminPackage, error) {
	var ctx = context.Background()

	result, err := a.codClient.CoordinatorViewPackage(ctx, &cpb.View{
		Id: p.Id,
	})
	if err != nil {
		return &adminpb.AdminPackage{}, errors.New("error while fetching all packages ")
	}

	var dstns = []*adminpb.AdminDestination{}
	
	var ctgry = adminpb.AdminCategory{
		Category: result.Category.CategoryName,
	}
	
	for _, dst := range result.Destinations {
		var dstn = adminpb.AdminDestination{}
		dstn.Description = dst.Description
		dstn.DestinationName = dst.DestinationName
		dstn.Image = dst.Image
		dstn.MaxCapacity = int64(dst.MaxCapacity)
		dstn.DestinationId = dst.DestinationId
		dstn.Minprice = int64(dst.Minprice)
		dstns = append(dstns, &dstn)
	}
	var pkg adminpb.AdminPackage
	pkg.Category = &ctgry
	pkg.PackageId = result.PackageId
	pkg.Destination = result.Destination
	pkg.DestinationCount = int64(result.DestinationCount)
	pkg.Enddate = result.Enddatetime
	pkg.Endlocation = result.Endlocation
	pkg.Image = result.Image
	pkg.Packagename = result.Packagename
	pkg.Price = int64(result.Price)
	pkg.Startdate = result.Startdatetime
	pkg.Startlocation = result.Startlocation
	pkg.Description = result.Description
	pkg.CoorinatorId = result.CoorinatorId
	pkg.MaxCapacity = result.MaxCapacity
	pkg.Destinations = dstns

	return &pkg, nil
}

func (a *AdminService) PackageStatusSVC(p *adminpb.AdminView) (*adminpb.AdminResponce, error) {
	var ctx = context.Background()

	result, err := a.codClient.AdminPacakgeStatus(ctx, &cpb.View{
		Id: p.Id,
	})
	if err != nil {
		return &adminpb.AdminResponce{
			Status:  result.Status,
			Message: result.Message,
		}, errors.New("error while fetching all packages ")
	}

	return &adminpb.AdminResponce{
		Status:  result.Status,
		Message: result.Message,
	}, nil
}

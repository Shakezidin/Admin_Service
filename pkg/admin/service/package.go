package service

import (
	"context"

	cpb "github.com/Shakezidin/pkg/admin/client/pb"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminService) ViewPackagesSVC(p *adminpb.AdminView) (*adminpb.AdminPackages, error) {
	var ctx = context.Background()

	result, err := a.codClient.AvailablePackages(ctx, &cpb.View{
		Status: p.Status,
		Page:   p.Page,
	})
	if err != nil {
		return &adminpb.AdminPackages{
			Packages: nil,
		}, err
	}
	var pkgs []*adminpb.AdminPackage
	for _, pakg := range result.Packages {
		var pkg adminpb.AdminPackage
		pkg.PackageId = pakg.PackageId
		pkg.Destination = pakg.Destination
		pkg.DestinationCount = int64(pakg.DestinationCount)
		pkg.Enddate = pakg.Enddate
		pkg.Image = pakg.Image
		pkg.Packagename = pakg.Packagename
		pkg.AvailableSpace = pakg.AvailableSpace
		pkg.Price = int64(pakg.Price)
		pkg.Startdate = pakg.Startdate
		pkg.Starttime = pakg.Starttime
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
		return &adminpb.AdminPackage{}, err
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
		dstn.DestinationId = dst.DestinationId
		dstns = append(dstns, &dstn)
	}
	var pkg adminpb.AdminPackage
	pkg.Category = &ctgry
	pkg.PackageId = result.PackageId
	pkg.Destination = result.Destination
	pkg.DestinationCount = int64(result.DestinationCount)
	pkg.Enddate = result.Enddate
	pkg.Image = result.Image
	pkg.Packagename = result.Packagename
	pkg.AvailableSpace = result.AvailableSpace
	pkg.Price = int64(result.Price)
	pkg.Startdate = result.Startdate
	pkg.Starttime = result.Starttime
	pkg.Startlocation = result.Startlocation
	pkg.Description = result.Description
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
		}, err
	}

	return &adminpb.AdminResponce{
		Status:  result.Status,
		Message: result.Message,
	}, nil
}

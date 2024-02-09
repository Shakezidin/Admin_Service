package service

import (
	"context"

	cpb "github.com/Shakezidin/pkg/admin/client/pb"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminService) ViewPackagesSVC(p *adminpb.AdminView) (*adminpb.AdminPackages, error) {
	var ctx = context.Background()
	result, err := a.CodClient.AvailablePackages(ctx, &cpb.View{
		Status: p.Status,
		Page:   p.Page,
	})
	if err != nil {
		return nil, err
	}

	var packages []*adminpb.AdminPackage
	for _, pack := range result.Packages {
		pkg := &adminpb.AdminPackage{
			PackageId:        pack.PackageId,
			Destination:      pack.Destination,
			DestinationCount: int64(pack.DestinationCount),
			Enddate:          pack.Enddate,
			Image:            pack.Image,
			Packagename:      pack.Packagename,
			AvailableSpace:   pack.AvailableSpace,
			Price:            int64(pack.Price),
			Startdate:        pack.Startdate,
			Starttime:        pack.Starttime,
			Startlocation:    pack.Startlocation,
			Description:      pack.Description,
			CoorinatorId:     pack.CoorinatorId,
		}
		packages = append(packages, pkg)
	}

	return &adminpb.AdminPackages{Packages: packages}, nil
}

func (a *AdminService) ViewPackageSVC(p *adminpb.AdminView) (*adminpb.AdminPackage, error) {
	var ctx = context.Background()
	result, err := a.CodClient.CoordinatorViewPackage(ctx, &cpb.View{
		Id: p.Id,
	})
	if err != nil {
		return nil, err
	}

	var destinations []*adminpb.AdminDestination
	for _, dst := range result.Destinations {
		dest := &adminpb.AdminDestination{
			Description:     dst.Description,
			DestinationName: dst.DestinationName,
			Image:           dst.Image,
			DestinationId:   dst.DestinationId,
		}
		destinations = append(destinations, dest)
	}

	category := &adminpb.AdminCategory{
		Category: result.Category.CategoryName,
	}

	return &adminpb.AdminPackage{
		Category:         category,
		PackageId:        result.PackageId,
		Destination:      result.Destination,
		DestinationCount: int64(result.DestinationCount),
		Enddate:          result.Enddate,
		Image:            result.Image,
		Packagename:      result.Packagename,
		AvailableSpace:   result.AvailableSpace,
		Price:            int64(result.Price),
		Startdate:        result.Startdate,
		Starttime:        result.Starttime,
		Startlocation:    result.Startlocation,
		Description:      result.Description,
		MaxCapacity:      result.MaxCapacity,
		Destinations:     destinations,
	}, nil
}

func (a *AdminService) PackageStatusSVC(p *adminpb.AdminView) (*adminpb.AdminResponse, error) {
	var ctx = context.Background()
	result, err := a.CodClient.AdminPacakgeStatus(ctx, &cpb.View{
		Id: p.Id,
	})
	if err != nil {
		return nil, err
	}

	return &adminpb.AdminResponse{
		Status:  result.Status,
		Message: result.Message,
	}, nil
}

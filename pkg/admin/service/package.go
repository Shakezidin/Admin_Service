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
			Package_ID:        pack.Package_ID,
			Destination:       pack.Destination,
			Destination_Count: int64(pack.Destination_Count),
			End_Date:          pack.End_Date,
			Image:             pack.Image,
			Packagename:       pack.Packagename,
			Available_Space:   pack.Available_Space,
			Price:             int64(pack.Price),
			Start_Date:        pack.Start_Date,
			Start_Time:        pack.Start_Time,
			Start_Location:    pack.Start_Location,
			Description:       pack.Description,
			Coordinator_ID:    pack.Coordinator_ID,
		}
		packages = append(packages, pkg)
	}

	return &adminpb.AdminPackages{Packages: packages}, nil
}

func (a *AdminService) ViewPackageSVC(p *adminpb.AdminView) (*adminpb.AdminPackage, error) {
	var ctx = context.Background()
	result, err := a.CodClient.CoordinatorViewPackage(ctx, &cpb.View{
		ID: p.ID,
	})
	if err != nil {
		return nil, err
	}

	var destinations []*adminpb.AdminDestination
	for _, dst := range result.Destinations {
		dest := &adminpb.AdminDestination{
			Description:      dst.Description,
			Destination_Name: dst.Destination_Name,
			Image:            dst.Image,
			Destination_ID:   dst.Destination_ID,
		}
		destinations = append(destinations, dest)
	}

	category := &adminpb.AdminCategory{
		Category: result.Category.Category_Name,
	}

	return &adminpb.AdminPackage{
		Category:          category,
		Package_ID:        result.Package_ID,
		Destination:       result.Destination,
		Destination_Count: int64(result.Destination_Count),
		End_Date:          result.End_Date,
		Image:             result.Image,
		Packagename:       result.Packagename,
		Available_Space:   result.Available_Space,
		Price:             int64(result.Price),
		Start_Date:        result.Start_Date,
		Start_Time:        result.Start_Time,
		Start_Location:    result.Start_Location,
		Description:       result.Description,
		Max_Capacity:      result.Max_Capacity,
		Destinations:      destinations,
	}, nil
}

func (a *AdminService) PackageStatusSVC(p *adminpb.AdminView) (*adminpb.AdminResponse, error) {
	var ctx = context.Background()
	result, err := a.CodClient.AdminPacakgeStatus(ctx, &cpb.View{
		ID: p.ID,
	})
	if err != nil {
		return nil, err
	}

	return &adminpb.AdminResponse{
		Status:  result.Status,
		Message: result.Message,
	}, nil
}

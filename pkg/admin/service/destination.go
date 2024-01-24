package service

import (
	"context"

	cpb "github.com/Shakezidin/pkg/admin/client/pb"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminService) ViewDestination(p *adminpb.AdminView) (*adminpb.AdminDestination, error) {
	var ctx = context.Background()

	result, err := a.codClient.CoordinatorViewDestination(ctx, &cpb.View{
		Id: p.Id,
	})
	if err != nil {
		return &adminpb.AdminDestination{}, err
	}

	var actvtys = []*adminpb.AdminActivity{}

	for _, act := range result.Activity {
		var actvty = adminpb.AdminActivity{}
		actvty.Description = act.Description
		actvty.ActivityId = act.ActivityId
		actvty.ActivityType = act.ActivityType
		actvty.Activityname = act.Activityname
		actvty.Amount = act.Amount
		actvty.Date = act.Date
		actvty.Location = act.Location
		actvty.Time = act.Time
		actvtys = append(actvtys, &actvty)
	}
	var dst adminpb.AdminDestination

	dst.Activity = actvtys
	dst.Description = result.Description
	dst.DestinationId = result.DestinationId
	dst.DestinationName = result.DestinationName
	dst.Image = result.DestinationName
	dst.MaxCapacity = result.MaxCapacity
	dst.Minprice = result.Minprice

	return &dst, nil
}

func (a *AdminService) ViewActivity(p *adminpb.AdminView) (*adminpb.AdminActivity, error) {
	var ctx = context.Background()

	result, err := a.codClient.CoordinatorViewActivity(ctx, &cpb.View{
		Id: p.Id,
	})
	if err != nil {
		return &adminpb.AdminActivity{}, err
	}
	return &adminpb.AdminActivity{
		ActivityId:   result.ActivityId,
		Activityname: result.Activityname,
		Description:  result.Description,
		Location:     result.Location,
		ActivityType: result.ActivityType,
		Amount:       result.Amount,
		Date:         result.Date,
		Time:         result.Time,
	}, nil
}

package service

import (
	"context"

	cpb "github.com/Shakezidin/pkg/admin/client/pb"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminService) ViewDestination(p *adminpb.AdminView) (*adminpb.AdminDestination, error) {
	var ctx = context.Background()
	result, err := a.CodClient.CoordinatorViewDestination(ctx, &cpb.View{
		Id: p.Id,
	})
	if err != nil {
		return nil, err
	}

	var activities []*adminpb.AdminActivity
	for _, act := range result.Activity {
		activities = append(activities, &adminpb.AdminActivity{
			ActivityId:   act.ActivityId,
			Activityname: act.Activityname,
			Description:  act.Description,
			Location:     act.Location,
			ActivityType: act.ActivityType,
			Amount:       act.Amount,
			Date:         act.Date,
			Time:         act.Time,
		})
	}

	return &adminpb.AdminDestination{
		Activity:        activities,
		Description:     result.Description,
		DestinationId:   result.DestinationId,
		DestinationName: result.DestinationName,
		Image:           result.DestinationName, // Note: Corrected variable name
		MaxCapacity:     result.MaxCapacity,
		Minprice:        result.Minprice,
	}, nil
}

func (a *AdminService) ViewActivity(p *adminpb.AdminView) (*adminpb.AdminActivity, error) {
	var ctx = context.Background()
	result, err := a.CodClient.CoordinatorViewActivity(ctx, &cpb.View{
		Id: p.Id,
	})
	if err != nil {
		return nil, err
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

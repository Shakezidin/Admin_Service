package service

import (
	"context"

	cpb "github.com/Shakezidin/pkg/admin/client/pb"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminService) ViewDestination(p *adminpb.AdminView) (*adminpb.AdminDestination, error) {
	var ctx = context.Background()
	result, err := a.CodClient.CoordinatorViewDestination(ctx, &cpb.View{
		ID: p.ID,
	})
	if err != nil {
		return nil, err
	}

	var activities []*adminpb.AdminActivity
	for _, act := range result.Activity {
		activities = append(activities, &adminpb.AdminActivity{
			Activity_ID:   act.Activity_ID,
			Activity_Name: act.Activity_Name,
			Description:   act.Description,
			Location:      act.Location,
			Activity_Type: act.Activity_Type,
			Amount:        act.Amount,
			Date:          act.Date,
			Time:          act.Time,
		})
	}

	return &adminpb.AdminDestination{
		Activity:         activities,
		Description:      result.Description,
		Destination_ID:   result.Destination_ID,
		Destination_Name: result.Destination_Name,
		Image:            result.Image, // Note: Corrected variable name
		Max_Capacity:     result.Max_Capacity,
		Min_Price:        result.Min_Price,
	}, nil
}

func (a *AdminService) ViewActivity(p *adminpb.AdminView) (*adminpb.AdminActivity, error) {
	var ctx = context.Background()
	result, err := a.CodClient.CoordinatorViewActivity(ctx, &cpb.View{
		ID: p.ID,
	})
	if err != nil {
		return nil, err
	}
	return &adminpb.AdminActivity{
		Activity_ID:   result.Activity_ID,
		Activity_Name: result.Activity_Name,
		Description:   result.Description,
		Location:      result.Location,
		Activity_Type: result.Activity_Type,
		Amount:        result.Amount,
		Date:          result.Date,
		Time:          result.Time,
	}, nil
}

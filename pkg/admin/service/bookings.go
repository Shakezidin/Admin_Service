package service

import (
	"context"

	clientpb "github.com/Shakezidin/pkg/admin/client/pb"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminService) ViewBookings(p *adminpb.AdminView) (*adminpb.AdminHistories, error) {
	var ctx = context.Background()
	result, err := a.CodClient.ViewHistory(ctx, &clientpb.View{
		Id:     p.Id,
		Status: p.Status,
	})
	if err != nil {
		return nil, err
	}

	var history []*adminpb.AdminHistory
	if result != nil && result.History != nil {
		for _, hstry := range result.History {
			history = append(history, &adminpb.AdminHistory{
				Id:              hstry.Id,
				PaymentMode:     hstry.PaymentMode,
				BookingStatus:   hstry.BookingStatus,
				CancelledStatus: hstry.CancelledStatus,
				TotalPrice:      hstry.TotalPrice,
				UserId:          hstry.UserId,
				BookingId:       hstry.BookingId,
				BookDate:        hstry.BookDate,
				StartDate:       hstry.StartDate,
				PaidAmount:      hstry.PaidAmount,
			})
		}
	}

	return &adminpb.AdminHistories{
		History: history,
	}, nil
}

func (a *AdminService) ViewBooking(p *adminpb.AdminView) (*adminpb.AdminHistory, error) {
	var ctx = context.Background()
	booking, err := a.CodClient.ViewBooking(ctx, &clientpb.View{
		Id: p.Id,
	})
	if err != nil {
		return nil, err
	}

	var traveller []*adminpb.AdminTravellerDetails
	if booking != nil && booking.Travellers != nil {
		for _, trvler := range booking.Travellers {
			traveller = append(traveller, &adminpb.AdminTravellerDetails{
				Id:     int64(trvler.Id),
				Name:   trvler.Name,
				Age:    trvler.Age,
				Gender: trvler.Gender,
			})
		}
	}

	return &adminpb.AdminHistory{
		Id:              int64(booking.Id),
		PaymentMode:     booking.PaymentMode,
		BookingStatus:   booking.BookingStatus,
		CancelledStatus: booking.CancelledStatus,
		TotalPrice:      int64(booking.TotalPrice),
		UserId:          int64(booking.UserId),
		BookingId:       booking.BookingId,
		BookDate:        booking.BookDate,
		StartDate:       booking.StartDate,
		Travellers:      traveller,
		PaidAmount:      int64(booking.PaidAmount),
	}, nil
}

func (c *AdminService) SearchBooking(p *adminpb.AdminBookingSearchCriteria) (*adminpb.AdminHistories, error) {
	var ctx = context.Background()
	booking, err := c.CodClient.SearchBooking(ctx, &clientpb.BookingSearchCriteria{
		PaymentMode:     p.PaymentMode,
		BookingStatus:   p.BookingStatus,
		CancelledStatus: p.CancelledStatus,
		UserEmail:       p.UserEmail,
		BookingId:       p.BookingId,
		BookDate:        p.BookDate,
		StartDate:       p.StartDate,
		CoordinatorId:   p.CoordinatorId,
		Page:            p.Page,
	})
	if err != nil {
		return nil, err
	}

	var history []*adminpb.AdminHistory
	if booking != nil && booking.History != nil {
		for _, hstry := range booking.History {
			history = append(history, &adminpb.AdminHistory{
				Id:              hstry.Id,
				PaymentMode:     hstry.PaymentMode,
				BookingStatus:   hstry.BookingStatus,
				CancelledStatus: hstry.CancelledStatus,
				TotalPrice:      hstry.TotalPrice,
				UserId:          hstry.UserId,
				BookingId:       hstry.BookingId,
				BookDate:        hstry.BookDate,
				StartDate:       hstry.StartDate,
				PaidAmount:      hstry.PaidAmount,
			})
		}
	}

	return &adminpb.AdminHistories{
		History: history,
	}, nil

}

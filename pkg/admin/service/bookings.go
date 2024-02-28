package service

import (
	"context"

	clientpb "github.com/Shakezidin/pkg/admin/client/pb"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminService) ViewBookings(p *adminpb.AdminView) (*adminpb.AdminHistories, error) {
	var ctx = context.Background()
	result, err := a.CodClient.ViewHistory(ctx, &clientpb.View{
		ID:     p.ID,
		Status: p.Status,
	})
	if err != nil {
		return nil, err
	}

	var history []*adminpb.AdminHistory
	if result != nil && result.History != nil {
		for _, hstry := range result.History {
			history = append(history, &adminpb.AdminHistory{
				ID:               hstry.ID,
				Payment_Mode:     hstry.Payment_Mode,
				Booking_Status:   hstry.Booking_Status,
				Cancelled_Status: hstry.Cancelled_Status,
				Total_Price:      hstry.Total_Price,
				User_ID:          hstry.User_ID,
				Booking_ID:       hstry.Booking_ID,
				Book_Date:        hstry.Book_Date,
				Start_Date:       hstry.Start_Date,
				Paid_Amount:      hstry.Paid_Amount,
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
		ID: p.ID,
	})
	if err != nil {
		return nil, err
	}

	var traveller []*adminpb.AdminTravellerDetails
	if booking != nil && booking.Travellers != nil {
		for _, trvler := range booking.Travellers {
			traveller = append(traveller, &adminpb.AdminTravellerDetails{
				ID:     int64(trvler.ID),
				Name:   trvler.Name,
				Age:    trvler.Age,
				Gender: trvler.Gender,
			})
		}
	}

	return &adminpb.AdminHistory{
		ID:               int64(booking.ID),
		Payment_Mode:     booking.Payment_Mode,
		Booking_Status:   booking.Booking_Status,
		Cancelled_Status: booking.Cancelled_Status,
		Total_Price:      int64(booking.Total_Price),
		User_ID:          int64(booking.User_ID),
		Booking_ID:       booking.Booking_ID,
		Book_Date:        booking.Book_Date,
		Start_Date:       booking.Start_Date,
		Travellers:       traveller,
		Paid_Amount:      int64(booking.Paid_Amount),
	}, nil
}

func (c *AdminService) SearchBooking(p *adminpb.AdminBookingSearchCriteria) (*adminpb.AdminHistories, error) {
	var ctx = context.Background()
	booking, err := c.CodClient.SearchBooking(ctx, &clientpb.BookingSearchCriteria{
		Payment_Mode:     p.Payment_Mode,
		Booking_Status:   p.Booking_Status,
		Cancelled_Status: p.Cancelled_Status,
		User_Email:       p.User_Email,
		Booking_ID:       p.Booking_ID,
		Book_Date:        p.Book_Date,
		Start_Date:       p.Start_Date,
		Coordinator_ID:   p.Coordinator_ID,
		Page:             p.Page,
	})
	if err != nil {
		return nil, err
	}

	var history []*adminpb.AdminHistory
	if booking != nil && booking.History != nil {
		for _, hstry := range booking.History {
			history = append(history, &adminpb.AdminHistory{
				ID:               hstry.ID,
				Payment_Mode:     hstry.Payment_Mode,
				Booking_Status:   hstry.Booking_Status,
				Cancelled_Status: hstry.Cancelled_Status,
				Total_Price:      hstry.Total_Price,
				User_ID:          hstry.User_ID,
				Booking_ID:       hstry.Booking_ID,
				Book_Date:        hstry.Book_Date,
				Start_Date:       hstry.Start_Date,
				Paid_Amount:      hstry.Paid_Amount,
			})
		}
	}

	return &adminpb.AdminHistories{
		History: history,
	}, nil

}

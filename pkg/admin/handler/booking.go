package handler

import (
	"context"

	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminHandler) AdminViewBookings(ctx context.Context, p *adminpb.AdminView) (*adminpb.AdminHistories, error) {
	bookings, err := a.svc.ViewBookings(p)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (c *AdminHandler) AdminViewBooking(ctx context.Context, p *adminpb.AdminView) (*adminpb.AdminHistory, error) {
	respnc, err := c.svc.ViewBooking(p)
	if err != nil {
		return nil, err
	}
	return respnc, nil
}

func (c *AdminHandler) AdminSearchBooking(ctx context.Context, p *adminpb.AdminBookingSearchCriteria) (*adminpb.AdminHistories, error) {
	respnc, err := c.svc.SearchBooking(p)
	if err != nil {
		return nil, err
	}
	return respnc, nil
}

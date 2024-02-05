package service

import (
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminService) AddWalletSVC(p *adminpb.AdminAddWallet) (*adminpb.AdminResponce, error) {
	err := a.Repo.UpdateaWallet(float64(p.Amount))
	if err != nil {
		return &adminpb.AdminResponce{}, err
	}
	return &adminpb.AdminResponce{
		Status: "success",
	}, nil
}

func (a *AdminService) ReduseWalletSVC(p *adminpb.AdminAddWallet) (*adminpb.AdminResponce, error) {
	amount := p.Amount
	err := a.Repo.UpdateaWallet(float64(-amount))
	if err != nil {
		return &adminpb.AdminResponce{}, err
	}
	return &adminpb.AdminResponce{
		Status: "success",
	}, nil
}

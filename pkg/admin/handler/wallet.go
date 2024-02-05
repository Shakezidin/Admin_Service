package handler

import (
	"context"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func (a *AdminHandler)AdminAddWalletRequest(ctx context.Context,p *adminpb.AdminAddWallet)(*adminpb.AdminResponce,error){
	rslt, err := a.svc.AddWalletSVC(p)
	if err != nil {
		return &adminpb.AdminResponce{}, err
	}
	return rslt, nil
}

func (a *AdminHandler)AdminReduseWalletRequesr(ctx context.Context,p *adminpb.AdminAddWallet)(*adminpb.AdminResponce,error){
	rslt, err := a.svc.ReduseWalletSVC(p)
	if err != nil {
		return &adminpb.AdminResponce{}, err
	}
	return rslt, nil
}
package handler

import (
	adminpb "github.com/Shakezidin/pkg/admin/pb"
	inter "github.com/Shakezidin/pkg/admin/service/interface"
)

type AdminHandler struct {
	svc inter.ServiceInterface
	adminpb.AdminServer
}

func NewAdminHandler(svc inter.ServiceInterface) *AdminHandler {
	return &AdminHandler{
		svc: svc,
	}
}

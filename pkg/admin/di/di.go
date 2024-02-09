package di

import (
	"fmt"
	"log"

	"github.com/Shakezidin/config"
	"github.com/Shakezidin/pkg/admin/client"
	"github.com/Shakezidin/pkg/admin/handler"
	"github.com/Shakezidin/pkg/admin/repository"
	"github.com/Shakezidin/pkg/admin/server"
	"github.com/Shakezidin/pkg/admin/service"
	"github.com/Shakezidin/pkg/admin/userclient"
	"github.com/Shakezidin/pkg/db"
)

func Init() {
	config := config.LoadConfig()
	db := db.Database(config)
	useClient, err := userclient.ClientDial(*config)
	codClient, err := client.ClientDial(*config)
	if err != nil {
		fmt.Println("error while dialing coordinator client")
		return
	}
	adminrepo := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(adminrepo, codClient, useClient)
	adminHandler := handler.NewAdminHandler(adminService)
	err = server.NewAdminGrpcServer(config, adminHandler)
	if err != nil {
		log.Fatalf("something went wrong", err)
	}
}

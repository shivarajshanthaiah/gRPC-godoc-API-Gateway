package main

import (
	"github.com/shivaraj-shanthaiah/godoc-API/config"
	"github.com/shivaraj-shanthaiah/godoc-API/pkg/admin"
	"github.com/shivaraj-shanthaiah/godoc-API/pkg/server"
	"github.com/shivaraj-shanthaiah/godoc-API/pkg/user"
)

func main() {
	server := server.Server()
	config.LoadConfig()
	user.NewUserRoute(server.R)
	admin.NewAdminRoute(server.R)
	server.StartServer()
}

package main

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"vitshop.vimfn.in/cmd/api"
	"vitshop.vimfn.in/configs"
	"vitshop.vimfn.in/db"
)

func main() {
	cfg := mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPass,
		Addr:                 configs.Envs.DBAddr,
		DBName:               configs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	dbConn, err := db.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	db.InitDB(dbConn)

	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), dbConn)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

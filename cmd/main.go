package main

import (
	. "Go-market-test/config"
	"Go-market-test/pkg/comment"
	"Go-market-test/pkg/database"
	"Go-market-test/pkg/deal"
	"Go-market-test/pkg/imagechild"
	route "Go-market-test/pkg/routes"
	"Go-market-test/pkg/user"
	"log"
)

func main() {
	var err error
	var cfg Config

	cfg, err = SetConfig()
	if err != nil {
		return
	}

	err = database.InitDatabase()
	if err != nil {
		log.Fatalln("could not create database", err)
	}

	err = database.GlobalDB.AutoMigrate(&user.User{}, &deal.Deal{}, &comment.Comment{}, &imagechild.ImageChild{})
	if err != nil {
		return
	}

	r := route.SetupRouter()
	err = r.Run(cfg.ServerHost)
	if err != nil {
		return
	}
}

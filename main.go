package main

import (
	"log"
	"webinar/api"
	"webinar/repo"
	"webinar/service"
)

func main() {
	rep, err := repo.NewGetter(&repo.Options{Environment: "debug"})
	if err != nil {
		//  в мейн паниковать можно
		log.Fatalln(err)
	}

	serv := service.New(rep)

	myapy := api.New(serv)

	log.Fatalln(myapy.Run())
}

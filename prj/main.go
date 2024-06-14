package main

import (
	"filmlib/auth"
	"filmlib/endpoints"
	"filmlib/repo"
	"log"
)

func main() {
	var dbURI string = "mongodb://azfastapiserver:jaWquzgp9MRClHHgCG58wAOs5vlz8azjsu7tpbvc48l6Gj3D4YKdnQGdcSD6ECkCOCkjFkkkdGgWACDbTUTJHg%3D%3D@azfastapiserver.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@azfastapiserver@"

	mongoRepo, err := repo.NewMongoRepo(dbURI)
	if err != nil {
		log.Fatal(err)
	}

	svc := auth.NewService(mongoRepo)
	authHandler := endpoints.NewAuthHandler(svc)

	e := authHandler.InitRoutes()

	if err := e.Start(":5003"); err != nil {
		mongoRepo.MongoShutdown()
		log.Fatal(err)
	}
}

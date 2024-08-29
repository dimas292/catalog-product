package main

import (
	"catalog-product/external/database"
	"log"
)

func main(){

	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPass := "root"
	dbName := "postgres"

	db, err := database.ConnectPostgsres(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("connected db")
	}
	

}

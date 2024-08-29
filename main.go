package main

import (
	"catalog-product/apps"
	"catalog-product/external/database"
	"log"
)

func main(){

	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPass := "root"
	dbName := "belajar_go"

	db, err := database.ConnectPostgsres(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("connected db")
	}

	APP_PORT := ":9000"

	apps.Run(APP_PORT, db)
	

}

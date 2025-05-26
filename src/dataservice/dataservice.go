package dataservice

import (
	"golang-backend-microservice/container/log"
	Gin "golang-backend-microservice/dataservice/gin"
	MySql "golang-backend-microservice/dataservice/mysql"
	Nats "golang-backend-microservice/dataservice/nats"
	"golang-backend-microservice/usecase"
	"os"
	"runtime"
	"time"
)

const RETRY_TIMER = 10

func EstablishDataService() {
	for {
		nc, svc := Nats.OpenNatsServerConnection()
		if nc != nil && svc != nil {
			log.Info("NATS server connection established!")
			for {
				mysql := MySql.OpenMySqlConnection()
				if mysql != nil {
					log.Info("MySQL connection established!")

					// Add services to NATS connection
					database := svc.AddGroup("database")
					{
						// Add MySQL endpoints
						MySqlEndpointGroup := database.AddGroup("mysql")
						if err := MySqlEndpointGroup.AddEndpoint("select", MySql.SelectRecord(mysql)); err != nil {
							log.Error("Error adding NATS service - %s", err.Error())
							return
						}
						if err := MySqlEndpointGroup.AddEndpoint("insert", MySql.InsertRecord(mysql)); err != nil {
							log.Error("Error adding NATS service - %s", err.Error())
							return
						}
						if err := MySqlEndpointGroup.AddEndpoint("update", MySql.UpdateRecord(mysql)); err != nil {
							log.Error("Error adding NATS service - %s", err.Error())
							return
						}
						if err := MySqlEndpointGroup.AddEndpoint("delete", MySql.DeleteRecord(mysql)); err != nil {
							log.Error("Error adding NATS service - %s", err.Error())
							return
						}
					}
					break
				}

				// Attempt reconnecting to MySQL if not successful
				log.Debug("Retry in %d seconds...", RETRY_TIMER)
				time.Sleep(RETRY_TIMER * time.Second)
			}

			// Set up routes using Gin
			r := Gin.SetupRoutes(nc)
			usecase.AddBookRoutes(nc, r)
			r.Run(":" + os.Getenv("PORT"))
			break
		}

		// Attempt reconnecting to NATS server if not successful
		log.Debug("Retry in %d seconds...", RETRY_TIMER)
		time.Sleep(RETRY_TIMER * time.Second)
	}

	log.Info("Service established successfully!")
	runtime.Goexit()
}

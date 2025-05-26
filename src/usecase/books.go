package usecase

import (
	"encoding/json"
	"fmt"
	"golang-backend-microservice/container/log"
	Nats "golang-backend-microservice/dataservice/nats"
	"golang-backend-microservice/model"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

func AddBookRoutes(nc *nats.Conn, r *gin.Engine) {
	r.GET("/books", func(c *gin.Context) {
		params := model.MySqlReqArgs{
			Table: "Books",
		}

		// Add "author" parameter
		author := c.Query("author")
		if author != "" {
			if params.Where == nil {
				params.Where = make(map[string]any)
			}
			params.Where["author"] = author
		}

		// Get mysql response from Nats
		res, err := Nats.Request(nc, "database.mysql.select", params)
		if err != nil {
			log.Error(err.Error())
		} else {
			var r Nats.DataResponse[model.Book]
			if err := json.Unmarshal(res.Data, &r); err != nil {
				fmt.Println("Error unmarshalling:", err)
				return
			}
			c.JSON(r.Status, r)
		}
	})
}

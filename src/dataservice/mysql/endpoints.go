package mysql

import (
	"encoding/json"
	"fmt"
	"golang-backend-microservice/container/log"
	"golang-backend-microservice/dataservice/nats"
	"golang-backend-microservice/model"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go/micro"
)

func SelectRecord(db *sqlx.DB) micro.HandlerFunc {
	return func(req micro.Request) {
		var data model.MySqlReqArgs
		json.Unmarshal(req.Data(), &data)
		query, args, err := BuildSelectQuery(&data)
		if err != nil {
			log.Error(err.Error())
			nats.StatusResponse{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}.Respond(req)
			return
		}

		switch data.Table {
		case "Books":
			var records []model.Book
			results, err := db.Queryx(query, args...)
			if err != nil {
				log.Error(err.Error())
				nats.StatusResponse{
					Status: http.StatusInternalServerError,
					Error:  err.Error(),
				}.Respond(req)
				return
			}
			for results.Next() {
				var record model.Book
				if err := results.StructScan(&record); err != nil {
					log.Error(err.Error())
					continue
				}
				records = append(records, record)
			}
			results.Rows.Close()
			log.Info("Output: %v", records)
			nats.DataResponse[model.Book]{
				Status: http.StatusOK,
				Data:   records,
			}.Respond(req)
		default:
			log.Error("Error: unknwon table %s", data.Table)
			nats.StatusResponse{
				Status: http.StatusNotFound,
				Error:  fmt.Sprintf("Error: unknown table %s", data.Table),
			}.Respond(req)
		}
	}
}

func InsertRecord(db *sqlx.DB) micro.HandlerFunc {
	return func(req micro.Request) {
		var data model.MySqlReqArgs
		json.Unmarshal(req.Data(), &data)
		query, args, err := BuildInsertQuery(&data)
		if err != nil {
			log.Error(err.Error())
			// TODO: send Nats response
			return
		}
		results, err := db.Queryx(query, args...)
		if err != nil {
			log.Error(err.Error())
			// TODO: send Nats response
			return
		}
		results.Rows.Close()
	}
}

func UpdateRecord(db *sqlx.DB) micro.HandlerFunc {
	return func(req micro.Request) {
		var data model.MySqlReqArgs
		json.Unmarshal(req.Data(), &data)
		query, args, err := BuildUpdateQuery(&data)
		if err != nil {
			log.Error(err.Error())
			// TODO: send Nats response
			return
		}
		results, err := db.Queryx(query, args...)
		if err != nil {
			log.Error(err.Error())
			// TODO: send Nats response
			return
		}
		results.Rows.Close()
	}
}

func DeleteRecord(db *sqlx.DB) micro.HandlerFunc {
	return func(req micro.Request) {
		var data model.MySqlReqArgs
		json.Unmarshal(req.Data(), &data)
		query, args, err := BuildDeleteQuery(&data)
		if err != nil {
			log.Error(err.Error())
			// TODO: send Nats response
			return
		}
		results, err := db.Queryx(query, args...)
		if err != nil {
			log.Error(err.Error())
			// TODO: send Nats response
			return
		}
		results.Rows.Close()
	}
}

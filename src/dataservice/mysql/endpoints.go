package mysql

import (
	"encoding/json"
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
		results, err := db.Queryx(query, args...)
		if err != nil {
			log.Error(err.Error())
			nats.StatusResponse{
				Status: http.StatusInternalServerError,
				Error:  err.Error(),
			}.Respond(req)
			return
		}

		switch data.Table {
		case "book":
			var records []model.Book
			for results.Next() {
				var record model.Book
				if err := results.StructScan(&record); err != nil {
					log.Error(err.Error())
					continue
				}
				records = append(records, record)
			}
			log.Info("Output: %v", records)
			// TODO: send Nats response
		default:
			log.Error("Error: unknwon table %s", data.Table)
			nats.StatusResponse{
				Status: http.StatusNotFound,
				Error:  "Error: unknwon table %s",
			}.Respond(req)
		}
		results.Rows.Close()
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

package db

import (
	"encoding/json"
	"fmt"

	"github.com/gocql/gocql"
	"go.uber.org/zap"
)

type mutant_data struct {
	first_name       string `json:"first_name"`
	last_name        string `json:"last_name"`
	address          string `json:"address"`
	picture_location string `json:"picture_location"`
}

func SelectQuery(session *gocql.Session, logger *zap.Logger, query string) []byte {
	logger.Info("Displaying Results:")
	q := session.Query(query)
	var jsonArray []map[string]string
	var fname, lname, addresss, picture_locationa string
	it := q.Iter()
	defer func() {
		if err := it.Close(); err != nil {
			logger.Warn("select", zap.Error(err))
		}
	}()
	for it.Scan(&fname, &lname, &addresss, &picture_locationa) {
		logger.Info("\t" + fname + " " + lname + ", " + addresss + ", " + picture_locationa)

		//var appData = &mutant_data{first_name: fname, last_name: lname, address: addresss, picture_location: picture_locationa}

		data := map[string]string{
			"first_name":       fname,
			"last_name":        lname,
			"address":          addresss,
			"picture_location": picture_locationa,
		}

		jsonArray = append(jsonArray, data)

	}

	b, err := json.Marshal(jsonArray)
	if err != nil {
		fmt.Println(err)
	}

	return b
}

func insertQuery(session *gocql.Session, logger *zap.Logger, query string) {
	logger.Info("Inserting")
	if err := session.Query(query).Exec(); err != nil {
		logger.Error("insert", zap.Error(err))
	}
}

func updateQuery(session *gocql.Session, logger *zap.Logger, query string) {
	logger.Info("Inserting")
	if err := session.Query(query).Exec(); err != nil {
		logger.Error("insert", zap.Error(err))
	}
}

func deleteQuery(session *gocql.Session, logger *zap.Logger, query string) {
	logger.Info("Deleting")
	if err := session.Query(query).Exec(); err != nil {
		logger.Error("delete error query", zap.Error(err))
	}
}

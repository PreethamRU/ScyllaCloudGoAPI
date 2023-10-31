package api

import (
	"golang-scyllacloud/db"
	"golang-scyllacloud/log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	session := db.Session
	logger := log.CreateLogger("info")
	val := db.SelectQuery(session, logger, "select * from catalog.mutant_data")

	w.Header().Set("Content-Type", "application/json")
	w.Write(val)
}

/*
	1. Generate random numbers of required length
	2. insert it into database
	3. assign a shortcode for a URL
	4. Store it in cache for shortcode against URL
	5. link click tracking -> stream it to webhook and store it in database
*/

package main

import (
	"database/sql"
	"fmt"
	"strings"
)

type ServeRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Year      string `json:"year"`
	Stat      string `json:"stat"` //not important
}

type ServeResponse struct {
	RES []float64 `json:"res"`
	OK  bool      `json:"ok"`
}

/*type Ace struct {
	value float64 `json:"value"`
}*/

const getServe1 = `SELECT %s from OverallServe O, Players P WHERE P.playerID=O.playerID AND P.lastName="%s";`

func GetServe(req ServeRequest, db *sql.DB) (ServeResponse, error) {
	var yearValue = "O." + mp[req.Year]
	var builtQuery = fmt.Sprintf(getServe1, yearValue, strings.ToLower(req.LastName))
	fmt.Println(builtQuery)
	res, err := getQueryResServe(builtQuery, db)
	if err != nil {
		return ServeResponse{OK: false}, fmt.Errorf("Could not retrieve overall serve stat.")
	}
	return ServeResponse{RES: res, OK: true}, nil
}

func getQueryResServe(builtQuery string, db *sql.DB) ([]float64, error) {
	rows, err := db.Query(builtQuery)
	logger.Info("Just queried database")
	fmt.Println(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var serve []float64

	for rows.Next() {
		logger.Info("In for loop for rows")
		var stat float64
		if err := rows.Scan(&stat); err != nil {
			return serve, err
		}
		serve = append(serve, stat)
	}
	return serve, nil
}

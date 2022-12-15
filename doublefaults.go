package main

import (
	"database/sql"
	"fmt"
	"strings"
)

type DFRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Year      string `json:"year"`
	Stat      string `json:"stat"` //not important
}

type DFResponse struct {
	RES []float64 `json:"res"`
	OK  bool      `json:"ok"`
}

/*type Ace struct {
	value float64 `json:"value"`
}*/

const getDF1 = `SELECT %s from DoubleFaults D, Players P WHERE P.playerID=D.playerID AND P.lastName="%s";`

func GetDF(req DFRequest, db *sql.DB) (DFResponse, error) {
	var yearValue = "D." + mp[req.Year]
	var builtQuery = fmt.Sprintf(getDF1, yearValue, strings.ToLower(req.LastName))
	fmt.Println(builtQuery)
	res, err := getQueryResDF(builtQuery, db)
	if err != nil {
		return DFResponse{OK: false}, fmt.Errorf("Could not retrieve double fault stat.")
	}
	return DFResponse{RES: res, OK: true}, nil
}

func getQueryResDF(builtQuery string, db *sql.DB) ([]float64, error) {
	rows, err := db.Query(builtQuery)
	logger.Info("Just queried database")
	fmt.Println(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doubleFaults []float64

	for rows.Next() {
		logger.Info("In for loop for rows")
		var stat float64
		if err := rows.Scan(&stat); err != nil {
			return doubleFaults, err
		}
		doubleFaults = append(doubleFaults, stat)
	}
	return doubleFaults, nil
}

package main

import (
	"database/sql"
	"fmt"
	"strings"
)

type BreakRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Year      string `json:"year"`
	Stat      string `json:"stat"` //not important
}

type BreakResponse struct {
	RES []float64 `json:"res"`
	OK  bool      `json:"ok"`
}

/*type Ace struct {
	value float64 `json:"value"`
}*/

const getBreak1 = `SELECT %s from BreakSaved B, Players P WHERE P.playerID=B.playerID AND P.lastName="%s";`

func GetBreak(req BreakRequest, db *sql.DB) (BreakResponse, error) {
	var yearValue = "B." + mp[req.Year]
	var builtQuery = fmt.Sprintf(getBreak1, yearValue, strings.ToLower(req.LastName))
	fmt.Println(builtQuery)
	res, err := getQueryResBreak(builtQuery, db)
	if err != nil {
		return BreakResponse{OK: false}, fmt.Errorf("Could not retrieve ace stat.")
	}
	return BreakResponse{RES: res, OK: true}, nil
}

func getQueryResBreak(builtQuery string, db *sql.DB) ([]float64, error) {
	rows, err := db.Query(builtQuery)
	logger.Info("Just queried database")
	fmt.Println(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var aces []float64

	for rows.Next() {
		logger.Info("In for loop for rows")
		var stat float64
		if err := rows.Scan(&stat); err != nil {
			return aces, err
		}
		aces = append(aces, stat)
	}
	return aces, nil
}

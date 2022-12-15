package main

import (
	"database/sql"
	"fmt"
	"strings"
)

type SecondServeWonRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Year      string `json:"year"`
	Stat      string `json:"stat"` //not important
}

type SecondServeWonResponse struct {
	RES []float64 `json:"res"`
	OK  bool      `json:"ok"`
}

/*type Ace struct {
	value float64 `json:"value"`
}*/

const getSecondWon1 = `SELECT %s from SecondServeWon S, Players P WHERE P.playerID=S.playerID AND P.lastName="%s";`

func GetSecondServeWon(req SecondServeWonRequest, db *sql.DB) (SecondServeWonResponse, error) {
	var yearValue = "S." + mp[req.Year]
	var builtQuery = fmt.Sprintf(getSecondWon1, yearValue, strings.ToLower(req.LastName))
	fmt.Println(builtQuery)
	res, err := getQueryResSecondWon(builtQuery, db)
	if err != nil {
		return SecondServeWonResponse{OK: false}, fmt.Errorf("Could not retrieve overall second serve won stat.")
	}
	return SecondServeWonResponse{RES: res, OK: true}, nil
}

func getQueryResSecondWon(builtQuery string, db *sql.DB) ([]float64, error) {
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

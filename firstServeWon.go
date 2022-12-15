package main

import (
	"database/sql"
	"fmt"
	"strings"
)

type FirstServeWonRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Year      string `json:"year"`
	Stat      string `json:"stat"` //not important
}

type FirstServeWonResponse struct {
	RES []float64 `json:"res"`
	OK  bool      `json:"ok"`
}

/*type Ace struct {
	value float64 `json:"value"`
}*/

const getFirstWon1 = `SELECT %s from FirstServeWon F, Players P WHERE P.playerID=F.playerID AND P.lastName="%s";`

func GetFirstServeWon(req FirstServeWonRequest, db *sql.DB) (FirstServeWonResponse, error) {
	var yearValue = "F." + mp[req.Year]
	var builtQuery = fmt.Sprintf(getFirstWon1, yearValue, strings.ToLower(req.LastName))
	fmt.Println(builtQuery)
	res, err := getQueryResFirstWon(builtQuery, db)
	if err != nil {
		return FirstServeWonResponse{OK: false}, fmt.Errorf("Could not retrieve overall first serve won stat.")
	}
	return FirstServeWonResponse{RES: res, OK: true}, nil
}

func getQueryResFirstWon(builtQuery string, db *sql.DB) ([]float64, error) {
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

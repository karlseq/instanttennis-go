package main

import (
	"database/sql"
	"fmt"
	"strings"
)

type AcesRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Year      string `json:"year"`
	Stat      string `json:"stat"` //not important
}

type AcesResponse struct {
	RES []float64 `json:"res"`
	OK  bool      `json:"ok"`
}

/*type Ace struct {
	value float64 `json:"value"`
}*/

const getAces1 = `SELECT %s from Aces A, Players P WHERE P.playerID=A.playerID AND P.lastName="%s";`

func GetAce(req AcesRequest, db *sql.DB) (AcesResponse, error) {
	var yearValue = "A." + mp[req.Year]
	var builtQuery = fmt.Sprintf(getAces1, yearValue, strings.ToLower(req.LastName))
	fmt.Println(builtQuery)
	res, err := getQueryRes(builtQuery, db)
	if err != nil {
		return AcesResponse{OK: false}, fmt.Errorf("Could not retrieve ace stat.")
	}
	return AcesResponse{RES: res, OK: true}, nil
}

func getQueryRes(builtQuery string, db *sql.DB) ([]float64, error) {
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

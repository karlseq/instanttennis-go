package main

import (
	"database/sql"
	"fmt"
	"strings"
)

type FirstServeInRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Year      string `json:"year"`
	Stat      string `json:"stat"` //not important
}

type FirstServeInResponse struct {
	RES []float64 `json:"res"`
	OK  bool      `json:"ok"`
}

/*type Ace struct {
	value float64 `json:"value"`
}*/

const getFirstIn1 = `SELECT %s from FirstServeIn F, Players P WHERE P.playerID=F.playerID AND P.lastName="%s";`

func GetFirstServeIn(req FirstServeInRequest, db *sql.DB) (FirstServeInResponse, error) {
	var yearValue = "F." + mp[req.Year]
	var builtQuery = fmt.Sprintf(getFirstIn1, yearValue, strings.ToLower(req.LastName))
	fmt.Println(builtQuery)
	res, err := getQueryResFirstIn(builtQuery, db)
	if err != nil {
		return FirstServeInResponse{OK: false}, fmt.Errorf("Could not retrieve first serve in stat.")
	}
	return FirstServeInResponse{RES: res, OK: true}, nil
}

func getQueryResFirstIn(builtQuery string, db *sql.DB) ([]float64, error) {
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

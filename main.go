package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"go.uber.org/zap"
)

type HandleResponse struct {
	OK    bool   `json:"ok"`
	ReqID string `json:"req_id"`
}

type HandleRequest struct {
	Event string          `json:"event"`
	Body  json.RawMessage `json:"body"`
}

var logger *zap.Logger
var db *sql.DB

// This function initializes the database connection
func initDatabaseConnection() {
	l, _ := zap.NewProduction()
	logger = l
	logger.Info("Getting DB connection")

	dbConnection, err := GetConnection()
	if err != nil {
		logger.Error("error connecting to database", zap.Error(err))
		panic(err)
	}

	logger.Info("Pinging Database")
	err = dbConnection.Ping()
	if err != nil {
		logger.Error("error pinging database", zap.Error(err))
		panic(err)
	}

	// Set global var
	db = dbConnection
}

// Handle the calls
func Handle(ctx context.Context, req HandleRequest) (interface{}, error) {
	var reqID string
	if lc, ok := lambdacontext.FromContext(ctx); ok {
		reqID = lc.AwsRequestID
	}

	select {
	case <-ctx.Done():
		return HandleResponse{OK: false, ReqID: reqID}, fmt.Errorf("request timeout: %w", ctx.Err())
	default:
	}

	//Initialize Database
	initDatabaseConnection()

	//This is the first row in the json request and will do certain things based on this variable
	logger.Info("before switch case")
	switch req.Event {
	// EMPLOYEE
	case "aces":
		var dest AcesRequest
		if err := json.Unmarshal(req.Body, &dest); err != nil {
			return nil, err
		}
		return GetAce(dest, db)
	case "doubleFaults":
		var dest DFRequest
		if err := json.Unmarshal(req.Body, &dest); err != nil {
			return nil, err
		}
		return GetDF(dest, db)
	case "overallServe":
		var dest ServeRequest
		if err := json.Unmarshal(req.Body, &dest); err != nil {
			return nil, err
		}
		return GetServe(dest, db)
	case "firstIn":
		var dest FirstServeInRequest
		if err := json.Unmarshal(req.Body, &dest); err != nil {
			return nil, err
		}
		return GetFirstServeIn(dest, db)
	case "firstWon":
		var dest FirstServeWonRequest
		if err := json.Unmarshal(req.Body, &dest); err != nil {
			return nil, err
		}
		return GetFirstServeWon(dest, db)
	case "secondWon":
		var dest SecondServeWonRequest
		if err := json.Unmarshal(req.Body, &dest); err != nil {
			return nil, err
		}
		return GetSecondServeWon(dest, db)
	case "break":
		var dest BreakRequest
		if err := json.Unmarshal(req.Body, &dest); err != nil {
			return nil, err
		}
		return GetBreak(dest, db)
	}
	db.Close()
	return HandleResponse{OK: false, ReqID: reqID}, fmt.Errorf("%s is an unknown event", req.Event)
}

func main() {
	lambda.Start(Handle)
}

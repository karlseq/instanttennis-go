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
	switch req.Event {
	// EMPLOYEE
	case "test":
		var dest TestRequest
		if err := json.Unmarshal(req.Body, &dest); err != nil {
			return nil, err
		}
		return TestFunc(dest)
	}
	db.Close()
	return HandleResponse{OK: false, ReqID: reqID}, fmt.Errorf("%s is an unknown event", req.Event)
}

func main() {
	lambda.Start(Handle)
}

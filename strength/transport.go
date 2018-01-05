package strength

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

// MakeIndexStrengthEndpoint is the endpoint for retrieving data.
func MakeIndexStrengthEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// TODO error handling...later
		// list, err := svc.Index()
		list := svc.Index(request)
		// if err != nil {
		//   return strengthResponse{list, err.Error()}, nil
		// }
		return strengthResponse{list, Workout{}, "", ""}, nil
	}
}

// MakeAddRowsEndpoint is the endpoint for adding rows.
func MakeAddRowsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// TODO error handling...later
		// list, err := svc.Index()
		svc.AddRows(request)
		list := svc.Index(request)
		return strengthResponse{list, Workout{}, "", ""}, nil
	}
}

// MakeSaveWorkoutEndpoint is the endpoint for saving a workout.
func MakeSaveWorkoutEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// TODO error handling...later
		// list, err := svc.Index()
		workout := svc.SaveWorkout(request)
		return strengthResponse{nil, workout, "", ""}, nil
	}
}

// MakeUpdateDateEndpoint is the endpoint for updating the date for records.
func MakeUpdateDateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		svc.UpdateRowsDate(request)

		return strengthResponse{nil, Workout{}, "", ""}, nil
	}
}

// MakeDeleteRowEndpoint the endpoint for deleting a row.
func MakeDeleteRowEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		svc.DeleteRow(request)

		return strengthResponse{nil, Workout{}, "", ""}, nil
	}
}

// DecodeStrengthRequest will decode the request paramters without a body.
func DecodeStrengthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var row Row
	var request strengthRequest

	vars := mux.Vars(r)

	switch urlPath := r.URL.Path; urlPath {
	case "/strength/deleteRow":
		q := r.URL.Query()
		request.UserID = q["userID"][0]
		rowIds := q["rowIds"][0]
		rowIDList := strings.Split(rowIds, ",")
		if len(rowIDList) > 0 {
			for i := 0; i < len(rowIDList); i++ {
				if id, err := strconv.ParseInt(rowIDList[i], 10, 32); err == nil {
					row.RowIds = append(row.RowIds, int32(id))
				}
			}
			request.Row = row
		}
		break
	default:
		request.UserID = vars["userID"]
		if request.UserID == "" {
			return nil, errors.New("userID missing")
		}
		break
	}

	return request, nil
}

// DecodeStrengthBodyRequest will decode the request with request bodies(PUT, POST etc).
func DecodeStrengthBodyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request strengthRequest
	json.NewDecoder(r.Body).Decode(&request)

	switch urlPath := r.URL.Path; urlPath {
	case "/strength/addRows":
		if request.UserID == "" {
			return nil, errors.New("userID missing")
		}
		break
	case "/strength/saveWorkout":
		if request.Workout == (Workout{}) {
			return nil, errors.New("no workout record passed to update")
		}
		break
	case "/strength/updateDate":
		if len(request.Row.RowIds) == 0 {
			return nil, errors.New("no rowIds passed")
		}
		break
	}

	return request, nil
}

// EncodeResponse will encode the results and return the response.
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type strengthRequest struct {
	UserID    string     `json:"userID"`
	Workout   Workout    `json:"workout"`
	List      []Strength `json:"list"`
	Row       Row        `json:"row"`
	StartDate int64      `json:"startDate"`
	EndDate   int64      `json:"endDate"`
	Amount    int        `json:"amount"`
}

type strengthResponse struct {
	List    []Strength `json:"list"`
	Workout Workout    `json:"workout"`
	Err     string     `json:"err, omitempty"`
	Code    string     `json:"code, omitempty"`
}

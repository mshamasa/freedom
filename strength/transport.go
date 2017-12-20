package strength

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

// MakeIndexStrengthEndpoint is the endpoint for retrieving data.
func MakeIndexStrengthEndpoint(svc StrengthService) endpoint.Endpoint {
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

// MakeSaveRowEndpoint is the endpoint for saving a row.
func MakeSaveRowEndpoint(svc StrengthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// TODO error handling...later
		// list, err := svc.Index()
		svc.SaveRow(request)
		return strengthResponse{nil, Workout{}, "", ""}, nil
	}
}

// MakeSaveWorkoutEndpoint is the endpoint for saving a workout.
func MakeSaveWorkoutEndpoint(svc StrengthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// TODO error handling...later
		// list, err := svc.Index()
		workout := svc.SaveWorkout(request)
		return strengthResponse{nil, workout, "", ""}, nil
	}
}

// MakeUpdateDateEndpoint is the endpoint for updating the date for records.
func MakeUpdateDateEndpoint(svc StrengthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		svc.UpdateRowsDate(request)

		return strengthResponse{nil, Workout{}, "", ""}, nil
	}
}

// DecodeStrengthGetRequest will decode the request paramters.
func DecodeStrengthGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	// TODO passing params through url
	// queries := r.URL.Query()
	// request := strengthRequest{vars["userId"], queries["startDate"][0], queries["endDate"][0]}
	request := strengthRequest{vars["userId"], Workout{}, nil, Row{}}

	if request.UserID == "" {
		return nil, errors.New("userId missing")
	}
	return request, nil
}

// DecodeStrengthRequest will decode the request paramters.
func DecodeStrengthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request strengthRequest
	json.NewDecoder(r.Body).Decode(&request)

	if request.UserID == "" {
		return nil, errors.New("userId missing")
	}
	if len(request.List) == 0 && request.Workout == (Workout{}) && len(request.Row.RowIds) == 0 {
		return nil, errors.New("Nothing to Save/Update")
	}

	return request, nil
}

// EncodeResponse will encode the results and return the response.
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type strengthRequest struct {
	UserID  string     `json:"userId"`
	Workout Workout    `json:"workout"`
	List    []Strength `json:"list"`
	Row     Row        `json:"row"`
}

type strengthResponse struct {
	List    []Strength `json:"list"`
	Workout Workout    `json:"workout"`
	Err     string     `json:"err, omitempty"`
	Code    string     `json:"code, omitempty"`
}

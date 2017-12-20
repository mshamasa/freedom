package strength

import (
  "context"
  "errors"
  "encoding/json"
  "net/http"

  "github.com/go-kit/kit/endpoint"
  "github.com/gorilla/mux"
)

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

func MakeSaveRowEndpoint(svc StrengthService) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (interface{}, error) {
    // TODO error handling...later
    // list, err := svc.Index()
    svc.SaveRow(request)
    return strengthResponse{nil, Workout{}, "", ""}, nil
  }
}

func MakeSaveWorkoutEndpoint(svc StrengthService) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (interface{}, error) {
    // TODO error handling...later
    // list, err := svc.Index()
    workout := svc.SaveWorkout(request)
    return strengthResponse{nil, workout, "", ""}, nil
  }
}

func DecodeStrengthGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
  vars := mux.Vars(r)
  // TODO passing params through url
  // queries := r.URL.Query()
  // request := strengthRequest{vars["userId"], queries["startDate"][0], queries["endDate"][0]}
  request := strengthRequest{vars["userId"], Workout{}, nil}

  if request.UserId == "" {
    return nil, errors.New("UserId missing!")
  }
  return request, nil
}

func DecodeStrengthRequest(_ context.Context, r *http.Request) (interface{}, error) {
  var request strengthRequest
  json.NewDecoder(r.Body).Decode(&request)

  if request.UserId == "" {
    return nil, errors.New("UserId missing!")
  }
  if len(request.List) == 0 && request.Workout == (Workout{}){
    return nil, errors.New("Nothing to Save/Update")
  }

  return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
  return json.NewEncoder(w).Encode(response)
}

type strengthRequest struct {
  UserId  string      `json:"userId"`
  Workout Workout     `json:"workout"`
  List    []Strength  `json:"list"`
}

type strengthResponse struct {
  List    []Strength  `json:"list"`
  Workout Workout     `json:"workout"`
  Err     string      `json:"err, omitempty"`
  Code    string      `json:"code, omitempty"`
}
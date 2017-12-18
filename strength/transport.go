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
    return strengthResponse{list, "", ""}, nil
  }
}

func MakeSaveStrengthEndpoint(svc StrengthService) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (interface{}, error) {
    // TODO error handling...later
    // list, err := svc.Index()
    svc.Save(request)
    return strengthResponse{nil, "", ""}, nil
  }
}

func DecodeStrengthGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
  vars := mux.Vars(r)
  // TODO passing params through url
  // queries := r.URL.Query()
  // request := strengthRequest{vars["userId"], queries["startDate"][0], queries["endDate"][0]}
  request := strengthRequest{vars["userId"], nil}

  if request.UserId == "" {
    return nil, errors.New("UserId missing!")
  }
  return request, nil
}

func DecodeStrengthPostRequest(_ context.Context, r *http.Request) (interface{}, error) {
  var request strengthRequest
  json.NewDecoder(r.Body).Decode(&request)

  if request.UserId == "" {
    return nil, errors.New("UserId missing!")
  }
  if len(request.List) == 0 {
    return nil, errors.New("List Empty!")
  }

  return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
  return json.NewEncoder(w).Encode(response)
}

type strengthRequest struct {
  UserId string `json:"userId"`
  // TODO future....
  // StartDate string `json:"startDate"`
  // EndDate string `json:"endDate"`
  List []Strength `json:"list"`
}

type strengthResponse struct {
  List []Strength `json:"list"`
  Err string `json:"err, omitempty"`
  Code string `json:"code, omitempty"`
}
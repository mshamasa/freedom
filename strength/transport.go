package strength

import (
  "context"
  "errors"
  "encoding/json"
  "net/http"

  "github.com/go-kit/kit/endpoint"
  "github.com/gorilla/mux"
)

func MakeStrengthEndpoint(svc StrengthService) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (interface{}, error) {
    // TODO error handling...later
    // list, err := svc.Index()
    list := svc.Index()
    // if err != nil {
    //   return strengthResponse{list, err.Error()}, nil
    // }
    return strengthResponse{list, ""}, nil
  }
}

func DecodeStrengthRequest(_ context.Context, r *http.Request) (interface{}, error) {
  vars := mux.Vars(r)
  // userId := vars["userId"]
  request := strengthRequest{vars["userId"]}

  if request.UserId == "" {
    return nil, errors.New("UserId missing!")
  }
  return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
  return json.NewEncoder(w).Encode(response)
}

type strengthRequest struct {
  UserId string `json:"userId"`
}

type strengthResponse struct {
  List []Strength `json:"list"`
  Err string `json:"err, omitempty"`
}
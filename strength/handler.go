package strength

import httptransport "github.com/go-kit/kit/transport/http"

var svc StrengthService
var service = strengthService{}

var IndexHandler = httptransport.NewServer(
  MakeIndexStrengthEndpoint(service),
  DecodeStrengthGetRequest,
  EncodeResponse,
)

var SaveHandler = httptransport.NewServer(
  MakeSaveStrengthEndpoint(service),
  DecodeStrengthPostRequest,
  EncodeResponse,
)
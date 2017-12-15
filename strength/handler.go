package strength

import httptransport "github.com/go-kit/kit/transport/http"

var svc StrengthService
var service = strengthService{}

var IndexHandler = httptransport.NewServer(
  MakeIndexStrengthEndpoint(service),
  DecodeStrengthRequest,
  EncodeResponse,
)

var AddHandler = httptransport.NewServer(
  MakeAddStrengthEndpoint(service),
  DecodeStrengthRequest,
  EncodeResponse,
)
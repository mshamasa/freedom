package strength

import httptransport "github.com/go-kit/kit/transport/http"

var svc StrengthService
var service = strengthService{}

var IndexHandler = httptransport.NewServer(
  MakeStrengthEndpoint(service),
  DecodeStrengthRequest,
  EncodeResponse,
)

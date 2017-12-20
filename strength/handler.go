package strength

import httptransport "github.com/go-kit/kit/transport/http"

var svc StrengthService
var service = strengthService{}

var IndexHandler = httptransport.NewServer(
	MakeIndexStrengthEndpoint(service),
	DecodeStrengthGetRequest,
	EncodeResponse,
)

var SaveRowHandler = httptransport.NewServer(
	MakeSaveRowEndpoint(service),
	DecodeStrengthRequest,
	EncodeResponse,
)

var SaveWorkoutHandler = httptransport.NewServer(
	MakeSaveWorkoutEndpoint(service),
	DecodeStrengthRequest,
	EncodeResponse,
)

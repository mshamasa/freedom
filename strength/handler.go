package strength

import httptransport "github.com/go-kit/kit/transport/http"

var svc StrengthService
var service = strengthService{}

// IndexHandler is the entry point to getting all user data.
var IndexHandler = httptransport.NewServer(
	MakeIndexStrengthEndpoint(service),
	DecodeStrengthGetRequest,
	EncodeResponse,
)

// SaveRowHandler is the entry point to saving rows.
var SaveRowHandler = httptransport.NewServer(
	MakeSaveRowEndpoint(service),
	DecodeStrengthRequest,
	EncodeResponse,
)

// SaveWorkoutHandler is the entry point to updating a single row.
var SaveWorkoutHandler = httptransport.NewServer(
	MakeSaveWorkoutEndpoint(service),
	DecodeStrengthRequest,
	EncodeResponse,
)

// UpdateDateHandler is the entry point to updating a single row.
var UpdateDateHandler = httptransport.NewServer(
	MakeUpdateDateEndpoint(service),
	DecodeStrengthRequest,
	EncodeResponse,
)

// DeleteRowHandler is the entry point to delete a row.
var DeleteRowHandler = httptransport.NewServer(
	MakeDeleteRowEndpoint(service),
	DecodeStrengthRequest,
	EncodeResponse,
)

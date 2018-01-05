package strength

import httptransport "github.com/go-kit/kit/transport/http"

var svc Service
var service = strengthService{}

// IndexHandler is the entry point to getting all user data.
var IndexHandler = httptransport.NewServer(
	MakeIndexStrengthEndpoint(service),
	DecodeStrengthRequest,
	EncodeResponse,
)

// DeleteRowHandler is the entry point to delete a row.
var DeleteRowHandler = httptransport.NewServer(
	MakeDeleteRowEndpoint(service),
	DecodeStrengthRequest,
	EncodeResponse,
)

// AddRowsHandler is the entry point to adding rows.
var AddRowsHandler = httptransport.NewServer(
	MakeAddRowsEndpoint(service),
	DecodeStrengthBodyRequest,
	EncodeResponse,
)

// SaveWorkoutHandler is the entry point to updating a single row.
var SaveWorkoutHandler = httptransport.NewServer(
	MakeSaveWorkoutEndpoint(service),
	DecodeStrengthBodyRequest,
	EncodeResponse,
)

// UpdateDateHandler is the entry point to updating a single row.
var UpdateDateHandler = httptransport.NewServer(
	MakeUpdateDateEndpoint(service),
	DecodeStrengthBodyRequest,
	EncodeResponse,
)

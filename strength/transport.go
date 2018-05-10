package strength

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// DecodeStrengthBodyRequest will decode the request with request bodies(PUT, POST etc).
func DecodeStrengthBodyRequest(r *http.Request) (Request, error) {
	var request Request
	json.NewDecoder(r.Body).Decode(&request)

	switch urlPath := r.URL.Path; urlPath {
	case "/strength/addRows":
		if request.UserID == "" {
			return request, errors.New("userID missing")
		}
		break
	case "/strength/saveWorkout":
		if request.Workout == (Workout{}) {
			return request, errors.New("no workout record passed to update")
		}
		break
	case "/strength/updateDate":
		if len(request.Row.RowIds) == 0 {
			return request, errors.New("no rowIds passed")
		}
		break
	}

	return request, nil
}

// EncodeResponse will encode the results and return the response.
func EncodeResponse(w http.ResponseWriter, response Response) {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Fatal(err)
	}
}

package query

import (
	"encoding/json"

	"github.com/ankitshah86/jsoniz/internal/types"
)

func ParseQuery(query string) (types.Request, error) {

	var request types.Request

	err := json.Unmarshal([]byte(query), &request)
	if err != nil {
		return types.Request{}, err
	}

	return request, nil
}

package core

import (
	"fmt"

	"github.com/ankitshah86/jsoniz/internal/types"
	"github.com/ankitshah86/jsoniz/pkg/query"
)

func ParseRequest(req string) (string, error) {

	request, err := query.ParseQuery(req)
	if err != nil {
		return "", err
	}

	switch request.Operation {
	case types.OperationType(types.SELECT):
		return "SELECT", nil
	case types.OperationType(types.INSERT):
		return "INSERT", nil
	case types.OperationType(types.UPDATE):
		return "UPDATE", nil
	case types.OperationType(types.DELETE):
		return "DELETE", nil
	default:
		fmt.Printf("invalid operation requested: %s\n", request.Operation)
		return "", fmt.Errorf("invalid operation requested: %s", request.Operation)
	}
}

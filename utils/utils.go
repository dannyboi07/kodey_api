package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	stdlog "log"
	"net/http"
	"os"
	"strings"
)

var Log stdlog.Logger

func InitLogger() {
	Log = *stdlog.New(os.Stdout, "", stdlog.Lshortfile|stdlog.Ldate|stdlog.Ltime)
}

func JsonParseErr(err error) (int, error) {
	if err == nil {
		return 0, nil
	}

	var syntaxError *json.SyntaxError
	var unmarshallTypeError *json.UnmarshalTypeError
	switch {
	case errors.As(err, &syntaxError):
		return http.StatusBadRequest, fmt.Errorf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)

	case errors.Is(err, io.ErrUnexpectedEOF):
		return http.StatusBadRequest, fmt.Errorf("Request body contains badly formed JSON")

	case errors.As(err, &unmarshallTypeError):
		return http.StatusBadRequest, fmt.Errorf("Request body contains an invalid value for field: %q, value: %q (at position: %d)", unmarshallTypeError.Field, unmarshallTypeError.Value, unmarshallTypeError.Offset)

	case strings.HasPrefix(err.Error(), "json: unknown field "):
		unknownFieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
		return http.StatusBadRequest, fmt.Errorf("Request body contains unknown field: %s", unknownFieldName)

	case errors.Is(err, io.EOF):
		return http.StatusBadRequest, errors.New("Request body cannot be empty")

	case err.Error() == "http: request body too large":
		return http.StatusRequestEntityTooLarge, errors.New("Request body cannot be larger than 1MB")
	default:
		return http.StatusBadRequest, err
	}
}

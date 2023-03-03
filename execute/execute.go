package execute

import (
	"context"
	"errors"
	"main/schema"
	"main/utils"
	"net/http"
	"time"
)

func ExecuteCode(code schema.ExecuteCode) ([]byte, int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var (
		output []byte
		err    error
	)
	switch *code.Language {
	case "python":
		output, err = RunPython(ctx, code)
	}

	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			utils.Log.Println("Code timed-out")
			return output, http.StatusBadRequest, errors.New("Timeout executing code")
		}

		return output, http.StatusInternalServerError, errors.New("Error executing code")
	}

	return output, 0, nil
}

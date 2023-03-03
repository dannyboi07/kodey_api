package controller

import (
	"encoding/json"
	"main/execute"
	"main/schema"
	"main/utils"
	"net/http"
)

func ExecuteCode(w http.ResponseWriter, r *http.Request) {
	var jDec json.Decoder = *json.NewDecoder(r.Body)
	jDec.DisallowUnknownFields()

	var (
		executeRequest schema.ExecuteCode
		statusCode     int
		err            error
	)

	statusCode, err = utils.JsonParseErr(jDec.Decode(&executeRequest))
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		utils.Log.Println("Cntrl/Client error: ", err)
		return
	}

	statusCode, err = executeRequest.Validate()
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		utils.Log.Println("Client error: ", err)
		return
	}

	var output []byte
	output, statusCode, err = execute.ExecuteCode(executeRequest)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		utils.Log.Println("Error executing user code, err: ", err, r.RemoteAddr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schema.ExecuteCodeResponse{
		Result:   string(output),
		Language: *executeRequest.Language,
	})
}

func FileSubmit(w http.ResponseWriter, r *http.Request) {

	var jDec json.Decoder = *json.NewDecoder(r.Body)
	jDec.DisallowUnknownFields()

	var (
		fileSubmit schema.FileSubmitReq
		statusCode int
		err        error
	)

	statusCode, err = utils.JsonParseErr(jDec.Decode(&fileSubmit))
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		utils.Log.Println("cntrl/client error: ", err)
		return
	}

	statusCode, err = fileSubmit.Validate()
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	utils.Log.Println("Request body", *fileSubmit.Code, *fileSubmit.FileName, *fileSubmit.Type)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
	<html>
		<body style="
			min-height:100vh;
			display:flex;
			justify-content:center;
			align-items:center"
		>
			<h2>
				Received request
			</h2>
		</body>
	</html>
	`))
}

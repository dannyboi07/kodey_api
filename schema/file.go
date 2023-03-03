package schema

import (
	"fmt"
	"net/http"
	"path"
)

type FileSubmitReq struct {
	Code      *string `json:"code"`
	Type      *string `json:"type"`
	FileName  *string `json:"file_name"`
	extension string
}

var typeExtMap = map[string]string{
	"python":          ".py",
	"javascript":      ".js",
	"javascriptreact": ".jsx",
	"typescript":      ".tsx",
	"typescriptreact": ".tsx",
}

func (f *FileSubmitReq) IsMissing() bool {
	return f.Code != nil || f.Type != nil || f.FileName != nil
}

func (f *FileSubmitReq) Validate() (int, error) {
	switch {
	case f.Code == nil:
		return http.StatusBadRequest, fmt.Errorf("Field 'code' is empty")
	case f.Type == nil:
		return http.StatusBadRequest, fmt.Errorf("Field 'type' is empty")
	case f.FileName == nil:
		return http.StatusBadRequest, fmt.Errorf("Field 'file_name' is empty")
	}

	if shouldBeExt, exists := typeExtMap[*f.Type]; !exists {
		return http.StatusBadRequest, fmt.Errorf("Invalid file type")
	} else {

		var fileExtension string = path.Ext(*f.FileName)
		if shouldBeExt != fileExtension {
			return http.StatusBadRequest, fmt.Errorf("Invalid file extension")
		}

		f.extension = fileExtension
	}

	return 0, nil
}

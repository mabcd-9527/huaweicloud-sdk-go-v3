package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ExportConsistencyResultsResponse Response Object
type ExportConsistencyResultsResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ExportConsistencyResultsResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ExportConsistencyResultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportConsistencyResultsResponse struct{}"
	}

	return strings.Join([]string{"ExportConsistencyResultsResponse", string(data)}, " ")
}

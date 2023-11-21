package metadata

import (
	"fmt"
	"net/http"
)

func (m *MetadataService) uploadHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("req: ", req)
}

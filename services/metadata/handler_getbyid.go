package metadata

import (
	"fmt"
	"net/http"
)

func (m *MetadataService) getByIDHandler(res http.ResponseWriter, req *http.Request) {
	value, ok := req.URL.Query()["id"]
	if !ok || len(value[0]) < 1 {
		fmt.Println("Url Param 'id' is missing")
	}
}

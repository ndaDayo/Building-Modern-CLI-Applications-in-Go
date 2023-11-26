package metadata

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ndaDayo/Audio-Metadata-CLI/storage"
)

func (m *MetadataService) getByIDHandler(res http.ResponseWriter, req *http.Request) {
	value, ok := req.URL.Query()["id"]
	if !ok || len(value[0]) < 1 {
		fmt.Println("Url Param 'id' is missing")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	id := string(value[0])
	fmt.Println("requesting audio by id: ", id)

	audio, err := m.Storage.GetByID(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	audioString, err := audio.JSON()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.WriteString(res, audioString)
}

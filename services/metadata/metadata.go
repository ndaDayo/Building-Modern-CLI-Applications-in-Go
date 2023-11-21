package metadata

import (
	"fmt"
	"net/http"

	"github.com/ndaDayo/Audio-Metadata-CLI/internal/interfaces"
	"github.com/ndaDayo/Audio-Metadata-CLI/storage"
)

type MetadataService struct {
	Server  *http.Server
	Storage interfaces.Storage
}

func CreateMetadataService(port int, storage interfaces.Storage) *MetadataService {
	mux := http.NewServeMux()

	metadataService := &MetadataService{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%v", port),
			Handler: mux,
		},
		Storage: storage,
	}

	mux.HandleFunc("/upload", metadataService.uploadHandler)

	return metadataService
}

func Run(port int) {
	flatfileStorage := storage.FlatFile{}
	service := CreateMetadataService(port, flatfileStorage)
	err := service.Server.ListenAndServe()
	if err != nil {
		fmt.Println("error starting api: ", err)
	}
}

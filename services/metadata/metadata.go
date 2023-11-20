package metadata

import "net/http"

type MetadataService struct {
	Server  *http.Server
	Storage interfaces.Storage
}

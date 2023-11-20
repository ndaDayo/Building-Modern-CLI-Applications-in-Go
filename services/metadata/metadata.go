package metadata

import (
	"net/http"

	"github.com/ndaDayo/Audio-Metadata-CLI/internal/interfaces"
)

type MetadataService struct {
	Server  *http.Server
	Storage interfaces.Storage
}

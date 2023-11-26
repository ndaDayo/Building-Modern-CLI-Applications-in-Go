package metadata

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/ndaDayo/Audio-Metadata-CLI/models"
)

func TestUploadHandler(t *testing.T) {
	tests := []struct {
		name       string
		setupMock  func(*MockStorage)
		wantStatus int
		body       *bytes.Buffer
	}{
		{
			name: "Successful Upload",
			setupMock: func(ms *MockStorage) {
				ms.UploadFunc = func(data []byte, filename string) (string, string, error) {
					return "id123", "/path/to/audio", nil
				}
				ms.SaveMetadataFunc = func(audio *models.Audio) error {
					return nil
				}
			},
			wantStatus: http.StatusOK,
			body:       bytes.NewBufferString("test file content"),
		},
	}
}

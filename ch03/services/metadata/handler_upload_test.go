package metadata

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUploadHandler(t *testing.T) {
	tests := []struct {
		name       string
		setupMock  func(*MockStorage)
		wantStatus int
		body       *bytes.Buffer
	}{
		{
			name:       "No file provided",
			setupMock:  func(ms *MockStorage) {},
			wantStatus: http.StatusBadRequest,
			body:       bytes.NewBuffer(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := &MockStorage{}
			tt.setupMock(mockStorage)
			service := &MetadataService{Storage: mockStorage}
			req, err := http.NewRequest(http.MethodPost, "/upload", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(service.uploadHandler)
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler return wrong status code: got %v want %v", status, tt.wantStatus)
			}
		})

	}
}

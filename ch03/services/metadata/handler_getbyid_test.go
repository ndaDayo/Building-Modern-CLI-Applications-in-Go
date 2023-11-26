package metadata

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ndaDayo/Audio-Metadata-CLI/models"
)

func TestGetByIDHandler(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		setupMock  func(*MockStorage)
		wantStatus int
	}{
		{
			name:       "Url Param 'id' is missing",
			url:        "/",
			setupMock:  func(ms *MockStorage) {},
			wantStatus: http.StatusInternalServerError,
		},
		{
			name: "Not Found Error",
			url:  "/?id=non-existent-id",
			setupMock: func(ms *MockStorage) {
				ms.GetByIDFunc = func(id string) (*models.Audio, error) {
					return nil, errors.New("not found")
				}
			},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := &MockStorage{}
			tt.setupMock(mockStorage)
			service := &MetadataService{Storage: mockStorage}
			req, err := http.NewRequest(http.MethodGet, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(service.getByIDHandler)
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler return wrong status code: got %v want %v", status, tt.wantStatus)
			}
		})
	}

}

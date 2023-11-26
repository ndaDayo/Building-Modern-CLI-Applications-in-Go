package metadata

import (
	"errors"

	"github.com/ndaDayo/Audio-Metadata-CLI/models"
)

type MockStorage struct {
	GetByIDFunc func(id string) (*models.Audio, error)
}

func (m *MockStorage) GetByID(id string) (*models.Audio, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id)
	}
	return nil, errors.New("GetByID not implemented")
}

func (m *MockStorage) Upload(bytes []byte, filename string) (string, string, error) {
	return "", "", errors.New("Upload not implemented")
}

func (m *MockStorage) SaveMetadata(audio *models.Audio) error {
	return errors.New("SaveMetadata not implemented")
}

func (m *MockStorage) List() ([]*models.Audio, error) {
	return nil, errors.New("List not implemented")
}

func (m *MockStorage) Delete(id string) error {
	return errors.New("Delete not implemented")
}

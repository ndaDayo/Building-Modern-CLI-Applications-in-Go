package storage

import "github.com/ndaDayo/Audio-Metadata-CLI/models"

type FlatFile struct {
	Name string
}

func (f FlatFile) GetByID(id string) (*models.Audio, error) {
	return nil
}

func (f FlatFile) SaveMetadata(audio *models.Audio) error {
	return nil
}

func (f FlatFile) Upload(bytes []byte, filename string) (string, string, error) {
	return nil
}

func (f FlatFile) List() ([]*models.Audio, error) {
	return nil
}

func (f FlatFile) Delete(id string) error {
	return nil
}

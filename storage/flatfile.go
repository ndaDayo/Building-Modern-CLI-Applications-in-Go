package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/ndaDayo/Audio-Metadata-CLI/models"
)

type FlatFile struct {
	Name string
}

func (f FlatFile) GetByID(id string) (*models.Audio, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	metadataFilePath := filepath.Join(dirname, "audiofile", id, "metadata.json")
	if _, err := os.Stat(metadataFilePath); errors.Is(err, os.ErrNotExist) {
		_ = os.Mkdir(metadataFilePath, os.ModePerm)
	}

	file, err := os.ReadFile(metadataFilePath)
	if err != nil {
		return nil, err
	}
	data := models.Audio{}

	err = json.Unmarshal([]byte(file), &data)
	return &data, err
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

package storage

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nuxxxcake/storage/internal/file"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	file, err := file.NewFile(filename, blob)

	if err != nil {
		return nil, err
	}

	s.files[file.ID] = file

	return file, nil
}

func (s *Storage) GetByID(fileID uuid.UUID) (*file.File, error) {
	file, ok := s.files[fileID]

	if !ok {
		return nil, errors.New("cannot find file by GUID")
	}

	return file, nil
}

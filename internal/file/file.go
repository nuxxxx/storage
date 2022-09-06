package file

import (
	"fmt"

	"github.com/google/uuid"
)

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, blob []byte) (*File, error) {
	guid, err := uuid.NewUUID()

	if err != nil {
		return &File{}, err
	}

	return &File{
		ID:   guid,
		Name: filename,
		Data: blob,
	}, nil
}

func (f *File) String() string {
	return fmt.Sprintf("File(%s, %v)", f.Name, f.ID)
}

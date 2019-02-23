package file

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// FileService #path:"/file/"#
type FileService struct {
	BasePath string
}

// NewFileService Create a new FileService
func NewFileService(basePath string) (*FileService, error) {
	err := os.MkdirAll(basePath, 0666)
	if err != nil {
		return nil, err
	}
	return &FileService{basePath}, nil
}

func (s *FileService) Handler() http.Handler {
	return http.FileServer(http.Dir(s.BasePath))
}

// Get a file #route:"GET /{filename}"#
func (s *FileService) Get(filename string) (file []byte /* #content:"application/octet-stream"# */, err error) {
	return ioutil.ReadFile(filepath.Join(s.BasePath, filename))
}

// Upload a file #route:"POST /"#
func (s *FileService) Upload(file io.Reader) (filename string, err error) {
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, file)
	data := buf.Bytes()
	h := md5.Sum(data)
	filename = hex.EncodeToString(h[:])

	err = ioutil.WriteFile(filepath.Join(s.BasePath, filename), data, 0666)
	if err != nil {
		return "", err
	}
	return filename, nil
}

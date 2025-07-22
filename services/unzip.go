package services

import (
	"archive/zip"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func ExtractZip(file multipart.File, dest string) (int, error) {
	// Read zip into buffer
	tempFile, err := os.CreateTemp("", "upload-*.zip")
	if err != nil {
		return 0, err
	}
	defer os.Remove(tempFile.Name())
	io.Copy(tempFile, file)
	tempFile.Close()

	r, err := zip.OpenReader(tempFile.Name())
	if err != nil {
		return 0, err
	}
	defer r.Close()

	var rootFolder string
	if len(r.File) > 0 {
		first := strings.Split(r.File[0].Name, "/")
		if len(first) > 1 {
			rootFolder = first[0] + "/"
		}
	}

	count := 0
	for _, f := range r.File {
		name := f.Name

		if rootFolder != "" && strings.HasPrefix(name, rootFolder) {
			name = strings.TrimPrefix(name, rootFolder)
		}

		if name == "" {
			continue
		}
		filePath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return 0, err
		}

		src, err := f.Open()
		if err != nil {
			return 0, err
		}
		defer src.Close()

		dst, err := os.Create(filePath)
		if err != nil {
			return 0, err
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return 0, err
		}
		count++
	}
	return count, nil
}

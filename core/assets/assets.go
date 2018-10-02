// Package assets is generated by github.com/omeid/go-resources
package assets

import (
	"bytes"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileSystem is an http.FileSystem implementation.
type FileSystem struct {
	files map[string]File
}

// String returns the content of the file as string.
func (fs *FileSystem) String(name string) (string, bool) {
	if filepath.Separator != '/' && strings.IndexRune(name, filepath.Separator) >= 0 ||
		strings.Contains(name, "\x00") {
		return "", false
	}

	file, ok := fs.files[name]

	if !ok {
		return "", false
	}

	return string(file.data), true
}

// Open implements http.FileSystem.Open
func (fs *FileSystem) Open(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.IndexRune(name, filepath.Separator) >= 0 ||
		strings.Contains(name, "\x00") {
		return nil, errors.New("http: invalid character in file path")
	}
	file, ok := fs.files[name]
	if !ok {
		files := []os.FileInfo{}
		for path, file := range fs.files {
			if strings.HasPrefix(path, name) {
				fi := file.fi
				files = append(files, &fi)
			}
		}

		if len(files) == 0 {
			return nil, os.ErrNotExist
		}

		//We have a directory.
		return &File{
			fi: FileInfo{
				isDir: true,
				files: files,
			}}, nil
	}
	file.Reader = bytes.NewReader(file.data)
	return &file, nil
}

// File implements http.File
type File struct {
	*bytes.Reader
	data []byte
	fi   FileInfo
}

// Close is a noop-closer.
func (f *File) Close() error {
	return nil
}

// Readdir implements http.File.Readdir
func (f *File) Readdir(count int) ([]os.FileInfo, error) {
	return nil, os.ErrNotExist
}

// Stat implements http.Stat.Readdir
func (f *File) Stat() (os.FileInfo, error) {
	return &f.fi, nil
}

// FileInfo implements the os.FileInfo interface.
type FileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
	sys     interface{}

	files []os.FileInfo
}

// Name implements os.FileInfo.Name
func (f *FileInfo) Name() string {
	return f.name
}

// Size implements os.FileInfo.Size
func (f *FileInfo) Size() int64 {
	return f.size
}

// Mode implements os.FileInfo.Mode
func (f *FileInfo) Mode() os.FileMode {
	return f.mode
}

// ModTime implements os.FileInfo.ModTime
func (f *FileInfo) ModTime() time.Time {
	return f.modTime
}

// IsDir implements os.FileInfo.IsDir
func (f *FileInfo) IsDir() bool {
	return f.isDir
}

// Readdir implements os.FileInfo.Readdir
func (f *FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return f.files, nil
}

// Sys returns the underlying value.
func (f *FileInfo) Sys() interface{} {
	return f.sys
}

var DEFAULTS *FileSystem

func init() {
	DEFAULTS = &FileSystem{
		files: map[string]File{
			"/core/config/mainnet.json": File{
				data: []byte{
					0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x22, 0x69, 0x64, 0x65, 0x6e, 0x74,
					0x69, 0x74, 0x79, 0x22, 0x3a, 0x20, 0x22, 0x6d, 0x61, 0x69, 0x6e, 0x6e,
					0x65, 0x74, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x22, 0x6e, 0x61,
					0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4e, 0x67, 0x69, 0x6e, 0x5f, 0x4d,
					0x61, 0x69, 0x6e, 0x6e, 0x65, 0x74, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20,
					0x20, 0x22, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x3a, 0x20, 0x7b, 0x0a,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x73, 0x74, 0x61,
					0x72, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x3a,
					0x20, 0x30, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x2c, 0x0a, 0x20, 0x20,
					0x20, 0x20, 0x22, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x3a,
					0x20, 0x35, 0x32, 0x35, 0x32, 0x30, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x22, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x22, 0x3a,
					0x20, 0x22, 0x4d, 0x30, 0x30, 0x4e, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20,
					0x20, 0x22, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x22, 0x3a, 0x20,
					0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x6e,
					0x6f, 0x6e, 0x63, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x30, 0x78, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x22, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
					0x3a, 0x20, 0x22, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x22, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73,
					0x68, 0x22, 0x3a, 0x20, 0x22, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x22, 0x65, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61,
					0x74, 0x61, 0x22, 0x3a, 0x20, 0x22, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d,
					0x69, 0x74, 0x22, 0x3a, 0x20, 0x22, 0x30, 0x78, 0x30, 0x33, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x22, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c,
					0x74, 0x79, 0x22, 0x3a, 0x20, 0x22, 0x30, 0x78, 0x30, 0x30, 0x30, 0x31,
					0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22,
					0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x22, 0x3a, 0x20, 0x22,
					0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22,
					0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x22, 0x3a, 0x20, 0x7b, 0x7d, 0x0a, 0x20,
					0x20, 0x20, 0x20, 0x7d, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x22, 0x63,
					0x68, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x3a,
					0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22,
					0x66, 0x6f, 0x72, 0x6b, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x0a, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7b, 0x0a,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20,
					0x22, 0x52, 0x61, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x22, 0x2c, 0x0a, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x22, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x3a, 0x20,
					0x30, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x72, 0x65, 0x71, 0x75,
					0x69, 0x72, 0x65, 0x64, 0x48, 0x61, 0x73, 0x68, 0x22, 0x3a, 0x20, 0x22,
					0x30, 0x78, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x22, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x3a,
					0x20, 0x5b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7b,
					0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x65, 0x69, 0x70, 0x31,
					0x35, 0x35, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
					0x73, 0x22, 0x3a, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x63,
					0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x22, 0x3a, 0x20, 0x31, 0x31, 0x31,
					0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x7d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d,
					0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x5d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x5d, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x22, 0x62, 0x61, 0x64, 0x48, 0x61, 0x73, 0x68,
					0x65, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x5d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x2c, 0x0a,
					0x20, 0x20, 0x20, 0x20, 0x22, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72,
					0x61, 0x70, 0x22, 0x3a, 0x20, 0x5b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x22, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x3a, 0x2f, 0x2f,
					0x30, 0x61, 0x61, 0x37, 0x33, 0x61, 0x38, 0x63, 0x63, 0x31, 0x35, 0x33,
					0x37, 0x32, 0x33, 0x33, 0x33, 0x34, 0x32, 0x63, 0x37, 0x38, 0x36, 0x39,
					0x32, 0x30, 0x64, 0x39, 0x35, 0x37, 0x34, 0x36, 0x61, 0x35, 0x61, 0x33,
					0x36, 0x37, 0x38, 0x38, 0x62, 0x37, 0x39, 0x35, 0x32, 0x39, 0x35, 0x65,
					0x32, 0x36, 0x33, 0x63, 0x61, 0x65, 0x37, 0x33, 0x35, 0x63, 0x39, 0x36,
					0x31, 0x64, 0x32, 0x36, 0x65, 0x34, 0x39, 0x32, 0x65, 0x35, 0x66, 0x36,
					0x62, 0x33, 0x65, 0x62, 0x37, 0x62, 0x38, 0x30, 0x65, 0x32, 0x65, 0x39,
					0x32, 0x37, 0x34, 0x34, 0x35, 0x66, 0x31, 0x39, 0x33, 0x31, 0x37, 0x63,
					0x61, 0x37, 0x63, 0x33, 0x62, 0x64, 0x37, 0x63, 0x33, 0x35, 0x32, 0x35,
					0x37, 0x38, 0x33, 0x65, 0x31, 0x38, 0x63, 0x35, 0x38, 0x34, 0x63, 0x61,
					0x37, 0x30, 0x35, 0x66, 0x31, 0x30, 0x38, 0x65, 0x40, 0x31, 0x34, 0x39,
					0x2e, 0x32, 0x38, 0x2e, 0x38, 0x33, 0x2e, 0x31, 0x31, 0x35, 0x3a, 0x35,
					0x32, 0x35, 0x32, 0x30, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x22, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x3a, 0x2f, 0x2f,
					0x64, 0x31, 0x37, 0x37, 0x36, 0x63, 0x34, 0x30, 0x61, 0x39, 0x33, 0x36,
					0x34, 0x36, 0x35, 0x35, 0x30, 0x36, 0x64, 0x64, 0x63, 0x39, 0x36, 0x30,
					0x31, 0x31, 0x63, 0x66, 0x39, 0x62, 0x39, 0x66, 0x61, 0x33, 0x30, 0x35,
					0x37, 0x61, 0x64, 0x64, 0x32, 0x31, 0x32, 0x32, 0x65, 0x39, 0x33, 0x37,
					0x32, 0x65, 0x65, 0x39, 0x36, 0x31, 0x63, 0x35, 0x63, 0x34, 0x66, 0x34,
					0x33, 0x66, 0x33, 0x32, 0x30, 0x39, 0x30, 0x34, 0x37, 0x63, 0x66, 0x39,
					0x34, 0x61, 0x38, 0x39, 0x38, 0x64, 0x31, 0x35, 0x65, 0x38, 0x32, 0x32,
					0x63, 0x64, 0x66, 0x36, 0x63, 0x33, 0x36, 0x61, 0x61, 0x35, 0x31, 0x31,
					0x66, 0x62, 0x31, 0x66, 0x30, 0x31, 0x35, 0x30, 0x66, 0x33, 0x66, 0x31,
					0x61, 0x64, 0x66, 0x31, 0x36, 0x30, 0x32, 0x34, 0x32, 0x64, 0x39, 0x61,
					0x31, 0x30, 0x62, 0x39, 0x36, 0x64, 0x66, 0x34, 0x40, 0x31, 0x39, 0x32,
					0x2e, 0x32, 0x32, 0x37, 0x2e, 0x32, 0x33, 0x32, 0x2e, 0x37, 0x35, 0x3a,
					0x35, 0x32, 0x35, 0x32, 0x30, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x22, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x3a, 0x2f,
					0x2f, 0x66, 0x37, 0x65, 0x34, 0x37, 0x38, 0x33, 0x64, 0x31, 0x62, 0x63,
					0x34, 0x30, 0x38, 0x63, 0x35, 0x62, 0x37, 0x30, 0x37, 0x39, 0x64, 0x65,
					0x63, 0x35, 0x37, 0x33, 0x39, 0x61, 0x63, 0x64, 0x62, 0x30, 0x30, 0x32,
					0x39, 0x34, 0x37, 0x36, 0x35, 0x62, 0x63, 0x66, 0x63, 0x38, 0x33, 0x30,
					0x63, 0x66, 0x36, 0x61, 0x38, 0x66, 0x62, 0x33, 0x65, 0x63, 0x33, 0x65,
					0x62, 0x32, 0x64, 0x65, 0x30, 0x35, 0x65, 0x32, 0x37, 0x35, 0x65, 0x36,
					0x39, 0x34, 0x65, 0x64, 0x65, 0x37, 0x66, 0x37, 0x38, 0x64, 0x63, 0x34,
					0x62, 0x36, 0x64, 0x32, 0x62, 0x34, 0x66, 0x65, 0x37, 0x64, 0x65, 0x61,
					0x35, 0x36, 0x34, 0x35, 0x35, 0x63, 0x39, 0x62, 0x36, 0x30, 0x36, 0x64,
					0x63, 0x38, 0x66, 0x66, 0x34, 0x61, 0x62, 0x66, 0x34, 0x61, 0x31, 0x65,
					0x33, 0x34, 0x35, 0x36, 0x66, 0x62, 0x65, 0x39, 0x34, 0x40, 0x31, 0x39,
					0x38, 0x2e, 0x31, 0x32, 0x2e, 0x37, 0x31, 0x2e, 0x32, 0x33, 0x30, 0x3a,
					0x35, 0x32, 0x35, 0x32, 0x30, 0x22, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x5d,
					0x0a, 0x7d, 0x0a,
				},
				fi: FileInfo{
					name:    "mainnet.json",
					size:    1455,
					modTime: time.Unix(0, 1538479444727798500),
					isDir:   false,
				},
			}, "/core/config/testnet.json": File{
				data: []byte{
					0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x22, 0x69, 0x64, 0x65, 0x6e, 0x74,
					0x69, 0x74, 0x79, 0x22, 0x3a, 0x20, 0x22, 0x61, 0x6c, 0x70, 0x68, 0x61,
					0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x22, 0x6e, 0x61, 0x6d, 0x65,
					0x22, 0x3a, 0x20, 0x22, 0x4e, 0x67, 0x69, 0x6e, 0x5f, 0x54, 0x65, 0x73,
					0x74, 0x6e, 0x65, 0x74, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x22,
					0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x3a, 0x20, 0x7b, 0x0a, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x73, 0x74, 0x61, 0x72, 0x74,
					0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x3a, 0x20, 0x31,
					0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x22, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x3a, 0x20, 0x32,
					0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x22, 0x63, 0x6f, 0x6e, 0x73, 0x65,
					0x6e, 0x73, 0x75, 0x73, 0x22, 0x3a, 0x20, 0x22, 0x4d, 0x30, 0x30, 0x4e,
					0x2d, 0x54, 0x45, 0x53, 0x54, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x22, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x22, 0x3a, 0x20, 0x7b,
					0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x6e, 0x6f,
					0x6e, 0x63, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x30, 0x78, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x34,
					0x31, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x22, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x3a,
					0x20, 0x22, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x22, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68,
					0x22, 0x3a, 0x20, 0x22, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x22, 0x65, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74,
					0x61, 0x22, 0x3a, 0x20, 0x22, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x22, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69,
					0x74, 0x22, 0x3a, 0x20, 0x22, 0x30, 0x78, 0x30, 0x33, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x22, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74,
					0x79, 0x22, 0x3a, 0x20, 0x22, 0x30, 0x78, 0x30, 0x30, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x31, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x22, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73,
					0x65, 0x22, 0x3a, 0x20, 0x22, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x22, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x22, 0x3a,
					0x20, 0x7b, 0x7d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x2c, 0x0a, 0x20,
					0x20, 0x20, 0x20, 0x22, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e,
					0x66, 0x69, 0x67, 0x22, 0x3a, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x22, 0x66, 0x6f, 0x72, 0x6b, 0x73, 0x22, 0x3a,
					0x20, 0x5b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x6e, 0x61,
					0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x44, 0x69, 0x65, 0x68, 0x61, 0x72,
					0x64, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x62, 0x6c, 0x6f,
					0x63, 0x6b, 0x22, 0x3a, 0x20, 0x31, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x48, 0x61, 0x73,
					0x68, 0x22, 0x3a, 0x20, 0x22, 0x30, 0x78, 0x30, 0x30, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x22,
					0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x66, 0x65, 0x61, 0x74, 0x75,
					0x72, 0x65, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x20,
					0x22, 0x65, 0x69, 0x70, 0x31, 0x35, 0x35, 0x22, 0x2c, 0x0a, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x6f,
					0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3a, 0x20, 0x7b, 0x0a, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x22, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x22,
					0x3a, 0x20, 0x31, 0x31, 0x31, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x5d, 0x0a,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x7d, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x6e, 0x61,
					0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x47, 0x6f, 0x74, 0x68, 0x61, 0x6d,
					0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x62, 0x6c, 0x6f, 0x63,
					0x6b, 0x22, 0x3a, 0x20, 0x35, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x2c,
					0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
					0x65, 0x64, 0x48, 0x61, 0x73, 0x68, 0x22, 0x3a, 0x20, 0x22, 0x30, 0x78,
					0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
					0x30, 0x30, 0x30, 0x30, 0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22,
					0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x3a, 0x20, 0x5b,
					0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7b, 0x0a, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22,
					0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
					0x22, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x22, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
					0x3a, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x65, 0x72, 0x61,
					0x22, 0x3a, 0x20, 0x35, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x2c, 0x0a,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x20,
					0x22, 0x65, 0x63, 0x69, 0x70, 0x31, 0x30, 0x31, 0x37, 0x22, 0x0a, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d,
					0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x5d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x5d, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x22, 0x62, 0x61, 0x64, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73,
					0x22, 0x3a, 0x20, 0x5b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x5d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x2c, 0x0a, 0x20, 0x20,
					0x20, 0x20, 0x22, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70,
					0x22, 0x3a, 0x20, 0x5b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x5d, 0x0a, 0x7d, 0x0a,
				},
				fi: FileInfo{
					name:    "testnet.json",
					size:    1445,
					modTime: time.Unix(0, 1538156190155053000),
					isDir:   false,
				},
			},
		},
	}
}

// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// FileSystem contains project templates.
var FileSystem = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 6, 6, 7, 58, 11, 389644409, time.UTC),
		},
		"/error.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "error.tmpl",
			modTime:          time.Date(2019, 6, 6, 8, 9, 43, 539983197, time.UTC),
			uncompressedSize: 824,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x92\xdf\x6e\x13\x3d\x10\xc5\xef\xf3\x14\xfe\xf6\xe2\xbb\x81\xd8\x49\xa0\xb4\xc0\x7a\x51\x81\xd2\x52\x54\xb5\x24\x6d\x69\x2f\xa7\xde\xc9\xae\x13\xff\x59\x79\x26\x49\x37\x51\x24\xde\x81\x37\xe4\x49\xd0\xb6\x2a\x88\x9b\x4a\x20\xae\xac\x73\xac\x39\xf3\x1b\xe9\xe4\xff\x95\xd1\x70\xdb\xa0\xa8\xd9\xbb\xa2\x97\x77\x8f\x70\x10\x2a\x8d\xa1\xe8\xe5\x1e\x19\x84\xa9\x21\x11\xb2\x5e\xf0\xb4\xbf\xf7\x60\x06\xf0\xa8\x97\x16\x57\x4d\x4c\x2c\x4c\x0c\x8c\x81\x75\xb6\xb2\x25\xd7\xba\xc4\xa5\x35\xd8\xbf\x13\x4f\x6d\xb0\x6c\xc1\xf5\xc9\x80\x43\x3d\xcc\x8a\x5e\xce\x96\x1d\x16\x9b\x8d\x9c\x30\xf0\x82\xde\xc5\x12\xb7\x5b\xb1\xd9\xd4\xcc\xcd\xbd\x75\x8e\xb7\x2c\x7e\xff\xfe\xfe\xf5\x9b\x98\x5b\x2f\xd9\x73\x9d\xaa\x52\x96\xb8\xcc\xd5\x7d\x52\x2f\x77\x36\xcc\x45\x42\xa7\x89\x5b\x87\x54\x23\xb2\xa8\x13\x4e\x75\x17\x49\xaf\x94\x32\x65\x98\x91\x34\x2e\x2e\xca\xa9\x83\x84\xd2\x44\xaf\x60\x06\xb7\xca\xd9\x1b\x52\x21\x26\x0f\xce\xae\x51\xed\xc9\x81\x1c\xfe\xd2\xd2\xdb\x20\x0d\x91\xb0\x81\xb1\x4a\x96\x5b\x9d\x51\x0d\xa3\x9d\x17\x7d\xb7\xb7\x73\xea\xcf\x4e\x67\x4b\xeb\xd4\xe4\x74\x79\xc9\xcf\x8e\xc2\x64\x32\x5b\x7f\x18\x9e\x5f\x9c\xb4\xe7\x2f\xf1\x72\x60\x46\x6f\xd7\x07\x87\xeb\x0b\x9d\x09\x93\x22\x51\x4c\xb6\xb2\x41\x43\x88\xa1\xf5\x71\x41\xff\x80\x9c\xe6\xe8\x90\x63\x50\x23\x39\x90\xcf\x7f\xca\x47\xb8\x47\xd7\x9f\xc7\xc7\x27\x57\xef\x77\x9b\x8f\xfb\x67\x47\xf6\x2a\x0d\xe8\xc9\xd2\x8d\xbf\xec\xef\x1e\x5e\x1f\x1f\x7c\x1a\xec\x8f\x03\xed\xce\x47\x74\x73\x74\xfd\x37\xdc\xd9\x03\xf8\x34\x06\x26\x59\xc5\x58\x39\x84\xc6\xd2\x1d\xb8\x21\x7a\x33\x05\x6f\x5d\xab\xc7\xe0\x70\x05\xed\xff\xe0\x9b\xd7\xa5\xa5\xc6\x41\xab\x69\x05\x4d\xf6\x58\xfa\x66\x03\x44\xc8\x67\xc0\xb5\xc8\x14\xa6\x14\x53\x77\x66\xb6\xdd\x76\xcd\x04\x1b\x84\x71\x40\xa4\xbb\x46\x82\x0d\x98\xba\x52\x0f\xff\xb0\x6b\xb9\xaa\x87\x45\x2f\x6f\xba\xb1\x13\x24\x82\xea\xce\x6c\x8a\x5e\xae\xba\x1d\xc5\x8f\x00\x00\x00\xff\xff\xd5\xe7\x05\x4f\x38\x03\x00\x00"),
		},
		"/post.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "post.tmpl",
			modTime:          time.Date(2022, 3, 6, 3, 46, 11, 414945480, time.UTC),
			uncompressedSize: 791,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x7c\x91\xcb\x6e\xdb\x3c\x10\x85\xf7\x7a\x0a\xfe\x5a\xfc\x68\x81\x5a\x4c\x52\xa4\x48\x53\x49\x8b\xde\xdd\x5b\x8c\xda\xc9\x9e\xa6\xc6\x12\x1d\x72\x46\x20\x47\x76\x64\x41\xdb\xee\xfb\x3c\x7d\x9b\x3e\x49\x41\xbb\x4e\xd1\x02\xed\x46\xd2\x99\x19\x7c\x3a\x33\x27\xff\xaf\x22\xcd\x7d\x0b\xa2\x61\x67\xcb\x24\x8f\x2f\x61\x15\xd6\x05\x60\x99\xe4\x0e\x58\x09\xdd\x28\x1f\x80\x8b\x8e\x57\x93\x8b\x63\x11\x95\x83\x62\x63\x60\xdb\x92\x67\xa1\x09\x19\x90\x8b\x74\x6b\x2a\x6e\x8a\x0a\x36\x46\xc3\x64\x2f\x1e\x19\x34\x6c\x94\x9d\x04\xad\x2c\x14\xa7\x69\x99\xe4\x6c\xd8\x42\xf9\xde\x38\xf1\x8e\xb0\x9e\x74\x28\x86\x21\x7b\xa1\x5a\x36\x84\xe3\x98\xcb\x43\x3f\xc9\xad\xc1\x5b\xe1\xc1\x16\x81\x7b\x0b\xa1\x01\x60\xd1\x78\x58\x15\x0d\x73\x1b\x2e\xa5\xd4\x15\xae\x43\xa6\x2d\x75\xd5\xca\x2a\x0f\x99\x26\x27\xd5\x5a\xdd\x49\x6b\x96\x41\x22\x79\xa7\xac\xd9\x81\xbc\xc8\x4e\xb2\xd3\x5f\x3a\x73\x06\x33\x1d\x82\x30\xc8\x50\x7b\xc3\x7d\x91\x86\x46\x9d\x9d\x3f\x99\xd8\x8b\xf3\x2b\x37\xbb\x5a\x6f\x8c\x95\xf3\xab\xcd\x0d\x3f\x7e\x8b\xf3\xf9\x7a\xf7\xfa\x74\x71\xfd\xb1\x5f\x3c\x85\x9b\x13\x7d\xf6\x7c\xf7\xea\xcd\xee\xba\x48\x85\xf6\x14\x02\x79\x53\x1b\x2c\x14\x12\xf6\x8e\xba\xf0\x2f\xe7\xc3\xa0\x42\x00\x9e\x29\x6e\x44\x2a\x5b\x0a\x1c\x7d\xa4\xe3\x58\x26\xf9\x7e\x56\x20\xa1\x86\x22\x8d\x17\x99\xcf\x3e\x45\x31\x8e\x69\xb9\xa4\xaa\x1f\x96\x4a\xdf\xd6\x9e\x3a\xac\x26\xc6\xa9\x1a\x2e\x3b\x6f\x1f\xc4\xc9\x69\x54\xe3\x98\x3e\x1c\x73\xb9\xa7\xc4\x28\x41\x55\xe0\xcb\x24\x6f\xcb\x5c\x1d\x7e\x9e\xca\x61\xc8\x66\x14\x78\xfa\x32\x32\xff\x1e\x80\x2a\x73\xd9\x96\x49\x2e\xef\x21\x2b\x22\xfe\x83\x26\xc5\x3e\xa7\x22\x45\xd8\x8a\xbd\x9f\xb4\x1c\x06\xb3\x12\xd9\x34\x4c\xb1\x82\xbb\x71\xfc\xfe\xe5\xdb\x30\x80\x0d\x10\x3f\xbf\x0e\x03\x60\x75\xc0\xff\x8f\xcb\xd0\x3e\x3b\x3c\x9d\xaa\x40\x2c\x7b\x71\xe4\x1e\xc3\x65\x72\xdc\x90\xa7\x9a\xa8\xca\x10\x78\x7f\x4f\xd5\xc5\x5a\xb9\x20\x27\x16\xc7\x66\x04\x8a\xad\xe1\xe6\xe0\x22\x88\x95\x27\x77\x8f\x8b\xf7\x99\x53\xe7\x35\x5c\x7f\xfe\x10\xd7\xbe\xd7\x8b\x68\xff\xb7\x75\x7f\x6e\xf9\x23\x00\x00\xff\xff\x42\xf7\xa5\xbc\x17\x03\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/error.tmpl"].(os.FileInfo),
		fs["/post.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}

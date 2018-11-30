// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package asset

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

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 11, 29, 5, 13, 10, 296716841, time.UTC),
		},
		"/etc": &vfsgen۰DirInfo{
			name:    "etc",
			modTime: time.Date(2018, 11, 29, 5, 14, 11, 748627076, time.UTC),
		},
		"/etc/docker": &vfsgen۰DirInfo{
			name:    "docker",
			modTime: time.Date(2018, 11, 29, 5, 14, 18, 386140001, time.UTC),
		},
		"/etc/docker/daemon.json": &vfsgen۰CompressedFileInfo{
			name:             "daemon.json",
			modTime:          time.Date(2018, 11, 29, 5, 16, 43, 339101512, time.UTC),
			uncompressedSize: 229,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\xe2\x54\xca\xcc\x2b\x4e\x4d\x2e\x2d\x4a\xd5\x2d\x4a\x4d\xcf\x2c\x2e\x29\xca\x4c\x2d\x56\xb2\x52\x88\xe6\x52\x50\x50\xe0\xe4\x54\x4a\xc9\x4f\xce\x4e\x2d\x52\xd2\xe1\x82\xb3\xf5\x32\xf3\x21\x5c\xa8\xfa\x4a\x5d\x43\x3d\x34\x99\xf4\x64\x04\x3b\xdb\xa2\x58\x0f\x99\x5f\x58\x9a\x58\x09\xe7\x18\x5a\x18\xe9\x19\x5a\x9a\xea\x59\x18\xea\x19\x1a\x1a\x5b\x59\x9a\x18\x18\xe2\x92\x30\xc2\x25\x61\x8c\x4b\xc2\x44\x89\x8b\x33\x96\xab\x96\x0b\x10\x00\x00\xff\xff\x4e\xc8\xf1\xb0\xe5\x00\x00\x00"),
		},
		"/etc/nginx": &vfsgen۰DirInfo{
			name:    "nginx",
			modTime: time.Date(2018, 11, 29, 5, 13, 10, 296716794, time.UTC),
		},
		"/etc/nginx/nginx.conf": &vfsgen۰CompressedFileInfo{
			name:             "nginx.conf",
			modTime:          time.Date(2018, 11, 29, 5, 13, 10, 296000000, time.UTC),
			uncompressedSize: 2514,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x92\x4d\x8f\x9b\x30\x10\x86\xef\xfe\x15\x73\xc8\x35\x10\x08\xad\xc8\x72\xef\xc7\xad\xea\xa9\x37\xcb\x8b\x07\x62\x01\x1e\x3a\x63\xf2\xd1\x2a\xff\xbd\x82\x90\x6a\xb3\x6d\x7a\xca\xa5\x6d\x90\x90\xc6\x7a\x9f\x99\xb1\xe4\x67\x10\x64\x00\x5f\x3b\x7f\x28\xd4\x9e\xb8\x41\xd6\x3d\x53\x89\x22\x28\x00\x66\x08\x54\x28\x85\xcc\xc4\xba\xa5\x1a\x20\xde\x19\x8e\x5b\xaa\xe3\xa9\x27\x9e\x92\x68\x4c\xf6\x86\x7d\xa1\x7a\x67\x61\xfe\x26\x92\x07\x7f\x26\xa3\xde\xd9\x71\xd2\x0e\x7d\x10\xf8\xae\x00\xe6\x6d\x25\x79\x8f\x65\x70\xe4\x05\x20\x59\xa5\x59\xa1\x4e\x4a\x6d\x43\xe8\x27\x4a\x90\x77\xc8\x53\x79\x39\x68\x6f\x3a\x04\xc6\xda\x49\xe0\xe3\x32\x89\x2c\x95\x0d\x72\xe4\xa8\x98\xa8\xd6\x49\x40\x0f\xf9\x6a\x3e\x52\x69\xc6\xf1\x10\xcf\x53\x00\x7a\xa6\xc3\x51\xf7\x46\x04\xc6\x45\x4f\x71\x9c\xe4\x69\x94\x6c\xde\x44\x79\x12\x25\xc9\xfa\x69\x93\xad\x92\xe2\x0a\x66\xb4\x8e\xb1\x0c\x40\x55\x75\x9d\x08\x06\xbd\x45\x63\x91\xe1\x03\x49\x80\xc5\x96\x24\xdc\x44\xbe\x2c\x3f\xa3\x69\x97\x1f\x3f\xc1\x82\xb1\xa3\x80\xda\x58\xcb\x7f\xc0\xdf\x11\xef\x0d\x5b\xb4\x63\x05\x8b\x33\x61\xac\xd5\x07\x5d\x5d\xa2\xb1\xba\x8c\x28\x5b\x87\x3e\xe8\xce\x1c\xf4\x33\xd9\xa3\x16\xf7\x0d\x21\x59\xbd\x7f\x95\x4f\xd9\xf3\x50\x55\xc8\x33\x92\xe6\xcd\xf5\x35\xe6\x97\xd1\xc1\x75\x48\x43\x80\xcd\xea\xf5\x35\xbd\xbd\x19\x32\x9a\xdb\xe1\xcb\xbd\x59\xf3\xbb\x4c\x20\x83\x75\xfa\x4b\x24\x3f\xf3\x73\xf3\xdb\xb1\x5b\x5d\x41\x01\xbb\x5e\x57\xae\x45\xbd\x67\x17\xf0\x05\x37\x32\x27\x75\xfe\x6f\x6a\x55\x97\xf7\x16\x29\x7d\x88\xf4\x97\x88\x74\x47\x8d\x9a\x5c\xa2\xfb\xab\xb4\x7e\xa8\xf4\xff\xa9\xf4\x75\x30\xc7\x3b\x7b\x94\x3d\x3c\xfa\x47\x3d\x3a\xa9\x1f\x01\x00\x00\xff\xff\x6d\xa8\x78\xe0\xd2\x09\x00\x00"),
		},
		"/etc/yum.repos.d": &vfsgen۰DirInfo{
			name:    "yum.repos.d",
			modTime: time.Date(2018, 11, 29, 5, 18, 31, 291890084, time.UTC),
		},
		"/etc/yum.repos.d/centos7.repo": &vfsgen۰CompressedFileInfo{
			name:             "centos7.repo",
			modTime:          time.Date(2018, 11, 29, 5, 12, 38, 890000000, time.UTC),
			uncompressedSize: 1696,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x93\x5f\x6b\xdb\x30\x14\xc5\xdf\xf5\x29\x2e\xa4\xaf\x8a\xea\xfd\x4b\x1a\xf0\xd3\x66\xfa\xb0\x76\x33\x09\x7b\x18\x61\x0f\x8a\x7c\x2d\x8b\xc8\x92\xd0\x9f\x62\x7f\xfb\xe1\xc4\x4d\x3b\x18\x2d\x24\x26\x6f\x32\xbe\xbe\xe7\xfc\x8e\x8f\xb6\xe8\x50\xff\x21\x86\xb7\x98\x17\x5d\xf4\x1c\x4a\x2e\xf6\x5c\x62\x80\xda\x7a\x28\x4c\x44\xef\xbc\x0a\x08\x0f\xca\xa4\x0e\x16\x40\xe1\x66\xc7\x03\x72\x2f\x1a\x32\x1c\x92\xd7\x79\x13\xa3\x5b\x31\x96\x2d\x3f\xcc\xb3\xbb\xcf\xf3\x65\x36\xcf\xb2\x8f\xab\xbb\x4f\x5f\x96\xcc\xa3\xb3\x41\x45\xeb\x7b\x26\x5a\x47\xfb\xd4\xd2\x41\x92\xdd\x78\xd4\xc8\x03\x3e\xa1\x67\x2f\x0b\x6b\xae\xb4\x7d\x42\xdf\x62\x6c\x6c\x95\x3b\xaf\xac\x57\xb1\x27\x68\xf8\x4e\x63\x95\x67\x44\x3a\x29\x1a\x14\xfb\xfc\x96\xcc\xa4\x93\x7b\xec\xf3\x5a\x69\x5c\x31\xc6\x30\x0a\xe6\xf6\x8a\x79\xd7\x52\xe9\x24\x5b\x97\x8f\xf4\xbe\xbc\xa7\xdf\x8b\xdf\xb4\x28\x8b\x07\xba\x20\xe4\xc0\x4b\x2b\xdc\x25\xa9\x4c\x6d\xcf\x25\x07\x0a\xdf\x86\x1d\x93\x27\xc0\x0e\xd6\xde\xcd\xe1\x76\x9a\x1c\x82\x4d\x5e\xe0\x05\x21\x6c\x0e\x0b\xa6\x4a\x61\xb3\x2e\x1f\x37\x57\x60\x1f\xec\x8e\xd0\x3f\xb0\x4b\x01\xd6\x27\x73\x67\xa2\x08\x34\xd1\x86\x7f\x60\x86\xc7\xd3\x5f\xbd\xbc\xc0\x5f\xd1\xc4\x9f\x1b\xba\x20\xcf\x79\xe4\x19\x21\xb3\x51\xaf\x82\xe4\x2a\x1e\x31\x90\xed\x78\x18\xf1\xc6\xaf\x5e\xf9\x02\x0a\xbf\xc6\xd9\xc9\x50\x47\xcd\xd7\xbc\x97\x52\x92\x19\xaf\x2a\x15\x95\x35\x5c\x83\x7b\xee\x64\x6c\x78\x84\x96\xf7\xb0\x43\x48\x01\xeb\xa4\xc9\x16\x87\xda\xbe\xc9\x7b\x28\xf6\x84\xb8\x47\xc5\x2b\xd1\x62\x17\xd1\x54\x50\x27\x23\x8e\xef\x55\xec\xc1\xd6\x80\x9d\x0a\x51\x19\x79\x1a\x27\xdb\xa3\x55\xa7\xd3\x9b\x69\x94\x3a\x4d\x98\xc5\x8b\xe6\xff\xda\x7e\xf6\x55\x3d\x25\xf3\x37\x00\x00\xff\xff\x95\xed\x6a\xb6\xa0\x06\x00\x00"),
		},
		"/etc/yum.repos.d/kubernetes-new.repo": &vfsgen۰FileInfo{
			name:    "kubernetes-new.repo",
			modTime: time.Date(2018, 11, 29, 5, 20, 9, 48242578, time.UTC),
			content: []byte("\x5b\x6b\x75\x62\x65\x72\x6e\x65\x74\x65\x73\x2d\x6e\x65\x77\x5d\x0a\x6e\x61\x6d\x65\x3d\x4b\x75\x62\x65\x72\x6e\x65\x74\x65\x73\x2d\x6e\x65\x77\x0a\x62\x61\x73\x65\x75\x72\x6c\x3d\x66\x69\x6c\x65\x3a\x2f\x2f\x2f\x76\x61\x72\x2f\x6c\x6f\x67\x2f\x72\x70\x6d\x73\x2f\x31\x2e\x31\x31\x2e\x32\x0a\x65\x6e\x61\x62\x6c\x65\x64\x3d\x31\x0a\x67\x70\x67\x63\x68\x65\x63\x6b\x3d\x30\x0a\x72\x65\x70\x6f\x5f\x67\x70\x67\x63\x68\x65\x63\x6b\x3d\x30"),
		},
		"/etc/yum.repos.d/kubernetes-old.repo": &vfsgen۰FileInfo{
			name:    "kubernetes-old.repo",
			modTime: time.Date(2018, 11, 29, 5, 19, 28, 282839863, time.UTC),
			content: []byte("\x5b\x6b\x75\x62\x65\x72\x6e\x65\x74\x65\x73\x2d\x6f\x6c\x64\x5d\x0a\x6e\x61\x6d\x65\x3d\x4b\x75\x62\x65\x72\x6e\x65\x74\x65\x73\x2d\x6f\x6c\x64\x0a\x62\x61\x73\x65\x75\x72\x6c\x3d\x66\x69\x6c\x65\x3a\x2f\x2f\x2f\x76\x61\x72\x2f\x6c\x6f\x67\x2f\x72\x70\x6d\x73\x2f\x31\x2e\x31\x30\x2e\x36\x0a\x65\x6e\x61\x62\x6c\x65\x64\x3d\x31\x0a\x67\x70\x67\x63\x68\x65\x63\x6b\x3d\x30\x0a\x72\x65\x70\x6f\x5f\x67\x70\x67\x63\x68\x65\x63\x6b\x3d\x30"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/etc"].(os.FileInfo),
	}
	fs["/etc"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/etc/docker"].(os.FileInfo),
		fs["/etc/nginx"].(os.FileInfo),
		fs["/etc/yum.repos.d"].(os.FileInfo),
	}
	fs["/etc/docker"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/etc/docker/daemon.json"].(os.FileInfo),
	}
	fs["/etc/nginx"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/etc/nginx/nginx.conf"].(os.FileInfo),
	}
	fs["/etc/yum.repos.d"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/etc/yum.repos.d/centos7.repo"].(os.FileInfo),
		fs["/etc/yum.repos.d/kubernetes-new.repo"].(os.FileInfo),
		fs["/etc/yum.repos.d/kubernetes-old.repo"].(os.FileInfo),
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
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{
			vfsgen۰FileInfo: f,
			Reader:          bytes.NewReader(f.content),
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

// vfsgen۰FileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type vfsgen۰FileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {}

func (f *vfsgen۰FileInfo) Name() string       { return f.name }
func (f *vfsgen۰FileInfo) Size() int64        { return int64(len(f.content)) }
func (f *vfsgen۰FileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰FileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰FileInfo) IsDir() bool        { return false }
func (f *vfsgen۰FileInfo) Sys() interface{}   { return nil }

// vfsgen۰File is an opened file instance.
type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	return nil
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
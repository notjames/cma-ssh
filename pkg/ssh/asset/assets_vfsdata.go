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
			modTime: time.Date(2018, 12, 1, 0, 36, 47, 585438738, time.UTC),
		},
		"/.DS_Store": &vfsgen۰CompressedFileInfo{
			name:             ".DS_Store",
			modTime:          time.Date(2018, 12, 2, 2, 58, 51, 687837542, time.UTC),
			uncompressedSize: 8196,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x4f\x6c\x14\xd5\x1f\xff\x7e\xb6\x05\xde\x0c\x3f\x7e\x7d\x94\x7f\x1d\xcb\x9f\xc5\x42\x05\xad\xb0\xdb\x16\x90\xf0\xc7\xed\x76\x09\x10\x44\x90\x6d\x69\xa1\x40\x3b\xbb\xf3\x64\x07\x66\x67\xea\xce\xec\x2e\x50\x5a\x7b\x53\x51\x2f\x26\x82\xa2\x07\x13\xf4\xa2\x26\x6a\x0c\x31\xd1\x70\xf1\x22\x9e\xb8\x90\x00\xf1\x62\x3c\x78\x21\xea\xc1\x0b\x47\x33\x33\x6f\xe9\x52\x29\x9e\x0c\x04\xe7\x93\x34\x9f\x99\xf7\x79\x7f\xe6\xbd\xf7\xe9\xf7\xe5\xfb\x96\x88\x90\x2e\x1b\x49\xa2\x16\x22\x62\x14\x32\x9f\x4b\xf7\x05\x93\x7f\x7f\x43\x4c\xb2\xe2\xf7\x47\xc4\xc9\x1d\xcd\x59\x4e\x8e\x88\x7e\xcc\x8d\xde\xbf\xaf\x08\x11\x22\x44\x88\x10\x21\xc2\x43\x06\x42\x62\x33\x9c\xfb\x11\x22\x44\xf8\x0f\xc3\x8f\x0f\x71\xc9\x29\xc9\x93\x21\x43\xea\x31\xc9\x8d\x75\x6d\xb8\xe4\xb8\xe4\x94\xe4\xc9\x90\x21\xeb\xc5\x24\x37\x4a\x66\x92\xb9\xe4\xb8\xe4\x94\xe4\xc9\x90\x65\xd0\x82\x4c\x3e\x20\x47\x86\xcc\x50\xc0\x25\xc7\x25\xa7\xfe\x9d\xb5\x89\x10\xe1\x71\x80\x9f\xbb\x37\x90\x20\x8f\xf2\xb9\x6a\x7d\xfe\x6e\x99\xae\x97\x48\xdc\x40\xac\xa1\x71\xd6\xec\x39\x8c\xb1\xb9\x6c\x1e\x3b\x9a\x2d\x38\xd5\xac\xa7\x7b\x65\x37\xad\x97\x86\xfc\xb7\xfd\xba\x57\xc8\xc9\xe7\x3e\xc7\xb1\xee\x3e\xeb\xb9\x83\xa6\xa8\x0e\xf3\x85\xbd\x8e\xed\xe9\xa6\x2d\x4a\x41\x63\xd3\x10\x39\xbd\x74\x64\xc0\xb4\x0d\xa7\x9a\x76\xca\xb6\xe1\x0e\xd5\x09\x8a\xa2\x30\x65\x98\xb7\x8c\x8d\x25\x37\x6c\xec\x88\x77\x75\x76\x8f\x77\xc4\xc7\x36\x6d\x4a\x74\xc4\xbb\xbb\x36\x8e\x8f\x2b\x6c\xc9\xea\xe4\xb6\xdd\x23\xc5\xd3\x67\xc6\xce\x8e\x4f\xbc\x1f\xce\xa2\x16\x15\xa8\x69\xda\xf4\x2e\x4e\x4d\xcf\x70\xf3\x56\xce\x71\x2c\x4c\x15\x59\xc7\x93\xd9\xbc\x53\x94\x77\x15\x07\xa8\x4e\x71\x2b\xbd\xe1\x5a\xc4\x8a\xb5\xb5\xb8\x59\x5b\x0b\x45\x55\xff\xb7\xa3\x59\xdd\x39\x68\xe6\x1d\x3b\x6b\x9e\x11\xc3\xbc\xc9\x2d\x38\xd5\xdd\x79\xc7\xde\x5f\x12\x95\x60\xda\xf3\xf3\xba\x95\x2f\x5b\xba\x27\x7a\x2c\xcb\xaf\xe4\x0e\xe4\x1d\xab\x5c\xb4\xdd\x41\x4f\x9c\xf2\xfc\x92\xc3\xae\x53\xf2\x7a\x83\xc2\x61\xce\xcb\xae\x38\x20\x2c\xdd\x33\x2b\x22\xa3\x7b\xc2\x1d\xe6\xcd\x7e\x4f\xfb\x46\x3d\xd3\xb1\xdd\x83\xa2\xe4\x9a\x8e\xdd\x96\x4a\xc8\x7d\x53\x3e\x9b\xb7\xb8\x35\xbe\x6a\xed\xba\xae\xe7\xb6\xf6\x5c\xff\x7f\x13\x9f\xdf\xbc\x40\x55\x0f\x9b\x86\xb0\x3d\xf3\x65\x53\x94\xfa\xab\xa6\xe1\x15\x0e\xe9\x6e\x5e\xd8\x86\x69\x1f\x1f\xa8\x98\xae\x99\xb3\x44\x9f\xad\x17\xc5\x7c\x74\x28\x4a\xd0\x6a\x49\x8b\xa6\x0d\x96\x73\xe6\x2b\x65\xd3\x3b\xcd\xdb\x18\x0b\x4a\x97\x2e\xd3\xd4\x23\x86\xee\x89\xbd\x8e\xe1\xf7\x66\xf0\xcb\x2c\xac\xbf\x72\x99\xa6\x0d\xf9\x4a\x6f\x49\xe8\x9e\x30\x64\x83\xd5\xed\x9a\xda\xe7\x9a\x67\x04\xd7\x65\xc5\xa7\x9f\x51\xd5\xbe\x93\xa6\x6d\x70\x57\x0e\xb5\x3e\xa1\x6a\xfd\x96\x9e\x13\x16\x37\x94\xb0\x59\xf7\x06\x55\x1b\xa8\x84\x73\xe3\x7b\x64\xe1\xe6\x05\xaa\x36\x98\x77\x8a\x45\x61\x7b\xae\x2c\xdb\xb6\x5d\xd3\x8e\xf9\xc3\xbe\xa0\xbb\xde\xbe\x51\x61\x0b\x83\x5f\x95\x63\xa7\x97\x69\xda\x21\x5f\xeb\x31\x0c\xff\x83\xda\x52\x6b\xe4\x22\x71\xff\x58\xd1\xe8\x49\xea\xa6\x5d\xb4\x9f\x0e\x91\x41\x55\x7a\x8b\xde\xa5\xf3\x74\x81\x2e\xd1\xe7\xf4\x0d\x5d\xa1\x1f\xe8\x1a\x5d\xa7\x1b\x74\x93\x6e\xd1\x2f\x74\x9b\x7e\xa3\xdf\xe9\x0f\xba\x83\x59\x98\x03\x06\x05\xcd\x58\x8e\x15\x88\x63\x2d\xd6\x21\x81\x24\x3a\xb1\x05\x29\xa4\xd1\x8b\x0c\xf6\x22\x8b\x7e\x1c\xc4\x00\x46\x50\xc0\x09\x9c\x84\x85\x32\x26\xf0\x2a\x26\xf1\x06\xde\xc3\x45\x7c\x80\x0f\xf1\x09\xbe\xc0\x97\xf8\x0a\xdf\xe2\xbb\xf0\xc3\x62\x35\xbf\xee\xba\xd7\xae\xb8\x72\x8f\x05\xe5\xbf\x63\xec\xc0\xfd\x2d\xb8\x48\x5d\xf3\x70\x2d\x78\x6b\x5e\xe8\xbb\x85\x8b\x16\xaf\x58\xb5\xf6\xd9\xce\x4d\x5b\x7b\xee\xee\xdb\xb4\xcd\xba\xc7\x4c\xf5\xfe\x09\x4c\x13\xba\x22\x70\x4b\xcd\x0d\x81\x4d\xaf\x2f\x69\xd1\x9e\x68\x5d\xaa\x2e\xaf\x99\x77\xba\xa9\xfb\x4d\xdb\x10\xa7\x98\x6f\x67\x3e\x27\xac\xbd\xb2\xb5\x8d\xf1\xab\x8c\xb3\xe0\x55\x6d\x6f\x5d\xa3\xf0\xcb\x8c\x23\x54\xdb\x5b\x3b\x18\xe3\xb1\x50\x5b\xdf\x9a\x54\xb8\xce\x78\x43\xa8\x75\xab\x1b\x19\x37\x14\x3e\x2b\x54\x37\xab\x5b\x14\xee\x2a\xbc\x31\x54\xb7\xab\x29\xc6\xf7\x28\x7c\x76\xa8\x2e\x55\x33\x8a\xc2\x69\xca\x6b\x0f\x70\xda\xa7\xf4\x35\x7d\x4f\xd7\xe8\x27\xfa\x99\x7e\xa5\xdb\xf4\x27\xdd\x01\xa0\xa0\x09\x5a\xe0\xac\x36\xac\x42\x3b\xd6\x23\x81\x4e\x74\x61\x03\xb6\xe3\x79\xf4\x20\x1d\xf8\xeb\x45\xec\xc3\x4b\x38\x8c\x21\x1c\xc5\x31\x8c\xc0\xc4\x09\x58\x28\xc2\xc1\x29\x9c\xc6\x18\xce\x62\x02\xaf\xe1\x75\x9c\xc3\x9b\x78\x1b\xe7\x71\x21\x70\xdd\xa5\x69\x3e\xdb\x39\xcd\x67\x1f\x4f\xf9\xac\xe8\x64\x32\x46\xd9\xcb\x13\xdd\x7c\xea\xdc\x3b\x13\x6e\xbd\x62\xcc\xa0\x8c\x16\xea\x03\xe7\x47\x75\x81\xb3\x92\x2d\xd9\x96\x63\x1f\x9f\x4a\x2f\x23\x3c\xaa\x68\x08\xa9\xc5\xcf\xff\x77\xcc\x7c\xff\x1f\x21\x42\x84\xc7\x18\x68\xcc\x64\x33\xe9\x07\x44\xec\x98\xbc\x08\x18\xa9\x35\xf8\x87\x8b\x00\xd4\xfd\x60\xf8\xc8\x5d\x04\x44\xe7\x7f\x74\xfe\x47\xa0\xbf\x02\x00\x00\xff\xff\x60\x6c\x02\xb3\x04\x20\x00\x00"),
		},
		"/etc": &vfsgen۰DirInfo{
			name:    "etc",
			modTime: time.Date(2018, 12, 4, 21, 45, 24, 298726378, time.UTC),
		},
		"/etc/.DS_Store": &vfsgen۰CompressedFileInfo{
			name:             ".DS_Store",
			modTime:          time.Date(2018, 12, 1, 0, 47, 20, 774453573, time.UTC),
			uncompressedSize: 10244,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5d\x6c\x14\xd5\x17\x3f\x67\x29\xe5\xce\x14\xe8\x2d\x9f\xed\x00\xdb\xfd\xa7\xfc\xa1\x48\xb3\x6c\xe5\x1b\x02\x6e\xdb\x85\xf0\x51\x02\xb1\x2d\x6d\x69\xa1\xcc\xee\xbd\x74\x07\x66\x67\x96\x9d\xd9\xdd\x96\x5a\x43\xd4\x00\x31\x46\xd0\x37\x89\x89\x0f\xc6\x44\x02\x26\xbe\x40\x62\x78\x41\x5f\x44\xc5\x04\x49\x13\x40\xa3\x3e\xe8\x8b\x8f\x46\xa2\x89\x4f\x66\x66\xee\xc2\xee\x76\x5b\x12\x13\x85\xe0\xfc\x92\xe6\xb7\x73\xcf\xb9\xf7\xcc\x3d\xf7\x77\x67\xf6\xf4\x2e\x00\x60\x7b\x96\xb5\x02\x84\x00\x80\x80\xc7\x74\x36\x54\x04\x11\x7f\x93\x10\x10\xbc\xd0\x19\x0f\x80\x02\x70\xc8\xe8\xc3\xad\x5d\x09\x33\x95\xae\x3c\x96\x0f\x1f\x3e\x7c\xf8\xf0\xe1\xe3\x09\x03\x3d\x22\x35\x4f\xfa\x46\x7c\xf8\xf0\xf1\xd4\x01\x45\x61\xe0\x70\x54\xf0\x69\x8f\x51\xd8\x03\x82\xab\x8a\xfa\x50\xc1\x21\xc1\x51\xc1\xa7\x3d\x46\xe1\x17\x10\x5c\x25\x98\x08\xa6\x82\x43\x82\xa3\x82\x4f\x7b\x2c\x1e\x5a\x28\x8a\x0f\x14\x91\x51\x54\x28\x48\x05\x87\x04\x47\xff\x99\xdc\xf8\xf0\xf1\x2c\xc0\xa9\xdd\xab\x81\x81\x09\x09\x38\x51\xa1\x7e\xff\xbe\xdc\x9e\x32\x63\x31\x96\xb5\x13\x00\xf7\x56\x1e\x0e\x5f\x8f\x4d\xb6\xb3\x69\xed\xe9\x64\xf1\xf8\xee\x6e\x9d\x09\x06\x0c\x83\x06\x06\x8c\x94\x46\x97\xbe\x2c\xb5\x96\xc6\xfe\xe4\x7a\xb9\x95\x4d\x63\xad\x10\x97\x80\x05\xa3\x60\x41\x02\x6c\xd0\x21\x0c\x2c\x9e\xb7\xd2\x71\xdd\x8c\x03\xc0\x17\xf1\xb4\xae\x59\x76\x24\x72\x17\x03\x33\xaa\x66\x56\xcf\x22\x84\xd4\x90\x39\xe4\x70\x57\xd2\xcc\x77\xd9\xaa\x9d\xb5\xda\xd5\xcc\x80\x73\x75\x40\xb5\x93\x71\xf1\xb9\xdb\x34\xf5\x87\x9f\xd5\xf8\x41\x8d\xe7\x87\xe8\x82\x0e\xd3\xb0\x55\xcd\xe0\x19\xb7\xb3\xc6\x78\x5c\xcd\x0c\xf6\x6a\x06\x33\xf3\xed\x66\xd6\x60\xd6\x40\x91\x41\x92\x24\x22\x0d\xd1\xfa\xb1\xb1\x75\x9b\x37\xb4\x84\x5a\x9f\x8f\x8c\xb7\x84\xc6\x36\x6e\x8c\xb4\x84\xd6\xad\xdd\x30\x3e\x2e\x91\xc5\xff\x6f\xdd\xb6\xfb\x68\x6a\xf4\xd4\xd8\x4b\xe3\x2f\xbf\xe3\x4d\xa8\xf0\x5c\x84\xda\xb2\x05\xbe\x58\x69\xa2\x15\xfe\x49\x33\xd9\xc7\xca\x75\x78\xc9\x08\xbc\x5d\x48\xc6\xbd\x42\x32\x24\x59\x9e\xbd\x6b\xf7\x1e\x79\x88\xce\xcb\x69\x3c\xbf\x3f\x6d\x6b\xa6\x61\x1d\xe4\x19\x4b\x33\x8d\x21\x5a\x6b\x25\xcd\xfc\xee\x84\x69\x1c\xc8\xf0\x9c\x9b\x83\xba\x84\xaa\x27\xb2\xba\x6a\xf3\x36\x5d\xef\xd2\x4e\x71\xab\x37\x61\xea\xd9\x94\x61\xf5\xd9\x7c\xc4\x76\x5a\x0e\x59\x66\xc6\xee\x70\x1b\xfb\xb4\x84\x69\x38\x6d\x43\x94\x66\x2d\xfe\x22\xd7\x55\x5b\xcb\xf1\x98\x6a\x73\x8b\xa2\x24\x5d\x9e\xb3\x28\xd8\xb4\xb2\xa5\x75\xc3\xd6\x17\x62\x13\x73\x6b\x69\xdd\xbc\xf9\xb2\x7c\x48\x63\xdc\xb0\xb5\x63\x1a\xcf\xf4\xe4\x35\x66\x27\xfb\x55\x2b\xc1\x0d\xa6\x19\xc3\xbd\x39\xcd\xd2\xe2\x3a\xef\x36\xd4\x14\xaf\xc3\x16\x49\x9a\x58\x5c\xdf\x30\x57\x59\xa2\x2c\x2b\x98\xca\xbb\x10\xda\x44\xfa\xb2\x71\xed\x64\x56\xb3\x47\xdd\x18\x8d\x21\x45\x1e\x64\xaa\xcd\xf7\x99\xcc\x89\xc2\xe8\x55\x22\xb9\x96\xe5\x21\x45\x19\x70\x2c\x1d\x19\xae\xda\x9c\x11\x32\x51\x57\x4b\xe7\xca\xab\x94\xd5\x12\x55\x49\xb7\xa5\x9d\xe2\x5e\x4b\x58\x8e\x48\xd4\x92\xba\x4f\x68\x06\x73\x5b\x94\xb5\xf2\x7a\x42\x99\xd4\xa3\xab\x71\xae\x7b\x4d\x9b\xe4\x2d\x84\xee\x95\x7a\x73\x5e\x3a\xbd\xc6\xf9\xf2\x76\x22\xf5\x25\xcc\x54\x8a\x1b\xb6\xe5\x86\x8d\xb6\x29\xca\x11\x27\x6c\xa7\x6a\xd9\xfb\xd3\xdc\xe0\x8c\xde\x24\xc4\xb5\xed\x08\x29\x4a\xbf\x63\x6b\x63\xcc\xb9\xa1\xa6\x68\xb3\xb7\xce\x6e\x0a\x9a\xa2\x11\xb1\xc3\x80\x40\x03\x84\x21\x0a\xdd\x30\x08\x1c\xd2\x30\x0a\x6f\xc0\x9b\x70\x1e\x2e\xc0\xbb\xf0\x01\x7c\x0c\xd7\xe0\x06\xdc\x84\xaf\xe1\x36\x7c\x03\x77\xe0\x3b\xf8\x09\x7e\x81\xdf\xe0\x01\xfc\x0e\x7f\x20\xa0\x84\x8b\xb0\x1e\x1b\x50\xc1\x26\x5c\x83\x11\x6c\xc5\x2d\xb8\x15\xb7\xe1\x76\xec\xc0\x4e\xdc\x87\xfb\xf1\x00\xf6\xe0\x11\x1c\x42\x15\xe3\x98\xc4\x93\x98\x41\x1b\xb3\x38\x8e\x67\xf0\x2c\x9e\xc3\xb7\xf0\x22\x5e\xc2\xcb\x78\x05\x3f\xc2\x6b\x78\x03\x3f\xc5\xcf\xf0\x2b\xbc\x8d\xf7\xbd\xdb\x0b\x14\x54\xdd\x59\x2a\x6a\xfc\x76\x0a\xc5\x8a\xed\x1b\xe8\xaf\xa8\xd8\xd8\x8e\x9d\x4f\x4a\xb1\xf7\xe7\x78\x32\x5d\xb0\x70\x51\xe3\xf2\x55\xe1\xb5\x9b\xb6\xb5\x3d\x5c\xcc\xb2\x35\x2c\x56\x92\x2b\x1d\x4f\x1b\xae\x66\x0a\x9a\x70\x57\xb1\x44\x8c\x8e\xa0\x95\x25\x4b\xe5\x60\x8f\x66\x30\x3e\x32\xd5\x0e\xa0\xb3\x1c\xf9\x93\x09\xa5\xbe\x61\x71\xf0\x7f\xc1\x26\x42\x6f\x12\x2a\x2e\x57\x04\x9b\x09\xbd\x4a\x68\xc0\x1d\xeb\xb9\xd5\x41\x99\xce\xa0\x2a\x71\xf7\x8a\xb2\x26\x22\x07\xe9\x4c\xca\x24\xe2\x5e\xae\x5b\x2f\xcb\xb4\x8a\x5a\xde\x4e\x52\x36\x6f\x91\x83\xb4\x9a\xee\x15\x23\xcb\x4b\xe5\xa8\x24\x51\xf0\x2e\x56\x04\x25\xe9\xef\x08\xf0\x7d\xb8\x02\xd7\xe1\x73\xb8\x05\x77\xe0\x2e\xfc\x08\x3f\xc3\xaf\xf0\x27\x02\x56\x23\xc5\x7a\x54\x70\x19\x06\xb1\x11\x9b\x71\x15\xae\xc6\x16\x0c\xe3\x46\xdc\x24\xe4\xb7\x13\x77\xe1\x1e\xdc\x8b\x9d\xd8\x83\xbd\xd8\x8f\x87\x70\x00\x19\x1e\xc3\x24\x6a\x78\x1c\x2d\xcc\x62\x1e\x47\x70\x14\x5f\xc1\x57\xf1\x35\x3c\x83\xe7\xf1\x82\x2b\xc5\xf7\xf0\xc3\x32\xe1\xed\x2a\x13\xde\xa5\x4a\xc2\x2b\x7e\x2f\xbd\x7e\xf6\xf8\x60\x65\x1f\xf6\x58\x9f\xd2\x77\x54\xe5\xc7\x72\xae\x2b\x63\xe8\xa6\x31\x2c\xbe\x8f\xd7\xc0\x28\x64\x21\x05\x61\xc8\xb8\xd9\x33\xc1\x72\xbc\x98\x95\xd0\xe3\xa6\xa9\x4f\xe3\x52\xfa\x0e\x68\x9c\x35\xb5\x67\xf1\xf4\xce\xfd\xf0\xe0\xd6\x74\x9e\x8f\x26\x79\x6e\xc5\xf2\x65\x53\x7b\x96\x4e\xd5\xff\x92\xee\xc3\x87\x0f\x1f\x3e\x7c\xfc\x5b\x98\xe1\x51\xa8\x06\x00\x76\x4c\x7d\xfe\xef\xc3\x87\x8f\x67\x18\x58\x15\xeb\x8a\xb5\x3f\x3a\x10\x9c\x84\x80\x38\x08\x38\x5a\xe8\xf0\x98\x83\x80\x80\xf7\x83\xa1\x66\xd1\x16\x81\x47\xfd\x9e\x8e\xc3\x00\xbf\xfe\xf7\xeb\xff\xff\x78\xfd\xff\x57\x00\x00\x00\xff\xff\x40\x4e\x28\x41\x04\x28\x00\x00"),
		},
		"/etc/docker": &vfsgen۰DirInfo{
			name:    "docker",
			modTime: time.Date(2018, 12, 4, 21, 46, 45, 219764659, time.UTC),
		},
		"/etc/docker/daemon.json": &vfsgen۰CompressedFileInfo{
			name:             "daemon.json",
			modTime:          time.Date(2018, 12, 4, 21, 46, 45, 218998851, time.UTC),
			uncompressedSize: 239,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\x52\x50\x50\xca\xcc\x2b\x4e\x4d\x2e\x2d\x4a\xd5\x2d\x4a\x4d\xcf\x2c\x2e\x29\xca\x4c\x2d\x56\xb2\x52\x88\xe6\x52\x50\x50\x50\x50\x4a\xc9\x4f\xce\x4e\x2d\x52\xd2\x41\xe6\xe9\x65\xe6\xc3\x04\xa0\x5a\x2a\x75\x0d\xf5\x30\xe4\xd2\x93\x91\x79\xd9\x16\xc5\x7a\xa8\x22\x85\xa5\x89\x95\x48\xdc\xea\x6a\xbd\x80\xa2\xfc\x8a\x4a\xcf\x82\xda\x5a\x2b\x4b\x13\x03\x43\x5c\x12\x46\xb8\x24\x8c\x71\x49\x98\x28\x71\x29\x28\xc4\x72\xd5\x02\x02\x00\x00\xff\xff\xc3\x6d\x8a\xf6\xef\x00\x00\x00"),
		},
		"/etc/nginx": &vfsgen۰DirInfo{
			name:    "nginx",
			modTime: time.Date(2018, 12, 4, 21, 48, 41, 776811351, time.UTC),
		},
		"/etc/nginx/nginx.conf": &vfsgen۰CompressedFileInfo{
			name:             "nginx.conf",
			modTime:          time.Date(2018, 12, 4, 21, 48, 41, 776043735, time.UTC),
			uncompressedSize: 2505,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x92\xcf\x8e\x9b\x30\x10\x87\xef\x7e\x8a\x39\xe4\x1a\x48\xb2\xa8\xca\x2e\xf7\xb6\x7b\x5b\xf5\xd4\x9b\xe5\xc5\x03\xb1\x00\x0f\x9d\x31\xf9\xd3\x88\x77\xaf\x4c\x48\xb5\xd9\x36\x3d\xe5\xd2\x36\x48\x48\x63\x7e\xdf\xcc\x58\xe2\xeb\x05\x19\xc0\x57\xce\xef\x73\xb5\x23\xae\x91\x75\xc7\x54\xa0\x08\x0a\x80\xe9\x03\xe5\x4a\x21\x33\xb1\x6e\xa8\x02\x48\xb7\x86\xd3\x86\xaa\x74\xec\x49\xc7\x24\x89\xc9\xce\xb0\xcf\x55\xe7\x2c\x4c\xcf\x48\x72\xef\x4f\x64\xd2\x39\x1b\x27\x6d\xd1\x07\x81\xa3\x02\x98\xb6\x15\xe4\x3d\x16\xc1\x91\x17\x80\xe5\x62\x95\xe5\x6a\x50\x6a\x13\x42\x37\x52\x82\xbc\x45\x1e\xcb\xf3\x41\x7b\xd3\x22\x30\x56\x4e\x02\x1f\xe6\xcb\xc4\x52\x51\x23\x27\x8e\xf2\x91\x6a\x9c\x04\xf4\xb0\x5e\x4c\x47\x2a\x4c\x1c\x0f\xe9\x34\x05\xa0\x63\xda\x1f\x74\x67\x44\x20\x2e\x7a\x4a\xd3\xe3\x31\x79\x89\x1f\x9f\xbb\x61\x78\x7a\xcc\x16\xcb\xfc\x02\x65\xb4\x8e\xb1\x08\x40\x65\x79\x99\x08\x06\xbd\x41\x63\x91\xe1\x33\x49\x80\xd9\x86\x24\x5c\x45\xbe\xce\xbf\xa0\x69\xe6\xcf\x2f\x30\x63\x6c\x29\xa0\x36\xd6\xf2\x1f\xf0\x8f\xc4\x3b\xc3\x16\x6d\xac\x60\x76\x22\x8c\xb5\x7a\xaf\xcb\x73\x14\xab\xf3\x88\xa2\x71\xe8\x83\x6e\xcd\x5e\xbf\x92\x3d\x68\x71\xdf\x11\x96\x8b\x4f\xef\xf2\x31\x7b\xed\xcb\x12\x79\x42\x56\xeb\xfa\xf2\x1a\xd3\x7f\xd1\xc1\xb5\x48\x7d\x80\xc7\xc5\xfb\x6b\x7a\x7b\x35\x64\x34\xd7\xc3\xb7\x7b\xb3\xfa\x77\x99\x40\x06\x0f\xab\x5f\x22\xf9\x99\x9f\x9a\x3f\xc4\x6e\x75\x01\x05\x6c\x3b\x5d\xba\x06\xf5\x8e\x5d\xc0\x37\x5c\x64\x06\x75\x7a\xaf\x4a\x55\x15\xb7\xd5\x68\x75\xd7\xe8\x2f\xd1\xe8\x86\x12\xd5\x6b\x49\x6e\x2d\xd2\xc3\x5d\xa4\xff\x4f\xa4\x6f\xbd\x39\xdc\xd4\xa2\xec\x6e\xd1\x3f\x6a\xd1\xf0\x23\x00\x00\xff\xff\x44\x5d\xa6\xe7\xc9\x09\x00\x00"),
		},
		"/etc/sysctl.d": &vfsgen۰DirInfo{
			name:    "sysctl.d",
			modTime: time.Date(2018, 12, 1, 0, 47, 46, 409823562, time.UTC),
		},
		"/etc/sysctl.d/k8s.conf": &vfsgen۰CompressedFileInfo{
			name:             "k8s.conf",
			modTime:          time.Date(2018, 12, 1, 0, 47, 46, 408992927, time.UTC),
			uncompressedSize: 78,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x4b\x2d\xd1\x4b\x2a\xca\x4c\x49\x4f\x85\x52\xba\x79\x69\xba\xc9\x89\x39\x39\xba\x99\x05\x66\x25\x89\x49\x39\xa9\xc5\x0a\xb6\x0a\x86\x5c\xf8\xd4\x21\x94\x01\x02\x00\x00\xff\xff\x7a\x4e\x4e\x73\x4e\x00\x00\x00"),
		},
		"/etc/yum.repos.d": &vfsgen۰DirInfo{
			name:    "yum.repos.d",
			modTime: time.Date(2018, 12, 3, 6, 35, 15, 445849337, time.UTC),
		},
		"/etc/yum.repos.d/.DS_Store": &vfsgen۰CompressedFileInfo{
			name:             ".DS_Store",
			modTime:          time.Date(2018, 12, 1, 0, 36, 54, 141570163, time.UTC),
			uncompressedSize: 6148,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\xd8\x31\x0a\x02\x31\x10\x85\xe1\x37\x31\x45\xc0\x26\xa5\x65\x1a\x0f\xe0\x0d\xc2\xb2\x9e\xc0\x0b\x58\x78\x05\xfb\x1c\x5d\x96\x79\x60\x60\xd5\x4e\x8c\xcb\xfb\x40\xfe\x05\x37\x2a\x16\x31\x23\x00\x9b\xee\xb7\x13\x90\x01\x24\x78\x71\xc4\x4b\x89\x8f\x95\xd0\x5d\x1b\x5f\x43\x44\x44\x44\xc6\x66\x9e\xb4\xff\xf5\x07\x11\x91\xe1\x2c\xfb\x43\x61\x2b\xdb\xbc\xc6\xe7\x03\x1b\xbb\x35\x99\x2d\x6c\x65\x9b\xd7\x78\x5f\x60\x23\x9b\xd8\xcc\x16\xb6\xb2\xcd\xcb\x4d\xcb\x38\x7c\x18\xdf\xd9\x38\xa1\x18\xa7\x10\x2b\x6c\xfd\xce\x77\x23\xf2\xef\x76\x9e\xbc\xfc\xfe\x9f\xdf\xcf\xff\x22\xb2\x61\x16\xe7\xcb\x3c\x3d\x07\x82\xf5\x0d\x00\xae\xdd\xf5\xa7\x43\x40\xf0\x3f\x0b\x0f\xdd\x5a\x1d\x04\x44\x06\xf3\x08\x00\x00\xff\xff\x6a\x00\x88\x6d\x04\x18\x00\x00"),
		},
		"/etc/yum.repos.d/bootstrap.repo": &vfsgen۰CompressedFileInfo{
			name:             "bootstrap.repo",
			modTime:          time.Date(2018, 12, 3, 6, 35, 15, 443816102, time.UTC),
			uncompressedSize: 169,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xae\xae\xd6\x73\xca\xcf\x2f\x29\x2e\x29\x4a\x2c\xf0\x2c\xa8\xad\x8d\x47\x16\x08\xc8\x2f\x2a\xa9\xad\x8d\xe5\xca\x4b\xcc\x4d\xb5\x4d\x4c\x49\x49\x4d\x51\x48\x2b\xca\xcf\xb5\x52\xc8\x28\x29\x29\xb0\xd2\xd7\x47\xd7\x6c\x85\xa9\x99\x2b\x29\xb1\x38\xb5\xb4\x28\xc7\x96\x04\x2d\xa9\x79\x89\x49\x39\xa9\x29\xb6\x86\x5c\xe9\x05\xe9\xc9\x19\xa9\xc9\xd9\xb6\x06\x80\x00\x00\x00\xff\xff\x5b\x55\x87\x86\xa9\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/.DS_Store"].(os.FileInfo),
		fs["/etc"].(os.FileInfo),
	}
	fs["/etc"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/etc/.DS_Store"].(os.FileInfo),
		fs["/etc/docker"].(os.FileInfo),
		fs["/etc/nginx"].(os.FileInfo),
		fs["/etc/sysctl.d"].(os.FileInfo),
		fs["/etc/yum.repos.d"].(os.FileInfo),
	}
	fs["/etc/docker"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/etc/docker/daemon.json"].(os.FileInfo),
	}
	fs["/etc/nginx"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/etc/nginx/nginx.conf"].(os.FileInfo),
	}
	fs["/etc/sysctl.d"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/etc/sysctl.d/k8s.conf"].(os.FileInfo),
	}
	fs["/etc/yum.repos.d"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/etc/yum.repos.d/.DS_Store"].(os.FileInfo),
		fs["/etc/yum.repos.d/bootstrap.repo"].(os.FileInfo),
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
			gr: gr,
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

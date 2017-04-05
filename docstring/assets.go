// Code generated by go-bindata.
// sources:
// docstring/async.md
// docstring/channel.md
// docstring/concat.md
// docstring/cond.md
// docstring/do.md
// docstring/filter.md
// docstring/future.md
// docstring/if.md
// docstring/map.md
// docstring/promise.md
// docstring/quote.md
// docstring/repl-cls.md
// docstring/repl-doc.md
// docstring/repl-help.md
// docstring/repl-quit.md
// docstring/repl-use.md
// DO NOT EDIT!

package docstring

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _asyncMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\x41\x6e\x2b\x21\x10\x05\xf7\x9c\xe2\xc9\xde\x18\x7d\x7d\xdf\x21\x0b\x5f\x20\x27\x70\x8b\x79\xe3\x41\xc2\xcd\x18\x1a\xdb\x73\xfb\x08\x26\x91\x92\x6d\xd3\x55\xd5\xe2\x88\x93\xd4\x4d\x03\xe6\x5c\xee\xff\x3c\x6e\x54\x16\x31\x56\x88\x62\xbc\x2c\x25\x6b\x6e\x15\x95\x8f\x46\x0d\x74\x97\xa7\xa4\x36\x56\x6c\x21\xea\xca\x10\xe7\xc8\x69\x18\x2a\xa2\x42\x50\xb9\x4a\xb7\xc0\x96\x42\x99\x90\x67\xf0\xcd\xd0\x2c\x66\x3d\xe3\x93\xd6\x8a\xd6\xb1\xb7\x3b\x61\x8b\x18\x5e\x31\x25\x44\x1b\x7d\xe4\x27\x0b\x44\xb7\xce\xf6\x4e\x8f\x8e\xa4\x18\xa4\x10\xbc\x47\x33\x4e\x68\x35\xea\x0d\x82\x94\x83\xa4\xb4\xa1\x86\xbc\xf6\x63\x9a\x86\x5e\xfb\xc1\xfb\x71\xb8\x9e\x3a\xb5\xab\xfc\xf5\xec\xdc\xf1\x88\x0f\xc5\xe5\x2d\xf7\x35\xd1\x39\xe0\x34\x71\x86\x54\x3e\xbe\xbf\xc5\x01\xc0\x4e\x1d\x0a\xa7\x83\xff\x3d\xc8\x45\xf4\xc6\xbf\xb3\x8d\x29\xe5\xd7\xc1\x7b\x3f\x74\x96\xff\x3f\x19\x2c\x97\x21\xf5\xee\x2b\x00\x00\xff\xff\x2a\xec\x27\x95\x6f\x01\x00\x00")

func asyncMdBytes() ([]byte, error) {
	return bindataRead(
		_asyncMd,
		"async.md",
	)
}

func asyncMd() (*asset, error) {
	bytes, err := asyncMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "async.md", size: 367, mode: os.FileMode(420), modTime: time.Unix(1491404447, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _channelMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x92\x3d\x6f\x1b\x3d\x10\x84\x7b\xfe\x8a\x81\x5d\xbc\x77\x07\x5b\xc5\x5b\xaa\x09\x8c\xc0\x45\x90\x2e\x48\x17\x04\xb8\x3d\xde\xd2\x47\x84\x47\xca\xe4\x9e\xfc\xf1\xeb\x83\x25\xf5\x11\xa9\x92\x96\xb3\x9c\x67\x86\xba\x47\x67\x17\x8a\x91\x03\x8a\xff\xe4\x2f\x3d\x6c\x66\x12\x2e\x20\x6c\xd1\xcf\x3e\xb3\x15\x9f\x22\x05\x9c\x74\xe6\xe9\xfc\x0d\x5e\x55\x33\x09\xa1\x48\xde\xac\x6c\x99\x21\x0b\x89\x1e\x6c\x85\x67\x48\xc2\x0b\x47\xce\x24\x0c\x42\xa0\xcf\x0f\x14\x7e\xdd\x38\x5a\x46\x72\x38\x52\xd8\xb8\xec\xf0\xcd\x81\xaa\xbd\x2e\x96\x03\x5b\xef\x3c\xcf\x0f\x90\x85\x2f\x5e\x6f\x3e\x04\x4c\x8c\x69\x73\x8e\x33\xcf\x3b\xfc\x5c\x18\x99\xcb\x16\xa4\x81\x2c\x54\x96\xc7\x95\x0e\xb0\x29\x16\x5f\xc4\xc7\x17\x35\xa1\x88\x91\x57\x2f\x23\xdc\x16\x6b\x96\x07\x10\x46\x1b\x52\xe1\x9b\x59\x9c\x95\xe2\x84\xb7\xc3\x0f\x96\xec\xf9\xa8\xb7\x50\x04\x07\x5e\x39\x0a\x5c\x4e\x6b\xe5\xba\xe4\xa8\x60\xc3\x14\x92\xfd\x33\x3c\xe0\x8d\x7c\x35\x76\x29\x57\x59\xe4\x77\x69\x39\xb5\x8c\x89\xa1\x28\xc2\x33\x52\xbe\x68\xce\x11\x9b\xa0\x82\xcd\x3b\x3c\xab\xb0\xba\x5f\xf7\xe9\xb6\x0e\x0a\x25\xa1\x3a\xc3\xbb\x7a\x55\x6b\x47\xab\x88\xff\x09\x26\xe6\x08\x17\xb6\xb2\xf0\x0c\xd2\x8e\x4e\x7d\x25\x07\x2f\xfa\x2c\xb5\xa3\x23\xe7\x9b\x48\x3b\x63\xee\xef\xf1\xf5\xe4\xf4\x9d\x3f\x8a\x31\xc3\xbe\xf0\xeb\x00\xfd\xdc\xa4\x6f\xcc\xe7\x47\x9e\xcd\xb0\xd7\x80\x55\xa8\xad\xd5\xb0\xf9\xd2\xb2\x1a\xeb\xba\x4b\x79\xc5\xd8\xe9\x71\x0b\xd7\x8f\x66\xd8\xd7\xe4\xba\x4a\xd7\x05\x49\xad\x90\x9b\xa2\xc6\xae\xce\xfa\xb1\x82\x3e\x45\x3c\xbf\xd3\x7a\x08\x6c\x0c\xd0\x05\x16\xfc\xb2\xd7\xbf\xf5\xff\x7d\xc5\x40\x57\xc9\x60\xfb\xd3\x85\x5d\xf3\x83\xed\x7f\x1b\xc5\x6d\x34\x77\x2e\xa5\xbb\xfe\xdf\xc1\x44\xf9\x3c\x68\xae\xa6\xfd\x90\xf4\x78\x64\x2b\x29\xa3\xd3\x6e\x60\xfb\xbe\x37\x7f\x03\x00\x00\xff\xff\x5c\xd5\x61\xd6\x4f\x03\x00\x00")

func channelMdBytes() ([]byte, error) {
	return bindataRead(
		_channelMd,
		"channel.md",
	)
}

func channelMd() (*asset, error) {
	bytes, err := channelMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "channel.md", size: 847, mode: os.FileMode(420), modTime: time.Unix(1491404450, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _concatMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\x4d\x6a\xc3\x30\x10\x85\xf7\x73\x8a\x07\x5e\xd4\xa6\x50\xe8\xef\xbe\x94\xde\xa0\xbb\x12\xb0\x90\x27\x91\x40\x91\x1c\x8d\x64\xc7\x39\x7d\x90\x0c\x22\xdb\xf7\xbd\xbf\x0e\xbd\x0e\x5e\xab\x04\xe1\xcb\xf3\x00\xa7\x6e\xd6\x6d\xd8\x35\xf6\x2a\xb1\x14\x92\xd9\x6b\x16\xfa\x89\x5c\x15\x55\x7c\x5b\x03\x58\x4d\x10\x2e\xa1\xc4\x3e\xc1\x0a\x92\x61\x44\x96\xec\x12\xc2\xf1\xa1\xcd\xfa\x53\x65\x8b\x72\x99\xa5\x30\x56\xda\x60\x8e\x61\xb1\x13\x4f\xad\xf1\x85\xa8\xeb\xf0\xed\xf1\x7b\x55\xe7\xd9\x31\x11\xda\xd1\xff\x57\xbc\xe1\xfd\x80\xa7\xfe\x03\x9f\xf8\x1a\x06\xa2\x3f\x63\x05\xab\x75\x0e\x91\x53\x8e\xbe\x6e\xb4\x77\x63\x5f\x13\xd8\xed\x23\xdd\x03\x00\x00\xff\xff\xdb\xca\x7f\x93\xf4\x00\x00\x00")

func concatMdBytes() ([]byte, error) {
	return bindataRead(
		_concatMd,
		"concat.md",
	)
}

func concatMd() (*asset, error) {
	bytes, err := concatMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "concat.md", size: 244, mode: os.FileMode(420), modTime: time.Unix(1491378041, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _condMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8f\xc1\x6e\xf2\x30\x10\x84\xef\x7e\x8a\x11\x1c\x48\x7e\xa1\x5f\xe1\xc0\x81\x0a\x21\xf5\xd0\x3e\x07\x26\x99\x10\x4b\x66\x1d\xad\x97\x26\x7d\xfb\xca\x09\xaa\x7a\xdb\xb5\xbf\xf9\x34\xbb\x45\xd5\x26\xe9\x70\x1e\x95\x1d\x6c\xa0\x5c\xfe\xd5\x18\xa9\x7d\xd2\x47\x46\xf9\x0b\x16\x92\xf8\x88\x9b\x7a\x69\x87\x20\x77\xf7\x99\x14\xf4\xed\x80\xeb\x9f\xd8\xf5\x05\xec\xcb\x8a\xf2\x1e\x5a\x6f\xc4\x14\x62\xc4\x8d\xe0\x97\x8f\x4f\x6f\xec\xf6\xf0\xd2\x21\xf4\x08\x86\x90\x61\xfa\xb4\xe1\x1b\x95\x24\x43\xef\x63\xe6\x1e\x65\x94\x10\xeb\xc5\xb4\x2b\xf6\x1d\x4a\x9f\x82\xff\x6a\x16\x8b\xd2\x9e\x2a\xc5\x99\x6c\xa0\x4e\x21\x73\x09\x09\x67\x7b\xf5\x29\xa1\x51\x53\xcb\x9c\xd9\xfd\x77\x6e\xbb\xc5\xbb\xe0\x63\xf6\x8f\x31\xd2\x39\xa0\xea\xd8\x63\xc6\xe9\x54\x2f\x5b\x39\xd9\x01\x40\x75\xc6\x8c\x63\x53\x03\x9b\xc9\x67\x44\xe6\x0c\x1b\xbc\xe0\xd8\x6c\x56\xe0\x82\x19\x87\xa6\xa9\x57\xe0\xae\xf4\x46\x5d\x99\x43\xf3\x82\xde\x18\x33\xcb\xb0\x42\x41\x70\xa3\x4d\xa4\x6c\x6a\xf7\x13\x00\x00\xff\xff\xdd\x52\x28\x8e\x7d\x01\x00\x00")

func condMdBytes() ([]byte, error) {
	return bindataRead(
		_condMd,
		"cond.md",
	)
}

func condMd() (*asset, error) {
	bytes, err := condMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cond.md", size: 381, mode: os.FileMode(420), modTime: time.Unix(1491408149, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _doMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xce\x3d\x8a\xc3\x30\x10\xc5\xf1\x5e\xa7\xf8\x83\x1b\x69\x77\x59\x76\xcb\x34\x81\x14\x39\x43\x6a\x91\x8c\xe3\x01\x79\x6c\xf4\x11\x74\xfc\x60\x07\x5c\xcd\x3c\x7e\xaf\x78\x03\xfe\xb1\x30\x2e\x79\xfe\x0a\xc8\x2b\xa6\x16\xab\x14\xe6\x96\xaa\xae\x49\x76\x29\xee\xa6\x29\x1d\x8a\xc4\xfb\xb4\x03\x6a\xd4\x96\xed\x87\x2c\xdb\x55\x7b\x52\x27\x61\x54\x8b\x47\x5f\x17\x23\x16\xb4\x16\xb2\x94\x96\xea\xaf\x73\xc3\xc0\xc5\xb8\xf6\x38\xaf\x49\x9c\x03\xaf\x23\xfe\x4c\xe7\x14\xb6\x41\x0e\xc0\xaf\xd9\xe8\xe1\xf3\x7f\xd3\xf9\xff\xdb\x42\x08\xee\x1d\x00\x00\xff\xff\x91\x6e\x1e\x77\xb5\x00\x00\x00")

func doMdBytes() ([]byte, error) {
	return bindataRead(
		_doMd,
		"do.md",
	)
}

func doMd() (*asset, error) {
	bytes, err := doMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "do.md", size: 181, mode: os.FileMode(420), modTime: time.Unix(1491410712, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filterMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xcb\x4e\xc3\x30\x10\x45\xf7\xfe\x8a\x2b\x65\x13\x0b\x54\xa9\x2d\x4b\x36\x08\xb1\x60\xcf\xae\xaa\x54\x93\x8e\xc9\x48\xc3\x38\xf8\xd1\x26\x7c\x3d\x72\x22\xc2\x63\x67\xcf\xbd\x3a\x67\x34\x0d\x5a\xcf\x92\x29\xc2\x17\xed\x90\xe8\xe3\xc6\x42\xdc\x27\xcb\x84\x25\x48\x75\x58\x48\x3b\x4a\xe6\x31\x92\xcb\x94\xe0\x6a\x65\x5a\x03\x5c\xfb\x90\x08\x5d\xd0\x4c\x9a\xc1\x09\xb9\x27\x44\x4a\x45\x32\x82\x87\x1b\x06\x99\x58\xdf\xe6\xf1\x10\xc3\x85\xcf\x74\x9e\x85\x99\x83\x22\x87\x39\xb8\x38\x29\x94\x6a\xff\x4f\x6d\xb5\x6f\xf0\xec\xff\x81\xeb\xaf\xc2\xb9\x73\x33\xa9\x9a\x63\xc9\xfd\x84\x56\x43\x86\x77\x92\xe8\x16\xf5\xa9\x2c\xb6\xd6\xf5\x47\x85\x2b\x8b\xe0\x95\xc0\xda\x49\xa9\x2a\xd6\x5f\xfc\xba\xef\xb7\x7b\x63\x4c\xd3\xe0\x41\xf1\x34\xba\xf7\x41\xc8\x18\xac\x77\x6b\xbd\xe2\x30\x1e\xd1\xde\x63\xc4\xde\x5a\x1c\xb6\xd8\x61\x8f\xbb\xa3\x35\xe6\xa5\xe7\xb4\x78\x22\xe5\x12\x17\xfe\x7a\xb6\x53\xbb\xc5\xce\x9e\xcc\x57\x00\x00\x00\xff\xff\x44\xaf\x37\x64\x85\x01\x00\x00")

func filterMdBytes() ([]byte, error) {
	return bindataRead(
		_filterMd,
		"filter.md",
	)
}

func filterMd() (*asset, error) {
	bytes, err := filterMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "filter.md", size: 389, mode: os.FileMode(420), modTime: time.Unix(1491378047, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _futureMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\x41\x4e\x03\x31\x0c\x45\xf7\x39\xc5\x57\xbb\x99\xa9\xa0\x77\x60\xc1\x05\x10\x17\x70\x13\x0f\x13\xe1\xda\x55\xe2\xb4\x9d\xdb\xa3\xc9\xa0\x56\x90\xe5\x73\xfe\x7f\x89\xf7\x18\xa6\xe6\xad\x30\x26\x2b\xe7\x11\x7c\x25\x69\xe4\x5c\x41\x9d\x80\xea\xa2\x71\x2e\xa6\xd6\xaa\x2c\xe1\x83\xbd\x15\xed\xd3\x2d\x96\x15\x3e\x6f\x69\xd8\xd4\xb9\x46\xcf\xa6\x47\x7c\xce\x8c\x4b\xb1\x6b\x4e\x9c\xb6\x0b\xb7\x2c\x82\x13\x3f\x2c\x69\x8d\x13\x2a\x5f\xa8\x90\x33\x7c\x2e\x4c\x69\xed\xe1\x3b\xc7\xb6\xd6\xbc\x80\x34\x81\x74\x41\x24\x91\x0a\xb7\xcd\xf7\x6b\xd9\x2a\x0f\x27\xb1\xf8\x7d\x40\x53\xcf\xf2\x7c\xcf\x4c\x15\x27\x66\x45\xb4\xf3\x45\xd8\x59\x96\xa7\xfa\x18\xc2\x7e\x8f\x37\xc5\xfb\x9d\xd6\x69\x08\xc0\x90\x78\x5a\x3f\xf6\x58\xca\x90\x2c\x00\xc0\xe0\xf6\x7a\xe5\xe8\x56\x30\xf4\x8d\x74\x0a\x0c\x7c\xce\x8e\x5d\xe1\xb4\x1b\xff\x22\x2b\xa4\x5f\xfc\x9f\x2e\x2c\x62\xb7\xdd\xd8\x4f\x37\x4e\xcd\xc7\xf0\x13\x00\x00\xff\xff\xa2\x72\x30\x2c\x86\x01\x00\x00")

func futureMdBytes() ([]byte, error) {
	return bindataRead(
		_futureMd,
		"future.md",
	)
}

func futureMd() (*asset, error) {
	bytes, err := futureMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "future.md", size: 390, mode: os.FileMode(420), modTime: time.Unix(1491378053, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ifMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\xbd\x4e\xc3\x30\x14\x46\x77\x3f\xc5\x51\x33\xc4\x96\x22\x24\x28\x7f\x13\x88\x81\x81\xc7\x70\x9a\xeb\xc6\x92\xe3\x44\xb6\x43\xd3\xb7\x47\x4e\x07\x58\xd8\x6c\xdd\xef\x9c\xfb\xdd\x06\xed\x1d\x4b\x92\x81\x32\x4a\x44\x42\x96\x77\xc3\x22\xc9\xcd\x69\xca\x64\x3f\x2d\x41\xe8\x93\x8d\xa7\xd1\xc7\xb3\xfa\x72\x35\x88\x7c\xdb\xb0\xda\x22\xc3\xce\xfa\x93\x2d\x82\xcf\x94\xb4\x96\xf1\x8a\x8e\x73\xc1\xd9\x90\xa5\xa3\x3e\xa3\x0f\xa6\xdb\xb9\xb6\x6e\x69\xa9\xf2\x9a\xff\xd5\xd8\x38\x90\xa4\xac\x29\xca\xd0\x31\x97\x51\xd2\xc5\x67\xb9\x41\xb5\xd5\x0d\xea\xf0\x0e\x1b\xaf\x1d\x17\x1f\x02\xbd\xfc\xa3\xb8\x53\xaa\x69\xf8\x88\x7c\x6e\xb6\x5e\xa0\x14\xe8\x41\x1c\x1b\xad\xbe\xe7\x81\x23\x8f\x3c\xf1\xcc\x0b\xaf\xc6\xec\x53\xef\xd0\x6f\xe8\x20\x91\xcd\x70\x34\x0a\xe0\xb0\xd5\x9a\xbd\x3f\x1f\xfe\x7c\xf3\x64\x43\x38\x18\xf5\x13\x00\x00\xff\xff\x25\xb1\x7f\x05\x3d\x01\x00\x00")

func ifMdBytes() ([]byte, error) {
	return bindataRead(
		_ifMd,
		"if.md",
	)
}

func ifMd() (*asset, error) {
	bytes, err := ifMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "if.md", size: 317, mode: os.FileMode(420), modTime: time.Unix(1491408137, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mapMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xcd\x4a\xc4\x40\x10\x84\xef\xfd\x14\x05\xb9\x64\x14\x04\xe3\x22\x5e\x45\x7c\x03\x6f\xcb\xc2\x36\xbb\x3d\x66\xa0\x33\x33\xce\xcf\x26\xf1\xe9\x25\xa3\x04\xbc\xd6\xd7\x5d\xf5\x75\xe8\x27\x8e\xb0\xd5\x5f\x90\xe5\xeb\xde\x40\xf9\xdb\xe9\x8a\x89\x63\xde\x92\x2a\xfe\x22\x99\xde\x92\x70\x91\x0c\xde\xf8\xba\x03\xcc\x63\xc8\x82\x1b\x6b\xdd\x60\x12\x94\x51\x90\x24\x57\x2d\x08\x16\x1c\xa3\xae\xce\x7f\xb6\x38\xa6\x70\x73\x57\xb9\xb6\xb5\xe2\x82\x47\x09\x0d\xfc\xbd\x07\xfb\xff\x6c\x5f\x7f\x20\xea\x3a\xbc\x7a\xbc\x2f\x3c\x45\x15\x22\xfc\x6a\xf7\xd6\xe3\xb8\x9c\xd0\xdf\x61\xc1\x60\x0c\x8e\x8f\x18\xf0\x84\xc3\xc9\x10\x7d\x8c\x2e\x63\x76\xaa\x48\x52\x6a\xf2\xad\x7c\x17\x3f\xf7\x03\x0e\x78\xc6\x8b\x39\xd3\x4f\x00\x00\x00\xff\xff\x2d\x7f\x03\x26\x05\x01\x00\x00")

func mapMdBytes() ([]byte, error) {
	return bindataRead(
		_mapMd,
		"map.md",
	)
}

func mapMd() (*asset, error) {
	bytes, err := mapMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "map.md", size: 261, mode: os.FileMode(420), modTime: time.Unix(1491378064, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _promiseMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\x4f\x6e\x2a\x31\x0c\xc6\xf7\x39\xc5\x27\xd8\x00\x42\xdc\x81\xc5\x5b\xbc\x6d\x6f\x60\x12\x0f\x63\xd5\x63\x47\xc1\x61\xda\xdb\x57\x51\xdb\x41\x5d\x26\xdf\x1f\xff\xec\x3d\x0e\xb5\xf9\x22\x0f\x3e\xa2\x36\x2f\x3d\xf3\x03\x84\xec\x16\x62\x9d\x42\xdc\xb0\xb8\x51\x49\x6f\x1c\xbd\xd9\x10\x7f\x02\x10\x43\xcc\x8c\xc9\xdb\x02\x9f\x40\x98\xba\xe5\x91\xb8\xe0\xff\x04\xaa\x55\x85\x0b\x56\x89\xd9\x7b\x80\x0c\xd4\xee\x7d\x61\x8b\x33\x62\x96\xc7\x66\xc7\x2a\xaa\x38\xdd\xd4\xf3\xfb\xe9\x8c\x95\x24\xc4\xee\xa3\x17\x84\x27\x69\x67\x84\xe3\xc6\x28\xac\xf2\xe4\xc6\x05\x4f\xa1\x41\x49\xaa\x43\x1a\x14\xbf\x50\x5b\x69\xcc\x14\x10\xcb\xda\xcb\x58\xe9\x35\xfd\x82\xeb\xe6\xce\x64\x70\xd3\xcf\xbf\xed\x6e\x99\x2f\x29\xed\xf7\xb8\x1a\xfe\x7d\xd0\x52\x95\x53\x02\x0e\x85\x27\xd4\xd7\xc5\x8e\xe3\xaf\x62\x37\xb3\xaa\xef\xbe\x5f\xc7\xf4\x15\x00\x00\xff\xff\x76\xa7\x10\xe9\x53\x01\x00\x00")

func promiseMdBytes() ([]byte, error) {
	return bindataRead(
		_promiseMd,
		"promise.md",
	)
}

func promiseMd() (*asset, error) {
	bytes, err := promiseMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "promise.md", size: 339, mode: os.FileMode(420), modTime: time.Unix(1491378071, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _quoteMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8f\x3b\x6e\xc3\x40\x0c\x44\xfb\x3d\xc5\x00\x2e\x6c\x03\x81\x01\x27\xb9\x40\x8a\x94\xe9\x72\x01\x5a\x1a\x45\x04\x76\xb9\x1b\x2e\xe5\xcf\xed\x03\xc9\x69\x49\x60\xde\x7b\x3b\x1c\x7e\x97\x1a\xc4\x54\xbd\x1c\xe1\x8c\xc5\xad\x23\x66\xa2\x37\x0e\x3a\x29\xc7\xed\x07\x35\x8c\x12\x82\x52\x47\xa6\x2f\x8a\xa9\xfd\x20\x66\x09\x64\xed\xd1\x21\x36\xa2\x3f\xca\xa5\xe6\x8e\x9b\xe6\x0c\xab\x81\x0b\xc1\xab\xe4\x45\x82\xe3\x09\xf8\x9e\xb5\xa3\xc8\xe0\x15\xda\xc1\x69\xe2\x10\x7a\x65\x7e\x3c\x81\x52\x08\xe9\x68\xce\x46\x1b\xd7\x7d\x31\xf0\xde\x9c\xbd\x6b\x35\xdc\x34\xe6\xf5\x24\xad\xf6\xf0\xda\x66\xe2\xb0\x3f\x9e\x52\xda\xed\xf0\x61\xf8\xbc\x4b\x69\x99\x29\x01\xfb\xc3\x19\xaf\x78\xc3\xfb\x31\xa5\x0d\xba\x19\x3d\xeb\x36\x56\xd6\xa0\x4b\xde\xdc\xe1\x12\x33\x7d\x8d\x31\x84\x3f\xb6\xb0\x0a\x69\xed\x5f\xcc\x96\x72\xa1\xe3\xfc\xb2\xda\xe9\x04\x0d\xdc\xe8\x84\x60\x5a\x6c\x08\xad\x76\x4a\x7f\x01\x00\x00\xff\xff\x85\xb4\xeb\x86\x49\x01\x00\x00")

func quoteMdBytes() ([]byte, error) {
	return bindataRead(
		_quoteMd,
		"quote.md",
	)
}

func quoteMd() (*asset, error) {
	bytes, err := quoteMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "quote.md", size: 329, mode: os.FileMode(420), modTime: time.Unix(1491410540, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _replClsMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xd0\x48\xce\x29\xd6\x54\x48\xce\x49\x4d\x2c\x2a\x56\x28\xc9\x48\x55\x28\x4e\x2e\x4a\x4d\xcd\xe3\x72\x06\x89\x80\x05\x92\x4b\x8b\x8a\x52\xf3\x4a\x14\x92\xf3\xf3\x4a\x40\x74\x7e\x1a\x92\x3a\x3d\x85\x90\x8c\xcc\x62\x85\xcc\x62\x85\x44\x85\x20\xd7\x00\x1f\xdd\xfc\xbc\x9c\x4a\x85\xb4\xd2\xbc\xe4\x92\xcc\xfc\x3c\x3d\x2e\x40\x00\x00\x00\xff\xff\x25\x68\x02\x02\x61\x00\x00\x00")

func replClsMdBytes() ([]byte, error) {
	return bindataRead(
		_replClsMd,
		"repl-cls.md",
	)
}

func replClsMd() (*asset, error) {
	bytes, err := replClsMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "repl-cls.md", size: 97, mode: os.FileMode(420), modTime: time.Unix(1491378075, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _replDocMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x41\x0a\xc2\x30\x10\x45\xf7\x73\x8a\x0f\xdd\x58\xd0\xde\x41\xb0\x3b\x17\x22\x5e\x60\x48\x26\x74\xa0\x9d\x04\x27\x85\xe6\xf6\xd2\x8a\x0b\xb7\xef\xf3\x1f\xaf\xc3\x29\xe6\x80\x94\xdf\x4b\x8f\xa8\x5e\x66\x6e\x8e\x98\xc3\xba\x88\x55\xae\x9a\x8d\x6e\x3f\x5c\x27\xf9\x9f\xf6\xdf\x41\xbd\x48\xd0\xa4\x12\x0f\xd3\x19\x9a\xc0\xd6\x20\x9b\x7a\xf5\x01\xaf\x49\x1d\xea\x60\x3c\xc7\xc7\xfd\x92\x6d\x6e\x48\xab\x85\xdd\x31\x10\x75\x1d\xae\x86\x71\xe3\xa5\xcc\x42\x84\x6f\x13\x7b\xb3\xd0\xd3\x27\x00\x00\xff\xff\x14\x0b\x79\x2c\xa2\x00\x00\x00")

func replDocMdBytes() ([]byte, error) {
	return bindataRead(
		_replDocMd,
		"repl-doc.md",
	)
}

func replDocMd() (*asset, error) {
	bytes, err := replDocMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "repl-doc.md", size: 162, mode: os.FileMode(420), modTime: time.Unix(1491378078, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _replHelpMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\x4b\x8e\xdb\x30\x10\x44\xf7\x3a\x45\x2d\x65\x21\xf0\x01\xb2\x0b\x1c\xef\xb2\x48\x82\x1c\xc0\x14\x55\xb2\x08\x34\x9b\x0c\x9b\xf4\xe7\xf6\x03\xc9\x33\xc6\x18\xd3\xdb\x7a\xaf\xba\x86\x7e\xa1\xe4\xdd\x80\xc7\xfd\x0c\x96\xc5\xdd\x51\x97\x60\x58\x13\x44\x9a\xb9\x33\xbb\xa1\x9f\x92\xc7\x9c\x4a\xdc\xe0\x0f\x70\x4a\xbe\x45\x6a\x75\x35\x24\xdd\xe3\x78\x73\x31\x0b\xbf\xe3\xb4\xe1\xce\xee\xea\x77\xa7\x6e\xe8\x9b\x11\x6a\xef\x7f\x0e\x8b\xd3\x33\xe1\x5b\x29\xd4\x0a\x75\x91\x96\x9d\xe7\x8b\xbf\x1a\x73\x4a\x9b\xed\xc5\x9e\x13\x71\x10\xba\x82\xba\x10\xe6\x0b\xa9\xdd\xd0\xff\x6f\xa1\x3e\x81\x3f\x2d\xd4\x2d\xfe\x7b\xfc\xfd\xab\xeb\xfe\x2d\x44\x2e\x29\xe6\x8a\x91\x92\xae\x08\x13\xb5\x86\x39\xd0\x36\xea\xcb\x8a\x6f\x98\x93\x48\xba\x72\xc2\x78\xdf\x10\x68\x8b\x23\x0b\xd2\x0c\xde\x72\xa1\x59\x48\xba\xda\xae\xe2\x1a\x44\xb0\xb8\x0b\x31\x92\x0a\x5e\x9c\x34\x57\x39\x21\xe8\x4b\xfd\x3a\x06\xf6\x50\xf7\xc0\x0f\x58\xd0\xb3\xf0\x53\x21\xbc\x53\x58\x76\x8a\xd8\xa4\x86\x2c\x84\x04\xa5\xed\xbb\xb7\x00\x00\x00\xff\xff\x68\x18\x87\x3a\xa5\x01\x00\x00")

func replHelpMdBytes() ([]byte, error) {
	return bindataRead(
		_replHelpMd,
		"repl-help.md",
	)
}

func replHelpMd() (*asset, error) {
	bytes, err := replHelpMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "repl-help.md", size: 421, mode: os.FileMode(420), modTime: time.Unix(1491391829, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _replQuitMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xd0\x28\x2c\xcd\x2c\xd1\x54\x00\x91\xc5\x0a\x25\x19\xa9\x0a\x41\xae\x01\x3e\x5c\x9e\xb9\xb9\xa9\x29\x99\x89\x25\xa9\x39\x95\x0a\xa9\x15\x30\xa9\xe0\x82\xd2\x92\x92\xd4\x22\xb0\x12\x3d\x85\x90\x8c\xcc\x62\x85\xcc\x62\x85\x44\x30\x5f\x37\x3f\x2f\xa7\x52\x21\xad\x34\x2f\xb9\x24\x33\x3f\x4f\x8f\x0b\x10\x00\x00\xff\xff\xe9\x14\xcc\xb5\x5a\x00\x00\x00")

func replQuitMdBytes() ([]byte, error) {
	return bindataRead(
		_replQuitMd,
		"repl-quit.md",
	)
}

func replQuitMd() (*asset, error) {
	bytes, err := replQuitMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "repl-quit.md", size: 90, mode: os.FileMode(420), modTime: time.Unix(1491327304, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _replUseMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcd\x41\x0a\xc2\x40\x0c\x46\xe1\x7d\x4e\xf1\xc3\x6c\xec\xc2\xde\x41\xa4\x3b\x17\x22\x5e\x20\x84\xd4\x19\x68\x93\xd2\x64\xa0\xde\xde\x85\x82\xdb\xc7\x83\xaf\xe0\xd4\x43\x61\x31\x40\x2a\xdb\x4b\x03\xc6\xab\xc6\xc6\xa2\x74\xfd\x95\xac\x0a\xe9\xfb\xae\x96\x10\xb7\xd4\x23\xff\xd7\x88\x67\x6d\x81\x16\x60\x3c\xa6\xfb\xed\xec\xb6\xbc\x31\x77\x93\x6c\x6e\x23\x51\x29\xb8\x18\xa6\x83\xd7\x6d\x51\x22\x7c\xc5\xd9\x7d\xa0\x4f\x00\x00\x00\xff\xff\xc5\x9a\xf9\x8e\x7e\x00\x00\x00")

func replUseMdBytes() ([]byte, error) {
	return bindataRead(
		_replUseMd,
		"repl-use.md",
	)
}

func replUseMd() (*asset, error) {
	bytes, err := replUseMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "repl-use.md", size: 126, mode: os.FileMode(420), modTime: time.Unix(1491327304, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"async.md": asyncMd,
	"channel.md": channelMd,
	"concat.md": concatMd,
	"cond.md": condMd,
	"do.md": doMd,
	"filter.md": filterMd,
	"future.md": futureMd,
	"if.md": ifMd,
	"map.md": mapMd,
	"promise.md": promiseMd,
	"quote.md": quoteMd,
	"repl-cls.md": replClsMd,
	"repl-doc.md": replDocMd,
	"repl-help.md": replHelpMd,
	"repl-quit.md": replQuitMd,
	"repl-use.md": replUseMd,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"async.md": &bintree{asyncMd, map[string]*bintree{}},
	"channel.md": &bintree{channelMd, map[string]*bintree{}},
	"concat.md": &bintree{concatMd, map[string]*bintree{}},
	"cond.md": &bintree{condMd, map[string]*bintree{}},
	"do.md": &bintree{doMd, map[string]*bintree{}},
	"filter.md": &bintree{filterMd, map[string]*bintree{}},
	"future.md": &bintree{futureMd, map[string]*bintree{}},
	"if.md": &bintree{ifMd, map[string]*bintree{}},
	"map.md": &bintree{mapMd, map[string]*bintree{}},
	"promise.md": &bintree{promiseMd, map[string]*bintree{}},
	"quote.md": &bintree{quoteMd, map[string]*bintree{}},
	"repl-cls.md": &bintree{replClsMd, map[string]*bintree{}},
	"repl-doc.md": &bintree{replDocMd, map[string]*bintree{}},
	"repl-help.md": &bintree{replHelpMd, map[string]*bintree{}},
	"repl-quit.md": &bintree{replQuitMd, map[string]*bintree{}},
	"repl-use.md": &bintree{replUseMd, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}


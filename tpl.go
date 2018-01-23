// Code generated by go-bindata.
// sources:
// tmpl/timeboard.tmpl
// DO NOT EDIT!

package main

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

var _tmplTimeboardTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x55\xdd\x6a\x1b\x3d\x10\xbd\xf7\x53\x0c\x8b\x2f\x63\x3f\xc0\x07\xbe\x08\x84\x2f\x14\xfa\x93\xa6\xc1\xbd\x28\xc5\xc8\xd6\xec\x5a\x54\xde\x55\xb4\x5a\x27\x8e\xd0\xbb\x17\x8d\xfe\xd6\xf6\x36\x14\xea\x2b\xcd\x39\x67\xa4\x33\xd2\x8c\x57\x63\xdf\x0d\x7a\x87\x50\x71\x66\x18\xef\x9a\x8d\x11\x07\xdc\x76\x4c\xf3\x0a\x2a\x6b\x61\xf9\x81\x83\x73\x15\xd8\x19\x80\x11\x46\x22\x84\xdf\x2a\xb0\x4f\x04\x39\x57\xcd\x00\x38\xf6\x3b\x2d\x94\x11\x5d\x9b\xe8\xbb\x11\x14\x44\x1a\x19\xdf\x74\xad\x3c\xd1\x1e\x5e\xf3\x88\x8c\x7f\xf1\x80\x73\x33\x00\x6b\x5f\x84\xd9\xc3\xf2\x5e\x33\xb5\xef\x61\x11\x41\xcd\xda\x06\x61\x19\x34\x8d\xe7\xc8\x51\xf2\x34\xe1\xa6\x6c\x75\x87\xb5\x68\x05\x79\x08\xdb\x01\x1c\xc5\x5b\xa9\x61\x2d\xde\x3c\x91\x92\x44\x0d\xcb\xdb\xc1\x74\xfd\x8e\x49\xf4\x04\xcb\x41\x4c\x28\xac\x73\x95\xb5\x0b\xc0\x96\xc7\x8d\x43\xfa\x83\xc6\x9d\xe8\xe3\x81\x2a\x07\x31\xbd\xb0\x7f\x48\xbf\xd7\xdd\xa0\xa8\xf6\xc6\xaf\x60\x05\x3f\xac\x9d\x37\x01\xfd\x6f\x95\x04\xce\xa5\x7b\x99\x8b\x96\xe3\xeb\x0d\xcc\x51\xe2\xe1\x42\x21\xea\x48\x3b\x77\x63\x2d\x1d\x55\x59\x4b\x4a\x5a\x11\xf2\x73\xca\xc6\xb7\x5d\xa7\x90\x6c\xf4\x7e\x15\x6d\xf4\x01\xf5\x87\x04\xc1\x7b\x36\x8a\xe2\x1f\x6c\x98\x53\x78\x07\xc2\x7a\x8a\xc2\xdb\x03\x28\x26\xd1\x18\x3c\x6b\x49\xd2\x2f\x1f\x22\x93\xde\x35\x6b\x37\xb5\x14\x6a\x52\xfb\xbf\x27\x92\xde\x5d\x5a\x59\x80\x37\xf3\x89\xe9\x5f\xa8\xe9\x52\x3c\x14\x1a\xec\x0c\x8c\x9d\x1a\xf3\x0e\x44\x65\xbf\xe6\xa4\xb0\xcc\x8e\x0f\x8a\xbf\x23\x93\x43\x6e\xb1\x35\x05\x85\xb4\x96\x8e\xff\xc8\xb6\x28\xfd\x39\x92\x16\x51\x1c\xd0\xab\x66\x2a\x25\x5c\x2f\x72\x51\xa1\x82\x47\x7c\x1e\xb0\x37\x93\x25\xe8\xc0\xe5\x1a\x9e\x47\x77\xfd\x75\x40\x7d\x2a\xc3\x93\x7d\x52\x69\x0b\xe7\xa8\xde\x8b\x72\xa3\x07\x88\xdb\xe7\x9c\xdb\xa6\xd1\xd8\x30\xd3\x69\x1a\xba\x12\xa5\xa9\x2b\xc8\x55\xa9\x79\x93\xdc\x2c\xa5\xb6\xf3\xfe\xb9\xec\xa0\x3c\xb2\xb1\x5f\xfc\xc0\xc6\x65\x1a\xd7\xd2\x49\x97\x87\xa6\xe4\xef\x82\x9b\xbd\x4f\x7d\xa1\x45\x4c\x0c\xe8\x3b\x69\x7f\x71\x4b\x39\xa7\xd4\x39\xf9\x94\xe5\x41\xcf\x9f\x77\x3a\x72\x71\xde\xc6\x58\xb9\xb0\x27\x3c\x28\xc9\x0c\xae\x99\x16\x6c\x2b\x31\x77\xfb\xe8\x4f\xd8\x7f\x0d\xa2\x6c\x73\x8c\xba\x78\xa7\x2d\x3b\xe0\xa8\x45\x3e\xfb\x30\x35\x88\xd2\x58\x8b\x57\x18\xfd\x13\xfa\x30\xb1\x1c\x6b\x36\x48\x53\x3e\x1e\x21\xf4\x9f\x9f\x49\xd3\x6e\xf6\x3b\x00\x00\xff\xff\x01\x64\x7e\x13\xbf\x06\x00\x00")

func tmplTimeboardTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplTimeboardTmpl,
		"tmpl/timeboard.tmpl",
	)
}

func tmplTimeboardTmpl() (*asset, error) {
	bytes, err := tmplTimeboardTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/timeboard.tmpl", size: 1727, mode: os.FileMode(420), modTime: time.Unix(1511796026, 0)}
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
	"tmpl/timeboard.tmpl": tmplTimeboardTmpl,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"timeboard.tmpl": &bintree{tmplTimeboardTmpl, map[string]*bintree{}},
	}},
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

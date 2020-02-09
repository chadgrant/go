// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// schema.json
package schema

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x92\xc1\x6e\xa3\x40\x10\x44\xef\x7c\x45\x6b\xec\xa3\x0d\x7b\x5b\x89\x3f\x58\x69\xa5\x48\x91\x72\x8a\x72\x98\x30\x85\x69\x0b\x33\xa4\xa7\xed\xc4\xb2\xf8\xf7\x68\x00\x13\x23\x92\x48\xb9\xe4\xda\xd5\xd5\xbc\x1a\xea\x92\x10\x11\x99\x35\x3b\x93\x93\xa9\x54\xdb\x3c\xcb\x42\x51\xe1\x60\x43\x1a\xd0\x28\xde\x52\xf6\xe3\xa4\xe6\xa0\xe9\x3e\xf8\xc6\x6c\x46\xdb\x30\xbf\xb1\x46\x75\x3b\x4c\x53\x2f\xbb\xcc\x89\x2d\x75\xfb\xe7\xef\x78\x61\x75\x75\x3a\x94\xdc\xb0\xb2\x6f\x82\xc9\x69\xa0\xe8\x05\xc1\x8e\x83\x42\xe0\xc6\xdb\x74\x2b\xf7\x2b\x7a\x6e\x11\xc7\xc6\x3f\xef\x51\xe8\x78\x72\x92\x5b\xf1\x2d\x44\x19\xf3\xcb\x93\x7e\x14\x5e\x1e\x9d\x54\x87\x50\x08\xb7\x91\x2c\xc6\x7a\xb8\xff\x47\xec\xd0\x28\x97\x0c\x21\x5f\xd2\x88\xb5\xf9\xdc\x3e\xb1\x05\x15\x6e\x76\x5f\xad\x95\x5e\x0e\x56\xe3\x07\x22\xcd\x62\xa7\x5b\xda\xcc\x51\xea\x9f\x60\xff\x8f\xac\xaf\x15\x04\xa4\x9e\x04\x2a\x8c\x13\x7e\x87\x3e\xf9\x26\x8b\x11\xbc\x1c\x59\x10\xfb\xf6\xd8\xfb\x37\x7d\xb4\xa7\x64\x6e\xef\xa6\xa2\xcc\x82\xc5\x0a\x7e\xfc\x85\x40\x5a\x81\x02\xe4\xc4\x05\xa8\xb2\x81\x6a\x6f\x1d\xdc\xb5\x65\x7d\x9e\x9c\x8c\x15\xb1\xe7\xeb\x90\x15\x87\x30\x7f\x4a\x63\xeb\xfa\xae\x8c\x44\x17\xb3\x16\x94\xfd\x1b\xac\xb2\x9b\x92\x66\xcb\x5e\x76\x03\x72\x97\x74\xef\x01\x00\x00\xff\xff\xd3\xf6\xc9\x68\x45\x03\x00\x00")

func schemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaJson,
		"schema.json",
	)
}

func schemaJson() (*asset, error) {
	bytes, err := schemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.json", size: 837, mode: os.FileMode(420), modTime: time.Unix(1581282343, 0)}
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
	"schema.json": schemaJson,
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
	"schema.json": &bintree{schemaJson, map[string]*bintree{}},
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

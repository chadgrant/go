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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x92\xc1\x6e\xa3\x40\x10\x44\xef\x7c\x45\x6b\xec\xa3\x0d\x7b\x5b\x89\x3f\x58\x69\xa5\x48\x91\x72\x8a\x72\x98\x30\x85\x69\x0b\x33\xa4\xa7\xed\xc4\xb2\xf8\xf7\x68\x00\x13\x23\x92\x48\xb9\xe4\xda\xd5\x55\xbc\x1e\xea\x92\x10\x11\x99\x35\x3b\x93\x93\xa9\x54\xdb\x3c\xcb\x42\x51\xe1\x60\x43\x1a\xd0\x28\xde\x52\xf6\x59\x80\x9c\xb8\xc0\xa8\xd4\x1c\x34\xdd\x07\xdf\x98\xcd\x68\x1f\xe6\x37\x11\x51\xdd\x0e\xd3\xd4\xcb\x2e\x73\x62\x4b\xdd\xfe\xf9\x3b\x26\xac\xae\x4e\x87\x92\x1b\x56\xf6\x4d\x30\x39\x0d\x34\xbd\x20\xd8\x71\x50\x08\xdc\x98\x4d\xb7\x72\xbf\xa2\xe7\x16\x71\x6c\xfc\xf3\x1e\x85\x8e\x91\x93\xdc\x8a\x6f\x21\xca\x98\x27\x4f\xfa\x51\x78\x19\x3a\xa9\x0e\xa1\x10\x6e\x23\x59\x3c\xeb\xe1\xfe\x1f\xb1\x43\xa3\x5c\x32\x84\x7c\x49\x23\xd6\xe6\x73\xfb\xc4\x16\x54\xb8\xd9\x7d\xb5\x56\x7a\x39\x58\x8d\x1f\x88\x34\x8b\x9d\x6e\x69\x33\x47\xa9\x7f\x82\xfd\x3f\xb2\xbe\x56\x10\x90\x7a\x12\xa8\x30\x4e\xf8\x1d\xfa\xe4\x9b\x5b\x8c\xe0\xe5\xc8\x82\xd8\xbb\xc7\xde\xbf\xe9\x4f\x7b\x4a\xe6\xf6\x6e\x2a\xca\xec\xb0\x58\xc1\x8f\xbf\x10\x48\x2b\xd0\x58\x52\xaa\x6c\xa0\xda\x5b\x07\x77\x6d\x59\x7f\x4f\x4e\xc6\x8a\xd8\xf3\x75\xc8\x8a\x43\x98\x3f\xa5\xb1\x75\x7d\x57\x46\xa2\x8b\x59\x0b\xca\xfe\x0d\x56\xd9\x4d\x49\xb3\x65\x2f\xbb\x01\xb9\x4b\xba\xf7\x00\x00\x00\xff\xff\x44\xa0\xff\xf2\x4d\x03\x00\x00")

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

	info := bindataFileInfo{name: "schema.json", size: 845, mode: os.FileMode(420), modTime: time.Unix(1581291520, 0)}
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

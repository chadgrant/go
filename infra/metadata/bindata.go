// Code generated for package metadata by go-bindata DO NOT EDIT. (@generated)
// sources:
// schema.json
package metadata

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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x95\x41\x6f\xdb\x3c\x0c\x86\xef\xfd\x15\x84\xbf\x1e\xe3\xa4\xc1\xb7\x61\x58\x6e\xc3\x4e\x3b\x6f\xc0\x80\xb5\x5d\x20\x5b\x4c\xac\x4e\x96\x34\x8a\x4e\x97\x06\xf9\xef\x83\x64\xbb\x49\x5a\xc5\xc5\xe0\xdc\x6c\x51\x7a\xde\x57\x12\x49\xed\xae\x00\x00\xb2\x6b\x25\xb3\x05\x64\x15\xb3\x5b\xcc\x66\xbe\xac\xb0\x16\x7e\xea\xd1\x30\xfe\x99\x2a\x3b\xf3\x48\x1b\x55\xe2\xac\x46\x16\x52\xb0\x98\x3e\x78\x6b\xb2\x49\xb7\xb8\x9d\x7f\x04\x08\xd1\xbc\x1d\x9d\x5a\x5a\xcf\x24\x89\x15\xe7\x37\x1f\x3a\xf2\x7f\xfd\x4a\x89\xbe\x24\xe5\x58\x59\x13\x56\x77\x2a\xd0\xab\xf4\xd3\x78\xeb\x30\xc4\x6d\xf1\x80\x25\xf7\xa3\x8e\xac\x43\x62\x85\x3e\x5b\x40\xbb\x8f\x38\xbe\x41\x23\x2d\x9d\x8c\x9d\x50\x3c\x93\x32\xeb\x8e\xf2\x1c\x7d\x61\xc5\x88\x1a\xc1\xae\x80\x2b\x84\x0e\xf8\x3c\x7f\x7f\x58\x9a\xad\xc9\x36\x6e\xa4\x56\x64\x00\xa1\x77\xd6\x78\x55\x68\x84\x95\xa5\xa8\xdc\x1d\x48\x5a\xba\x0f\x5e\x6e\xa3\x1d\x71\x02\xc6\x52\x2d\xb4\x7a\x42\x09\x8f\x8a\x2b\x30\x16\xbc\x13\x25\xfa\xb4\x93\x15\x29\x34\x52\x6f\x47\x5a\xe9\x31\x90\xf0\x94\x16\x3e\x05\x8c\xd1\xb6\xf1\x4b\x68\xf0\x95\x25\x86\xa3\x68\x30\x32\x68\xa2\x68\x94\x96\x4b\x42\x67\x2f\xe5\xa1\x21\x7d\xa4\x0a\x11\x3d\x20\x6d\x9a\xba\xc0\xb1\x09\xef\x55\x0d\x1b\x24\x88\x48\x68\x91\x13\x68\x7c\x23\xb4\xde\x82\x23\xbb\x51\x12\x25\x14\x5b\xf8\xfc\xa5\x9d\xe4\x5f\x12\x9d\x60\x46\x8a\xb4\x9f\xb7\x37\xf9\xc7\xfb\xdd\x7c\xf2\xff\xfe\xee\x6e\xba\x9b\xef\x4f\xfe\x0f\x3f\xd7\xe7\xf7\xc5\xcb\x62\x6c\x42\x7d\xab\x10\x1a\x8f\xa1\x96\x04\x83\x54\x32\x66\x54\x34\x3f\x74\x9e\xac\xea\x7f\xae\xaa\x55\xa8\x18\x0e\x71\x29\x18\xf3\x88\x18\x36\xf7\xbd\x42\x73\xf0\x03\x8f\xc2\x83\xb4\xe6\x4c\x8e\xad\x15\x2f\x2b\xe1\xab\x0b\x1c\xc8\x5a\x31\xf8\x4a\xcc\x21\xf0\xfa\x2a\x6b\xcf\x64\xe8\x3e\x3f\xe5\x3f\x44\xfe\x14\x2f\xee\xfd\xe4\xdd\xcd\xb9\x9b\x0b\x46\x0b\x12\xa6\xbc\x94\xd5\x16\x76\xea\x33\xa9\x5c\xda\xda\x29\x8d\xb4\xdc\x20\xf9\xf1\x0d\xa1\xc7\x41\x87\x0b\x89\x24\x81\x6d\x77\x5d\x6f\xb6\xa6\x37\x5d\x74\x55\x3b\xec\xe2\x6b\x53\x87\xad\x0b\xad\xa1\xac\xb0\xfc\x05\xb2\x21\x11\x82\xc7\x9d\xf8\xea\x48\x3b\x23\xfc\xdd\x28\xc2\xf0\x9c\xdf\xbe\x7a\x13\x5f\x3d\x5c\x89\xe7\x24\xd5\xd7\xcf\xb4\xdc\x74\x13\x4c\xd5\x71\xba\xc6\x52\x09\x9e\xce\xa5\xa1\x7b\x4e\x1c\xfa\xb9\x16\x19\x87\xef\xaf\xf6\x7f\x03\x00\x00\xff\xff\xd5\x81\x76\x7f\xf6\x08\x00\x00")

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

	info := bindataFileInfo{name: "schema.json", size: 2294, mode: os.FileMode(420), modTime: time.Unix(1581291747, 0)}
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

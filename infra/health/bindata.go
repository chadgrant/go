// Code generated for package health by go-bindata DO NOT EDIT. (@generated)
// sources:
// schema.json
package health

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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x55\x4d\x6f\xd3\x40\x10\xbd\xe7\x57\x8c\xb6\x3d\x80\xe4\xc6\x15\x12\x42\xea\x5f\x40\x82\x03\xdc\x68\x15\x4d\xec\x71\x3c\xad\x77\x37\xcc\x8e\xfb\x21\xd4\xff\x8e\x76\x9d\xb8\x76\x62\x9c\xc0\x8d\xeb\xce\xd7\x7b\xe3\xf7\xc6\xbf\x16\x00\x00\xe6\x92\x4b\x73\x03\xa6\x56\xdd\xde\xe4\x79\x28\x6a\xb2\x18\x96\x81\x9c\xd2\xf3\x92\x7d\x5e\x13\x36\x5a\x2f\xef\x83\x77\x26\xdb\x95\x74\x59\x83\xb2\x18\xbd\xea\x5e\x97\x5e\x36\x79\x29\x58\xe9\xd5\xf5\xa7\x5d\xbf\x8b\x7d\x65\x49\xa1\x10\xde\x2a\x7b\x97\xaa\x53\x6f\x28\x6a\x2a\x1e\x40\x68\xeb\x45\xc3\x5b\x6a\xc5\x8e\x63\x66\x30\x37\xd0\x81\x4d\x81\xae\x28\xd5\x18\x18\x46\x52\xd4\xa1\xa5\xe3\xe7\x14\xd2\x97\x2d\xc5\xa9\x41\x85\xdd\x66\x37\x67\x94\x71\x00\x2f\xf6\x02\x5f\x81\xd6\x04\x83\xa9\xa0\x14\x74\xaa\xdc\xb2\x63\xdb\xda\x38\xfe\xc3\x54\x18\x9f\xf7\xe1\x8f\xd7\xa3\xf0\xeb\x38\xdb\xb0\x53\x92\x47\x6c\x56\x36\xcc\x71\x81\x88\xb1\xb5\x6b\x92\x33\xc8\x70\x05\x6b\x2c\x1e\x36\xe2\x5b\x57\x26\x0a\x19\xd4\xfe\x09\x7c\xa5\xe4\x12\xc5\xf8\x06\x1c\x40\x3d\xac\x09\xa4\x75\x66\x16\x64\xd9\x0a\xc6\xe6\x27\x40\xfe\x0d\xc6\x2f\x29\x31\xae\xdc\x72\xd3\x70\xa0\xc2\xbb\x32\x80\xe2\x43\x84\xe8\x23\xa6\x1e\xe9\x3c\xb8\xa0\xa8\xed\x29\x5c\x67\x0b\xe1\x7b\x4d\x10\x3b\xf6\x6a\xe8\xd6\xc7\x15\x38\xaf\x70\x6b\xbe\x7e\xbe\x35\x59\x5c\x1d\x3b\xa8\x90\x1b\x2a\x73\x12\xf1\xd2\x15\x9d\x92\xca\x2c\x91\x38\x89\xca\x15\xea\xaa\xd5\xe2\x9f\xf9\x54\x5e\x2c\x6a\xcc\x29\x51\xe9\x4a\xd9\x4e\xa2\x9a\xa0\xdd\x60\x50\x88\xf9\x6f\x12\x79\xc2\x70\x5a\x1d\x42\x3f\x5b\x16\x2a\x23\xe4\x1f\x9d\x2b\xb3\xb1\x66\xb2\xfe\x2b\x65\x87\x34\xef\xfa\x5e\xaf\x8b\x41\xf3\x9e\xa8\x5f\xdf\x53\xb1\xb7\xa0\xd9\x8a\xdf\x92\x28\xd3\xc1\xa1\x68\xf8\x91\x1c\x85\xf1\xeb\xa8\x0f\x8a\xe0\x8b\x39\x34\x9f\x92\x3d\xae\x49\xa1\x4b\xa1\x2a\xd6\x5d\xe4\x83\xeb\x94\x0f\x0f\xd2\xbc\x63\xc6\xeb\x4d\xd3\xa3\xa2\xf6\x40\x8f\x8e\x4c\x7f\x13\x17\x13\x2d\x8d\x10\x96\xfc\x5f\x10\xec\x91\x9e\xc5\xf0\x10\xd5\x90\x71\x4c\x5e\x61\x58\xf9\x2a\x09\xe5\x8f\xc4\x27\xad\x70\x86\x0d\xa6\x2c\x90\xd4\x8f\x0a\x4f\x35\x17\x35\x68\xcd\x61\x87\x3a\x39\x61\x43\x8e\x04\x95\x4a\x78\x97\x42\x16\x5f\xd2\x55\x58\x77\x8e\x29\x5a\x11\x72\x9d\x85\xde\x4f\x7f\xc7\xd9\x43\x7a\xe2\xd2\x1f\xe2\xfd\xd6\xda\xb8\x70\x6c\x9a\xdd\x5f\x75\xdf\x7c\xb8\xe1\x91\xa7\x7a\xa3\x46\x9f\x1e\x2d\x38\x1b\x5b\xf6\xcd\x53\xd9\x40\x7e\x77\x8b\xd7\xdf\x01\x00\x00\xff\xff\x4e\x8d\x88\x25\x4d\x08\x00\x00")

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

	info := bindataFileInfo{name: "schema.json", size: 2125, mode: os.FileMode(420), modTime: time.Unix(1581095424, 0)}
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

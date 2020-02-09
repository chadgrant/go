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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x55\xcb\x6e\xdb\x40\x0c\xbc\xfb\x2b\x88\x4d\x0e\x2d\xe0\x58\x41\x81\xa2\x80\x7f\xa1\x40\x7b\x68\x6f\x4d\x60\xd0\x12\x65\x31\x91\x76\x5d\x2e\x95\x07\x8a\xfc\x7b\xb1\x2b\x59\x91\xec\x8d\x92\x16\xbd\xf4\xe8\x1d\x3e\x86\x14\x67\xfc\x6b\x01\x00\x60\xce\xb9\x30\x6b\x30\x95\xea\x7e\x9d\x65\x3e\xaf\xa8\x41\xbf\xf2\x64\x95\x1e\x56\xec\x32\x4f\x72\xc7\x39\x65\x15\x61\xad\xd5\xea\xc6\x3b\x6b\x96\x7d\x6a\x17\x3d\x4a\x0f\xe8\x45\xf7\xba\x72\xb2\xcb\x0a\xc1\x52\x2f\x2e\x3f\xf5\x75\xcf\x0e\x99\x05\xf9\x5c\x78\xaf\xec\x6c\xcc\x8e\xb5\x21\xaf\x28\xbf\x05\xa1\xbd\x13\xf5\xcf\xa1\x25\x5b\x0e\x91\xde\xac\xa1\x23\x1d\x81\x2e\x29\xe6\x4c\x80\x08\x5a\x6c\xe8\xe4\x35\x22\xfa\xb8\x0f\x88\xf1\x2a\x6c\x77\x7d\x97\x49\xc4\x11\xb9\x50\x0a\x5c\x09\x5a\x11\x8c\x7a\x82\x92\xd7\x54\x7a\xc3\x96\x9b\xb6\x31\x6b\xf8\x90\x42\xf1\xa1\x47\x3f\x5e\x4e\xd0\xa7\x69\xb0\x61\xab\x24\x77\x58\x6f\x1a\x3f\x3f\x88\x6d\x9b\x2d\xc9\x1b\x06\xe1\x12\xb6\x98\xdf\xee\xc4\xb5\xb6\x88\xf4\x97\x50\xb9\x7b\x70\xa5\x92\x8d\xe3\x85\x37\x60\x0f\xea\x60\x4b\x20\xad\x35\xb3\x14\x8b\x56\x30\x14\xff\x77\x14\xbf\xc4\xc0\xb0\xed\x86\xeb\x9a\x3d\xe5\xce\x16\x1e\x14\x6f\x03\x43\x17\x28\x0d\x44\xe7\xb9\x79\x45\x6d\x5f\xa1\xf5\xe6\x13\xf8\x5e\x11\x84\x82\xc3\x1d\x74\xcb\xe3\x12\xac\x53\xb8\x32\x5f\x3f\x5f\x99\x65\x58\x1c\x5b\x28\x91\x6b\x2a\x32\x12\x71\xd2\x25\xbd\x72\x24\xb3\x63\x84\x46\x54\x6c\x50\x37\xad\xe6\x7f\x3b\x4d\xe9\xa4\x41\x0d\x31\x05\x2a\x5d\x28\x37\x49\x4e\x89\xa1\x6b\xf4\x0a\x21\xfe\xf9\x3c\xee\xd1\xbf\x7e\x19\x42\x3f\x5b\x16\x0a\xce\xf2\xe3\xb4\x51\x14\x67\x8a\xc0\xe8\x9e\x12\x70\xff\x49\x13\xc8\x74\x4b\x13\xfc\x7a\xf8\xf5\xb4\x18\x31\x1d\xb6\xe6\xb6\x37\x94\x1f\x74\x6c\xf6\xe2\xf6\x24\xca\x74\xe4\x35\x35\xdf\x91\x25\x7f\x7a\x4f\x43\x1d\x14\xc1\x47\x73\xac\x61\xa5\x97\xa4\x71\x2e\x54\x86\xbc\xb3\x6c\x64\x70\xd9\xd8\xd3\xe6\xa5\x37\xfd\x56\xb1\x7b\x38\xce\x03\xd1\x13\xa7\x1a\x6c\x75\x91\x28\x69\x84\xb0\xe0\xff\x62\xc0\x81\xe9\x1f\x4f\x18\xc0\x0d\xfa\x8d\x2b\x93\x5a\x9a\xd7\xd1\x1b\x34\x94\xd2\x4f\x94\x0e\x2a\xdc\x57\x9c\x57\xa0\x15\xfb\x9e\x65\x94\xd1\x8e\x2c\x09\x2a\x15\xf0\x2e\x42\x0d\x3e\x46\x43\xd9\x76\x72\xcb\x5b\x11\xb2\x9d\xfe\xde\xa7\xa7\x9a\x73\xe0\x79\xf7\x3d\x66\xfb\xad\x6d\xc2\x7a\xb1\xae\xfb\xbf\xe1\x43\xe9\xf1\x3e\x27\x0a\x4a\x6a\xfc\x74\xcf\x2f\xb0\x5d\x26\xc4\x95\x3c\xc8\xf8\x76\xbd\x78\xfa\x1d\x00\x00\xff\xff\xa6\x90\xbd\x4f\xaf\x08\x00\x00")

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

	info := bindataFileInfo{name: "schema.json", size: 2223, mode: os.FileMode(420), modTime: time.Unix(1581291741, 0)}
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

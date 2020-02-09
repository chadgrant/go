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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x55\x4d\x6f\xd3\x40\x10\xbd\xe7\x57\x8c\xb6\x3d\x80\xe4\xc6\x15\x12\x42\xea\x5f\x40\x82\x03\xdc\x68\x15\x4d\xec\x71\x3c\xad\x77\x37\xcc\x8e\xfb\x21\xd4\xff\x8e\x76\x9d\xb8\x76\x62\x9c\xc0\x8d\xeb\xce\xd7\x7b\xe3\xf7\xc6\xbf\x16\x00\x00\xe6\x92\x4b\x73\x03\xa6\x56\xdd\xde\xe4\x79\x28\x6a\xb2\x18\x96\x81\x9c\xd2\xf3\x92\x7d\x1e\x48\x1e\xb9\xa0\xbc\x26\x6c\xb4\x5e\xde\x07\xef\x4c\xb6\x2b\xed\xb2\x07\xe5\x31\x7a\xd5\xbd\x2e\xbd\x6c\xf2\x52\xb0\xd2\xab\xeb\x4f\xbb\xbe\x17\xfb\xca\x92\x42\x21\xbc\x55\xf6\x2e\x55\xa7\xde\x50\xd4\x54\x3c\x80\xd0\xd6\x8b\x86\xb7\xd4\x8a\x1d\xc7\xcc\x60\x6e\xa0\x03\x9d\x02\x5d\x51\xaa\x31\x30\x8c\xa4\xa8\x43\x4b\xc7\xcf\x29\xa4\x2f\x5b\x8a\x53\x83\x0a\xbb\xcd\x6e\xce\x28\xe3\x00\x5e\xec\x05\xbe\x02\xad\x09\x06\x53\x41\x29\xe8\x54\xb9\x65\xc7\xb6\xb5\x71\xfc\x87\xa9\x30\x3e\xef\xc3\x1f\xaf\x47\xe1\xd7\x71\xb6\x61\xa7\x24\x8f\xd8\xac\x6c\x98\xe3\x02\x11\x63\x6b\xd7\x24\x67\x90\xe1\x0a\xd6\x58\x3c\x6c\xc4\xb7\xae\x4c\x14\x32\xa8\xfd\x13\xf8\x4a\xc9\x25\x8a\xf1\x0d\x38\x80\x7a\x58\x13\x48\xeb\xcc\x2c\xc8\xb2\x15\x8c\xcd\x4f\x80\xfc\x1b\x8c\x5f\x52\x62\x5c\xb9\xe5\xa6\xe1\x40\x85\x77\x65\x00\xc5\x87\x08\xd1\x47\x4c\x3d\xd2\x79\x70\x41\x51\xdb\x53\xb8\xce\x16\xc2\xf7\x9a\x20\x76\xec\xd5\xd0\xad\x8f\x2b\x70\x5e\xe1\xd6\x7c\xfd\x7c\x6b\xb2\xb8\x3a\x76\x50\x21\x37\x54\xe6\x24\xe2\xa5\x2b\x3a\x25\x95\x59\x22\x71\x12\x95\x2b\xd4\x55\xab\xc5\x3f\xf3\xa9\xbc\x58\xd4\x98\x53\xa2\xd2\x95\xb2\x9d\x44\x35\x41\xbb\xc1\xa0\x10\xf3\xdf\x24\xf2\x84\xe1\xb4\x3a\x84\x7e\xb6\x2c\x54\x46\xc8\x3f\x3a\x57\x66\x63\xcd\x64\xfd\x57\xca\x0e\x69\xde\xf5\xbd\x5e\x17\x83\xe6\x3d\x51\xbf\xbe\xa7\x62\x6f\x41\xb3\x15\xbf\x25\x51\xa6\x83\x43\xd1\xf0\x23\x39\x0a\xe3\xd7\x51\x1f\x14\xc1\x17\x73\x68\x3e\x25\x7b\x5c\x93\x42\x97\x42\x55\xac\xbb\xc8\x07\xd7\x29\x1f\x1e\xa4\x79\xc7\x8c\xd7\x9b\xa6\x47\x45\xed\x81\x1e\x1d\x99\xfe\x26\x2e\x26\x5a\x1a\x21\x2c\xf9\xbf\x20\xd8\x23\x3d\x8b\xe1\x21\xaa\x21\xe3\x98\xbc\xc2\xb0\xf2\x55\x12\xca\x1f\x89\x4f\x5a\xe1\x0c\x1b\x4c\x59\x20\xa9\x1f\x15\x9e\x6a\x2e\x6a\xd0\x9a\xc3\x0e\x75\x72\xc2\x86\x1c\x09\x2a\x95\xf0\x2e\x85\x2c\xbe\xa4\xab\xb0\xee\x1c\x53\xb4\x22\xe4\x3a\x0b\xbd\x9f\xfe\x8e\xb3\x87\xf4\xc4\xa5\x3f\xc4\xfb\xad\xb5\x71\xe1\xd8\x34\xbb\xbf\xea\xbe\xf9\x70\xc3\x23\x4f\xf5\x46\x8d\x3e\x3d\x5a\x70\x36\xb6\xec\x9b\xa7\xb2\x81\xfc\xee\x16\xaf\xbf\x03\x00\x00\xff\xff\x5a\x16\x19\x13\x55\x08\x00\x00")

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

	info := bindataFileInfo{name: "schema.json", size: 2133, mode: os.FileMode(420), modTime: time.Unix(1581291515, 0)}
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

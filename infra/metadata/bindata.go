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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x4d\x8f\xd3\x3c\x10\xbe\xef\xaf\x18\xe5\xdd\xa3\xd3\x6e\xf5\x82\x10\xbd\x21\x4e\x9c\x41\x42\x62\x77\xa9\x9c\x78\xda\x78\x49\x6c\x33\x9e\x74\xe9\x56\xfd\xef\xc8\xf9\x68\xdc\x25\x0d\x48\xed\x2d\xf1\x78\x9e\xe7\xf1\x7c\xee\x6f\x00\x00\x92\x5b\xad\x92\x25\x24\x05\xb3\x5b\xce\xe7\x3e\x2f\xb0\x92\x33\x8f\x86\xf1\xd7\x4c\xdb\x79\x85\x2c\x95\x64\x39\x7b\xf2\xd6\x24\xa2\xf3\x69\xaf\x45\x7e\xc1\x9a\x76\xce\x96\x36\x73\x45\x72\xcd\xe9\xdd\xbb\x0e\xf0\xbf\xde\x53\xa1\xcf\x49\x3b\xd6\xd6\x04\x6f\x8f\xb4\xd5\x39\x42\xcf\xd2\x5f\xe3\x9d\xc3\x60\xb7\xd9\x13\xe6\xdc\x9f\x3a\xb2\x0e\x89\x35\xfa\x64\x09\xad\xfc\xe6\x7c\x8b\x46\x59\x4a\x20\x3e\x1c\x60\x20\xf0\x30\x69\xb3\xe9\x70\x8e\xe6\x57\x62\x8c\xac\x10\xec\x1a\xb8\x40\xe8\x20\x8f\xf7\x0f\x83\x6b\xb2\x21\x5b\xbb\x8b\xd9\x1a\x14\x20\xf4\xce\x1a\xaf\xb3\x12\x61\x6d\xa9\xe1\xee\x82\x32\x4e\xde\x1b\xaf\xf9\xd8\x0e\x53\x80\xb1\x54\xc9\x52\xbf\xa0\x82\x67\xcd\x05\x18\x0b\xde\xc9\x1c\xfd\xb8\x96\x35\x69\x34\xaa\xdc\x5d\x2c\xa6\x07\x82\x11\x55\xe3\xd4\x31\xc0\xa5\xec\xb6\xf9\x92\x25\xf8\xc2\x12\x43\x64\x0d\x52\x26\x65\x64\xb5\x2e\xd5\x8a\xd0\xd9\xeb\xa9\xa8\xa9\x8c\x78\xa1\x01\x9f\x20\x37\x75\x95\xe1\xe5\xc5\xef\x75\x05\x5b\x24\x68\x40\xa1\x05\x15\x50\xfb\x5a\x96\xe5\x0e\x1c\xd9\xad\x56\xa8\x20\xdb\xc1\xc7\x4f\xed\x25\xff\x1a\xd1\x49\x66\xa4\x06\xed\xfb\xfd\x5d\xfa\xfe\x71\xbf\x10\xff\x1f\x1e\x1e\x66\xfb\xc5\xe1\xe4\x7f\xf8\xb9\x3d\xff\x32\x5e\x65\x97\x17\xd6\x97\x02\xa1\xf6\x18\xfa\x4a\x32\x28\xad\x9a\xca\x6a\xe4\x27\xb1\xdf\x44\x7c\x59\x57\x98\x8c\xcb\x38\xab\x62\x1d\xfa\x88\x83\x5d\x49\xc6\xb4\x81\x98\x16\xfa\xb5\x40\x33\x68\x83\x67\xe9\x41\x59\x73\xa6\xea\x36\x9a\x57\x85\xf4\xc5\x55\xc2\xb3\xd1\x0c\xbe\x90\x0b\x08\x88\x7d\xef\xb5\x11\x9a\xca\xef\x87\xf4\x9b\x4c\x5f\x9a\x44\xbe\x15\x6f\xee\xce\x65\x32\x48\xcd\x48\x9a\xfc\x7a\x62\x5b\xb8\x53\xa5\xa3\xdc\xb9\xad\x9c\x2e\x91\x56\x5b\x24\x7f\x8d\x41\xd1\x03\x42\x07\x18\x4a\x4b\x01\xdb\x2e\x69\x7f\x1d\x5a\xff\xa0\xa3\xeb\xe6\x69\x1d\x9f\xeb\x2a\x3c\x5f\x96\x25\xe4\x05\xe6\x3f\x40\xd5\x24\x83\x31\x9e\xd3\x37\x11\x7b\x42\xf8\xb3\xd6\x84\x61\xd3\xdf\xf7\xeb\x52\x74\x8b\x4c\x1c\x77\x8a\x18\x26\xba\x38\xe1\x14\xf1\xa0\x13\x43\x7b\x8a\xb8\x47\xc4\x50\x97\x22\xce\xbb\xf8\x33\x11\x22\x19\xbe\x4e\xa6\xd8\xe3\xcd\xe1\x77\x00\x00\x00\xff\xff\xe3\xf6\xcb\xd7\x97\x08\x00\x00")

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

	info := bindataFileInfo{name: "schema.json", size: 2199, mode: os.FileMode(420), modTime: time.Unix(1580837231, 0)}
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

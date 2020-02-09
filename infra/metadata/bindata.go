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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x95\x4f\x8f\xd3\x3c\x10\xc6\xef\xfb\x29\x46\x79\xf7\xe8\xb4\x5b\xbd\x20\x44\x6f\x88\x13\x67\x90\x90\xd8\x5d\x2a\x27\x9e\x36\x5e\x12\xdb\x8c\x27\x5d\xba\x55\xbf\x3b\x72\xfe\x34\xee\x92\x06\xa4\xf6\x96\x78\xec\xdf\xf3\xd8\x9e\x19\xef\x6f\x00\x00\x92\x5b\xad\x92\x25\x24\x05\xb3\x5b\xce\xe7\x3e\x2f\xb0\x92\x7e\xe6\xd1\x30\xfe\x9a\x69\x3b\xf7\x48\x5b\x9d\xe3\xbc\x42\x96\x4a\xb2\x9c\x3d\x79\x6b\x12\xd1\x2d\x6e\xe7\x47\x80\x10\x4d\xdb\xd1\x99\xa5\xcd\x5c\x91\x5c\x73\x7a\xf7\xae\x23\xff\xd7\xaf\x54\xe8\x73\xd2\x8e\xb5\x35\x61\x75\xa7\x02\xbd\x4a\x3f\x8d\x77\x0e\x43\xdc\x66\x4f\x98\x73\x3f\xea\xc8\x3a\x24\xd6\xe8\x93\x25\xb4\xfb\x68\xc6\xb7\x68\x94\xa5\x04\xe2\xc1\x01\x03\x41\x87\x49\x9b\x4d\xc7\x39\x86\x5f\x99\x31\xb2\x42\xb0\x6b\xe0\x02\xa1\x43\x1e\xe7\x1f\x86\xa5\xc9\x86\x6c\xed\x2e\x56\x6b\x28\x40\xe8\x9d\x35\x5e\x67\x25\xc2\xda\x52\xa3\xdd\x1d\xca\xb8\x78\x1f\xbc\xe6\x66\x3b\xa6\x00\x63\xa9\x92\xa5\x7e\x41\x05\xcf\x9a\x0b\x30\x16\xbc\x93\x39\xfa\x71\x2f\x6b\xd2\x68\x54\xb9\xbb\xd8\x4c\x0f\x82\x11\x57\xe3\xd2\x31\xe0\x52\x75\xdb\x7c\xc9\x12\x7c\x61\x89\x21\x8a\x06\x2b\x93\x36\xb2\x5a\x97\x6a\x45\xe8\xec\xf5\x5c\xd4\x54\x46\xba\xd0\xc0\x27\xc4\x4d\x5d\x65\x78\x79\xf2\x7b\x5d\xc1\x16\x09\x1a\x28\xb4\x50\x01\xb5\xaf\x65\x59\xee\xc0\x91\xdd\x6a\x85\x0a\xb2\x1d\x7c\xfc\xd4\x4e\xf2\xaf\x89\x4e\x32\x23\x35\xb4\xef\xf7\x77\xe9\xfb\xc7\xfd\x42\xfc\x7f\x78\x78\x98\xed\x17\x87\x93\xff\xe1\xe7\xf6\xfc\xce\x78\x95\x5d\x9e\x58\x5f\x0a\x84\xda\x63\xa8\x2b\xc9\xa0\xb4\x6a\x32\xab\xb1\x9f\xc4\xeb\x26\xce\x97\x75\x85\xc9\xb8\x8d\xb3\x2e\xd6\xa1\x8e\x38\xc4\x95\x64\x4c\x1b\xc4\xb4\xd1\xaf\x05\x9a\xc1\x1b\x3c\x4b\x0f\xca\x9a\x33\x59\xb7\xd1\xbc\x2a\xa4\x2f\xae\x72\x3c\x1b\xcd\xe0\x0b\xb9\x80\x40\xec\x6b\xaf\x3d\xa1\xa9\xfb\xfd\x90\x7e\x93\xe9\x4b\x73\x91\x6f\xc5\x9b\xbb\x73\x37\x19\xac\x66\x24\x4d\x7e\x3d\xb3\x2d\xee\xd4\xe9\xa8\x76\x6e\x2b\xa7\x4b\xa4\xd5\x16\xc9\x5f\xa3\x51\xf4\x40\xe8\x80\x21\xb5\x14\xb0\xed\x2e\xed\xaf\x4d\xeb\x1f\x7c\x74\xd5\x3c\xed\xe3\x73\x5d\x85\xed\xcb\xb2\x84\xbc\xc0\xfc\x07\xa8\x9a\x64\x08\xc6\x7d\xfa\x26\x52\x4f\x08\x7f\xd6\x9a\x30\x3c\xf9\xf7\xfd\x73\x29\xba\x87\x4c\x1c\xdf\x14\x31\x74\x74\x71\xa2\x29\xe2\x46\x27\x86\xf2\x14\x71\x8d\x88\x21\x2f\x45\x7c\xef\xe2\xcf\x8b\x10\xc9\xf0\x75\xd2\xc5\x1e\x6f\x0e\xbf\x03\x00\x00\xff\xff\x4a\x3a\x59\xdd\xa0\x08\x00\x00")

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

	info := bindataFileInfo{name: "schema.json", size: 2208, mode: os.FileMode(420), modTime: time.Unix(1581291507, 0)}
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

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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x92\xcf\x6a\xf3\x40\x0c\xc4\xef\x7e\x0a\xb1\xc9\x31\xb1\xbf\xdb\x07\x79\x83\x42\xa1\x50\xe8\xa9\xf4\xb0\xb5\xc7\xb1\x82\xe3\x75\xb5\x4a\xda\x50\xfc\xee\x65\xe3\x75\xfe\xd4\x4e\xa1\x97\x5e\x35\x1a\xe9\xa7\xdd\xf9\x4c\x88\x88\xcc\x9c\x0b\xb3\x22\x53\xa9\xb6\xab\x2c\xf3\x79\x85\xad\xf5\xa9\x47\xa3\xf8\x48\xd9\x65\x1e\xb2\xe7\x1c\x51\xa9\xd9\x6b\xba\xf1\xae\x31\x8b\x68\xef\xeb\x17\x23\x82\xba\xec\xab\xa9\x93\x75\x56\x88\x2d\x75\xf9\xef\x7f\x9c\x30\x1b\x9c\x05\x4a\x6e\x58\xd9\x35\xde\xac\xa8\xa7\x39\x0a\x82\x35\x7b\x85\xa0\x38\xcd\x3e\xab\xc7\x0e\x3d\xb4\x08\x1b\xdd\xeb\x06\xb9\xc6\x81\x27\xb5\x15\xd7\x42\x94\xe1\x47\xce\xa3\xbe\x13\x9e\x14\x22\x95\xcf\x85\xdb\x80\x15\x36\x3c\x3d\xde\x11\x17\x68\x94\x4b\x86\x90\x2b\x29\x32\x2d\xa6\xed\x03\x99\x57\xe1\x66\x7d\xab\xab\x74\xb2\xb5\x1a\xfa\x02\xcb\xa8\xa7\x1b\xdb\xcc\x4e\xea\x5f\x40\xdf\x07\xd2\xf7\x0a\x02\x52\x47\x02\x15\xc6\x1e\x7f\xc2\x9e\xfc\x70\x89\x11\xbc\xed\x58\x10\x12\xf7\x3c\xfd\x31\x37\x4e\xbf\xaa\xbe\x24\xd7\xdb\xba\x53\xa2\xae\x9e\x21\x64\xf5\xfc\x63\x9e\xb4\x02\xc5\x34\x53\x65\x3d\xd5\xce\x16\x28\x86\x38\x0e\xd7\x5b\x11\x7b\x18\x8a\xac\xd8\x7e\x4b\xa7\xad\xeb\x87\x72\xc4\x3f\x11\xb3\xb9\x20\xf4\x99\x59\x76\x91\xf4\x6c\x14\xee\x64\xfa\xf5\xfa\x1b\xbb\xa4\xfb\x0a\x00\x00\xff\xff\xb8\x88\x91\x07\xa7\x03\x00\x00")

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

	info := bindataFileInfo{name: "schema.json", size: 935, mode: os.FileMode(420), modTime: time.Unix(1581291752, 0)}
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

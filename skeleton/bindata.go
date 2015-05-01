package skeleton

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _resource_tmpl_common_changelog_md_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\x56\xa8\xae\x56\xd0\x0b\x4b\x2d\x2a\xce\xcc\xcf\x53\xa8\xad\x55\xd0\x00\xf2\x53\x12\x4b\x52\x81\x6c\x4d\x2e\x2e\xcf\xbc\xcc\x92\xcc\xc4\x1c\x85\xa2\xd4\x9c\xd4\xc4\xe2\x54\x2e\x2e\x65\xa0\x0e\xc7\x94\x94\xd4\x14\x2e\x2e\x5d\x10\x43\xc1\xad\x34\x2f\x25\x31\x37\x35\xaf\x04\xa8\x2a\x2d\x35\xb1\xa4\xb4\x28\xb5\x18\xa2\xcc\x25\xb5\xa0\x28\x35\x19\x68\x14\x58\xad\x5f\x7e\x49\x46\x66\x5e\x3a\x44\x2a\x28\x35\x37\xbf\x0c\x8b\xb8\x5b\x66\x05\xaa\x28\x20\x00\x00\xff\xff\xf6\x0c\x35\xa7\xa0\x00\x00\x00")

func resource_tmpl_common_changelog_md_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_resource_tmpl_common_changelog_md_tmpl,
		"resource/tmpl/common/CHANGELOG.md.tmpl",
	)
}

func resource_tmpl_common_changelog_md_tmpl() (*asset, error) {
	bytes, err := resource_tmpl_common_changelog_md_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resource/tmpl/common/CHANGELOG.md.tmpl", size: 160, mode: os.FileMode(420), modTime: time.Unix(1430466720, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resource_tmpl_common_readme_md_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x50\x4d\x6b\xe3\x30\x10\xbd\xeb\x57\x0c\x64\x0f\x09\xec\xda\xec\x35\xb7\x25\x4b\xa1\x97\xb6\x84\xf6\x14\x0a\x1a\x3b\x63\x4b\xd4\x92\x52\x69\x44\x28\x21\xff\xbd\x23\xb9\xa5\xe9\xa9\xed\xc1\x58\x7a\xa3\xf7\x31\x6f\x01\xa7\x13\x34\x37\xe8\x08\xce\x67\xa5\xca\xe5\x3f\xa5\x3e\xda\x03\xdb\xe0\x2b\xb6\x58\xc0\x05\x54\xef\x0f\x09\x47\xaa\xa7\x6b\x9f\x18\xa7\x49\xa9\xfb\x00\x76\x3e\xff\x86\x9c\x08\xf4\x18\x60\x24\xd6\x6b\xa5\xb4\xd6\x1d\x26\xa3\x7e\xc1\x8c\xc1\x9f\x3d\x8c\x96\x4d\xee\x9a\x3e\xb8\xb6\x78\xde\x1e\x3d\x45\x71\x6b\x2f\xd3\x08\xaf\x7a\x6c\x82\xe7\x68\xbb\x3c\xdb\xff\x6d\xe0\x2a\xc4\x27\x58\xee\x0c\xf3\x21\xad\xdb\xf6\x1b\x5a\xed\x20\x94\xc7\xe5\x4f\x19\xab\x55\xb1\xdb\x44\x42\x26\x40\x18\xe4\x9f\x23\x41\x17\xd1\xf7\xa6\x8e\x82\x73\x96\xe1\x25\xe4\x08\xbd\x41\x3f\x52\x2a\xf0\x96\x64\x61\x9a\xe1\x29\xf4\x38\xbd\x0f\x01\x47\x2c\x35\x01\x1b\x02\x87\x89\xc5\xf7\x43\x6d\x9b\x3d\x30\xc9\x34\x65\x2b\x86\x47\x49\x59\x1f\x96\x2e\x2b\xde\xb4\x4d\xd3\x68\x90\xdc\x0e\xfd\x1e\xca\xd7\x07\x3f\xd8\xe8\xe4\x1d\x32\x48\x94\x03\xa6\xf4\x16\x42\xd4\x84\x39\x38\x29\x3c\xe9\x4f\x8b\x78\x3a\xc2\x5d\x9e\x26\x09\xfa\x9c\x45\xb8\xd6\xfc\x2f\xb3\x09\x51\xa9\xdd\x65\x25\x5f\x76\xb6\x52\xaf\x01\x00\x00\xff\xff\xb5\x89\x6c\xef\x43\x02\x00\x00")

func resource_tmpl_common_readme_md_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_resource_tmpl_common_readme_md_tmpl,
		"resource/tmpl/common/README.md.tmpl",
	)
}

func resource_tmpl_common_readme_md_tmpl() (*asset, error) {
	bytes, err := resource_tmpl_common_readme_md_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resource/tmpl/common/README.md.tmpl", size: 579, mode: os.FileMode(420), modTime: time.Unix(1430468872, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resource_tmpl_common_version_go_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xc8\x4d\xcc\xcc\xe3\xe2\x4a\xce\xcf\x2b\x2e\x51\xf0\x4b\xcc\x4d\x55\x28\x2e\x29\xca\xcc\x4b\x57\xb0\x55\x50\xaa\xae\x56\xd0\x03\x0b\xd5\xd6\x2a\x41\x55\x84\xa5\x16\x15\x67\xe6\xe7\xa1\x29\x82\x89\x82\xd4\x01\x02\x00\x00\xff\xff\x61\x03\x18\x7c\x58\x00\x00\x00")

func resource_tmpl_common_version_go_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_resource_tmpl_common_version_go_tmpl,
		"resource/tmpl/common/version.go.tmpl",
	)
}

func resource_tmpl_common_version_go_tmpl() (*asset, error) {
	bytes, err := resource_tmpl_common_version_go_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resource/tmpl/common/version.go.tmpl", size: 88, mode: os.FileMode(420), modTime: time.Unix(1430469109, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"resource/tmpl/common/CHANGELOG.md.tmpl": resource_tmpl_common_changelog_md_tmpl,
	"resource/tmpl/common/README.md.tmpl": resource_tmpl_common_readme_md_tmpl,
	"resource/tmpl/common/version.go.tmpl": resource_tmpl_common_version_go_tmpl,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"resource": &_bintree_t{nil, map[string]*_bintree_t{
		"tmpl": &_bintree_t{nil, map[string]*_bintree_t{
			"common": &_bintree_t{nil, map[string]*_bintree_t{
				"CHANGELOG.md.tmpl": &_bintree_t{resource_tmpl_common_changelog_md_tmpl, map[string]*_bintree_t{
				}},
				"README.md.tmpl": &_bintree_t{resource_tmpl_common_readme_md_tmpl, map[string]*_bintree_t{
				}},
				"version.go.tmpl": &_bintree_t{resource_tmpl_common_version_go_tmpl, map[string]*_bintree_t{
				}},
			}},
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}


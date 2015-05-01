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

var _resource_tmpl_common_readme_md_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x90\x41\x6b\xe3\x30\x10\x85\xef\xfa\x15\x03\xd9\x43\x02\xbb\x32\x7b\xcd\x6d\xc9\x52\xe8\xa5\x2d\xa1\x3d\x85\x82\xc6\xce\xd8\x12\xb5\xa4\x54\x1a\x11\x4a\xc8\x7f\xef\x58\x6e\x69\x7a\x6a\x7b\x30\x96\x66\x34\xef\x7b\xf3\x16\x70\x3a\x81\xbe\x41\x4f\x70\x3e\x2b\xb5\x58\xc0\x7f\xca\x5d\x72\x07\x76\x31\xd4\xfb\x43\xc6\x81\xea\xe9\x3a\x64\xc6\x71\x54\xea\x3e\x82\x9b\xcf\xbf\xa1\x64\x02\x33\x44\x18\x88\xcd\x5a\x29\x63\x4c\x8b\xd9\xaa\x5f\x30\xd7\xe0\xcf\x1e\x06\xc7\xb6\xb4\xba\x8b\xbe\x99\x60\xb7\xc7\x40\x49\x68\xcd\x25\x59\xe6\x2a\x63\x13\x03\x27\xd7\x96\x19\xff\x57\xc3\x55\x4c\x4f\xb0\xdc\x59\xe6\x43\x5e\x37\xcd\x37\xb4\x9a\x5e\x46\x1e\x97\x3f\x9d\x58\xad\x26\xdc\x26\x11\x32\x01\x42\x2f\xff\x92\x08\xda\x84\xa1\xb3\xb5\x15\xbd\x77\x0c\x2f\xb1\x24\xe8\x2c\x86\x81\xf2\x54\xde\x92\x2c\x4c\x73\x79\x8c\x1d\x8e\xef\x4d\xc0\x01\xa7\x98\x80\x2d\x81\xc7\xcc\xc2\xfd\x50\xdb\x96\x00\x4c\xd2\xcd\xc5\x09\xf0\x28\x2e\xeb\xc3\x29\xcb\x5a\xd7\x8d\xd6\xda\x80\xf8\xf6\x18\xf6\x30\x7d\x5d\x0c\xbd\x4b\x5e\xde\x21\x83\x58\x39\x60\xce\x6f\x26\x44\x4d\x26\x7b\x2f\x81\x67\xf3\x69\x91\x40\x47\xb8\x2b\xe3\x28\x46\x9f\x8b\x08\xd7\x98\xff\x15\xb6\x31\x29\xb5\xbb\x8c\xe4\xcb\xcc\x56\xea\x35\x00\x00\xff\xff\x50\xb7\xe3\x55\x2f\x02\x00\x00")

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

	info := bindata_file_info{name: "resource/tmpl/common/README.md.tmpl", size: 559, mode: os.FileMode(420), modTime: time.Unix(1430460412, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resource_tmpl_common_version_go_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xc8\x4d\xcc\xcc\xe3\xe2\x4a\xce\xcf\x2b\x2e\x51\xf0\x4b\xcc\x4d\x55\x28\x2e\x29\xca\xcc\x4b\x57\xb0\x55\x50\xaa\xae\x56\xd0\x03\x0b\xd5\xd6\x2a\x41\x55\x84\xa5\x16\x15\x67\xe6\xe7\x21\x29\x32\xd0\x33\xd4\x33\x50\x02\x04\x00\x00\xff\xff\x25\x62\xab\xbd\x4e\x00\x00\x00")

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

	info := bindata_file_info{name: "resource/tmpl/common/version.go.tmpl", size: 78, mode: os.FileMode(420), modTime: time.Unix(1430460357, 0)}
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


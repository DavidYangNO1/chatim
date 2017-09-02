package utility

import (
	"os"

	"github.com/spf13/afero"
)

var freeFS afero.Fs

func init() {
	freeFS = afero.NewOsFs()
}

func Mkdir(name string, perm os.FileMode) error {

	err := freeFS.Mkdir(name, perm)
	return err
}

func Chmod(name string, mode os.FileMode) error {
	return freeFS.Chmod(name, mode)
}

func Create(name string) (afero.File, error) {
	fs, err := freeFS.Create(name)
	return fs, err
}

func Open(name string) (afero.File, error) {
	return freeFS.Open(name)
}

func Remove(name string) error {
	return freeFS.Remove(name)
}

func RemovePath(path string) error {
	return freeFS.RemoveAll(path)
}

func DirExists(path string) (bool, error) {
	return afero.DirExists(freeFS, path)
}

func Exists(path string) (bool, error) {
	return afero.DirExists(freeFS, path)
}

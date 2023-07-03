package helper

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/lolorenzo777/verbose"
	"github.com/otiai10/copy"
)

// Copy copies src to dst, can be a single file or a full path
func CopyFiles(dstDir string, src string) (err error) {
	if dstDir == src {
		return os.ErrInvalid
	}

	infosrc, errstat := os.Stat(src)
	if errstat != nil {
		return errstat
	}

	infodst, errstat := os.Stat(dstDir)
	if errstat != nil {
		return errstat
	}
	if !infodst.IsDir() {
		return errors.New("destination must be an directory")
	}

	if infosrc.IsDir() {
		err = copy.Copy(src, dstDir)
	} else {
		srcF, erro := os.Open(src)
		if erro != nil {
			return erro
		}
		defer srcF.Close()

		_, filename := filepath.Split(src)
		dst := path.Join(dstDir, filename)

		dstF, errof := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, infosrc.Mode())
		if errof != nil {
			return errof
		}
		defer dstF.Close()

		_, err = io.Copy(dstF, srcF)
	}

	return err
}

// MustCheckOutputPath checks if output is a valid path and returns its corresponding absolute path if successuf.
// exit if the output is not nil and an error occurs.
func MustCheckOutputPath(output *string) string {
	apath, err := CheckOutputPath(output)
	if err != nil {
		os.Exit(1)
	}
	return apath
}

// CheckOutputPath checks if output is an existing path and returns its corresponding absolute path if successful.
// if output is nil or empty then returns an empty path an no error.
// output is a reference to a string to enable to pass directly the returned value of a flag.String result.
// Returns an error of the path is not a valid dir.
func CheckOutputPath(output *string) (validpath string, err error) {
	if output != nil && *output != "" {
		validpath, _ = filepath.Abs(*output)
		fileInfo, errf := os.Stat(validpath)
		if errf != nil || !fileInfo.IsDir() {
			verbose.Error("makedoc", fmt.Errorf("output %s is not a valid path", *output))
			fmt.Println("makedoc fails.")
			if !verbose.IsOn {
				fmt.Println("use the verbose flag to get more info.")
			}
			return "", errf
		}
	}
	return validpath, nil
}

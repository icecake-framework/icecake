package html

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sunraylab/verbose"
)

// MustCheckOutputPath checks if output is a valid path and returns its corresponding absolute path if successuf.
// exit if the output is not nil and an error occurs.
func MustCheckOutputPath(output *string) string {
	apath, err := CheckOutputPath(output)
	if err != nil {
		os.Exit(1)
	}
	return apath
}

// CheckOutputPath checks if output is a valid path and returns its corresponding absolute path if successuf.
// if output is nil or empty then returns an empty path an no error.
// Returns an error of the path is not a valid dir.
func CheckOutputPath(output *string) (validpath string, err error) {
	if output != nil && *output != "" {
		validpath, _ = filepath.Abs(*output)
		fileInfo, err := os.Stat(validpath)
		if err != nil || !fileInfo.IsDir() {
			verbose.Error("makedoc", fmt.Errorf("output %s is not a valid path", *output))
			fmt.Println("makedoc fails.")
			if !verbose.IsOn {
				fmt.Println("use the verbose flag to get more info.")
			}
		}
	}
	return validpath, err
}

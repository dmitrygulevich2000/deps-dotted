package finder

import (
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/Masterminds/semver/v3"
)

func EqualityMatchFunc(target string) MatchFunc {
	return func(str string) bool {
		return str == target
	}
}

var majorVerRegexp = regexp.MustCompile("^v[0-9]+$")

func ModuleMatchFunc(targetMod string, verConstraint *semver.Constraints) func(string) bool {
	return func(fullPath string) bool {
		modPath, verStr, err := splitModPath(fullPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return false
		}

		ver, err := semver.NewVersion(verStr)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return false
		}

		return modPath == targetMod && verConstraint.Check(ver)
	}
}

func stripMajorVer(modPath string) string {
	if majorVerRegexp.FindString(path.Base(modPath)) != "" {
		return path.Dir(modPath)
	} else {
		return modPath
	}
}

func splitModPath(modPath string) (path, ver string, err error) {
	entries := strings.Split(modPath, "@")
	if len(entries) != 2 {
		err = errors.New("invalid path " + modPath + ". expected format: module/path@v1.2.3")
		return
	}
	path = stripMajorVer(entries[0])
	ver = entries[1]
	return
}

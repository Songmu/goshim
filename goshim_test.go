package goshim

import (
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestSrcList(t *testing.T) {
	files := []string{"main.go", "type.go"}

	dir := "testdata/testproj"
	list, _ := srcList(dir)
	if !reflect.DeepEqual(list, files) {
		t.Errorf("failed to get srcList. got: %+v, expected: %+v", list, files)
	}

	dir = "./" + dir
	list, _ = srcList(dir)
	if !reflect.DeepEqual(list, files) {
		t.Errorf("failed to get srcList. got: %+v, expected: %+v", list, files)
	}

	dir, _ = filepath.Abs(dir)
	list, _ = srcList(dir)
	if !reflect.DeepEqual(list, files) {
		t.Errorf("failed to get srcList. got: %+v, expected: %+v", list, files)
	}

	dir = "testdata/non-existence"
	_, err := srcList(dir)
	if err == nil || !strings.HasPrefix(err.Error(), "can't load package: package ./testdata/non-existence: open") {
		t.Errorf("something went wrong: %v", err)
	}

	dir = "testdata/notgo"
	_, err = srcList(dir)
	if err == nil || !strings.HasPrefix(err.Error(), "can't load package: package ./testdata/notgo: no buildable Go source files in") {
		t.Errorf("something went wrong %v:", err)
	}
}

func TestRun(t *testing.T) {
	dir := "testdata/testproj"
	ret := Run([]string{dir})
	if ret != 0 {
		t.Errorf("something went wrong: %v", ret)
	}
}

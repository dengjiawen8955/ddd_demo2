package main

import (
	"dc2/internal/common/utils"
	"fmt"
	"path/filepath"
	"testing"
)

func Test_P2(t *testing.T) {
	GoToJSON("test.go", "json")
}

func TestP3(t *testing.T) {
	jsonBaseDir := "json/"
	files, err := utils.MatchFiles("./../../../common/", "*.go")

	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Printf("files: %v\n", files)
	// 遍历文件
	for _, filePath := range files {
		jsonDir := jsonBaseDir + filepath.Dir(filePath)
		GoToJSON(filePath, jsonDir)
	}
}

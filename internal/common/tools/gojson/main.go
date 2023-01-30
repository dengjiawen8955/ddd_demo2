package main

import (
	"dc2/internal/common/utils"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

var Usage = `
Usage: gojson [root] [pattern] [json dir]
Example: gojson src "*.go" json/
将 src 目录下所有匹配 *.go 的文件结构体转换为 json 文件，保存到 json 目录下
`

func main() {
	params := os.Args
	if len(params) != 4 {
		println(Usage)
		return
	}

	root := params[1]
	pattern := params[2]
	jsonDir := params[3]

	if root == "" {
		root = "."
	}
	if pattern == "" {
		pattern = "*.go"
	}
	if jsonDir == "" {
		jsonDir = "json/"
	}

	files, err := utils.MatchFiles(root, pattern)

	if err != nil {
		panic(err)
	}

	// 遍历文件
	for _, filePath := range files {
		jsonDir := jsonDir + filepath.Dir(filePath)
		// 打印 filePath 和 jsonDir
		fmt.Printf("[filePath]: %v ==> [jsonDir]: %v \n", filePath, jsonDir)
		GoToJSON(filePath, jsonDir)
	}
}

func GoToJSON(goFile string, jsonDir string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, goFile, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		if _, ok := typeSpec.Type.(*ast.StructType); !ok {
			return true
		}

		// create json data
		data := map[string]interface{}{}
		for _, field := range typeSpec.Type.(*ast.StructType).Fields.List {
			// 优先使用 json tag 作为 key
			var key string
			if field.Tag != nil {
				tagStr := field.Tag.Value
				// 使用正则表达式提取 json tag
				ss := regexp.MustCompile(`json:"(.+?)"`).FindStringSubmatch(tagStr)
				var jsonTag string
				if len(ss) == 2 {
					jsonTag = ss[1]
				}

				if jsonTag != "-" && jsonTag != "" {
					key = jsonTag
				}
			}
			// 如果没有 json tag 则使用字段名作为 key
			if key == "" {
				key = field.Names[0].Name
			}

			data[key] = fieldValue(field)
		}

		// marshal json data
		jsonData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			panic(err)
		}

		// write json data to file
		if _, err := os.Stat(jsonDir); os.IsNotExist(err) {
			if err := os.MkdirAll(jsonDir, os.ModePerm); err != nil {
				panic(err)
			}
		}
		jsonFile := filepath.Join(jsonDir, typeSpec.Name.Name+".json")

		// 如果文件已存在，则跳过
		if _, err := os.Stat(jsonFile); !os.IsNotExist(err) {
			return true
		}

		if err := ioutil.WriteFile(jsonFile, jsonData, os.ModePerm); err != nil {
			panic(err)
		}

		return true
	})
}

// 获取字段的值
// 如果有结构体嵌套，则递归获取
func fieldValue(f *ast.Field) interface{} {
	switch f.Type.(type) {
	case *ast.Ident:
		switch f.Type.(*ast.Ident).Name {
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float32", "float64":
			return 0
		case "string":
			return ""
		case "bool":
			return false
		case "[]string", "[]int", "[]int8", "[]int16", "[]int32", "[]int64", "[]uint", "[]uint8", "[]uint16", "[]uint32", "[]uint64", "[]float32", "[]float64":
			return []interface{}{}
		case "map[string]string", "map[string]int", "map[string]int8", "map[string]int16", "map[string]int32", "map[string]int64", "map[string]uint", "map[string]uint8", "map[string]uint16", "map[string]uint32", "map[string]uint64", "map[string]float32", "map[string]float64":
			return map[string]interface{}{}
		default:
			return ""
		}
	case *ast.ArrayType:
		return []interface{}{}
	case *ast.MapType:
		return map[string]interface{}{}
	case *ast.StructType:
		// 如果是结构体类型，则递归获取
		value := map[string]interface{}{}
		for _, field := range f.Type.(*ast.StructType).Fields.List {
			value[field.Names[0].Name] = fieldValue(field)
		}
		return value
	case *ast.StarExpr:
		value := map[string]interface{}{}
		return value
	}

	return ""
}

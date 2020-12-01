package main

import (
	"GolangParser/model"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
)

// Check Public or Private
func isPublic(temp string) bool {
	return temp[0] >= 'A' && temp[0] <= 'Z'
}

// Check Go File like xx.go

func isGoFile(path string) bool {
	pathLenth := len(path)
	if pathLenth < 3 {
		return false
	}
	return path[pathLenth-1] == 'o' && path[pathLenth-2] == 'g' && path[pathLenth-3] == '.'
}

func GetParamDetail(param *ast.Field) model.FunctionParam {
	var res model.FunctionParam
	for _, name := range param.Names {
		res.ParamName = name.Name
	}
	res.ParamType = param.Type
	return res
}

func GetFuncDetail(funcDecl *ast.FuncDecl) model.FunctionInfo {
	var res model.FunctionInfo
	res.FunctionName = funcDecl.Name.Name

	// Get The Function Params
	ast.Inspect(funcDecl, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.FuncType:
			{
				if x.Params != nil {
					for _, param := range x.Params.List {
						res.FuncParam = append(res.FuncParam, GetParamDetail(param))
					}
				}
				if x.Results != nil {
					for _, param := range x.Results.List {
						res.FuncReturn = append(res.FuncReturn, GetParamDetail(param))
					}
				}
			}

		}
		return true
	})

	// Get The Body Of The Function
	// to do

	// Get The Type Of Function
	if isPublic(res.FunctionName) {
		res.FunctionPublicOrPrivate = "Public"
	} else {
		res.FunctionPublicOrPrivate = "Private"
	}
	return res
}

func main() {
	// File Processing
	var GoFiles []model.GoSourceFile
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if isGoFile(path) {
				GoFiles = append(
					GoFiles,
					model.GoSourceFile{FilePath: "./" + path},
				)
			}
			return nil
		})
	if err != nil {
		fmt.Print(err)
	}

	for _, GoFile := range GoFiles {
		var FunctionSet []model.FunctionInfo
		fset := token.NewFileSet()
		path, _ := filepath.Abs(GoFile.FilePath)
		f, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
		if err != nil {
			log.Println(err)
			return
		}
		ast.Inspect(f, func(node ast.Node) bool {
			switch x := node.(type) {
			case *ast.FuncDecl:
				FunctionSet = append(FunctionSet, GetFuncDetail(x))
			}
			return true
		})

		fmt.Println("-------------------------------")
		fmt.Println(GoFile.FilePath + " :: as follow")
		for _, temp := range FunctionSet {
			model.ShowFunctionInfo(temp)
			fmt.Println("-------------------------------")
		}
	}

	// _ = ast.Print(fset, f)
}

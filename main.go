package main

import (
	"GolangParser/model"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
)

func isPublic(temp string) bool {
	return temp[0] >= 'A' && temp[0] <= 'Z'
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

	// Get The Type Of Function
	if isPublic(res.FunctionName) {
		res.FunctionPublicOrPrivate = "Public"
	} else {
		res.FunctionPublicOrPrivate = "Private"
	}
	return res
}

func main() {
	var FunctionSet []model.FunctionInfo
	fset := token.NewFileSet()
	path, _ := filepath.Abs("./admin.go")
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
	for _, temp := range FunctionSet {
		model.ShowFunctionInfo(temp)
		fmt.Println("-------------------------------")
	}
	_ = ast.Print(fset, f)
}

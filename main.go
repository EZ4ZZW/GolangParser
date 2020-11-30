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

func GetDetail(funcDecl *ast.FuncDecl)  model.FunctionInfo{
	var res model.FunctionInfo
	res.FunctionName = funcDecl.Name.Name
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
		var s string
		switch x := node.(type) {
		case *ast.FuncDecl:
			FunctionSet = append(FunctionSet, GetDetail(x))
		}
		d := len(s) > 0
		if d {
			fmt.Println(s)
		}
		return true
	})

	fmt.Println("-------------------------------")
	for _, temp := range FunctionSet {
		model.ShowFunctionInfo(temp)
		fmt.Println("-------------------------------")
	}

}

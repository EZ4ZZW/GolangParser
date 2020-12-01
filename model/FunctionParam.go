package model

import (
	"fmt"
	"go/ast"
)

type FunctionParam struct {
	ParamName string
	ParamType interface{}
}

func ShowParamInfo(param FunctionParam) {
	var TypeName string
	// assert the type of param
	switch x := param.ParamType.(type) {
	case *ast.Ident:
		TypeName = x.Name
		break
	case *ast.MapType:
		{
			switch mapType := x.Key.(type) {
			case *ast.Ident:
				TypeName = "map[" + mapType.Name + "]"
				break
			}
			switch mapValue := x.Value.(type) {
			case *ast.InterfaceType:
				mapValue.Pos()
				TypeName = TypeName + "interface{}"
				break
			}
		}
	case *ast.ArrayType:
		break
	case *ast.ChanType:
		break
	case *ast.InterfaceType:
		break

	}
	fmt.Println("ParamType :" + TypeName + "   ParamName :" + param.ParamName)
}

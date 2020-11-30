package model

import "fmt"

type FunctionParam struct {
	ParamName string
	ParamType string
}

func ShowParamInfo(param FunctionParam)  {
	fmt.Println("ParamType :" + param.ParamType + "   ParamName :" + param.ParamName)
}
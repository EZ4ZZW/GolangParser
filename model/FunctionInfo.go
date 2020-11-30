package model

import "fmt"

type FunctionInfo struct {
	FunctionName string
	FunctionPublicOrPrivate string
	FuncParam []FunctionParam
	FuncReturn []FunctionParam
}

func ShowFunctionInfo(info FunctionInfo)  {
	fmt.Println("FunctionName: " + info.FunctionName)
	fmt.Println("Public Or Private: " + info.FunctionPublicOrPrivate)
	fmt.Println("FuncParam :")
	for _, temp := range info.FuncParam {
		ShowParamInfo(temp)
	}
	fmt.Println("FuncReturn :")
	for _, temp := range info.FuncParam {
		ShowParamInfo(temp)
	}

}
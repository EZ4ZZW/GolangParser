package model

import "fmt"

type FunctionInfo struct {
	FunctionName            string
	FunctionPublicOrPrivate string
	FuncParam               []FunctionParam
	FuncReturn              []FunctionParam
	FunctionCalling         []string
}

func ShowFunctionInfo(info FunctionInfo) {
	fmt.Println("FunctionName: " + info.FunctionName)
	fmt.Println("Public Or Private: " + info.FunctionPublicOrPrivate)
	fmt.Println("FuncParam :")
	cnt := 0
	for _, temp := range info.FuncParam {
		fmt.Print(cnt)
		fmt.Print(" : ")
		cnt = cnt + 1
		ShowParamInfo(temp)
	}
	cnt = 0
	fmt.Println("FuncReturn :")
	for _, temp := range info.FuncParam {
		fmt.Print(cnt)
		fmt.Print(" : ")
		cnt = cnt + 1
		ShowParamInfo(temp)
	}

}

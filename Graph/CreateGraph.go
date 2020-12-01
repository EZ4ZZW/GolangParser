package Graph

import (
	"GolangParser/model"
)

func FindApiByGraph(functions []model.FunctionInfo) []string {
	var graph map[string]int
	var ApiList []string
	for _, function := range functions {
		for _, calling := range function.FunctionCalling {
			graph[calling] = 0
		}
	}
	for _, function := range functions {
		for _, calling := range function.FunctionCalling {
			f := calling
			graph[f]++
		}
	}

	for key, value := range graph {
		if value == 0 {
			ApiList = append(ApiList, key)
		}
	}

	return ApiList

}

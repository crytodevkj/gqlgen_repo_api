/// [date] 2022-04-28
/// [ref] https://stackoverflow.com/questions/49448302/convert-interface-to-struct

package main

import (
	"fmt"
	"strings"
)

/// @param raw: unstructured graphql response
/// @return: array of nodes
func GetNodes(raw interface{}) []interface{} {
	res, _ := raw.(map[string]interface{})["projects"].(map[string]interface{})["nodes"].([]interface{})
	return res
}

/// @param node: unstructured node element of graphql response
/// @return: name field of node
func GetName(node interface{}) string {
	res, ok := node.(map[string]interface{})["name"]
	if !ok {
		return "_"
	}
	return res.(string)
}

/// @param node: unstructured node element of graphql response
/// @return: forksCount field of node
func GetForksCount(node interface{}) float64 {
	res, ok := node.(map[string]interface{})["forksCount"]
	if !ok {
		return 0
	}
	return res.(float64)
}

func GetNamesFromApi(num string) string {
	// call to repository layer
	res := Api(Url, strings.Replace(Req, "DISPLAY_NUM", num, -1))
	// parse into array of node
	nodes := GetNodes(res)

	names := []string{}
	// loop through array to get names | sum of all forks
	for i := 0; i < len(nodes); i++ {
		names = append(names, GetName(nodes[i]))
	}

	// return results
	return strings.Join(names, ", ")
}

func GetSumFromApi(num string) string {
	// call to repository layer
	res := Api(Url, strings.Replace(Req, "DISPLAY_NUM", num, -1))
	// parse into array of node
	nodes := GetNodes(res)

	var sumOfAllForks float64 = 0
	// loop through array to get names | sum of all forks
	for i := 0; i < len(nodes); i++ {
		sumOfAllForks += GetForksCount(nodes[i])
	}

	// return results
	return fmt.Sprint(sumOfAllForks)
}

package pertemuanbagian2

import (
	"fmt"
	"strings"
)

func filter(data []string, callbasck func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callbasck(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func main() {
	var data = []string{"gajah", "buaya", "ular"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "u")
	})
	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t : ", data)

	fmt.Println("filter ada huruf \"u\"\t : ", dataContainsO)

	fmt.Println("filter ada huruf \"5\"\t : ", dataLength5)

}

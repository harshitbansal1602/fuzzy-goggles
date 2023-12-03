package helpers

import "fmt"

func ConvertToInt(s string) int {
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	if err != nil {
		fmt.Println("Error converting to int:", err)
	}
	return result
}

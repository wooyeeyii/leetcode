package string

import (
	"fmt"
	"testing"
)

func Test_problem1436(t *testing.T) {
	// Sao Paulo
	fmt.Println(destCity([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}))
	// A
	fmt.Println(destCity([][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}))
}

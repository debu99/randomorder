package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func init() {
	testing.Init()
}
func TestprintNums(t *testing.T) {
	result := printNums()
	slice := strings.Split(result, ",")
	if len(slice) != 10 {
		t.Errorf("Failed! Expected length=%v, got %v\n", 10, len(slice))
	} else {
		fmt.Printf("length=%v\n", len(slice))
	}
	if len(slice) != len(Unique(slice)) {
		t.Errorf("Failed! Expected unique length=%v, got %v\n", len(Unique(slice)), len(slice))
	} else {
		fmt.Printf("unique length=%v\n", len(slice))
	}
	for _, v := range slice {
		i, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			panic(err)
		}
		if i < 1 || i > 10 {
			t.Errorf("Failed! Expected %v, got %v\n", "1<=x<=10", v)
		}
	}
}

func TestgenRandIndex(t *testing.T) {
	result := genRandIndex(0, 0)
	if result != 0 {
		t.Errorf("Failed! Expected %v, got %v\n", 0, result)
	} else {
		fmt.Printf("0 %v\n", result)
	}

	result = genRandIndex(0, 1)
	if result < 0 || result > 1 {
		t.Errorf("Failed! Expected %v, got %v\n", "0<=x<=1", result)
	} else {
		fmt.Printf("0<=x<=1 %v\n", result)
	}

	result = genRandIndex(1, 1)
	if result != 1 {
		t.Errorf("Failed! Expected %v, got %v\n", 1, result)
	} else {
		fmt.Printf("1 %v\n", result)
	}

	result = genRandIndex(0, 10)
	if result < 0 || result > 10 {
		t.Errorf("Failed! Expected %v, got %v\n", "0<=x<=10", result)
	} else {
		fmt.Printf("0<=x<=10 %v\n", result)
	}

	result = genRandIndex(1, 10)
	if result < 1 || result > 10 {
		t.Errorf("Failed! Expected %v, got %v\n", "1<=x<=10", result)
	} else {
		fmt.Printf("1<=x<=10 %v\n", result)
	}

	result = genRandIndex(-20, 0)
	if result < -20 || result > 0 {
		t.Errorf("Failed! Expected %v, got %v\n", "-20<=x<=0", result)
	} else {
		fmt.Printf("-20<=x<=0 %v\n", result)
	}

	result = genRandIndex(-1, 0)
	if result < -1 || result > 0 {
		t.Errorf("Failed! Expected %v, got %v\n", "-1<=x<=0", result)
	} else {
		fmt.Printf("-1<=x<=0 %v\n", result)
	}

	result = genRandIndex(-1, -1)
	if result != -1 {
		t.Errorf("Failed! Expected %v, got %v\n", -1, result)
	} else {
		fmt.Printf("-1 %v\n", result)
	}

	result = genRandIndex(-2, -1)
	if result < -2 || result > -1 {
		t.Errorf("Failed! Expected %v, got %v\n", "-2<=x<=-1", result)
	} else {
		fmt.Printf("-2<=x<=-1 %v\n", result)
	}

	result = genRandIndex(-10, 10)
	if result < -10 || result > 10 {
		t.Errorf("Failed! Expected %v, got %v\n", "-10<=x<=10", result)
	} else {
		fmt.Printf("-10<=x<=10 %v\n", result)
	}

}

func Unique(slice []string) []string {
	// create a map with all the values as key
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}

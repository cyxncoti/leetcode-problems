package main

import "fmt"

func main() {
	fmt.Println(compareVersion("1.01", "1.001"))
	fmt.Println(compareVersion("1.0", "1.0.0"))
	fmt.Println(compareVersion("0.1", "1.1"))
	fmt.Println(compareVersion("1.1", "1.10"))
}

func compareVersion(version1 string, version2 string) int {
	len1, len2 := len(version1), len(version2)
	for ptr1, ptr2 := 0, 0; ptr1 < len1 || ptr2 < len2; {
		num1, num2 := 0, 0
		for ; ptr1 < len1 && version1[ptr1] != '.'; ptr1++ {
			num1 = num1*10 + int(version1[ptr1]-'0')
		}
		for ; ptr2 < len2 && version2[ptr2] != '.'; ptr2++ {
			num2 = num2*10 + int(version2[ptr2]-'0')
		}
		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}
		ptr1++
		ptr2++
	}
	return 0
}

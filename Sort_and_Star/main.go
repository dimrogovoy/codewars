package main

import (
	"fmt"
	"sort"
	"strings"
)

//type Strings []string
//
//func (s Strings) Len() int {
//	return len(s)
//}
//
//func (s Strings) Less(i, j int) bool {
//	return s[i] < s[j]
//}
//
//func (s Strings) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}

func TwoSort(arr []string) string {
	//sort.Sort(Strings(arr))
	sort.Strings(arr)
	a := strings.Split(arr[0], "")

	return strings.Join(a, "***")
	//for i := 1; i < len(a); i += 2 {
	//	a = append(a, "")
	//	copy(a[i+1:], a[i:])
	//	a[i] = "***"
	//}

}

func main() {
	s := []string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}
	fmt.Println(TwoSort(s))
}

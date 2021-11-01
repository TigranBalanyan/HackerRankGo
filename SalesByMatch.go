// package main

// import (
// 	"fmt"
// 	"sort"
// )

// /*
//  * Complete the 'sockMerchant' function below.
//  *
//  * The function is expected to return an INTEGER.
//  * The function accepts following parameters:
//  *  1. INTEGER n
//  *  2. INTEGER_ARRAY ar
//  */

// func sockMerchant(n int, ar []int) int {

// 	// Write your code here

// 	var k, _ = SortArray(ar)
// 	var result int = 0

// 	for i := 0; i < len(k)-1; i++ {

// 		if k[i] == k[i+1] {
// 			result += 1
// 			i++
// 		}

// 	}

// 	fmt.Println(result)

// 	return result
// }

// func SortArray(ar []int) ([]int, error) {

// 	sort.Ints(ar[:])

// 	return ar, nil
// }

// func main() {

// 	var i int

// 	_, err := fmt.Scan(&i)

// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}

// 	var array = make([]int, i)

// 	var j int = 0

// 	for ; j < i; j++ {
// 		var element int
// 		_, err := fmt.Scan(&element)

// 		if err != nil {
// 			fmt.Print(err.Error())
// 			break
// 		}
// 		array[j] = element
// 	}

// 	sockMerchant(i, array)
// }

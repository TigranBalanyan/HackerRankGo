package main

import (
	"errors"
	"fmt"
)

/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func sockMerchant(n int32, ar []int32) int32 {

	// Write your code here

	var k, _ = SortArray(ar)
	var result int32 = 0

	for i := 0; i < len(k) -1; i++ {

		fmt.Println("Sorted Array")
		fmt.Print((k[i]))
		
		if(k[i] == k[i+1]){
			result +=1;
			i ++
		} 
	}
	fmt.Println(result)

	return result
}

func SortArray(ar []int32) ([]int32, error) {

	var isDone = false

	var emptyArray error

	if(ar == nil){
		emptyArray = errors.New("Empty array")
	}

	for !isDone {
		isDone = true
		var i = 0

		for i < len(ar) - 1 {
			if ar[i] > ar[i+1] {
				ar[i], ar[i+1] = ar[i+1], ar[i]
			}
			i++
		}
	}

	return ar, emptyArray
}

func main() {

	var i int32

	_, err := fmt.Scan(&i)


	if err != nil {
		fmt.Print(err.Error())
	}

	var array = make([]int32, i)  


	var j int32 = 0
	
	for ; j < i; j++ {
		var element int32
		_, err := fmt.Scan(&element)

		if err != nil {
			fmt.Print(err.Error())
			break
		}
		array[j] = element
	}

	sockMerchant(i, array)
}

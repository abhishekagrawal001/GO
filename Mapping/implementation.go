package main

import "fmt"

func main() {
	// empty map declaration. It is equal to nil. Cannot add keys to an empty map.(CTE)
	// var mymap map[int] string

	// if mymap == nil {
	// 	fmt.Println("True")
	// }else {
	// 	fmt.Println("false")
	// }

	// map declaration with keys-values
	// mymap := map[int]string{1:"hello", 2:"gm"}
	// fmt.Println(mymap)

	//map declaration using make keyword
	var my_map = make(map[int]string)
    fmt.Println(my_map)
	my_map[91] = "Rohit"
    my_map[11] = "Sumit"
	my_map[66] = "Saurabh"
    // fmt.Println(my_map)

	//iterate over map
	for k,v := range my_map {
		fmt.Println(k,v)
	}

	//direct print
	fmt.Println(my_map[11])

	//delete a pair
	delete(my_map,11)

	//update a key-value
	my_map[11] = "Abhishek"

	//check whether value for a key exist or not
	value,present := my_map[66]
	fmt.Println("Value of mymap[2]=",value)
	fmt.Println("value is present or not=",present)

}
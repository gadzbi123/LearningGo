package main

import "fmt"

func randomMap() {
	myMap := map[int]bool{1: true, 2: false, 3: true}
	for k, v := range myMap {
		if v {
			myMap[k+10] = true
		}
	}
	fmt.Println(myMap)
}
func main() {
	// Results can change
	randomMap() // map[1:true 2:false 3:true 11:true 13:true 23:true 33:true 43:true 53:true 63:true]
	randomMap() // map[1:true 2:false 3:true 11:true 13:true 21:true]
	randomMap() // map[1:true 2:false 3:true 11:true 13:true 21:true]

}

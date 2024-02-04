package main

import (
	"fmt"
)

func main() {

	//Declare map
	languages := make(map[string]string)

	//append key and value to map
	languages["Go"] = "Golang"
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	//Accessing the key and values in the map
	fmt.Println(languages)
	fmt.Println("Go:", languages["Go"])
	fmt.Println("JS:", languages["JS"])
	fmt.Println("PY:", languages["PY"])
	fmt.Println("RB:", languages["RB"])

	//Deleting key and value from map
	delete(languages, "JS")
	fmt.Println(languages)

	//loop through maps
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	//we can make implicit like this
	maps := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println(maps)

	//Declare the map the values are slice
	sliceMap := make(map[string][]string)
	fmt.Println(sliceMap)

	sliceMap["Car"] = append(sliceMap["Car"], "BMW", "Benz", "Audi")
	fmt.Println(sliceMap["Car"][0])

	vehicles := map[string][]string{
		"Car":  {"BMW", "Audi", "Benz"},
		"Bike": {"R15", "Splendor", "KTM"},
	}
	fmt.Println(vehicles)

	//append
	vehicles["Others"] = []string{"Activa", "Dio"}
	fmt.Println(vehicles)
	delete(vehicles, "Others")

	temp := vehicles["Others"]
	fmt.Println(temp) //this is empty because no key Others in the map

	temp, ok := vehicles["Others"] //ok will check the key there then give bool value true or false
	fmt.Println(temp, ok)
}

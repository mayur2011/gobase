package main

import "fmt"

var data map[string]string //pkg level variable

func isKeyExists(key string) bool {
	_, ok := data[key]
	return ok
}

func getAllData() {
	for k, v := range data {
		fmt.Println(k, ":", v)
	}
}

func deleteDataElement(key string) {
	if isKeyExists(key) {
		delete(data, key)
		fmt.Println(key, "- Data Element is removed")
		return
	}
	fmt.Println(key, "- Data Element doesn't exist")
}

func updateDataElement(key, value string) {
	if isKeyExists(key) {
		data[key] = value
		fmt.Println(key, " is updated with value =", value)
		return
	}
	fmt.Println(key, "does not exist")
}

func addData(key, value string) {
	if isKeyExists(key) {
		fmt.Println("key already exists")
		return
	}
	data[key] = value
	fmt.Println(key, "is added")
}

func init() {
	data = make(map[string]string) //initialize map with make function
}

func main() {
	fmt.Println("adding data elements..")
	//add new data as key and value
	addData("IN", "INDIA")
	addData("US", "UNITED STATES")
	addData("GB", "UNITED KINGDOM")
	addData("SG", "SINGAPORE")
	addData("AU", "AUSTRALIA")
	addData("US", "UNITED STATES")
	addData("JP", "JAPAN")
	addData("LU", "JAPAN")
	addData("XY", "XXXX")
	addData("YZ", "YYYY")
	fmt.Println("\nAll initially loaded data elements")
	getAllData()

	fmt.Println("\nDeleting data elements -")
	//delete data element
	deleteDataElement("XY")
	deleteDataElement("YZ")

	fmt.Println("\nUpdating data elements -")
	//update data element
	updateDataElement("LU", "LUXEMBURG")
	updateDataElement("PK", "PPPPP")

	fmt.Println("\nDisplaying the latest data elements -")
	//get all of data printed
	getAllData()
}

package main

import "fmt"

func main(){
	region, continent := location("Bangalore")
	fmt.Printf("Bikash lives in %s, %s", region, continent)
}

func location(loc string) (string, string){
	var region string
	var continent string
	switch loc {
	case "Los Almos":
		region, continent = "California", "North America"
	case "Bangalore":
		region, continent = "India","Asia"
	default:
		region,continent = "Unknown", "Unknown"
	}
	return region, continent


}
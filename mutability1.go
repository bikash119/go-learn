package main

import "fmt"

type Artist struct {
	Name, Genre string
	Songs int
}

func newRelease(a Artist) int {
	a.Songs++
	return a.Songs
}

func main(){
	me := Artist{Name: "Bikash", Genre:"Electro", Songs:2}
	fmt.Printf("%s release their %dth song", me.Name,newRelease(me))
	fmt.Printf("%s released their %dth song", me.Name, me.Songs)
}
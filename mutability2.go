package main

import "fmt"

type Artist struct {
	Name , Genre string
	Songs int
}

func main(){
	me := &Artist{Name: "Bikash", Genre: "Electro", Songs : 42}
	fmt.Printf( "%s released their %dth song", me.Name, newRelease(me))
	fmt.Println()
	fmt.Printf( "%s released their %dth song", me.Name, me.Songs)
}

func newRelease(a *Artist) int{
	a.Songs++
	return a.Songs
}
/*
/*
@Time : 2020/11/23 3:43 下午
@Author : chengqunzhong
@File : ga
@Software: GoLand
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year  int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}
var movies = []Movie {
	{Title: "Casablance", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisser"}},
}

func ta() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}

func tb() {
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed %s", err)
	}
	fmt.Printf("%s\n", data)
}
var titles []struct{Title string}

func tc() {
	data, _ := json.MarshalIndent(movies, "", "    ")
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}

func main() {
	//ta()
	//tb()
	tc()
}
















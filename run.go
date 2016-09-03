package main

import (
	"log"
	"os"
	"strconv"
)

func counter() int {
	i, err := strconv.Atoi(getArgParameter("chance", "100"))
	if err != nil {
		log.Fatal(err)
	}
	return i
}
func getArgParameter(name string, defaultValue string) string {
	argsWithoutProg := os.Args[1:]
	for index, element := range argsWithoutProg {
		if element == "--"+name {
			return argsWithoutProg[index+1]
		}
	}
	return defaultValue
}

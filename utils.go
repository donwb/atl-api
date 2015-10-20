package main

import (
	"log"
)

func panicIf(err error) {
	if err != nil {
		log.Println("---AN ERROR OCCURED WHILE RUNNING atl-api---")
		log.Println("---COMPLETE DETAIL FOLLOWS BELOW---")
		log.Println(err)

		panic(err)
	}
}

func logIf(err error) {
	if err != nil {
		log.Println("---AN ERROR OCCURED WHILE RUNNING atl-api---")
		log.Println("---COMPLETE DETAIL FOLLOWS BELOW---")
		log.Println(err)

		log.Println("----- CONTINUING ----- ")
	}
}

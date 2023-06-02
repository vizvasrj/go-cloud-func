package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("time.list")
	if err != nil {
		log.Fatal(err)
	}
	var total_min int64
	var total_sec int64
	for _, x := range strings.Split(string(file), "\n") {
		// fmt.Println(x)
		time := strings.Split(x, "min")
		min := strings.TrimSpace(time[0])
		sec := strings.TrimSpace(strings.Split(time[1], "sec")[0])
		fmt.Printf("min = %s sec = %s\n", min, sec)
		minute, err := strconv.ParseInt(min, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		second, err := strconv.ParseInt(sec, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		total_min += minute
		total_sec += second

	}
	// fmt.Println("total min", total_min)
	fmt.Println("total min", total_sec/60+total_min+70, "sec", total_sec%60)

}

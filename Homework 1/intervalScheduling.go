package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var allTimeSlot []RequestTimeSlot

type RequestTimeSlot struct {
	startTime int
	endTime   int
}

func main() {

	defer fmt.Println("-=-=Done Executing=-=-")

	allTimeSlot = readCsv()
	fmt.Println("Here's all the requests read from requests.csv (sorted by endtime)")

	//Sort all of the timeslots by its end time
	sort.Slice(allTimeSlot, func(i, j int) bool {
		return allTimeSlot[i].endTime < allTimeSlot[j].endTime
	})

	fmt.Println(allTimeSlot)
	fmt.Print("===Begin interval scheduling===\n\n")

	fmt.Println("The accepted time slots are")
	fmt.Println(intervalSchedule())
}

func intervalSchedule() []RequestTimeSlot {

	acceptedTimeSlots := []RequestTimeSlot{}
	latestEndTime := 0

	for _, curTimeSlot := range allTimeSlot {
		if curTimeSlot.startTime >= latestEndTime {
			latestEndTime = curTimeSlot.endTime
			acceptedTimeSlots = append(acceptedTimeSlots, curTimeSlot)
		}
	}

	return acceptedTimeSlots
}

func readCsv() []RequestTimeSlot {

	//Values gotten from exercise_2 lecture_2 slides
	file, err := os.Open("requests5.csv") //change file here !!
	if err != nil {
		log.Fatal("Unable to read requests.csv")
	}
	defer file.Close()
	fmt.Println("Loaded requests.csv")

	csvReader := csv.NewReader(file)
	timeSlotList, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse requests.csv")
	}

	var allTimeSlot []RequestTimeSlot

	for _, currentTimeSlot := range timeSlotList {

		startTime, _ := strconv.Atoi(currentTimeSlot[0])
		endTime, _ := strconv.Atoi(currentTimeSlot[1])

		timeSlot := RequestTimeSlot{
			startTime: startTime,
			endTime:   endTime,
		}

		allTimeSlot = append(allTimeSlot, timeSlot)
	}

	return allTimeSlot
}

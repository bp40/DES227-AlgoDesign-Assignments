package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

var allTimeSlot []RequestTimeSlot

type RequestTimeSlot struct {
	startTime int
	endTime   int
}

func main() {

	allTimeSlot = readCsv()
	fmt.Println("Here's all the requests read from requests.csv")
	fmt.Println(allTimeSlot)
	fmt.Print("===Begin interval scheduling===\n\n")

	fmt.Println(intervalSchedule())
}

func intervalSchedule() []RequestTimeSlot {

	currentEndTime := 0
	var acceptedSlots []RequestTimeSlot

	for len(allTimeSlot) > 0 { //Loop unti no request left

    //Clear the non compatible slots before proceeding to find the fastest request 
    //in order to prevent the initial value for findFastestFinish() to be impossible
    clearImpossibleSlot(currentEndTime) // REMOVE
    if len(allTimeSlot) == 0 { //If the removal removes all the remaining time slot, exit loop
      break
    }
		fastestRequest := findFastestFinish()
		fmt.Println("The request ", fastestRequest, "is the fastest (accepted)")
		acceptedSlots = append(acceptedSlots, fastestRequest) // ACCEPT

		currentEndTime = fastestRequest.endTime
		
	}

  //results
  fmt.Print("\n\n ==The accepted slots are==\n")
	return acceptedSlots

}

//Searches through the remaining time slot list to find a slot that starts after the end time of the 
//latest accepted request and ends the fastest
func findFastestFinish() RequestTimeSlot {

  //allTimeSlot[0] are all compatible because of deletion first in intervalSchedule()
	fastestEndSlot := allTimeSlot[0]

	for _, curTimeSlot := range allTimeSlot {
		if  curTimeSlot.endTime < fastestEndSlot.endTime {
			fastestEndSlot = curTimeSlot
		}
	}

	return fastestEndSlot

}

//deletes any request that ends before current end time
//creates a temp array to store the non deleted request and 
//replaces the allTimeSlot array
func clearImpossibleSlot(currentFinishTime int) {

	var tempAllTimeSlot []RequestTimeSlot

	for _, curTimeSlot := range allTimeSlot {
    
		if curTimeSlot.startTime >= currentFinishTime {
			tempAllTimeSlot = append(tempAllTimeSlot, curTimeSlot)
		} else {
      fmt.Println(curTimeSlot, " is now either accepted or imppossible (deleted)")
    }
    
	}

	allTimeSlot = tempAllTimeSlot

}

func readCsv() []RequestTimeSlot {

	//Values gotten from exercise_2 lecture_2 slides
	file, err := os.Open("requests.csv") //change file here !!
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

package main
import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
	"regexp"
	)


func sanitizeInput(time []string) []string {
	var period string = ""
	var newTime []string = []string{""}
	calcTime := time[0]
	if len(time) > 1 {
		period = time[1]
	}
        
   	//Split the time into hours and minutes and assign to variables
        t := strings.Split(calcTime,":")
        hour,_ := strconv.Atoi(t[0])
        min,_ := strconv.Atoi(t[1])
	regEx1,_ := regexp.Compile("\\d+:\\d{2}")

	if len(time) > 2 {
		fmt.Println("Entered more than 2 arguments, Please try again entering time + AM/PM, example 12:00PM")
                os.Exit(0)
        }		
	
	//Check to ensure user entered proper time (numeric instead of alpha, presence of colon delimiter, inform user and exit if not)
        if (regEx1.MatchString(calcTime) == false) {
                fmt.Println("Invalid Time")
                os.Exit(0)
        }
	
	 //Check to ensure user has entered a proper time in 12 hour format, inform user and exit if not
        if (hour > 12) || (min > 59) {
                fmt.Println("Invalid Time")
                os.Exit(0)
        }
	
	//Check to see if user put more than 1 argument (this is to fix a shell issue where arguments are tokenized by space so the output looks correct (('9:00am is': instead of '9:00am  is))
	if len(time) == 1 {
		newTime = []string{calcTime}
	} else {
		newTime = []string{calcTime,period}
	}
	
	return newTime

}

func findDegrees(time []string) {

        calcTime := time[0]

        //Split the time into hours and minutes and assign to variables
        t := strings.Split(calcTime,":")
        hour,_ := strconv.Atoi(t[0])
        min,_ := strconv.Atoi(t[1])


        //Get the degrees for the hour hand
        hourDeg := (60*hour+min)/2
        //Get the degrees for the minute hand
        minDeg := 6*min
        //Get the angle of the hour hand in relation to the minute hand
        fullDeg := hourDeg - minDeg

        //Check to see if angle from previous equation is > 180, if so subtract that value from 360 to get smallest angle
        if fullDeg > 180 {
                fullDeg -= 360
        }
        //Print formatted string listing input time and angle
        fmt.Printf("The angle of %v is: %.0f degrees\n", strings.Join(time, " "),math.Abs(float64(fullDeg)))
}

func main() {
	//Call sanitizeInput method with input time as parameter.  We use 1: here to capture all arguments after the time (i.e if someone puts 9:00 AM instead of 9:00AM)
	time := sanitizeInput(os.Args[1:])
	//Call findDegrees method with return value from sanitizeInput
	findDegrees(time)
}

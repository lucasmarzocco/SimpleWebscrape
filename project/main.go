package project

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"unicode"
	"time"
)

// Struct to hold all of the information being returned
type JobListing struct {
	Title string `json:"title"`
	Location string `json:"location"`
	Company string `json:"company"`
	Url string `json:"url"`
}

// Struct to pass around the channels
type Result struct {
	Url string
	Contents string
}

// Return Job title - ALL CORNER CASES ARE NOT TAKEN CARE OF
func getTitleAndLocation(contents string) (string, string) {

	startIndex := strings.Index(contents, "<title>")

	// Amount of chars in <title>
	startIndex += 7
	endIndex := strings.Index(contents, "</title>")
               
	jobAndLocation := strings.Split(contents[startIndex:endIndex], "-")

	length := len(jobAndLocation)

	fmt.Println(length)
	fmt.Println(jobAndLocation)

	job := ""
	location := ""

	// Special number of dash-delimited items when checking for job/location
	if length == 3 {

		// Grab only the job
		job = strings.TrimSpace(jobAndLocation[0])

		// Grab only the location
		location = strings.TrimSpace(jobAndLocation[1])

	}else {

		for i := 0; i <= (length - 3); i++ {

			job += strings.TrimSpace(jobAndLocation[i])

			if i != (length - 3) {

				job += " - "
			}
		}

		// Location will always be the second to last item in the array
		location = strings.TrimSpace(jobAndLocation[length-2])
	}

	// CHecking last char for a number
	lastChar := location[len(location)-1]

	// Remove ZIP code
	if unicode.IsDigit(rune(lastChar)) {
		
		// Take care of stripping the zip code if need to
		newLocation := location[0:len(location)-6]
		location = newLocation
	}

	return job, location
}

// Return company
func getCompany(contents string) string {

	test := strings.Index(contents, "Get job updates from")

	// Amount of chars to add to reach end of string
	test += 21

	testEnd := strings.Index(contents, "duplicateEmailMessage")

	// Weird amount of number to backtrack, but it reaches the end of the company name
	testEnd -= 3

	company := strings.TrimSpace(contents[test:testEnd])

	return company
}

// HTTP GET
func get(link string, ch chan<-Result) {
	response, _ := http.Get(link)
	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	
	// Send all the data over the channel
	result := Result {
		Url: link,
		Contents: string(contents),
	}

	ch <- result
}

// Does all the heavy lifting
func GetAllJobs(vals []string) []*JobListing {

	start := time.Now()
	var list []*JobListing
	ch := make(chan Result)

	for _, element := range vals {
		go get(element, ch)
	}

	for range vals {

		val := <-ch

		job, location := getTitleAndLocation(val.Contents)
		company := getCompany(val.Contents)

		jobPosting := &JobListing {
			Title: job,
			Location: location,
			Company: company,
			Url: val.Url,
		}

		list = append(list, jobPosting)
	}

	elapsed := time.Since(start)
	fmt.Printf("Completed in: %s", elapsed)
	return list
}
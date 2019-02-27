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

// Return Job title
func getTitleAndLocation(contents string) (string, string) {

	startIndex := strings.Index(contents, "<title>")

	// Amount of chars in <title>
	startIndex += 7
	endIndex := strings.Index(contents, "</title>")

	jobAndLocation := strings.Split(contents[startIndex:endIndex], "-")

	// Grab only the job
	job := strings.TrimSpace(jobAndLocation[0])

	// Grab only the location
	location  := strings.TrimSpace(jobAndLocation[1])

	// Deals with the case when there's a - in the job title (Ex: Software Engineer - iOS Devices)
	if len(jobAndLocation) > 3 {

		newJob := strings.TrimSpace(jobAndLocation[0] + " - " + jobAndLocation[1])
		newLocation := strings.TrimSpace(jobAndLocation[2])

		job = newJob
		location = newLocation
	}

	lastChar := location[len(location)-1]

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
	contents, _ := ioutil.ReadAll(response.Body)
	
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
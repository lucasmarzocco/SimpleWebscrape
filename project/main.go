package project

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

// Struct to hold all of the information being returned
type JobListing struct {
	Title string `json:"title"`
	Location string `json:"location"`
	Company string `json:"company"`
	Url string `json:"url"`
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
func get(link string) string {
	response, err := http.Get(link)
	if err != nil {
		fmt.Println("%s", err)
		os.Exit(1)

	}else {

		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		return string(contents)
	}
	return ""
}

// Does all the heavy lifting
func GetAllJobs(vals []string) []*JobListing {

	var list []*JobListing

	for _, element := range vals {

		contents := get(string(element))
		job, location := getTitleAndLocation(contents)
		company := getCompany(contents)

		jobPosting := &JobListing {
			Title: job,
			Location: location,
			Company: company,
			Url: string(element),
		}

 
		list = append(list, jobPosting)
	}

	return list
}
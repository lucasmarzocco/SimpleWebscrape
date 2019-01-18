package project

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"
	"strings"
)

//Struct to store the result of fetching a name
type Name struct {
	Name string
	Surname string
	Gender string
	Region string
}

//Struct to store the result of fetching a joke
type ChuckNorris struct {
	Type string
	Value struct {
		Id int
		Joke string
		Categories []string
	}
}

//HTTP get to retrieve the data from the given links
func get(link string) []byte {

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

		return contents
	}

	slice := make([]byte, 0)
	return slice
}

//Main method, does all the heavy lifting
func Main() string {

	names := "http://uinames.com/api/"
	chuck_norris := "http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=[nerdy]"

	name_input := get(names)
	var p Name
	err := json.Unmarshal(name_input, &p)

	if err != nil {
		panic(err)
	}

	//Grabs name from the output
	name := p.Name
	surname := p.Surname

	/**************************************************************/

	joke_input := get(chuck_norris)
	var q ChuckNorris
	err1 := json.Unmarshal(joke_input, &q)

	if err1 != nil {
		panic(err1)
	}

	//Grabs joke from the output
	joke := q.Value.Joke

	//Replaced joke
	result := strings.Replace(joke, "John Doe", name + " " + surname, -1)
	return result + "."
	/**************************************************************/
}
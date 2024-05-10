package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}                 //Added a wait group to wait for all the threads to return back
	data, err := os.ReadFile("./input.txt") //read the input file from it's location

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n") //split the text line and break line by line

	for _, line := range lines {
		wg.Add(1)   //Increase thread count for each thread
		go func() { //used go routine to call the thread
			data, err := getData(line)
			if err != nil {
				log.Printf("unable to get status: %v", err)
				return
			}

			if data == "foo" {
				fmt.Printf("data found: %s", data)

			}
			wg.Done()
		}()
	}
	wg.Wait()  //Waits till thread returns
	os.Exit(0) //close the file
}

func getData(line string) (string, error) {
	var location struct {
		URL string
	}
	if err := json.Unmarshal([]byte(line), &location); err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, location.URL, nil)
	if err != nil {
		return "", err
	}

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close() // Close the body after reading

	decoder := json.NewDecoder(res.Body)
	var payload struct {
		Data string `json:"data"`
	}
	if err := decoder.Decode(&payload); err != nil {
		return "", err
	}

	return payload.Data, nil
}

// seond way:- Do simple get call

// func getData(line string) (string, error) {

// 	// Define a struct to represent the JSON data
// 	type Data struct {
// 		Location string `json:"location"`
// 	}

// 	// Unmarshal the JSON string into the Data struct
// 	var data Data
// 	err := json.Unmarshal([]byte(line), &data) //Unmarshal is a method. We don't want to send a copy so we use pointer as it passes reference
// 	if err != nil {
// 		return "", err
// 	}

// 	// Extract the value of the "location" field
// 	location := data.Location

// 	response, err := http.Get(location)
// 	if err != nil {
// 		fmt.Println("err:", err)
// 	}

// 	defer response.Body.Close()

// 	content, _ := io.ReadAll(response.Body)

// 	return location, nil
// }

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

type Origin struct {
	A       int `json:"a"`
	B       int `json:"b"`
	Average int
}

func main() {

	//123123123123
	var wg sync.WaitGroup
	var origin []Origin
	block := []int{}
	result := []int{}
	var sum int = 0

	content, err := ioutil.ReadFile("./input.json")
	if err != nil {
		fmt.Println("Cannot read file!")
	}

	err = json.Unmarshal([]byte(content), &origin)
	if err != nil {
		fmt.Println("Cannot parse file!")
	}

	for i, value := range origin {
		block = append(block, value.A, value.B)
		if i%100000 == 0 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				sum = 0
				for _, v := range block {
					sum += v
				}
				result = append(result, sum/len(block))

			}()
			wg.Wait()
			block = nil
		}
	}

	sum = 0
	for _, v := range result {
		sum += v
	}

	fmt.Println(sum / len(result))
}

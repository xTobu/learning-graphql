package helpers

import (
	"encoding/json"
	"fmt"
	"log"
)

func OutputGormResult(data interface{}) {

	a, err := json.Marshal(data) //get json byte array
	if err != nil {
		log.Fatalf("OutputGormResult err: %v", err)
	}
	n := len(a)        //Find the length of the byte array
	s := string(a[:n]) //convert to string
	fmt.Println(s)
}

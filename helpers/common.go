package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Common ...
type Common struct {
	NameType    string
	VersionType int
}

// TestGreater ...
func TestGreater(a int, b int) {
	if a > b {
		fmt.Println(a, "is greater")
	} else {
		fmt.Println(b, "is greater")
	}
}

// RandomNum ...
func RandomNum(n int) int {

	rand.Seed(time.Now().UnixNano())
	val := rand.Intn(n)
	return val
}

// Person ...
// struct for json unmarshalling...
type Person struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	HairColour string `json:"hair_colour"`
	HasDog     bool   `json:"has_dog"`
}

// UnmarshallToStructFromJSON ...
func UnmarshallToStructFromJSON() {

	myJSON := ` 
		[
			{
				"first_name": "Clark",
				"last_name": "Kent",
				"hair_colour":"Black",
				"has_dog": true
			},
			{
				"first_name": "Bruce",
				"last_name": "Wayne",
				"hair_colour":"Black",
				"has_dog": false
			}
		]
	`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJSON), &unmarshalled)

	if err != nil {
		log.Fatal("Error unmarshalling", err)
	} else {
		log.Printf("Unmarshalled : %v", unmarshalled)
	}

}

// MarshallToJSON ...
func MarshallToJSON() {

	var p1 Person
	p1.FirstName = "Peter"
	p1.LastName = "Parker"
	p1.HairColour = "Black"
	p1.HasDog = false

	var p2 Person
	p2.FirstName = "Diana"
	p2.LastName = "Prince"
	p2.HairColour = "Black"
	p2.HasDog = true

	var persons []Person

	persons = append(persons, p1, p2)

	myJSON, err := json.MarshalIndent(persons, "", "   ")

	if err != nil {
		log.Fatal("error marshaling to json", err)
	} else {
		fmt.Printf("Marshalled json : %v", string(myJSON))
	}

}

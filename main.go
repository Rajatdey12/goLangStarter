package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/Rajatdey12/goLangStarter/helpers"
)

var whtsUp string = "fella, tellya!"

func main() {
	fmt.Print("**************************************\n")
	fmt.Println("*                                    *")
	fmt.Println("*                                    *")
	fmt.Println("* LET US START TRAINING IN GOLANG!!! *")
	fmt.Println("*                                    *")
	fmt.Println("*                                    *")
	fmt.Print("**************************************\n")

	// whatToSay := "Bye, cruel world!"
	// fmt.Println(whtsUp)
	// fmt.Println(whatToSay)

	// fmt.Println("What was said", something("Ravi"))

	// firstVal, secondval := returnTwoVals()

	// fmt.Println("The two vals are ", firstVal, secondval)
	// x := []int{3, 2, 1, 0, 4, 5}
	// swap(x)
	// calc([]int{1, 1, 1, 2, 3, 4, 1, 5})
	// structTest()
	// customMap()
	// sliceFunc()
	// iterationWithRange()
	// testInterfaces()
	// printExtPackage()
	// helpers.TestGreater(9, 10)

	intChan := make(chan int)
	defer close(intChan)

	/* Executing this function as goroutine */
	go calculateRandom(intChan)
	num := <-intChan
	log.Println(num)
}

func something(name string) string {
	return "Something na! " + name
}

func returnTwoVals() (string, string) {
	return "abc", "def"
}

func swap(sw []int) {

	for a, b := 0, len(sw)-1; a < b; a, b = a+1, b-1 {

		sw[a], sw[b] = sw[b], sw[a]
	}
	fmt.Print(sw)

}

func calc(data []int) {
	myMap := make(map[int]int)
	for i := 0; i < len(data); i++ {
		if myMap[data[i]] == 0 {
			myMap[data[i]] = 1
		} else {
			myMap[data[i]] = myMap[data[i]] + 1
		}
	}
	fmt.Println("The count for the elements is : ", myMap)
}

/* Test structs & functions */
type person struct {
	firstName string
	lastname  string
}

// Receiver function to pass a struct as function
func (m *person) receiverFunc() string {
	return m.firstName
}

func structTest() {

	person := person{
		firstName: "John",
	}
	fmt.Println("The firstname is :", person.firstName)
	fmt.Println("The firstname is :", person.receiverFunc())
}

func customMap() {
	personMap := make(map[string]person)

	person1 := person{
		firstName: "Rajat",
		lastname:  "Dey",
	}

	personMap["firstPerson"] = person1

	fmt.Println(personMap["firstPerson"].firstName)

}

func sliceFunc() {

	// var mySlc []string
	// mySlc = append(mySlc, "Trevor", "Mike", "Hussain", "John")

	mySlc := []string{"Trevor", "Mike", "Hussain", "John"}
	sort.Strings(mySlc)
	fmt.Println(mySlc)
	fmt.Println(mySlc[0:2])
	fmt.Println(len(mySlc))
	fmt.Println(mySlc[len(mySlc)-1])

	person1 := person{
		firstName: "Rajat",
		lastname:  "Dey",
	}

	person2 := person{
		firstName: "Raghav",
		lastname:  "Juniyal",
	}

	mySlc1 := []person{person1, person2}

	fmt.Println(mySlc1)
}

func iterationWithRange() {
	/* Range over slices */
	animals := []string{"dog", "cat", "cow", "giraffe", "yak", "snail"}

	for i, v := range animals {
		log.Printf("The iteration/index is %d and the respective value is %v", i, v)
	}

	/* range over maps */
	animalMap := make(map[string]string)

	animalMap["dog"] = "Fido"
	animalMap["cat"] = "fluffy"

	for k, v := range animalMap {
		log.Println(k, v)
	}
}

// Understanding interfaces --

type animal interface {
	voice() string
	noOfLegs() int
}

type dog struct {
	name   string
	colour string
}

type gorilla struct {
	noOfTeeth int
	name      string
	colour    string
}

func testInterfaces() {

	myDog := dog{
		name:   "Samson",
		colour: "Black",
	}
	printFeaturesOfAnimal(&myDog)

	myGorrilla := gorilla{
		noOfTeeth: 38,
		name:      "jack",
		colour:    "Grey",
	}
	printFeaturesOfAnimal(&myGorrilla)
}

func (d *dog) voice() string {
	return "woof"
}

func (d *dog) noOfLegs() int {
	return 4
}

func (g *gorilla) noOfLegs() int {
	return 2
}

func (g *gorilla) voice() string {
	return "aaagghhhh"
}

func printFeaturesOfAnimal(a animal) {
	log.Println("The animal has a voice", a.voice(), "and the legs are", a.noOfLegs())
}

func printExtPackage() {

	commonType := helpers.Common{
		NameType:    "Krypton",
		VersionType: 665688,
	}

	fmt.Println(commonType.NameType, commonType.VersionType)
}

// Understanding channels.....
func calculateRandom(intChan chan int) {
	const numPool = 100
	randomNum := helpers.RandomNum(numPool)
	intChan <- randomNum
}

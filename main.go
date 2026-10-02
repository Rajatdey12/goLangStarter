package main

import (
	"errors"
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

	// intChan := make(chan int)
	// defer close(intChan)

	// /* Executing this function as goroutine */
	// go calculateRandom(intChan)
	// num := <-intChan
	// log.Println(num)

	/*marshalling & unmarshalling*/
	// helpers.UnmarshallToStructFromJSON()
	// helpers.MarshallToJSON()
	result, err := divide(100.0, 10.0)

	if err != nil {
		log.Println(err)
		return
	}
	log.Println(`The result is :::`, result)

	testSyntax()
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

/* Map in go */
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

func divide(x, y float32) (float32, error) {

	var res float32

	if y == 0 {
		return res, errors.New("Cannot divide by 0")
	}
	res = x / y

	return res, nil
}

func testSyntax() {

	x := 10
	y := 20

	var a int = 5
	var b int = 15

	var c int = a + b
	fmt.Printf("The sum of %d and %d is %d\n", a, b, c)

	sum := x + y
	fmt.Printf("The sum of %d and %d is %d\n", x, y, sum)
	fmt.Println("Hello, World!")

	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	for i := range []string{"Go", "Python", "Java"} {
		fmt.Printf("Language: %s\n", []string{"Go", "Python", "Java"}[i])
	}

	for i, lang := range []string{"Go", "Python", "Java"} {
		fmt.Printf("Language %d: %s\n", i, lang)
	}

	fmt.Println(`Old FMT======================================================`)

	for i, j := 0, 10; i < 5; i, j = i+1, j-1 {
		fmt.Printf("i: %d, j: %d\n", i, j)
	}

	i := 0
	fmt.Println(`New Fmt======================================================`)
	for j := 10; j > 0; j-- {
		fmt.Printf("i: %d, j: %d\n", i, j)
		i++
	}

	for i, j := range []int{1, 2, 3, 4, 5} {
		fmt.Printf("Index: %d, Value: %d\n", i, j)
	}

	for i := 0; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	fruits := [3]string{"apple", "orange", "banana"}
	animals := [3]string{"Cat", "Dog", "Monkey"}

	printFruits()

	testSwitch(4)

	testRangeFunc(fruits, animals)

}

func printFruits() {

	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}

	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}
}

func testSwitch(val int) {
	value := `Value is`
	switch val {
	case 1:
		fmt.Println("value is 1")
	case 2:
		fmt.Println("Value is 2")
	default:
		fmt.Println(value + ` default`)
	}
}

func testRangeFunc(fruits [3]string, anim [3]string) {
	for idx, val := range fruits {
		fmt.Printf("for  %v the index is %v and the type is %T and %T\n", val, idx, val, idx)
	}
	for idx, val := range anim {
		fmt.Printf(`for  %v the index is %v and the type is %T and %T `, val, idx, val, idx)
	}
}

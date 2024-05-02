# go-ecommerce-app
Series : Learn go


###
```sh 
go mod init go-ecommerce-app
```
```sh
mkdir configs infra internal pkg
```
```sh
touch main.go
```

### Basic Types: int, float64, string, bool
```go

	// var age int
	// var height float64
	// var firstName string
	// var isEmployed bool

	// age = 29
	// height = 185
	// firstName = "Nam"
	// isEmployed = true

	age := 29
	height := 185.4
	firstName := "Nam"
	isEmployed := true

	// fmt.Println(age, height, firstName, isEmployed)

	fmt.Printf("Age: %d \n", age)
	fmt.Printf("FirstName: %s \n", firstName)
	fmt.Printf("Height: %f \n", height)
	fmt.Printf("Employer?: %t \n", isEmployed)

	if age > 65 {
		fmt.Println("Senior Software Engineer")
	} else if age > 10 {
		fmt.Println("Adult")
	}
```
### Composite Types: array, slice
```go
    var myFamily [3]string
	myFamily[0] = "Nam"
	myFamily[1] = "Nhinh"
	myFamily[2] = "Tran"

	myFamily := [3]string{"Nam", "Tran", "Tien"}
	myFamily[2] = "Kate"

	fmt.Println("My Family: %v", myFamily)

	// Slice
	var myFriends []string
	myFriends = append(myFriends, "Mike", "Adam")
	myFriends = append(myFriends, "Kira")
	fmt.Println("My Friends: %v", myFriends)

	myCourses := [3][2]string{
		{"Go", "NodeJs"},
	 	{"AWS", "Azure"},
		{"CDK", "GPC"},
	}
	fmt.Println("Available Courses %v", myCourses)

	mySliceCourses := [][]string{
		{"Go", "NodeJs"},
	 	{"AWS", "Azure"},
	 	{"CDK", "GPC"},
	}

	course := []string{"AWS Cloud", "Azure Cloud"}

	mySliceCourses = append(mySliceCourses, course)

	fmt.Println("Available Slice Courses %v", mySliceCourses)

	fmt.Println(len(mySliceCourses))
```
### Composite Types: map, struct
```go
    myWishlist := make(map[string]string)
	myWishlist["first"] = "MacPro"
	myWishlist["second"] = "900 Billion Dollar"
	myWishlist["third"] = "a beautiful car"

	delete(myWishlist, "third")

	firstWish := myWishlist["first"]

	log.Println(firstWish)

	fmt.Println("My wish list %v", myWishlist)

	// struct

	type Details struct {
		Description string `json:"description"`
		Images      string `json:"images"`
	}

	type Product struct {
		Name    string  `json:"product_name"`
		Price   float64 `json:"price"`
		Details `json:"details"`
	}

	// var product Product

	// product = Product{
	// 	Name:  "MacPro",
	// 	Price: 9000,
	// }

	product := Product{
		Name:  "MacPro",
		Price: 9000,
		Details: Details{
			Description: "An incredible machine",
			Images:      "http://macproimage.jpg",
		},
	}
	fmt.Println("Product struct: %v", product)
```
### Pointer
```go
    /*
		- if else
		- switch case
		- select
	*/

	age := 29
	if age > 65 {
		fmt.Println("Senior ")
	} else if age > 10 {
		fmt.Println("Action")
	} else {
		fmt.Println("Khahaha")
	}

	seatClass := "FirstClass"

	switch seatClass {
	case "FirstClass":
		fmt.Println("You will get free drinks")
	case "BusinessClass":
		fmt.Println("You will get more legrooms")
	default:
		fmt.Println("You need to pay for services")
	}

	var myFriends []string
	for i := 0; i < 10; i++ {
		myNewFriend := fmt.Sprintf("Friend %d", i)
		myFriends = append(myFriends, myNewFriend)
	}

	for index, value := range myFriends {
		fmt.Println(index, value)
	}

	// isOver := 0
	// for {
	// 	isOver++
	// 	fmt.Println(isOver)
	// 	if isOver > 99 {
	// 		fmt.Println("It's really over now")
	// 		return
	// 	}
	// }

	// Jay => Laptop

	// Guest => Jay => Laptop
	jay := "laptop" // 0xc0000ac260
	fmt.Println(&jay)

	var guest *string // <nil>
	guest = &jay      // 0xc0000ac260

	fmt.Println(*guest)
```
### Func
```go
package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// // Basic Types: int, float64, string, bool
	// // Composite Types: array, slice, map, struct
	// // Pointer types: *

	sayHello()
	var name = getUserName()

	id, age := getUserById(1, 16)

	fmt.Println("Name: %v", name)
	fmt.Println("User: %s, Age: %d", id, age)

	fmt.Println("Total Amount: %f", calculateTotal(1.2, 4.4, 5.1, 7.4, 8.2, 9.1))

	concateUserName := func(fname string, lname string) string {
		return fmt.Sprintf("%s %s", fname, lname)
	}

	fmt.Printf("User FullName is: %s", concateUserName("Nam", "Tran"))

	log.Fatal(app.Listen(":8000"))
}

func sayHello() {
	fmt.Println("Hello World!")
}

func getUserName() string {
	return "My Name is Nam"
}

func getUserById(id, ageInput int) (name string, age int) {
	println(id, ageInput)
	return "Jay", 36
}

func calculateTotal(products ...float64) float64 {
	totalAmount := 0.0
	for _, price := range products {
		totalAmount += price
	}
	return totalAmount
}

```
```go
package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	Name  string
	Price float64
	Stock int
}

func (p *Product) Calculate(qty int) float64 {
	return p.Price * float64(qty)
}

func (p *Product) ReduceStock(qty int) {
	if p.Stock >= qty {
		p.Stock -= qty
	}
}

func main() {

	app := fiber.New()

	p := Product{
		Name:  "Macbook Pro",
		Price: 9000,
		Stock: 1,
	}

	fmt.Printf("Total Amount %f", p.Calculate(2))

	p.ReduceStock(1)
	fmt.Println(p)

	log.Fatal(app.Listen(":8000"))
}

```
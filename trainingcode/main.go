package main

import (
	"fmt"

	"github.com/pluralsight/webservice/models"
	//add more packages on new lines in here as shown below. no commas needed.
	//"os"
)

/* The main function, when part of the main package, identifies  the entry point
of an application.
*/

//for iota, constants must be at package level.
const (
	longerpi = 3.14159
	first    = 1
	second   = "second"
	one      = iota
	two      = iota
	eleven   = iota + 6
	timestwo = 2 << iota // 2 << 6 = 2x(iota i.e. 6 times two i.e. 2x 2x 2x 2x 2x x2) = 128. see https://bit-calculator.com/bit-shift-calculator for 2 << 6
	seven                // assumes incremented iota with last used expression i.e. 2 << 7 this time which is 256 this is a constant expression

	// note that iota resets in between new constant blocks i.e. every new constant block.
)

func main() {

	//constants
	const pi = 3.1415 //constants delcaration and initialization has to be on same line. the type must be inferable - it's inferred at the point where it's used.
	fmt.Println(pi)
	// pi = 1.2459 - errors if we try to redeclare

	// implicitly typed constant - compiler intterprets a  type every time it's used.
	const somenum = 3
	// it's inferred somenum as int by this point.
	fmt.Println(somenum + 3)
	// here somenum is inferred as a float
	fmt.Println(somenum + 1.2)

	// explicitly typed constant.
	const typednum int = 3
	// fmt.Println(typednum + 1.2) //not allowed without type conversion
	fmt.Println(float32(typednum) + 1.2) //allowed

	// iota and constants
	fmt.Println(longerpi, first, second, one, two, eleven, timestwo, seven)
	//prints "3.14159 1 second 3 4 11 128". note that one and two were iota type. it just prints its position value from zero index.

	// variables
	var intnum int
	intnum = 45

	fmt.Printf("hello from module go, critters\n")
	fmt.Println(intnum)

	var f float32 = 3.141 //or float64
	fmt.Println(f)

	// using := forces implicit initializattion.
	name := "arthur"
	fmt.Println(name)
	//c is a complex type
	c := complex(3, 4)
	fmt.Println(c) //prints 3 + 4i
	// multiple declarations. r is real no, img is imaginary no.
	r, img := real(c), imag(c)
	fmt.Println(r, img) //prints 3 4

	// working with pointers
	var firstName *string = new(string) //firstname is a pointer to a string
	fmt.Println(firstName)              //prints <nil> if new(string) wasnt used with the pointer in line above OR address like 0xc00010a230 if memory was allocated.
	*firstName = "Arthur"
	fmt.Println(*firstName) // errors saying invalid memory address or nil pointer dereference if new(string) wasnt used with the pointer

	// address operator
	addressOfName := &name
	fmt.Println(addressOfName, *addressOfName) //prints 0xc000096220 arthur
	// modify content i.e. arthur to tricia. pointers point to memory so we see new values auttomatically.
	name = "tricia"
	fmt.Println(addressOfName, *addressOfName)

	// Collections - linkedlists etc not built into the language. below are some of the key ones built into the language.
	// Arrays - fixed size
	// Slices - dynamically sizing collection
	// maps - like slices, dynamic but has key-value pairs
	// structs - the data side of a class ypically defining one concept. eg: accountholder
	// there's no class concept in go.

	var arr [3]int //arr is a array of 3 elements that are integers
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr[0], arr[1], arr[2])
	//arrays are bounded and there's compile-time bounds checking.

	// array - implicit initialization syntax
	//all of the above can be done in one line like so to declare an array and assign:
	arr2 := [3]int{1, 2, 3}

	fmt.Println(arr2[0], arr2[1], arr2[2])

	// Slices - they're built on top of arrays

	arr3 := [3]int{1, 2, 3}
	slice := arr3[:]         // : means beginning to end of  the array arr
	fmt.Println(arr3, slice) //prints [1 2 3] [1 2 3]
	arr3[1] = 42
	slice[2] = 27
	fmt.Println(arr3, slice) // prints [1 42 27] [1 42 27]
	//slice's values are related to the array - think of it as like a pointer. any updates to a slice updates array itself and vice versa.

	// Now try creating a slice without an underlying array - this means GO will internally create and manage the underlying array.
	slice2 := []int{1, 2, 3}
	fmt.Println(slice2) // prints [1 2 3]

	slice2 = append(slice2, 4, 5, 6)
	fmt.Println(slice2) //prints [1 2 3 4 5 6]

	slice3 := slice[1:]
	slice4 := slice[:2]
	slice5 := slice[1:2]

	fmt.Println(slice3, slice4, slice5)

	// Using map data type
	m := map[string]int{"age": 38} //maps strings to integers. age:38 is implicit initialization.
	fmt.Println(m)                 //prints map[age:38]
	fmt.Println(m["age"])          // prints 38
	m["age"] = 39
	fmt.Println(m) // prints map[age:39]
	//deleting elements from a map
	delete(m, "age")
	fmt.Println(m) // prints map[]
	//add a new item to map m
	m["year"] = 2020 // prints map[year:2020]
	fmt.Println(m)
	//update a map
	m["year"] = 2021
	fmt.Println(m) // prints map[year:2021]

	//make a string to float map
	m2 := make(map[string]float64)

	m2["pi"] = 3.14   // Add a new key-value pair
	m2["pi"] = 3.1416 // Update value
	fmt.Println(m)    // Print map: "map[pi:3.1416]"

	v := m2["pi"] // Get value: v == 3.1416
	v = m2["pie"] // Not found: v == 0 (zero value)
	fmt.Println(v)
	_, found := m2["pi"] // found == true
	fmt.Println(found)
	_, found = m2["pie"] // found == false
	fmt.Println(found)

	if x, found := m2["pi"]; found {
		fmt.Println(x)
	} // Prints "3.1416"

	delete(m2, "pi") // Delete a key-value pair
	fmt.Println(m2)  // prints map[]
	// looping through maps
	m3 := map[string]float64{
		"pi": 3.1416,
		"e":  2.71828,
	}
	fmt.Println(m3) // "map[e:2.71828 pi:3.1416]"

	for key, value := range m3 { // Order not specified
		fmt.Println(key, value)
	} // prints pi 3.1416
	//        e 2.71828

	// ===============================================================================
	// Structs - allows grouping together disparate data types unlike arrays and maps.
	// There's no concept of classes in go lang. use structs instead.
	// two step - define the struct and then initialize object using that struct's definition.
	// user 'type' keyword to define a struct.

	//Struct can be defined either within main /any function or higher up at the package level outside of main.
	type LocalUser struct {
		ID        int
		FirstName string
		LastName  string
	}

	var usr LocalUser
	usr.ID = 1
	usr.FirstName = "Ace"
	usr.LastName = "McCloud"
	fmt.Println(usr)
	//if struct variable is uninitialised prints '{0  }' because default value for int is 0 and for the two strings it's two empty strings.
	//else prints structs value i.e. '{1 Ace McCloud}'

	//inline initialization of a struct
	usr2 := LocalUser{ID: 2, FirstName: "Thunder", LastName: "Cats"}
	fmt.Println(usr2) // prints '{2 Thunder Cats}'

	//multi line struct initialization - go compiler inserts semicolon - not going to work till removed as it things struct ends before closing curly braces
	//way to fix it - end it with a comma after "cats" below.
	usr3 := LocalUser{ID: 2,
		FirstName: "Wonder",
		LastName:  "Cats",
	}
	fmt.Println(usr3) // {2 Wonder Cats}
	//Therefore in Go, always terminate multi-line entries with a comma to make it consistent with the rest of the elements.

	// MODULES and PACKAGES
	// *There are Go modules and within a module there are packages. Packages are discrete units of code that are associated with some sort of a concept within a module.
	// In a go project, create a directory and that declares a package for us.
	// If you need to interact with the package, the name of the folder matters.
	//Add folder named models then add a file in it called user.go
	//normally in go.mod you'd need to add module github.com/pluralsight/webservice/models but go takes care of all that for us.

	//create variable of type user which is defined in module package user.go
	u := models.User{ID: 1, FirstName: "Lee", LastName: "Bruce"}
	fmt.Println(u)
}

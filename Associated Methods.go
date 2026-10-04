package main

import (
	"bufio"
	"fmt"

	"os"
	"strings"
)

// Person defines the structure for individual records
type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

// ReadData is a pointer receiver method (*Person) because it modifies the struct fields.
func (p *Person) ReadData(reader *bufio.Reader) {
	fmt.Print("Enter Name: ")
	name, _ := reader.ReadString('\n')
	p.Name = strings.TrimSpace(name)

	fmt.Print("Enter Age: ")
	fmt.Scanln(&p.Age)

	// Consume the remaining newline character from Scanln
	reader.ReadString('\n')

	fmt.Print("Enter Job: ")
	job, _ := reader.ReadString('\n')
	p.Job = strings.TrimSpace(job)

	fmt.Print("Enter Salary: ")
	fmt.Scanln(&p.Salary)

	// Consume the remaining newline character from Scanln
	reader.ReadString('\n')
}

// PrintDetails is a value receiver method (Person) as it only reads and displays the struct fields.
func (p Person) PrintDetails() {
	fmt.Println("----------------------------------------")
	fmt.Printf("Name   : %s\n", p.Name)
	fmt.Printf("Age    : %d\n", p.Age)
	fmt.Printf("Job    : %s\n", p.Job)
	fmt.Printf("Salary : $%.2f\n", p.Salary)
	fmt.Println("----------------------------------------")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Create first Person object
	fmt.Println("=== Enter details for Person 1 ===")
	var p1 Person
	p1.ReadData(reader)

	// Create second Person object
	fmt.Println("\n=== Enter details for Person 2 ===")
	var p2 Person
	p2.ReadData(reader)

	// Display details using PrintDetails method
	fmt.Println("\n=== Person Details ===")
	p1.PrintDetails()
	p2.PrintDetails()
}

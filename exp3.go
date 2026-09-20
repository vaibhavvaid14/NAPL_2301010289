package main

import "fmt"

func main() {
	fmt.Println("==================================================")
	fmt.Println(" PART 1: SLICE OPERATIONS (Student Names)")
	fmt.Println("==================================================")

	// Initial Slice
	students := []string{"Alice", "Bob", "Charlie", "David"}
	fmt.Printf("Initial Slice:            %v\n\n", students)

	// 1. ADD OPERATION
	newStudent := "Eve"
	students = append(students, newStudent)
	fmt.Printf("[ADD] Added '%s':         %v\n", newStudent, students)

	// 2. UPDATE OPERATION
	updateIndex := 1
	updatedName := "Robert"
	if updateIndex >= 0 && updateIndex < len(students) {
		students[updateIndex] = updatedName
		fmt.Printf("[UPDATE] Index %d -> '%s': %v\n", updateIndex, updatedName, students)
	}

	// 3. REMOVE-BY-INDEX OPERATION
	removeIndex := 2 // Removes "Charlie"
	if removeIndex >= 0 && removeIndex < len(students) {
		removedElement := students[removeIndex]
		students = append(students[:removeIndex], students[removeIndex+1:]...)
		fmt.Printf("[REMOVE] Index %d ('%s'): %v\n", removeIndex, removedElement, students)
	}

	fmt.Println("\n==================================================")
	fmt.Println(" PART 2: MAP OPERATIONS (Subject -> Marks)")
	fmt.Println("==================================================")

	// Initial Map
	marks := map[string]int{
		"Math":    90,
		"Physics": 85,
	}
	fmt.Printf("Initial Map:                   %v\n\n", marks)

	// 1. INSERT OPERATION
	marks["Chemistry"] = 92
	fmt.Printf("[INSERT] Subject 'Chemistry':  %v\n", marks)

	// 2. UPDATE OPERATION
	marks["Physics"] = 88
	fmt.Printf("[UPDATE] Subject 'Physics':    %v\n", marks)

	// 3. LOOKUP OPERATION
	searchSubject := "Math"
	if val, exists := marks[searchSubject]; exists {
		fmt.Printf("[LOOKUP] Found '%s': %d marks\n", searchSubject, val)
	} else {
		fmt.Printf("[LOOKUP] '%s' not found\n", searchSubject)
	}

	// 4. DELETE OPERATION
	deleteSubject := "Physics"
	delete(marks, deleteSubject)
	fmt.Printf("[DELETE] Subject '%s':  %v\n", deleteSubject, marks)

	fmt.Println("\nFinal Map State:               ", marks)
}
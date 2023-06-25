package main

// All the imports used
import "fmt"
import "math"
import "time"


// Nursery struct
type Nursery struct {
	children []Child
	teachers []Teacher
	schedule Schedule
}

// Child struct
type Child struct {
	name string
	age int
	likes []string
	dislikes []string
}

// Teacher struct
type Teacher struct {
	name string
	qualifications []string
}

// Schedule struct
type Schedule struct {
	days []string
	start time.Time
	end timeout.Time
}

func (n * Nursery) enrollChild(name string, age int, likes []string, dislikes []string) {
	// Create a Child object with given params
	child := Child{
		name: name,
		age: age,
		likes: likes,
		dislikes: dislikes,
	}
	// Append the child to the list of children
	n.children = append(n.children, child)
}

func (n * Nursery) addTeacher(name string, qualifications []string) {
	// Create a Teacher object with given params
	teacher := Teacher{
		name: name,
		qualifications: qualifications,
	}
	// Append the teacher to the list of teachers
	n.teachers = append(n.teachers, teacher)
}

func (n * Nursery) setSchedule(days []string, start time.Time, end time.Time) {
	// Create a Schedule object with given params
	schedule := Schedule{
		days: days,
		start: start,
		end: end,
	}
	// Set the nursery schedule
	n.schedule = schedule
}

func (n * Nursery) getAges() []int {
	// Create a slice for storing the ages
	ages := []int{}

	// Iterate over all the children
	for _, c := range n.children {
		// Append the age of each child to ages
		ages = append(ages, c.age)
	}

	// Return the list of ages
	return ages
}

func (n * Nursery) getAverageAge() float64 {
	// Get the list of ages
	ages := n.getAges()

	// Initialize the sum and count
	sum := 0.0
	count := 0.0

	// Iterate over the list of ages
	for _, age := range ages {
		// Add the age to sum and increment the count
		sum += float64(age)
		count += 1.0
	}

	// Calculate and return the average age
	return math.Ceil(sum / count)
}

func main() {
	// Create a Nursery object
	nursery := Nursery{}

	// Enroll a few children into the nursery
	nursery.enrollChild("John", 5, []string{"painting", "drawing"}, []string{"maths"})
	nursery.enrollChild("Mary", 6, []string{"dancing", "swimming"}, []string{"history"})

	// Add a few teachers to the nursery
	nursery.addTeacher("Alice", []string{"early childhood education"})
	nursery.addTeacher("Bob", []string{"drama", "music"})

	// Set the nursery schedule
	nursery.setSchedule([]string{"Monday", "Tuesday", "Wednesday"}, time.Now(), time.Now().Add(2*time.Hour))

	// Print the average age of children
	fmt.Println("The average age of children in the nursery is:", nursery.getAverageAge())
}
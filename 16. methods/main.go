package main

import "fmt"

func main() {
	// 1. Your original example
	fmt.Println("=== 1. Basic User Example ===")
	user1 := User{"Utsho", "utshodey.tech@gmail.com", true, 26}
	fmt.Printf("User details: %+v\n", user1)
	user1.GetStatus()
	user1.NewMail()

	// 2. Simple Counter example
	fmt.Println("\n=== 2. Counter Example ===")
	counter := Counter{value: 0}
	counter.Increment()
	counter.Increment()
	counter.Display()
	counter.Decrement()
	counter.Display()

	// 3. Student Grade example
	fmt.Println("\n=== 3. Student Grade Example ===")
	student := Student{
		Name:   "Alex",
		Grades: []int{85, 92, 78, 95},
	}
	student.DisplayGrades()
	fmt.Printf("Average grade: %.2f\n", student.CalculateAverage())

	// 4. Rectangle example
	fmt.Println("\n=== 4. Rectangle Example ===")
	rect := Rectangle{width: 5, height: 3}
	rect.Display()
	fmt.Printf("Area: %d\n", rect.GetArea())
	rect.Scale(2)
	fmt.Printf("After scaling - Area: %d\n", rect.GetArea())
}

// 1. Original User struct and methods
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("New email of this user is", u.Email)
}

// 2. Simple Counter
type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) Decrement() {
	c.value--
}

func (c Counter) Display() {
	fmt.Printf("Counter value: %d\n", c.value)
}

// 3. Student with Grades
type Student struct {
	Name   string
	Grades []int
}

func (s Student) DisplayGrades() {
	fmt.Printf("Grades for %s: %v\n", s.Name, s.Grades)
}

func (s Student) CalculateAverage() float64 {
	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.Grades))
}

// 4. Rectangle with simple methods
type Rectangle struct {
	width  int
	height int
}

func (r Rectangle) Display() {
	fmt.Printf("Rectangle: %d x %d\n", r.width, r.height)
}

func (r Rectangle) GetArea() int {
	return r.width * r.height
}

func (r *Rectangle) Scale(factor int) {
	r.width *= factor
	r.height *= factor
}

/*
Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s =½ a t2 + vot + so

Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo,
and initial displacement so. GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given
values acceleration, initial velocity, and initial displacement. The function returned by GenDisplaceFn() should take one float64 argument t,
representing time, and return one float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial velocity, and
initial displacement: a = 10, vo = 2, so = 1. I can use the following statement to call GenDisplaceFn() to generate a function fn
which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5))

Submit your Go program source code. */
package main

import "fmt"

func main() {
	var (
		acceleration        float64
		initialVelocity     float64
		initialDisplacement float64
		time                float64
	)
	fmt.Print("Enter the value for acceleration: ")
	fmt.Scan(&acceleration)
	fmt.Print("Enter the value for initial velocity: ")
	fmt.Scan(&initialVelocity)
	fmt.Print("Enter the value for initial displacement: ")
	fmt.Scan(&initialDisplacement)
	fmt.Print("Enter the value for time: ")
	fmt.Scan(&time)
	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Print("The displacment is ")
	fmt.Println(fn(time))
}

// GenDisplaceFn is a function which returns a function for computing displacement
func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return ((0.5 * acceleration) * (time * time)) + (time * initialVelocity) + initialDisplacement
	}
}

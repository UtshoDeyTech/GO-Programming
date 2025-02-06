package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"time"
)

func main() {
	// Initialize and clear screen
	fmt.Print("\033[H\033[2J")
	fmt.Println("\033[1;36m=== Advanced Mathematics in Go ===\033[0m")
	fmt.Println("\033[1;33mExploring mathematical operations and calculations\033[0m")
	time.Sleep(1 * time.Second)

	// 1. Basic Arithmetic Operations
	fmt.Println("\n\033[1;32m1. BASIC ARITHMETIC\033[0m")
	fmt.Println("------------------")
	x, y := 25, 7
	fmt.Printf("Numbers: x = %d, y = %d\n", x, y)
	fmt.Printf("Addition:       %d + %d = %d\n", x, y, x+y)
	fmt.Printf("Subtraction:    %d - %d = %d\n", x, y, x-y)
	fmt.Printf("Multiplication: %d * %d = %d\n", x, y, x*y)
	fmt.Printf("Division:       %d / %d = %.3f\n", x, y, float64(x)/float64(y))
	fmt.Printf("Modulus:        %d %% %d = %d\n", x, y, x%y)
	fmt.Printf("Integer Division: %d / %d = %d\n", x, y, x/y)

	// 2. Advanced Math Functions
	fmt.Println("\n\033[1;32m2. ADVANCED MATH FUNCTIONS\033[0m")
	fmt.Println("------------------------")
	num := 16.0
	fmt.Printf("Number: %.2f\n", num)
	fmt.Printf("Square Root:    √%.0f = %.3f\n", num, math.Sqrt(num))
	fmt.Printf("Cube Root:      ∛%.0f = %.3f\n", num, math.Cbrt(num))
	fmt.Printf("Power:          %.0f² = %.0f\n", num, math.Pow(num, 2))
	fmt.Printf("Power:          %.0f³ = %.0f\n", num, math.Pow(num, 3))
	fmt.Printf("Natural Log:    ln(%.0f) = %.3f\n", num, math.Log(num))
	fmt.Printf("Log10:          log₁₀(%.0f) = %.3f\n", num, math.Log10(num))
	fmt.Printf("Exp:           e^2 = %.3f\n", math.Exp(2))

	// 3. Trigonometry
	fmt.Println("\n\033[1;32m3. TRIGONOMETRY\033[0m")
	fmt.Println("---------------")
	angles := []float64{0, 30, 45, 60, 90}
	fmt.Println("Common Angles:")
	for _, deg := range angles {
		rad := deg * math.Pi / 180
		fmt.Printf("\nAngle: %.0f°\n", deg)
		fmt.Printf("Sin(%.0f°) = %.3f\n", deg, math.Sin(rad))
		fmt.Printf("Cos(%.0f°) = %.3f\n", deg, math.Cos(rad))
		fmt.Printf("Tan(%.0f°) = %.3f\n", deg, math.Tan(rad))
	}

	// 4. Random Numbers using crypto/rand
	fmt.Println("\n\033[1;32m4. CRYPTOGRAPHIC RANDOM NUMBERS\033[0m")
	fmt.Println("-----------------------------")
	fmt.Println("Generating secure random numbers:")
	for i := 0; i < 5; i++ {
		randNum, _ := rand.Int(rand.Reader, big.NewInt(100))
		fmt.Printf("Random %d: %v\n", i+1, randNum)
	}

	// 5. Working with Big Numbers
	fmt.Println("\n\033[1;32m5. BIG NUMBER OPERATIONS\033[0m")
	fmt.Println("-----------------------")

	// Create big numbers
	bigNum1 := new(big.Int)
	bigNum2 := new(big.Int)
	bigNum1.SetString("9999999999999999999999999999999999", 10)
	bigNum2.SetString("1111111111111111111111111111111111", 10)

	// Perform operations
	result := new(big.Int)
	fmt.Println("Number 1:", bigNum1)
	fmt.Println("Number 2:", bigNum2)

	result.Add(bigNum1, bigNum2)
	fmt.Println("Addition:", result)

	result.Mul(bigNum1, bigNum2)
	fmt.Println("Multiplication:", result)

	result.Sub(bigNum1, bigNum2)
	fmt.Println("Subtraction:", result)

	// 6. Mathematical Constants and Precision
	fmt.Println("\n\033[1;32m6. MATHEMATICAL CONSTANTS\033[0m")
	fmt.Println("------------------------")
	fmt.Printf("π (Pi):           %.15f\n", math.Pi)
	fmt.Printf("e (Euler):        %.15f\n", math.E)
	fmt.Printf("φ (Golden Ratio): %.15f\n", math.Phi)
	fmt.Printf("√2:              %.15f\n", math.Sqrt2)
	fmt.Printf("ln(2):           %.15f\n", math.Ln2)
	fmt.Printf("log₂(e):         %.15f\n", math.Log2E)
	fmt.Printf("log₁₀(e):        %.15f\n", math.Log10E)

	// 7. Advanced Calculations
	fmt.Println("\n\033[1;32m7. ADVANCED CALCULATIONS\033[0m")
	fmt.Println("----------------------")

	// Area calculations
	radius := 5.0
	circle_area := math.Pi * math.Pow(radius, 2)
	circle_circumference := 2 * math.Pi * radius
	fmt.Printf("Circle (r=%.1f):\n", radius)
	fmt.Printf("  Area:         %.3f\n", circle_area)
	fmt.Printf("  Circumference: %.3f\n", circle_circumference)

	// 3D calculations
	side := 3.0
	cube_volume := math.Pow(side, 3)
	cube_surface := 6 * math.Pow(side, 2)
	fmt.Printf("\nCube (side=%.1f):\n", side)
	fmt.Printf("  Volume:        %.3f\n", cube_volume)
	fmt.Printf("  Surface Area:  %.3f\n", cube_surface)

	// 8. Number Analysis
	fmt.Println("\n\033[1;32m8. NUMBER ANALYSIS\033[0m")
	fmt.Println("------------------")
	testNum := -15.7
	fmt.Printf("Number: %.2f\n", testNum)
	fmt.Printf("Absolute:        %.2f\n", math.Abs(testNum))
	fmt.Printf("Ceiling:         %.2f\n", math.Ceil(testNum))
	fmt.Printf("Floor:           %.2f\n", math.Floor(testNum))
	fmt.Printf("Round:           %.2f\n", math.Round(testNum))
	fmt.Printf("Truncate:        %.2f\n", math.Trunc(testNum))

	// 9. High Precision Calculations
	fmt.Println("\n\033[1;32m9. HIGH PRECISION CALCULATIONS\033[0m")
	fmt.Println("----------------------------")

	// Calculate π with high precision
	pi := new(big.Float)
	pi.SetPrec(200)
	pi.SetString("3.14159265358979323846264338327950288419716939937510582097494459")

	// Calculate area of circle with high precision
	r := new(big.Float).SetFloat64(10.0)
	area := new(big.Float).Mul(pi, new(big.Float).Mul(r, r))
	fmt.Printf("High-precision circle area (r=10):\n%.50f\n", area)

	// Program completion
	fmt.Println("\n\033[1;36m=== Program Complete ===\033[0m")
	fmt.Println("\033[1;33mPress Enter to exit...\033[0m")
	fmt.Scanln()
}

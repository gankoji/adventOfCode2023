package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"os"
	"log"
	"strconv"
	"strings"
	"time"
)

type hailstone struct {
	x, y, z, dx, dy, dz float64
}

func parseEquations(in []string) ([]hailstone, error) {
	//280342553482089, 252447176665186, 220895081604913 @ 127, 12, 96
	var hailstones []hailstone
	for _, line := range in {
		parsedLine := strings.TrimSpace(strings.Replace(line, " ", "", -1))
		a := strings.Split(parsedLine, "@")
		if len(a) != 2 {
			return nil, errors.New("invalid input")
		}
		coordinates := strings.Split(a[0], ",")
		velocity := strings.Split(a[1], ",")
		if len(coordinates) != 3 || len(velocity) != 3 {
			return nil, errors.New("invalid input")
		}
		x, ex := strconv.Atoi(coordinates[0])
		y, ey := strconv.Atoi(coordinates[1])
		z, ez := strconv.Atoi(coordinates[2])
		if ex != nil || ey != nil || ez != nil {
			return nil, errors.New("invalid input")
		}
		dx, edx := strconv.Atoi(velocity[0])
		dy, edy := strconv.Atoi(velocity[1])
		dz, edz := strconv.Atoi(velocity[2])
		if edx != nil || edy != nil || edz != nil {
			return nil, errors.New("invalid input")
		}
		hailstones = append(hailstones, hailstone{
			x:  float64(x),
			y:  float64(y),
			z:  float64(z),
			dx: float64(dx),
			dy: float64(dy),
			dz: float64(dz),
		})
	}
	return hailstones, nil
}

func getSlope(h hailstone) float64 {
	return h.dy / h.dx
}

func areParallel(h1, h2 hailstone) bool {
	return getSlope(h1) == getSlope(h2)
}

func areCrossed(h1, h2 hailstone) bool {
	cx := (h2.y - (h2.dy/h2.dx)*h2.x - (h1.y - (h1.dy/h1.dx)*h1.x)) / (h1.dy/h1.dx - h2.dy/h2.dx)
	return !((cx > h1.x) == (h1.dx > 0) && (cx > h2.x) == (h2.dx > 0))
}

func calculateEquation(h1, h2 hailstone) (float64, float64) {
	if areParallel(h1, h2) {
		return -1, -1
	}
	// Line Equation X*dy - Y*dx = x*dy -y*dx
	if areCrossed(h1, h2) {
		return -1, -1
	}
	// Solve equation
	y := ((-h1.dy/h2.dy)*(h2.x*h2.dy-h2.y*h2.dx) + h1.x*h1.dy - h1.y*h1.dx) / ((h2.dx * h1.dy / h2.dy) - h1.dx)
	x := y*h2.dx/h2.dy + (h2.x*h2.dy-h2.y*h2.dx)/h2.dy
	return x, y
}

func getCrossedHailstones(in []string, min, max float64) (int, error) {
	if hailstones, err := parseEquations(in); err != nil {
		return -1, err
	} else {
		count := 0
		for i := 0; i < len(hailstones)-1; i++ {
			for j := i + 1; j < len(hailstones); j++ {
				x, y := calculateEquation(hailstones[i], hailstones[j])
				if x >= min && x <= max && y >= min && y <= max {
					count++
				}
			}
		}
		return count, nil
	}
}

func Dot(a, b [3]float64) float64 {
    return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

func Cross(a, b [3]float64) [3]float64 {
    return [3]float64{a[1]*b[2] - a[2]*b[1], a[2]*b[0] - a[0]*b[2], a[0]*b[1] - a[1]*b[0]}
}

func Sub(a, b [3]float64) [3]float64 {
	return [3]float64{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func solvePart2(in []string) int {

	hailstones, err := parseEquations(in)

	if (err != nil) {
		return -1
	}

	// Start building equations
	// 3 hailstones, p1, p2, p3
	// (vi - vj)x(pi-pj)*P = (vi-vj)*(pi x pj)
	p := [3][3]float64{}
	v := [3][3]float64{}

	for i:= 0; i<3; i++ {
		h := hailstones[50+i]

		newp := [3]float64{h.x, h.y, h.z}
		newv := [3]float64{h.dx, h.dy, h.dz}
		p[i] = newp
		v[i] = newv
	}

	indices := [3][2]int{{0,1}, {0,2}, {2,1}}

	var w,x,y,z [3]float64
	var d float64
	var P mat.Dense
	C := mat.NewDense(3,3, nil)
	Ci := mat.NewDense(3,3, nil)
	D := mat.NewVecDense(3, nil)
	for i, idx := range indices {
		w = Cross(p[idx[0]], p[idx[1]])
		x = Sub(v[idx[0]], v[idx[1]])
		y = Sub(p[idx[0]], p[idx[1]])
		z = Cross(x,y)
		d = Dot(x,w)

		for j:=0;j<3;j++{
			C.Set(i,j,z[j])
		}

		D.SetVec(i, d)
	}

	Ci.Inverse(C)
	P.Mul(Ci, D)

	// The 1 Norm of a matrix/vector is the sum of elements
	sum := P.Norm(1)
	return int(sum)
}

func main() {
	useExample := flag.Bool("example", false, "Whether to use the example input or the real problem input")
	flag.Parse()

	fileStr := "./input.txt"
	if *useExample {
		fmt.Println("Using example input")
		fileStr = "./exampleInput.txt"
	}

	file, err := os.Open(fileStr)
	if err != nil {
	    log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		input = append(input, line)
	}

	startTime := time.Now()
	sum, _ := getCrossedHailstones(input, 200000000000000.0, 400000000000000.0)
	endTime := time.Now()
	fmt.Println("Answer to part 1: ", sum, ". Took", endTime.Sub(startTime), "to calculate.")
	// About 1.5 ms on my desktop

	startTime = time.Now()
	sum = solvePart2(input)
	endTime = time.Now()
	fmt.Println("Answer to part 2: ", sum, ". Took", endTime.Sub(startTime), "to calculate.")
	// About 500 us on my desktop, Ryzen 7900X

}

package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type Point struct {
	X, Y float64
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//generatePoints array
func generatePoints(s string) ([]Point, error) {

	points := []Point{}

	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	vals := strings.Split(s, ",")
	if len(vals) < 2 {
		return []Point{}, fmt.Errorf("Point [%v] was not well defined", s)
	}

	var x, y float64

	for idx, val := range vals {

		if idx%2 == 0 {
			x, _ = strconv.ParseFloat(val, 64)
		} else {
			y, _ = strconv.ParseFloat(val, 64)
			points = append(points, Point{x, y})
		}
	}
	return points, nil
}

// getArea gets the area inside from a given shape
func getArea(points []Point) float64 {
	var area float64 = 0.0
	for index, point := range points {
		if index <= len(points)-2 {
			area += getDeterminant(point, points[index+1])
		} else {
			area += getDeterminant(points[len(points)-1], points[0])
		}
	}
	area = math.Abs(area / 2)
	return area
}

// getPerimeter gets the perimeter from a given array of connected points
func getPerimeter(points []Point) float64 {
	var perimeter float64 = 0.0
	for index, point := range points {
		if index <= len(points)-2 {
			perimeter += getDistance(point, points[index+1])
		} else {
			perimeter += getDistance(points[0], points[len(points)-1])
		}
	}
	return perimeter
}

func getDistance(start, end Point) float64 {
	x := math.Pow((end.X - start.X), 2)
	y := math.Pow((end.Y - start.Y), 2)
	var distance float64 = math.Sqrt(x + y)
	return distance
}

func getDeterminant(p1, p2 Point) float64 {
	det := p1.X*p2.Y - p1.Y*p2.X
	return det
}

// handler handles the web request and reponds it
func handler(w http.ResponseWriter, r *http.Request) {

	var vertices []Point
	for k, v := range r.URL.Query() {
		if k == "vertices" {
			points, err := generatePoints(v[0])
			if err != nil {
				fmt.Fprintf(w, fmt.Sprintf("error: %v", err))
				return
			}
			vertices = points
			break
		}
	}

	response := fmt.Sprintf("Welcome to the Remote Shapes Analyzer\n")
	if len(vertices) > 2 {
		// Results gathering
		area := getArea(vertices)
		perimeter := getPerimeter(vertices)

		// Logging in the server side
		log.Printf("Received vertices array: %v", vertices)

		// Response construction
		response += fmt.Sprintf(" - Your figure has : [%v] vertices\n", len(vertices))
		response += fmt.Sprintf(" - Vertices        : %v\n", vertices)
		response += fmt.Sprintf(" - Perimeter       : %v\n", perimeter)
		response += fmt.Sprintf(" - Area            : %v\n\n", area)
	} else {
		response += fmt.Sprintf(" - Error. Your input has two points or less, and can't make a figure. Please try again.\n\n ")
	}
	// Send response to client
	fmt.Fprintf(w, response)
}

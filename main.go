package main

import (
	"github.com/carterjs/algorithms/insertionSort"
	"github.com/carterjs/algorithms/mergeSort"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"log"
	"math/rand"
	"time"
)

func main() {
	p := plot.New()

	// Setup
	p.Title.Text = "runtime vs # of elements"
	p.X.Label.Text = "# of elements"
	p.Y.Label.Text = "execution time (nanoseconds)"

	// store points
	insertionSortPoints := make(plotter.XYs, 50)
	mergeSortPoints := make(plotter.XYs, 50)

	// Run the algorithms
	start := time.Now()

	for i := 0; i < 50; i++ {
		nums := rand.Perm(i + 1)

		// Insertion sort
		start = time.Now()
		insertionSort.Sort(nums)
		insertionSortPoints[i].X = float64(i + i)
		insertionSortPoints[i].Y = float64(time.Since(start).Nanoseconds())

		// Merge sort
		start = time.Now()
		mergeSort.Sort(nums)
		mergeSortPoints[i].X = float64(i + 1)
		mergeSortPoints[i].Y = float64(time.Since(start).Nanoseconds())
	}

	// Draw the lines
	err := plotutil.AddLinePoints(p, "Insertion Sort", insertionSortPoints, "Merge Sort", mergeSortPoints)
	if err != nil {
		return
	}

	if err := p.Save(6*vg.Inch, 4*vg.Inch, "graph.png"); err != nil {
		log.Panicf("Failed to save graph: %v", err)
	}
}
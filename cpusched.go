package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/LgoLgo/cpusched/core"
)

var (
	n      *int
	total  *int64
	resol  *int64
	help   *bool
	worker *bool
	id     *int
	start  *int64
)

func init() {
	// Initialize parameters
	n = flag.Int("n", 1, "Number of processes to run simultaneously")
	total = flag.Int64("total", 5000, "Total runtime of the program (in milliseconds)")
	resol = flag.Int64("resol", 1000, "Interval for collecting statistics (in milliseconds)")
	help = flag.Bool("h", false, "Display help information")

	// Flags for worker mode
	worker = flag.Bool("worker", false, "Run in worker mode")
	id = flag.Int("id", 0, "Worker ID")
	start = flag.Int64("start", 0, "Start time in nanoseconds")
}

func main() {
	// Parse parameters
	flag.Parse()

	// Display help information
	if *help {
		helpInfo()
		return
	}

	// Check if running in worker mode
	if *worker {
		startTime := time.Unix(0, *start)
		processor := &core.Processor{}
		processor.WorkerMain(*id, *total, *resol, startTime)
		return
	}

	// Initialize processor
	processor := core.Processor{
		N:     *n,
		Total: *total,
		Resol: *resol,
		Start: time.Now(),
	}

	// Check parameters
	err := processor.Check()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Execute
	err = processor.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func helpInfo() {
	fmt.Println("Usage:")
	fmt.Println("  This program runs multiple processes simultaneously and periodically collects statistics.")
	fmt.Println("\nParameters:")
	flag.PrintDefaults()
}

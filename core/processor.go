package core

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

type Processor struct {
	N     int
	Total int64
	Resol int64
	Start time.Time
	Wg    sync.WaitGroup
}

func (p *Processor) Check() error {
	// Param Check
	if p.N <= 0 || p.Total <= 0 || p.Resol <= 0 {
		return fmt.Errorf("[cpusched] param error")
	}
	if p.Resol > p.Total {
		return fmt.Errorf("[cpusched] resol over total, resol: %d, total: %d", p.Resol, p.Total)
	}
	return nil
}

func (p *Processor) Execute() error {
	fmt.Printf("Starting CPU scheduling experiment\n")
	fmt.Printf("Number of processes: %d\n", p.N)
	fmt.Printf("Total runtime: %d ms\n", p.Total)
	fmt.Printf("Sampling interval: %d ms\n", p.Resol)
	fmt.Printf("Start time: %s\n\n", p.Start.Format(time.RFC3339Nano))

	fmt.Printf("Estimating workload which takes just one millisecond...\n")
	loopsPerMs := loopsPerMsec()
	fmt.Printf("Loops per millisecond: %d\n", loopsPerMs)
	fmt.Printf("Loops per interval: %d\n\n", loopsPerMs*p.Resol)

	fmt.Printf("%-10s %-10s %-15s %-15s %-30s\n", "Worker ID", "PID", "Elapsed (ms)", "Progress (%)", "Current Time")

	errChan := make(chan error, p.N)
	for i := 0; i < p.N; i++ {
		go func(id int) {
			errChan <- p.runWorker(id)
		}(i)
	}

	for i := 0; i < p.N; i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	fmt.Printf("\nExperiment completed at %s\n", time.Now().Format(time.RFC3339Nano))
	fmt.Printf("Total duration: %s\n", time.Since(p.Start))
	return nil
}

func (p *Processor) runWorker(id int) error {
	executable, err := os.Executable()
	if err != nil {
		return fmt.Errorf("[cpusched] failed to get executable path: %v", err)
	}

	cmd := exec.Command(executable,
		"-worker",
		"-id", strconv.Itoa(id),
		"-total", strconv.FormatInt(p.Total, 10),
		"-resol", strconv.FormatInt(p.Resol, 10),
		"-start", strconv.FormatInt(p.Start.UnixNano(), 10))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (p *Processor) WorkerMain(id int, total, resol int64, start time.Time) {
	end := start.Add(time.Duration(total) * time.Millisecond)
	ticker := time.NewTicker(time.Duration(resol) * time.Millisecond)
	pid := os.Getpid()
	defer ticker.Stop()

	loopsPerMs := loopsPerMsec()
	loopsPerInterval := loopsPerMs * resol

	for {
		select {
		case <-ticker.C:
			now := time.Now()
			elapsed := now.Sub(start).Milliseconds()
			progress := float64(elapsed) / float64(total) * 100
			if progress > 100 {
				progress = 100
			}
			fmt.Printf("%-10d %-10d %-15d %-15.2f %-30s\n", id, pid, elapsed, progress, now.Format(time.RFC3339Nano))

			for i := int64(0); i < loopsPerInterval; i++ {
				// Do nothing
			}
		default:
			if time.Now().After(end) {
				return
			}
		}
	}
}

const (
	NloopForEstimation = 1000000000
)

func loopsPerMsec() int64 {
	start := time.Now()
	for i := 0; i < NloopForEstimation; i++ {
		// Do nothing
	}
	elapsed := time.Since(start)
	return NloopForEstimation * int64(time.Millisecond) / elapsed.Nanoseconds()
}

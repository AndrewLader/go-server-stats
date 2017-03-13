package stats

import (
	"encoding/json"
	"sync/atomic"
	"time"
)

// Stats maintains a set of status values for the service
type Stats struct {
	startTime             time.Time
	numberOfCalls         uint64
	numberOfFailures      uint64
	numberOfBytesReceived uint64
	numberOfBytesWritten  uint64
}

// Output is what is returned when a stats inquiry is made of the server
type Output struct {
	UpTime                string
	NumberOfCalls         uint64
	NumberOfFailures      uint64
	NumberOfBytesReceived uint64
	NumberOfBytesWritten  uint64
}

// New returns a pointer to a new instance of Stats
func New() *Stats {
	return &Stats{
		startTime:             time.Now(),
		numberOfCalls:         0,
		numberOfFailures:      0,
		numberOfBytesReceived: 0,
		numberOfBytesWritten:  0,
	}
}

// GetOutput returns the stats in the form of Output
func (stats *Stats) GetOutput() *Output {
	numberOfCalls := atomic.LoadUint64(&stats.numberOfCalls)
	numberOfFailures := atomic.LoadUint64(&stats.numberOfFailures)
	numberOfBytesReceived := atomic.LoadUint64(&stats.numberOfBytesReceived)
	numberOfVytesWritten := atomic.LoadUint64(&stats.numberOfBytesWritten)

	return &Output{
		UpTime:                time.Since(stats.startTime).String(),
		NumberOfCalls:         numberOfCalls,
		NumberOfFailures:      numberOfFailures,
		NumberOfBytesReceived: numberOfBytesReceived,
		NumberOfBytesWritten:  numberOfVytesWritten,
	}
}

// Bytes returns the stats in the form of a JSON series of bytes
func (stats *Stats) Bytes() []byte {
	outputInstance := stats.GetOutput()

	output, _ := json.Marshal(outputInstance)

	return output
}

// Update takes the provide values and increments the stats values using atomic calls
func (stats *Stats) Update(WasSuccessful bool, numberOfBytesReceived uint64, numberOfBytesWritten uint64) {
	atomic.AddUint64(&stats.numberOfCalls, 1)
	atomic.AddUint64(&stats.numberOfBytesReceived, numberOfBytesReceived)
	atomic.AddUint64(&stats.numberOfBytesWritten, numberOfBytesWritten)

	if !WasSuccessful {
		atomic.AddUint64(&stats.numberOfFailures, 1)
	}
}

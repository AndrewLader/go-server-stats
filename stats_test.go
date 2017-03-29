package stats

import "testing"
import "time"
import "encoding/json"

type simpleStruct struct {
	serverStats *Stats
}

func TestInitializingStatsSuccess(t *testing.T) {
	simple := createTestInstance()

	if simple.serverStats.numberOfBytesReceived != 0 {
		t.Errorf("The number of bytes received was expected to be initialized to 0, but was actually: %d", simple.serverStats.numberOfBytesReceived)
	}

	if simple.serverStats.numberOfBytesWritten != 0 {
		t.Errorf("The number of bytes written was expected to be initialized to 0, but was actually: %d", simple.serverStats.numberOfBytesWritten)
	}

	if simple.serverStats.numberOfCalls != 0 {
		t.Errorf("The number of calls was expected to be initialized to 0, but was actually: %d", simple.serverStats.numberOfCalls)
	}

	if simple.serverStats.numberOfFailures != 0 {
		t.Errorf("The number of failures was expected to be initialized to 0, but was actually: %d", simple.serverStats.numberOfFailures)
	}

	if simple.serverStats.startTime.Sub(time.Now()) > time.Minute {
		t.Errorf("The start time was not within one minute of now, but was actually: %v", simple.serverStats.startTime)
	}
}

func TestUpdatingStatsSuccess(t *testing.T) {
	simple := createTestInstance()

	simple.serverStats.Update(false, 100, 200)

	if simple.serverStats.numberOfBytesReceived != 100 {
		t.Errorf("The number of bytes received was expected to be updated to 100, but was actually: %d", simple.serverStats.numberOfBytesReceived)
	}

	if simple.serverStats.numberOfBytesWritten != 200 {
		t.Errorf("The number of bytes written was expected to be updated to 200, but was actually: %d", simple.serverStats.numberOfBytesWritten)
	}

	if simple.serverStats.numberOfCalls != 1 {
		t.Errorf("The number of calls was expected to be updated to 1, but was actually: %d", simple.serverStats.numberOfCalls)
	}

	if simple.serverStats.numberOfFailures != 1 {
		t.Errorf("The number of failures was expected to be updated to 1, but was actually: %d", simple.serverStats.numberOfFailures)
	}
}

func TestGetOutputSuccess(t *testing.T) {
	simple := createTestInstance()

	simple.serverStats.Update(true, 150, 250)
	simple.serverStats.Update(false, 100, 125)

	output := simple.serverStats.GetOutput()

	if output.NumberOfBytesReceived != 250 {
		t.Errorf("The number of bytes received was expected to be 250, but was actually: %d", output.NumberOfBytesReceived)
	}

	if output.NumberOfBytesWritten != 375 {
		t.Errorf("The number of bytes written was expected to be 375, but was actually: %d", output.NumberOfBytesWritten)
	}

	if output.NumberOfCalls != 2 {
		t.Errorf("The number of calls was expected to be 2, but was actually: %d", output.NumberOfCalls)
	}

	if output.NumberOfFailures != 1 {
		t.Errorf("The number of failures was expected to be 1, but was actually: %d", output.NumberOfFailures)
	}
}

func TestGetBytesSuccess(t *testing.T) {
	simple := createTestInstance()

	simple.serverStats.Update(true, 50, 66)
	simple.serverStats.Update(false, 5, 125)

	originalOutput := simple.serverStats.GetOutput()

	bytes := simple.serverStats.Bytes()

	var actualOutput Output
	err := json.Unmarshal(bytes, &actualOutput)
	if err != nil {
		t.Errorf("An error occurred while unmarshaling the data back into an Output struct: %v", err)
	}

	if actualOutput.NumberOfBytesReceived != originalOutput.NumberOfBytesReceived {
		t.Errorf("The actual number of bytes received did not match the original: %d vs %d", actualOutput.NumberOfBytesReceived, originalOutput.NumberOfBytesReceived)
	}

	if actualOutput.NumberOfBytesWritten != originalOutput.NumberOfBytesWritten {
		t.Errorf("The actual number of bytes written did not match the original: %d vs %d", actualOutput.NumberOfBytesWritten, originalOutput.NumberOfBytesWritten)
	}
	if actualOutput.NumberOfCalls != originalOutput.NumberOfCalls {
		t.Errorf("The actual number of calls did not match the original: %d vs %d", actualOutput.NumberOfCalls, originalOutput.NumberOfCalls)
	}

	if actualOutput.NumberOfFailures != originalOutput.NumberOfFailures {
		t.Errorf("The actual number of failures did not match the original: %d vs %d", actualOutput.NumberOfFailures, originalOutput.NumberOfFailures)
	}
}

func createTestInstance() *simpleStruct {
	return &simpleStruct{
		serverStats: New(),
	}
}

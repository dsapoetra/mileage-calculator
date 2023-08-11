package main

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"
)

func TestInputGetter(t *testing.T) {

	line := "00:01:00.0 10.0"

	var dataArr []Data
	dataArr, _ = InputGetter(dataArr, line)

	if len(dataArr) != 1 {
		t.Errorf("InputGetter was incorrect, got: %d, want: %d.", len(dataArr), 2)
	}
}

// mockStdin is a helper function that lets the test pretend dummyInput as os.Stdin.
// It will return a function for `defer` to clean up after the test.
//
// Note: `ioutil.TempFile` should be replaced to `os.CreateTemp` for Go v1.16 or higher.
func mockStdin(t *testing.T, dummyInput string) (funcDefer func(), err error) {
	t.Helper()

	oldOsStdin := os.Stdin

	tmpfile, err := ioutil.TempFile(t.TempDir(), t.Name())
	if err != nil {
		return nil, err
	}

	content := []byte(dummyInput)

	if _, err := tmpfile.Write(content); err != nil {
		return nil, err
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		return nil, err
	}

	// Set stdin to the temp file
	os.Stdin = tmpfile

	return func() {
		// clean up
		os.Stdin = oldOsStdin
		os.Remove(tmpfile.Name())
	}, nil
}

// test for checkTimeLatest(dataArr []Data, line string) error
func TestCheckTimeLatest(t *testing.T) {
	var dataArr []Data
	dataArr = append(dataArr, Data{"00:05:00.0", "0.0", 0.0})
	//err := checkTimeLatest(dataArr, "00:07:00.0 0.0")
	//if err != nil {
	//	t.Errorf("checkTimeLatest was incorrect, got: %s, want: %v.", err, nil)
	//}
	//
	//err = checkTimeLatest(dataArr, "00:20:00.0 0.0")
	//if err == nil {
	//	t.Errorf("checkTimeLatest was incorrect, got: %s, want: %v.", err, errors.New("error: time should be within 5 minutes"))
	//}

	err := checkTimeLatest(dataArr, "00:01:00.0 0.0")
	if err == nil {
		t.Errorf("checkTimeLatest was incorrect, got: %s, want: %v.", err, errors.New("error: time should be in ascending order"))
	}
}

// test for compareMileage(mileage1 string, mileage2 string) error
func TestCompareMileage(t *testing.T) {
	err := compareMileage("0.0", "0.0")
	if err != nil {
		t.Errorf("compareMileage was incorrect, got: %s, want: %v.", err, nil)
	}

	err = compareMileage("0.1", "0.0")
	if err == nil {
		t.Errorf("compareMileage was incorrect, got: %s, want: %v.", err, errors.New("error: mileage is not in ascending order"))
	}
}

// test checkInputFormat(line string) error
func TestCheckInputFormat(t *testing.T) {
	err := checkInputFormat("00:00:00.0 0.0")
	if err != nil {
		t.Errorf("checkInputFormat was incorrect, got: %s, want: %v.", err, nil)
	}

	err = checkInputFormat("00:00:00.0 0.0 0.0")
	if err != nil {
		t.Errorf("checkInputFormat was incorrect, got: %s, want: %v.", err, errors.New("error: input format is incorrect"))
	}
}

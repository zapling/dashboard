package calendar

import (
	"testing"
	"time"
)

func TestGetMonthDays(t *testing.T) {
	testCases := []struct {
		month           time.Time
		expectedResults int
	}{
		{month: date(2022, 1, 1), expectedResults: 42},
		{month: date(2022, 2, 1), expectedResults: 35},
		{month: date(2022, 3, 1), expectedResults: 35},
		{month: date(2022, 4, 1), expectedResults: 35},
		{month: date(2022, 5, 1), expectedResults: 42},
		{month: date(2022, 6, 1), expectedResults: 35},
		{month: date(2022, 7, 1), expectedResults: 35},
		{month: date(2022, 8, 1), expectedResults: 35},
		{month: date(2022, 9, 1), expectedResults: 35},
		{month: date(2022, 10, 1), expectedResults: 42},
		{month: date(2022, 11, 1), expectedResults: 35},
		{month: date(2022, 12, 1), expectedResults: 35},
	}

	for index, testCase := range testCases {
		days := getMonthDays(testCase.month)
		if len(days) != testCase.expectedResults {
			t.Errorf("[%d] expected: %d, got: %d", index, testCase.expectedResults, len(days))
		}
	}
}

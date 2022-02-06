package calendar

import "testing"

func TestGetMonth(t *testing.T) {
	testCases := []struct {
		year            int
		month           int
		expectedResults int
	}{
		{year: 2022, month: 1, expectedResults: 42},
		{year: 2022, month: 2, expectedResults: 35},
		{year: 2022, month: 3, expectedResults: 35},
		{year: 2022, month: 4, expectedResults: 35},
		{year: 2022, month: 5, expectedResults: 42},
		{year: 2022, month: 6, expectedResults: 35},
		{year: 2022, month: 7, expectedResults: 35},
		{year: 2022, month: 8, expectedResults: 35},
		{year: 2022, month: 9, expectedResults: 35},
		{year: 2022, month: 10, expectedResults: 42},
		{year: 2022, month: 11, expectedResults: 35},
		{year: 2022, month: 12, expectedResults: 35},
	}

	for index, testCase := range testCases {
		days := GetMonthView(testCase.year, testCase.month)
		if len(days) != testCase.expectedResults {
			t.Errorf("[%d] expected: %d, got: %d", index, testCase.expectedResults, len(days))
		}
	}
}

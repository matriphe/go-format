package datetime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_To(t *testing.T) {
	testCases := map[string]struct {
		toFunc   func(time.Time) string
		expected string
	}{
		ISODate: {
			toFunc:   ToISODate,
			expected: "2023-03-05",
		},
		ISOTime: {
			toFunc:   ToISOTime,
			expected: "18:45:21",
		},
		ISOHourMin: {
			toFunc:   ToISOHourMin,
			expected: "18:45",
		},
		ISODateTime: {
			toFunc:   ToISODateTime,
			expected: "2023-03-05 18:45:21",
		},
	}

	theTime := time.Date(2023, 3, 5, 18, 45, 21, 0, time.UTC)

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := tc.toFunc(theTime)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func Test_From(t *testing.T) {
	testCases := map[string]struct {
		time      string
		fromFunc  func(s string) (time.Time, error)
		expected  time.Time
		assertion func(t *testing.T, tt time.Time)
	}{
		ISODate: {
			time:     "2023-03-05",
			fromFunc: FromISODate,
			assertion: func(t *testing.T, tt time.Time) {
				assert.Equal(t, 2023, tt.Year())
				assert.Equal(t, time.Month(3), tt.Month())
				assert.Equal(t, 5, tt.Day())
			},
		},
		ISOTime: {
			time:     "18:45:21",
			fromFunc: FromISOTime,
			assertion: func(t *testing.T, tt time.Time) {
				assert.Equal(t, 18, tt.Hour())
				assert.Equal(t, 45, tt.Minute())
				assert.Equal(t, 21, tt.Second())
			},
		},
		ISOHourMin: {
			time:     "18:45",
			fromFunc: FromISOHourMin,
			assertion: func(t *testing.T, tt time.Time) {
				assert.Equal(t, 18, tt.Hour())
				assert.Equal(t, 45, tt.Minute())
				assert.Equal(t, 0, tt.Second())
			},
		},
		ISODateTime: {
			time:     "2023-03-05 18:45:21",
			fromFunc: FromISODateTime,
			assertion: func(t *testing.T, tt time.Time) {
				assert.Equal(t, 2023, tt.Year())
				assert.Equal(t, time.Month(3), tt.Month())
				assert.Equal(t, 5, tt.Day())
				assert.Equal(t, 18, tt.Hour())
				assert.Equal(t, 45, tt.Minute())
				assert.Equal(t, 21, tt.Second())
			},
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.fromFunc(tc.time)

			assert.NoError(t, err)

			tc.assertion(t, actual)
		})
	}
}

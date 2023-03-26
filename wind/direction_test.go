package wind

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Direction(t *testing.T) {
	testCases := []struct {
		deg      float32
		expected string
	}{
		{0, "N"},
		{90, "E"},
		{180, "S"},
		{270, "W"},
		{360, "N"},
		{720, "N"},

		{11, "N"},
		{13, "NNE"},
		{33, "NNE"},
		{34, "NE"},
		{34, "NE"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%f returns %s", tc.deg, tc.expected), func(t *testing.T) {
			t.Parallel()

			actual := Direction(tc.deg)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

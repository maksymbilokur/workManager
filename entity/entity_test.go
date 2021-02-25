package entity

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_stringToDuration(t *testing.T) {

	resultTime := time.Date(2012, 1, 2, 15, 12, 12, 0, time.UTC)
	tests := []struct {
		name   string
		args   string
		result time.Time
	}{
		{
			name:   "1",
			args:   "01/02/2012 15:12:12",
			result: resultTime,
		},
		{
			name:   "2",
			args:   "01/2/2012 15:12:12",
			result: resultTime,
		},
		{
			name:   "3",
			args:   "1/02/2012 15:12:12",
			result: resultTime,
		},
		{
			name:   "4",
			args:   "1/2/2012 15:12:12",
			result: resultTime,
		},
	}
	//01 month
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			te, err := time.Parse("1/_2/2006 15:04:05", tt.args)
			require.Nil(t, err)
			require.Equal(t, te, time.Date(2012, 1, 2, 15, 12, 12, 0, time.UTC))

		})
	}
}

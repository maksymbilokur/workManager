package entity

import "testing"

func Test_stringToDuration(t *testing.T) {
	tests := []struct {
		name string
		args string
	}{{
		name: "1",
		args: "01/02/2012 15:12:12",
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StringToData(tt.args)
		})
	}
}

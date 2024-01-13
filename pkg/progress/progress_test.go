package progress

import (
	"testing"
)

func TestPrintProgressBar(t *testing.T) {
	type args struct {
		progress int
		width    int
	}

	tests := []struct {
		name string
		args args
	}{
		{"Test 1", args{0, 50}},
		{"Test 2", args{25, 50}},
		{"Test 3", args{35, 50}},
		{"Test 4", args{45, 50}},
		{"Test 5", args{50, 50}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintProgressBar(tt.args.progress, tt.args.width)
		})
	}

}

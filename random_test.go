package deagon

import (
	"testing"
)

func TestRandomName(t *testing.T) {
	type args struct {
		formatter Formatter
	}
	tests := []struct {
		args args
		want string
	}{
		{args{NewCapitalizedSpaceFormatter()}, "Anna Plemons"},
		{args{NewCapitalizedSpaceFormatter()}, "Alisa Delorge"},
		{args{NewCapitalizedSpaceFormatter()}, "Sonja Luisi"},
		{args{NewCapitalizedSpaceFormatter()}, "Inez Alviar"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := RandomName(tt.args.formatter); got != tt.want {
				t.Errorf("RandomName() = %v, want %v", got, tt.want)
			}
		})
	}
}

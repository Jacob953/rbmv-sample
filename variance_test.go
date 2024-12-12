package rbmv

import "testing"

func TestVariance(t *testing.T) {
	type args struct {
		vals   []float64
		avgVal float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "getZero",
			args: args{
				vals:   []float64{3, 3, 3},
				avgVal: 3,
			},
			want: 0,
		},
		{
			name: "getNormal",
			args: args{
				vals:   []float64{1, 3, 5},
				avgVal: 3,
			},
			want: 8.0 / 3,
		},
		{
			name: "getWithSpecificAvg",
			args: args{
				vals:   []float64{1, 3},
				avgVal: 3,
			},
			want: 4.0 / 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Variance(tt.args.vals, tt.args.avgVal); got != tt.want {
				t.Errorf("Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}

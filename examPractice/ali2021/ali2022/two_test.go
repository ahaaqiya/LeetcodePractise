package main

import "testing"

func Test_findMandN(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findMandN(tt.args.s)
			if got != tt.want {
				t.Errorf("findMandN() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findMandN() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

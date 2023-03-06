package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestDo(t *testing.T) {
	tests := []struct {
		name    string
		params  DoParams
		want    []string
		wantErr bool
	}{
		{
			name: "happy path",
			params: DoParams{
				int1:  3,
				int2:  5,
				limit: 100,
				str1:  "fizz",
				str2:  "buzz",
			},
			want: []string{
				"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz", "fizz", "22", "23", "fizz", "buzz", "26", "fizz", "28", "29", "fizzbuzz", "31", "32", "fizz", "34", "buzz", "fizz", "37", "38", "fizz", "buzz", "41", "fizz", "43", "44", "fizzbuzz", "46", "47", "fizz", "49", "buzz", "fizz", "52", "53", "fizz", "buzz", "56", "fizz", "58", "59", "fizzbuzz", "61", "62", "fizz", "64", "buzz", "fizz", "67", "68", "fizz", "buzz", "71", "fizz", "73", "74", "fizzbuzz", "76", "77", "fizz", "79", "buzz", "fizz", "82", "83", "fizz", "buzz", "86", "fizz", "88", "89", "fizzbuzz", "91", "92", "fizz", "94", "buzz", "fizz", "97", "98", "fizz", "buzz",
			},
		},
		{
			name: "invalid int1",
			params: DoParams{
				int1:  0,
				int2:  5,
				limit: 100,
				str1:  "fizz",
				str2:  "buzz",
			},
			wantErr: true,
		},
		{
			name: "invalid int2",
			params: DoParams{
				int1:  3,
				int2:  0,
				limit: 100,
				str1:  "fizz",
				str2:  "buzz",
			},
			wantErr: true,
		},
		{
			name: "invalid limit",
			params: DoParams{
				int1:  3,
				int2:  5,
				limit: 0,
				str1:  "fizz",
				str2:  "buzz",
			},
			wantErr: true,
		},
		{
			name: "invalid str1",
			params: DoParams{
				int1:  3,
				int2:  5,
				limit: 100,
				str1:  "",
				str2:  "buzz",
			},
			wantErr: true,
		},
		{
			name: "invalid str2",
			params: DoParams{
				int1:  3,
				int2:  5,
				limit: 100,
				str1:  "fizz",
				str2:  "",
			},
			wantErr: true,
		},
		{
			name: "int1 and int2 are equal",
			params: DoParams{
				int1:  2,
				int2:  2,
				limit: 2,
				str1:  "fizz",
				str2:  "buzz",
			},
			want: []string{
				"1", "fizzbuzz",
			},
		},
		{
			name: "str1 and str2 are equal",
			params: DoParams{
				int1:  2,
				int2:  3,
				limit: 6,
				str1:  "fizz",
				str2:  "fizz",
			},
			want: []string{
				"1", "fizz", "fizz", "fizz", "5", "fizzfizz",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Do(tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

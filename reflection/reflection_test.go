package reflection

import (
	"reflect"
	"testing"
)

type Charecteristics struct {
	Age         int
	HP          int
	Description string
}

func TestWalk(t *testing.T) {
	type args struct {
		x any
	}
	tests := []struct {
		name          string
		args          args
		expectedCalls []string
	}{
		{
			name: "struct with one string field",
			args: args{
				x: struct {
					Name string
				}{
					"Liz",
				},
			},
			expectedCalls: []string{"Liz"},
		},
		{
			name: "two string fields",
			args: args{
				x: struct {
					FirstName  string
					SecondName string
				}{
					"Liz",
					"Chipi",
				},
			},
			expectedCalls: []string{"Liz", "Chipi"},
		},
		{
			name: "no string fields",
			args: args{
				x: struct {
					num int
				}{
					16,
				},
			},
			expectedCalls: []string(nil),
		},
		{
			name: "one string, one non-string fueld",
			args: args{
				x: struct {
					Name string
					Age  int
				}{"Liz", 24},
			},
			expectedCalls: []string{"Liz"},
		},
		{
			name: "non-flat struct",
			args: args{
				struct {
					Name            string
					Charecteristics Charecteristics
				}{
					Name: "Elder Titan",
					Charecteristics: Charecteristics{
						14214214,
						2800,
						"Best Dota hero",
					},
				},
			},
			expectedCalls: []string{"Elder Titan", "Best Dota hero"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			walk(tt.args.x, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, tt.expectedCalls) {
				t.Errorf("got %v want %v", got, tt.expectedCalls)
			}
		})
	}
}

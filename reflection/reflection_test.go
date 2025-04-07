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
					Num int
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
		{
			name: "x is a pointer to struct",
			args: args{
				&Charecteristics{1415, 1155, "kinda old"},
			},
			expectedCalls: []string{"kinda old"},
		},
		{
			name: "x is a slice of structs",
			args: args{
				[]Charecteristics{
					{1, 2, "three"},
					{4, 5, "sechs"},
				},
			},
			expectedCalls: []string{"three", "sechs"},
		},
		{
			name: "x is an array of structs",
			args: args{
				[2]Charecteristics{
					{1, 2, "three"},
					{4, 5, "sechs"},
				},
			},
			expectedCalls: []string{"three", "sechs"},
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

	t.Run("x is a map", func(t *testing.T) {
		aMap := map[string]string{"1": "ein", "2": "zwei", "3": "drei"}
		expectedCalls := []string{"ein", "zwei", "drei"}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		for _, expCall := range expectedCalls {
			assertContains(t, got, expCall)
		}
	})

}

func assertContains(t testing.TB, got []string, val string) {
	t.Helper()
	contains := false
	for _, elem := range got {
		if elem == val {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", got, val)
	}
}

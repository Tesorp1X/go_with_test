package reflection

import (
	"reflect"
	"testing"
)

/* func TestWalk(t *testing.T) {

	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}

	if got[0] != expected {
		t.Errorf("got %q want %q", got[0], expected)
	}
} */

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
		/* 		{
		   			name: "no string fields",
		   			args: args{
		   				x: struct {
		   					num int
		   				}{
		   					16,
		   				},
		   			},
		   			expectedCalls: []string{},
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
		   		}, */
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

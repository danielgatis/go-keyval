package keyval

import (
	"reflect"
	"testing"
)

func TestToMap(t *testing.T) {
	var tests = []struct {
		name string
		args []interface{}
		want map[string]interface{}
	}{
		{"even args", []interface{}{1, 2, 3, 4}, map[string]interface{}{"1": 2, "3": 4}},
		{"odd args", []interface{}{1, 2, 3}, map[string]interface{}{"1": 2, "3": nil}},
		{"dup args", []interface{}{1, 1, 1, 0}, map[string]interface{}{"1": 0}},
		{"no args", []interface{}{}, map[string]interface{}{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ans := ToMap(tt.args...); !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

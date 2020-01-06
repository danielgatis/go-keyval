package keyval

import (
	"fmt"
)

// ToMap returns a map by grouping arguments two by two where the first is the key and the second is the value
func ToMap(keyvals ...interface{}) map[string]interface{} {
	len := len(keyvals)
	kvm := make(map[string]interface{}, len/2)

	for i := 0; i < len; i += 2 {
		key := keyvals[i]

		if i+1 < len {
			kvm[fmt.Sprint(key)] = keyvals[i+1]
		} else {
			kvm[fmt.Sprint(key)] = nil
		}
	}

	return kvm
}

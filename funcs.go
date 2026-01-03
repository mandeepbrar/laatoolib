package lib

import (
	"fmt"

	"laatoo.io/sdk/utils"
)

// Merge performs a shallow merge of two maps (no deep nesting)
// Works with any map type: map[string]string, map[string]int, etc.
func Merge[K comparable, V any](m1, m2 map[K]V) map[K]V {
	if m1 == nil {
		return copyMap(m2)
	}
	if m2 == nil {
		return copyMap(m1)
	}

	result := make(map[K]V, len(m1)+len(m2))

	// Copy m1
	for k, v := range m1 {
		result[k] = v
	}

	// Override with m2
	for k, v := range m2 {
		result[k] = v
	}

	return result
}

// copyMap creates a shallow copy of a map
func copyMap[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

// DeepMerge performs a deep merge for maps with interface{} values
// Recursively merges nested maps
func DeepMerge[K comparable](m1, m2 map[K]interface{}) map[K]interface{} {
	if m1 == nil {
		return deepCopy(m2)
	}
	if m2 == nil {
		return deepCopy(m1)
	}

	result := make(map[K]interface{}, len(m1))

	// Copy everything from m1 first
	for k, v := range m1 {
		result[k] = deepCopyValue(v)
	}

	// Merge from m2 (with override)
	for k, v := range m2 {
		if existing, exists := result[k]; exists {
			// If both are maps, merge them recursively
			if existingMap, ok1 := existing.(map[K]interface{}); ok1 {
				if vMap, ok2 := v.(map[K]interface{}); ok2 {
					result[k] = DeepMerge(existingMap, vMap)
					continue
				}
			}
			// Also handle generic map[string]interface{} case
			if existingMap, ok1 := existing.(map[string]interface{}); ok1 {
				if vMap, ok2 := v.(map[string]interface{}); ok2 {
					result[k] = DeepMerge(existingMap, vMap)
					continue
				}
			}
		}
		// Override or add new value
		result[k] = deepCopyValue(v)
	}

	return result
}

func deepCopy[K comparable](m map[K]interface{}) map[K]interface{} {
	if m == nil {
		return nil
	}
	result := make(map[K]interface{}, len(m))
	for k, v := range m {
		result[k] = deepCopyValue(v)
	}
	return result
}

func deepCopyValue(v interface{}) interface{} {
	switch val := v.(type) {
	case map[string]interface{}:
		return deepCopy(val)
	case map[interface{}]interface{}:
		result := make(map[interface{}]interface{}, len(val))
		for k, v := range val {
			result[k] = deepCopyValue(v)
		}
		return result
	case []interface{}:
		result := make([]interface{}, len(val))
		for i, item := range val {
			result[i] = deepCopyValue(item)
		}
		return result
	default:
		// For primitive types, direct assignment is fine
		return v
	}
}

func ConvertToStringsMap(val map[string]interface{}) utils.StringsMap {
	if val == nil {
		return nil
	}
	strMp := make(utils.StringsMap)
	for k, v := range val {
		strV := fmt.Sprintf("%v", v)
		strMp[k] = strV
	}
	return strMp
}

package convert

func IsStringMap(i interface{}) (map[string]interface{}, bool) {
	switch i.(type) {
	case map[string]interface{}:
		return i.(map[string]interface{}), true
	}
	return nil, false
}

func IsArray(i interface{}) ([]interface{}, bool) {
	switch i.(type) {
	case []interface{}:
		return i.([]interface{}), true
	}
	return nil, false
}

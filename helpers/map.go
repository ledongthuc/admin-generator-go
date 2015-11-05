package helpers

func GetKeys(mapper map[interface{}]interface{}) []interface{} {
	var result []interface{}
	for key, _ := range mapper {
		result = append(result, key)
	}
	return result
}

func GetValues(mapper map[interface{}]interface{}) []interface{} {
	var result []interface{}
	for _, value := range mapper {
		result = append(result, value)
	}
	return result
}

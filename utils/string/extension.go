package stringextension

func CastMap(m map[interface{}]interface{}) map[string]interface{} {
	m2 := make(map[string]interface{})
	for key, value := range m {
		switch key := key.(type) {
		case string:
			m2[key] = value
		}
	}
	return m2
}

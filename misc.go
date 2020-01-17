package go_util

func MergeMap(a map[string]interface{}, b map[string]interface{}) {
	for k, v := range b {
		a[k] = v
	}
}

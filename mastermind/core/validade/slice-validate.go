package validate

import "reflect"

func Contains(slice interface{}, val interface{}) bool {
	if slice != nil {
		sv := reflect.ValueOf(slice)

		for i := 0; i < sv.Len(); i++ {
			if sv.Index(i).Interface() == val {
				return true
			}
		}
	}
	return false
}

func SliceEquals(x, y []int) (bool) {
	return reflect.DeepEqual(x, y)
}

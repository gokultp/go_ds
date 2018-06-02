package arrays

type Array []interface{}

type iterator func(interface{}, int) interface{}
type compare func(interface{}, interface{}) int

func (array Array) Map(iterator func(interface{}, int) interface{}) Array {
	results := Array{}
	for index, value := range array {
		results = append(results, iterator(value, index))
	}
	return results
}

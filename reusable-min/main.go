package main

import "reflect"

// Attempt 1 - Repeat ...
func min(a int, b int) int {
	if a < b { return a }
	return b
}

// Attempt 2 - use interface{}

// The function - Compile error

//func reusableMin(a interface{}, b interface{}) interface{} {
//	if a < b { // INVALID!
//	// Invalid operation: a < b (operator < is not defined on interface{})
//		return a
//	}
//	return b
//}


// The function - Use Comparator interface

type Comparator interface {
	Compare(v interface{}) int
}

func minByComparator(a Comparator, b Comparator) interface{} {
	if a.Compare(b) < 0 {
		return a
	}
	return b
}

// The client

type MyInt64 int64

func (u MyInt64) Compare(v interface{}) int {
	vv := v.(MyInt64)
	if u < vv 	{ return -1 }
	if u > vv 	{ return 1 }
	return 0
}

func fooClient() {
	a, b := MyInt64(3), MyInt64(4)
	minByComparator(a, b)
}

// Attempt 3 - Reflection

func minByInterface(a interface{}, b interface{}) interface{} {
	if a.(int) < b.(int) { return a }
	return b
}

func minByReflect(a interface{}, b interface{}) interface{} {
	va := reflect.ValueOf(a)
	switch va.Kind() {
	case reflect.Int:
		if a.(int) < b.(int) { return a }
		return b
	default:
		return a
	}
}

func main() {
	//minByComparator(3, 4) // INVALID!
	// Cannot use '3' (type untyped int) as type Comparator
	//Type does not implement 'Comparator' as some methods are missing:
	//Compare(v interface{}) int

	a, b := MyInt64(3), MyInt64(4)
	minByComparator(a, b)
}

package main

import (
	"testing"
)

/*
== go 1.15 ==
goos: darwin
goarch: amd64
pkg: june.yuan.me/practice.perfect.go/a-bite-of-generics/reusable-min
BenchmarkMin-8                     	1000000000	         0.321 ns/op
BenchmarkMinByInterface1-8         	1000000000	         0.317 ns/op
BenchmarkMinByInterface2-8         	1000000000	         0.317 ns/op
BenchmarkMinByInterfaceMyInt64-8   	374377816	         3.19 ns/op
BenchmarkMinByReflect1-8           	253392189	         4.72 ns/op
BenchmarkMinByReflect2-8           	248952228	         4.78 ns/op
PASS
*/

/*
== go 1.13 ==
goos: darwin
goarch: amd64
pkg: june.yuan.me/practice.perfect.go/a-bite-of-generics/reusable-min
BenchmarkMin-8                     	914771989	         1.34 ns/op
BenchmarkMinByInterface1-8         	1000000000	         0.738 ns/op
BenchmarkMinByInterface2-8         	899828476	         1.35 ns/op
BenchmarkMinByInterfaceMyInt64-8   	181246814	         6.84 ns/op
BenchmarkMinByReflect1-8           	121341340	         9.99 ns/op
BenchmarkMinByReflect2-8           	120099645	        10.0 ns/op
PASS
*/

var (
	a, b = 101, 102 // int
	c, d interface{}
)

func init() {
	a, b = 101, 102
	c, d = 101, 102
}

func BenchmarkMin(bb *testing.B) {
	for i := 0; i < bb.N; i++ {
		min(5, 6)
	}
}

func BenchmarkMinByInterface1(bb *testing.B) {
	for i := 0; i < bb.N; i++ {
		minByInterface(5, 6)
	}
}

func BenchmarkMinByInterface2(bb *testing.B) {
	for i := 0; i < bb.N; i++ {
		minByInterface(interface{}(5), interface{}(6))
	}
}

func BenchmarkMinByInterfaceMyInt64(bb *testing.B) {
	for i := 0; i < bb.N; i++ {
		minByComparator(MyInt64(5), MyInt64(6))
	}
}

func BenchmarkMinByReflect1(bb *testing.B) {
	for i := 0; i < bb.N; i++ {
		minByReflect(5, 6)
	}
}

func BenchmarkMinByReflect2(bb *testing.B) {
	for i := 0; i < bb.N; i++ {
		minByReflect(interface{}(5), interface{}(6))
	}
}

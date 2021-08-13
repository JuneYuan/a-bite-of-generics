package main

import (
	"testing"
)

/*
goos: darwin
goarch: amd64
pkg: a-bite-of-generics/slice-of-general-type
BenchmarkSln1_ExplicitlyConvert-8   	 3930650	       298 ns/op
BenchmarkSln2_RankByReflect-8       	 2074530	       589 ns/op
*/

var (
	pois []*Place
)

func init() {
	pois = []*Place{
		{PoiID: "POIID_1"},
		{PoiID: "POIID_2"},
		{PoiID: "POIID_3"},
		{PoiID: "POIID_4"},
		{PoiID: "POIID_5"},
	}
}

func TestExplicitlyConvert(t *testing.T) {
	fooClient()
}

func TestReflect(t *testing.T) {
	//pretty.Printf("before Rank: %v\n", pois)
	RankByReflect(pois)
	//pretty.Printf("after Rank: %v\n", pois)
}

func BenchmarkSln1_ExplicitlyConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var candidates = make([]Candidate, len(pois), len(pois))
		for i, poi := range pois { candidates[i] = poi }
		Rank(candidates)
		for i, cand := range candidates { pois[i] = cand.(*Place) }
	}
}

func BenchmarkSln2_RankByReflect(b *testing.B) {
	for i := 0; i < b.N; i++ { RankByReflect(pois) }
}

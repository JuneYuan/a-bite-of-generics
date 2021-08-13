package main

import (
	"reflect"
	"sort"
)

type Candidate interface {
	GetPoiId() string
}

type Place struct {
	PoiID string
}

func (p *Place) GetPoiId() string { return p.PoiID }
func (p *Place) String() string { return p.PoiID }


func Rank(candidates []Candidate) {
	// long score computing process ...
	var scores = make([]float64, len(candidates))
	for i := range candidates { scores[i] = float64(i) }
	// call system sort
	cmp := func(i, j int) bool { return scores[i] < scores[j] }
	sort.Slice(candidates, cmp)
}

// Solution 1. Explicitly convert
func fooClient() {
	var pois = []*Place{
		{PoiID: "POIID_1"},
		{PoiID: "POIID_2"},
		{PoiID: "POIID_3"},
		{PoiID: "POIID_4"},
		{PoiID: "POIID_5"},
	}
	// []*Place -> []Candidate
	var candidates = make([]Candidate, len(pois), len(pois))
	for i, poi := range pois { candidates[i] = poi }
	// call Rank function
	Rank(candidates)
	// []Candidate -> []*Place
	for i, cand := range candidates { pois[i] = cand.(*Place) }
	// Rank finished.
}

// Solution 2. Reflect
// https://blog.golang.org/laws-of-reflection
func RankByReflect(candidates interface{}) {
	// validate argument type
	tp := reflect.TypeOf(candidates)
	if tp.Kind() != reflect.Slice { return }

	// interface{} -> []Candidate
	var gCandidates []Candidate
	rv := reflect.ValueOf(candidates)
	// fmt.Printf("rv.Type()=%v, rv.Value=%v\n", rv.Type(), rv.String())
	for i := 0; i < rv.Len(); i++ {
		item := rv.Index(i)
		c, _ := item.Interface().(Candidate) // ignore assert error
		gCandidates = append(gCandidates, c)
	}
	// pretty.Printf("gCandidates=%v\n", gCandidates)

	// long, concurrent score computing process ...
	var scores = make([]float64, len(gCandidates))
	for i := range gCandidates { scores[i] = float64(i) }

	// call system sort
	cmp := func(i, j int) bool { return scores[i] < scores[j] }
	sort.Slice(candidates, cmp)
}

func main() {
	var pois = []*Place{
		{PoiID: "POIID_1"},
	}
	//Rank(pois) // Cannot use 'results' (type []*Place) as type []Candidate

	var candidates = make([]Candidate, len(pois), len(pois))
	for i := range pois {
		candidates[i] = pois[i]
	}
	Rank(candidates)

	RankByReflect(pois)
}



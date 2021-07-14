package main

import (
	"reflect"
	"testing"
	// "testing"
	// "fmt"
	// "context"
	// "flag"
	// "fmt"
	// "io/ioutil"
	// "log"
	// "os"
	// "strings"
	// "text/template"
	// "reflect"
	// "github.com/gocolly/colly"
)

func TableTest1(t *testing.T) {
	var vegFoodsSlice []VegFood
	if len(GetData(vegFoodsSlice)) != 5 {
		t.Error("Expected len(GetData(vegFoodsSlice)) == 5")
	}
	var uint8Var []uint8
	if reflect.TypeOf(GetData(vegFoodsSlice)) != reflect.TypeOf(uint8Var) {
		t.Error("Expected reflect.TypeOf(GetData(vegFoodsSlice)) == []uint8")
	}
}

func Benchmark(b *testing.B) {
	var vegFoodsSlice []VegFood
	for i := 0; i < b.N; i++ {
		GetData(vegFoodsSlice)
	}
}

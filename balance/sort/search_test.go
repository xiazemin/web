package sort

import (
	"testing"
	"fmt"
	"sort"
)

func TestSearch(t *testing.T){
	var a []int
	for i:=0;i<20;i++{
	a=append(a,i)
	}
	fmt.Println(sort.Search(len(a),func(i int)bool{return a[i]<=4}),a[19])
	fmt.Println(sort.Search(len(a),func(i int)bool{return a[i]>=4}))
}

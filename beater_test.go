package main

import (
	"fmt"
	"testing"
)

func TestBeater_calculate(t *testing.T) {
	match1 := match{
		team{"A", "B"},
		team{"C", "D"},
		17,
		21,
		1,
	}
	beater := Beater{[]match{match1}}
	res := beater.calculate()
	for _, item := range res {
		fmt.Printf("%v\n", item)
	}
}

func TestBeater_calculate2(t *testing.T) {
	matchList := []match{
		{
			team{"权", "泰"},
			team{"毅", "豪"},
			17,
			21,
			1,
		},
		{
			team{"豪", "权"},
			team{"毅", "泰"},
			19,
			21,
			1,
		},
		{
			team{"权", "雷"},
			team{"真", "毅"},
			19,
			21,
			1,
		},
		{
			team{"毅", "权"},
			team{"权", "豪"},
			22,
			20,
			1,
		},
		{
			team{"泰", "毅"},
			team{"权", "真"},
			21,
			17,
			2,
		},
	}
	beater := Beater{matchList: matchList}
	res := beater.calculate()
	for _, item := range res {
		fmt.Printf("%s 转给 %s %d 瓶水\n", item.from, item.to, item.water)
	}
}

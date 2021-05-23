package config

import "encoding/json"

//go:embed embed/chair_condtion.json
var rawChairConditionJSON []byte

type ChairSearchCondition struct {
	Width   RangeCondition `json:"width"`
	Height  RangeCondition `json:"height"`
	Depth   RangeCondition `json:"depth"`
	Price   RangeCondition `json:"price"`
	Color   ListCondition  `json:"color"`
	Feature ListCondition  `json:"feature"`
	Kind    ListCondition  `json:"kind"`
}

type RangeCondition struct {
	Prefix string   `json:"prefix"`
	Suffix string   `json:"suffix"`
	Ranges []*Range `json:"ranges"`
}

type Range struct {
	ID  int32 `json:"id"`
	Min int32 `json:"min"`
	Max int32 `json:"max"`
}

type ListCondition struct {
	List []string `json:"list"`
}

func getChairSearchCondition() (*ChairSearchCondition, error) {
	var cond ChairSearchCondition
	if err := json.Unmarshal(rawChairConditionJSON, &cond); err != nil {
		return nil, err
	}
	return &cond, nil
}

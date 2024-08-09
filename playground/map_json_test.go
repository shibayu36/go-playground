package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type BoneName string

const (
	BoneNameHips  BoneName = "hips"
	BoneNameSpine BoneName = "spine"
	BoneNameChest BoneName = "chest"
	BoneNameNeck  BoneName = "neck"
	BoneNameHead  BoneName = "head"
)

type HumanoidBone struct {
	Node int `json:"node"`
}

type HumanBones map[BoneName]*HumanoidBone

func TestMapJSON(t *testing.T) {
	jsonStr := `{
        "hips": {"node": 1},
        "spine": {"node": 2},
        "chest": {"node": 3}
    }`

	var bones HumanBones
	if err := json.Unmarshal([]byte(jsonStr), &bones); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Unmarshaled:", bones)

	for name, bone := range bones {
		fmt.Printf("%s: %+v\n", name, bone)
	}
}

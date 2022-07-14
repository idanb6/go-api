package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type Execise struct {
	BodyPart  string `json:"bodyPart"`
	Equipment string `json:"equipment"`
	GifUrl    string `json:"gifUrl"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	Target    string `json:"target"`
}

func getBodyPart() []string {
	return []string{"waist",
		"upper legs",
		"back",
		"lower legs",
		"chest",
		"upper arms",
		"cardio",
		"shoulders",
		"lower arms",
		"neck",
	}
}

func getTarget() []string {
	return []string{"abs",
		"quads",
		"lats",
		"calves",
		"pectorals",
		"glutes",
		"hamstrings",
		"adductors",
		"triceps",
		"cardiovascular system",
		"spine",
		"upper back",
		"biceps",
		"delts",
		"forearms",
		"traps",
		"serratus anterior",
		"abductors",
		"levator scapulae"}
}

func getEquipment() []string {
	return []string{"body weight",
		"cable",
		"leverage machine",
		"assisted",
		"medicine ball",
		"stability ball",
		"band",
		"barbell",
		"rope",
		"dumbbell",
		"ez barbell",
		"sled machine",
		"upper body ergometer",
		"kettlebell",
		"olympic barbell",
		"weighted",
		"bosu ball",
		"resistance band",
		"roller",
		"skierg machine",
		"hammer",
		"smith machine",
		"wheel roller",
		"stationary bike",
		"tire",
		"trap bar",
		"elliptical machine",
		"stepmill machine",
	}
}

func getExecise() ([]Execise, error) {
	jsonFile, err := os.Open("exerciseDB.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	var execises []Execise
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(byteValue, &execises)
	return execises, nil
}

func findById(id string, execises []Execise) Execise {
	for _, val := range execises {
		if val.Id == id {
			return val
		}
	}
	return Execise{}
}

func findByType(t, k string, execises []Execise) []Execise {
	var ex []Execise
	for _, val := range execises {
		switch t {
		case "bodyPart":
			if val.BodyPart == k {
				ex = append(ex, val)
			}
		case "equipment":
			if val.Equipment == k {
				ex = append(ex, val)
			}
		case "target":
			if val.Target == k {
				ex = append(ex, val)
			}
		case "serch":
			if strings.Contains(val.Name, k) {
				ex = append(ex, val)
			}

		}

	}
	return ex
}

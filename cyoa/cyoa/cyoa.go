package cyoa

import (
	"encoding/json"
	"fmt"
	"os"
)

type Next struct {
	Text string `json:"text,omitempty" bson:"text" yaml:"text"`
	Arc  string `json:"arc,omitempty" bson:"arc" yaml:"arc"`
}

type Cyoa struct {
	Title   string   `json:"title,omitempty" bson:"title" yaml:"title"`
	Story   []string `json:"story,omitempty" bson:"story" yaml:"story"`
	Options []Next   `json:"options,omitempty" bson:"options" yaml:"options"`
}

type Adventure map[string]Cyoa

func ParseJson(fileName string) (Adventure, error) {
	var adv Adventure
	data, err := os.ReadFile(fileName)
	if err != nil {
		return adv, err
	}
	err = json.Unmarshal(data, &adv)
	if err != nil {
		return adv, err
	}
	return adv, nil
}

func StartCyoa(data Adventure) error {
	arc := "intro"
	for {
		d := Cyoa{}
		if _, ok := data[arc]; !ok {
			return fmt.Errorf("invalid data")
		}
		d = data[arc]
		fmt.Printf("\nTitle: %s\n", d.Title)
		fmt.Printf("Story: %+v\n", d.Story[0])
		if len(d.Options) == 0 {
			break
		}
		fmt.Println("\n\nChoose any of the arc below")
		options := map[int]string{}
		for idx, value := range d.Options {
			fmt.Println("Option:", idx)
			fmt.Printf("\t\tArc: %s\n", value.Arc)
			fmt.Printf("\t\tDescription: %s\n", value.Text)
			options[idx] = value.Arc
		}
		var num int
		fmt.Print("\nType your options here>")
		fmt.Scanf("%d\n", &num)
		if _, ok := options[num]; !ok {
			return fmt.Errorf("invalid option selected")
		}
		arc = options[num]
	}
	return nil
}

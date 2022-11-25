package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var URL = "https://github.com/Disploy/create-disploy-app/archive/refs/heads/main.zip"
var project string

func main() {

	inputModel := InputModel()

	if m, ok := inputModel.(InputStruct); ok && m.textInput.Value() != "" {
		project = m.textInput.Value()
	}

	fmt.Print("> Downloading list of templates...\n\n")

	err := DownloadFile(".disploy.zip", URL)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = UnzipFile(".disploy.zip", ".disploy")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	files, err := ioutil.ReadDir(".disploy/create-disploy-app-main/assets")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, f := range files {
		// read the readme.md file and get the first heading

		readme, err := ioutil.ReadFile(".disploy/create-disploy-app-main/assets/" + f.Name() + "/README.md")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		heading := GetFirstHeading(string(readme))

		Choices = append(Choices, Template{
			Name: heading,
			Path: f.Name(),
		})

	}

	choiceModel := ChoiceModel()

	if m, ok := choiceModel.(OptionStruct); ok && m.choice.Path != "" {
		Copy(m.choice.Path, project)
	}
}

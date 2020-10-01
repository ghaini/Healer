package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Project struct {
	Up   Up   `json:"up"`
	Down Down `json:"down"`
}

type Up struct {
	Commands []string `json:"commands"`
}

type Down struct {
	Commands []string `json:"commands"`
}

func ReadProject(name string) (Project, error) {
	var project Project
	home, err := os.UserHomeDir()
	if err != nil {
		return Project{}, err
	}
	path := home + "/.healer/" + name + ".json"
	file, err := ioutil.ReadFile(path)
	err = json.Unmarshal(file, &project)
	return project, err
}

func SaveProject(name string, project Project) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := home + "/.healer/" + name + ".json"
	jsonString, err := json.Marshal(project)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, jsonString, 0777)
	return err
}

func CreateProject(name string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := home + "/.healer/" + name + ".json"
	err = ioutil.WriteFile(path, []byte("{}"), 0777)
	return err
}
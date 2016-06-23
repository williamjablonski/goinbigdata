package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Cluster struct {
	Name       string   `json:"name"`
	DataCentre string   `json:"datacentre"`
	Nodes      []string `json:"nodes"`
}

type Configuration struct {
	Clusters    []Cluster `json:"clusters"`
	MinReplicas int       `json:"min_replicas"`
	MaxReplicas int       `json:"max_replicas"`
}

func saveConfig(c Configuration, filename string) error {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0644)
}

func loadConfig(filename string) (Configuration, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return Configuration{}, err
	}

	var c Configuration
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return Configuration{}, err
	}

	return c, nil
}

func main() {
	err := saveConfig(createMockConfig(), "config.json")
	if err != nil {
		panic(err)
	}

	c, err := loadConfig("config.json")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", c)
}

func createMockConfig() Configuration {
	return Configuration{
		Clusters: []Cluster{
			Cluster{
				Name:       "Dev",
				DataCentre: "Local",
				Nodes:      []string{"dev1.company.com", "dev2.company.com"},
			},
			Cluster{
				Name:       "Prod",
				DataCentre: "Amazon",
				Nodes:      []string{"prd1.company.com", "prd2.company.com", "prd3.company.com"},
			},
		},
		MinReplicas: 1,
		MaxReplicas: 5,
	}
}

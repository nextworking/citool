package cit

import (
	"encoding/json"
	"fmt"
	"github.com/blang/semver"
	"io/ioutil"
	"os"
)

type Metadata struct {
	Name                   string      `json:"name"`
	Version                string      `json:"version"`
	Author                 string      `json:"author"`
	Summary                string      `json:"summary"`
	License                string      `json:"license"`
	Source                 string      `json:"source"`
	Dependencies           interface{} `json:"dependencies"`
	OperatingsystemSupport interface{} `json:"operatingsystem_support"`
	Requirements           interface{} `json:"requirements"`
	PdkVersion             string      `json:"pdk-version"`
	TemplateUrl            string      `json:"template-url"`
	TemplateRef            string      `json:"template-ref"`
}


var MyJson Metadata


func GetMetadataVersion(f string) semver.Version {

	file, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in metadata.json: %v\n", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &MyJson)
	metaVer, err := semver.Parse(MyJson.Version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in metadata.json: %v\n", err)
		os.Exit(1)
	}
	return metaVer

}

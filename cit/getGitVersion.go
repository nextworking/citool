package cit

import (
	"fmt"
	"os/exec"
	"strings"
)

var ReturnGitTag string


func GetGitTag(d string) string {

	cmd := exec.Command("git", "describe", "--abbrev=0")
	cmd.Dir = d
	out, _ := cmd.CombinedOutput()
	if strings.Contains(string(out), "No names found") {
		fmt.Println("no tags!")
		ReturnGitTag = "0.0.0"
	} else {
		ReturnGitTag = strings.TrimRight(string(out), "\n")
	}

	return ReturnGitTag
}

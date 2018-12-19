package cit

import (
	"fmt"
	"os"
	"os/exec"
)


func SetGitTag(d string, ver string) string {

	cmdTag := exec.Command("git", "tag", "-a", ver, "-m", "\"gitlab ci tag\"")
	cmdTag.Dir = d
	outCmdTag, err := cmdTag.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in setTag: %v\n", outCmdTag)
		os.Exit(1)
	}

	cmdPushTag := exec.Command("git", "push", "--tags")
	cmdPushTag.Dir = d
	outPushTag, err := cmdPushTag.CombinedOutput()
	fmt.Printf("path: %v\n", string(outPushTag))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in pushTag: %v\n", err)
		os.Exit(1)
	}

	return string(outPushTag)
}

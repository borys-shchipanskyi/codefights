package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var (
	cmd = new(cmdHelper)
	git = new(gitHelper)
)

type cmdHelper struct {
}

func (cmd cmdHelper) run(command string, args []string) ([]byte, error) {
	strArgs := fmt.Sprintf("%v", args)
	log.Printf("%s %s\n", command, strArgs[1:len(strArgs)-1])

	return exec.Command(command, args...).CombinedOutput()
}

type gitHelper struct {
	cmd string
}

func (h *gitHelper) init() {
	h.cmd = "git"

}

func (h *gitHelper) fatchAll() error {
	args := []string{"fetch", "--all"}

	if out, err := cmd.run(h.cmd, args); err != nil {
		return fmt.Errorf("%s %v", string(out), err)
	}
	return nil
}

func (h *gitHelper) checkout(branch string) error {
	args := []string{"checkout", branch}
	if out, err := cmd.run(h.cmd, args); err != nil {
		return fmt.Errorf("%s %v", string(out), err)
	}
	// pull last changes
	if err := h.pull(); err != nil {
		return err
	}
	return nil
}

func (h *gitHelper) branch() error {
	args := []string{"branch"}

	out, err := cmd.run(h.cmd, args)
	if err != nil {
		return fmt.Errorf("%s %v", string(out), err)
	}
	return nil
}

func (h *gitHelper) pull() error {
	args := []string{"pull"}

	out, err := cmd.run(h.cmd, args)
	if err != nil {
		return fmt.Errorf("%s %v", string(out), err)
	}
	return nil
}

func (h *gitHelper) getLastCommitInBranch(branch string) (string, error) {
	// checkout to the branch
	if err := h.checkout(branch); err != nil {
		return "", err
	}

	args := []string{"log", "--format=%h", "-n", "1"}
	out, err := cmd.run(h.cmd, args)
	if err != nil {
		return "", fmt.Errorf("%s %v", string(out), err)
	}

	return string(out), nil
}

func (h *gitHelper) getCommitsHistory(branch, commit string, skipMsg []string) ([]string, error) {
	if err := h.checkout(branch); err != nil {
		return []string{}, err
	}
	args := []string{"log", commit + "..HEAD", "--format=%s"}
	out, err := cmd.run(h.cmd, args)
	if err != nil {
		return nil, fmt.Errorf("%s %v", string(out), err)
	}
	result := []string{}
	for _, line := range strings.Split(string(out), "\n") {
		if len(line) > 0 && !isSkip(skipMsg, line) {
			result = append(result, line)
		}
	}
	return result, nil
}

func isSkip(skipMsg []string, src string) bool {
	for _, msg := range skipMsg {
		if strings.Contains(src, msg) {
			return true
		}
	}
	return false
}

func main() {
	// init git and fetch all changes
	git.init()
	git.fatchAll()

	// parse flags
	var releaseBranch string
	flag.StringVar(&releaseBranch, "release", "", "Required. Specifies the release branch name.")

	var developBranch string
	flag.StringVar(&developBranch, "develop", "", "Required. Specifies the develop branch name.")

	var skipMsgsFlag string
	flag.StringVar(&skipMsgsFlag, "skipMsgs", "", "Optimal. Provide list of skip msgs list, comma separator.")
	// Parse command-line flags
	flag.Parse()

	skipMsgs := getSkipMsgs(skipMsgsFlag)

	// log input flags values
	log.Printf("Release branch: %s", releaseBranch)
	log.Printf("Develop branch: %s", developBranch)
	log.Printf("Masages for skip: %v", skipMsgs)

	changeLog := getChangeLog(releaseBranch, developBranch, skipMsgs)
	printChangeLog(changeLog)

}

func getSkipMsgs(flagValue string) []string {
	msgs := []string{}
	for _, val := range strings.Split(flagValue, ",") {
		if len(val) > 0 {
			msgs = append(msgs, strings.TrimSpace(val))
		}
	}
	return msgs
}

func getChangeLog(releaseBranch, developBranch string, skipMsgs []string) []string {
	commit, err := git.getLastCommitInBranch(releaseBranch)
	if err != nil {
		log.Fatal(err)
	}

	commit = strings.TrimSuffix(commit, "\n")
	log.Printf("Branch: '%s', last commit: '%s'", releaseBranch, commit)

	// get commits history
	res, err := git.getCommitsHistory(developBranch, commit, skipMsgs)
	if err != nil {
		log.Fatal(err)
	}

	return res

}

func printChangeLog(changeLog []string) {
	log.Println("####################### CHANGE LOG #######################")
	for _, msg := range changeLog {
		log.Println(msg)
	}
	log.Println("####################### CHANGE LOG END ####################")
}

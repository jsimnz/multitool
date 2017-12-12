package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/rigelrozanski/common"
	"github.com/spf13/cobra"
)

// Lock2yamlCmd represents the lock2yaml command
var (
	GitCmd = &cobra.Command{
		Use:   "git",
		Short: "git tricks",
	}
	SetPullCmd = &cobra.Command{
		Use:   "setpull",
		Short: "set the pull to origin upstream head",
		RunE:  setPullCmd,
	}
	AddCommitPushCmd = &cobra.Command{
		Use:   "acp [message]",
		Short: "add -u, commit -m [message], push origin [cur branch]",
		RunE:  addCommitPushCmd,
	}
	DuplicateCmd = &cobra.Command{
		Use:   "dup",
		Short: "duplicate the repo to [thisreponame]2, cd there",
		RunE:  duplicateCmd,
	}
)

func init() {
	GitCmd.AddCommand(
		SetPullCmd,
		AddCommitPushCmd,
		DuplicateCmd,
	)
	RootCmd.AddCommand(GitCmd)
}

func setPullCmd(cmd *cobra.Command, args []string) error {
	command := fmt.Sprintf("git branch --set-upstream-to=origin/%[1]v %[1]v", getBranch())
	output, err := common.Execute(command)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n%v\n", command, output)
	return nil
}

func addCommitPushCmd(cmd *cobra.Command, args []string) error {

	if len(args) < 1 {
		return fmt.Errorf("Please include a commit message")
	}

	//combine args into one commit message
	message := strings.Join(args[:], " ")

	command1 := fmt.Sprintf("git add -u")
	output1, err := common.Execute(command1)
	fmt.Printf("%v\n%v\n", command1, output1)
	if err != nil {
		return err
	}

	command2 := fmt.Sprintf("git commit -m \"%v\"", message)
	output2, err := common.Execute(command2)
	fmt.Printf("%v\n%v\n", command2, output2)
	if err != nil {
		return err
	}

	command3 := fmt.Sprintf("git push origin %v", getBranch())
	output3, err := common.Execute(command3)
	fmt.Printf("%v\n%v\n", command3, output3)
	if err != nil {
		return err
	}
	return nil
}

func duplicateCmd(cmd *cobra.Command, args []string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	newDir := dir + "2"

	command2 := fmt.Sprintf("cp -a %v %v", dir, newDir)
	output2, err := common.Execute(command2)
	fmt.Printf("%v\n%v\n", command2, output2)
	if err != nil {
		return err
	}

	return nil
}

//_______________________________________________________________________________________

func getBranch() string {
	branch, err := common.Execute("git rev-parse --abbrev-ref HEAD")
	if err != nil {
		panic(err)
	}
	return branch
}

package commands

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/rigelrozanski/common"
	"github.com/spf13/cobra"
)

// file commands
var (
	FileCmd = &cobra.Command{
		Use:   "file",
		Short: "file processing commands",
	}
	MirrorCmd = &cobra.Command{
		Use:   "mirror [prefix] [suffix]",
		Short: "mirror files with a number in them",
		Args:  cobra.ExactArgs(2),
		RunE:  mirrorCmd,
	}
)

func init() {
	FileCmd.AddCommand(MirrorCmd)
	RootCmd.AddCommand(FileCmd)
}

func mirrorCmd(cmd *cobra.Command, args []string) error {

	prefix, suffix := args[0], args[1]

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	// get the number max and min
	min, max := 10000000, 0
	opGetMinMax := func(path string) error {
		str := strings.Split(path, prefix)
		str = strings.Split(str[len(str)-1], suffix)
		if len(str) >= 1 {
			fileNo, err := strconv.Atoi(str[0])
			if err != nil { // just skip if it's not a number
				return nil
			}
			if min > fileNo {
				min = fileNo
			}
			if max < fileNo {
				max = fileNo
			}
		}
		return nil
	}
	common.OperateOnDir(dir, opGetMinMax)

	// iterate through the files from largest to smallest and add names
	newFileNo := max
	for fileNo := max; fileNo >= min; fileNo-- {
		newFileNo++
		cpFileName := fmt.Sprintf("%v%v%v", prefix, fileNo, suffix)
		newFileName := fmt.Sprintf("%v%v%v", prefix, newFileNo, suffix)
		cpPath := path.Join(dir, cpFileName)
		newPath := path.Join(dir, newFileName)
		common.Copy(cpPath, newPath)
	}

	fmt.Println("completed copy")
	return nil
}

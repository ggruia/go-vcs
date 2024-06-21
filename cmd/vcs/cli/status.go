package cli

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"go-vcs/cmd/vcs/utils"
	"io"
	"os"
	"sort"
)

const (
	noChangesMessage = "No changes in working directory."
	notStagedMessage = "Changes not staged for commit:"
	stagedMessage    = "Changes to be committed:"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "This allows you to display all tracked files status",
	Example: "vcs status",
	RunE: func(_ *cobra.Command, _ []string) error {
		return runStatusCommand(os.Stdout, indexPath)
	},
}

func runStatusCommand(writer io.Writer, stagingPath string) error {
	statsMap, err := utils.ReadFilesFromWorkingDir(".")
	if err != nil {
		return err
	}

	metadata := utils.UpdateMetadata(statsMap)
	if err = utils.WriteMetadata(stagingPath, metadata); err != nil {
		return err
	}

	stagingFilesInfo, err := getStagingInfo(stagingPath)
	if err != nil {
		return err
	}

	displayStatus(writer, stagingFilesInfo)
	return nil
}

// Display Format: | modifiedAt | filepath | status |
func displayStatus(writer io.Writer, filesInfo utils.FileInfoArr) {
	if len(filesInfo) == 0 {
		fmt.Println(noChangesMessage)
		return
	}

	notStagedTable := createTable(writer)
	stagedTable := createTable(writer)

	for _, f := range filesInfo {
		row := []string{string(f.Status), f.Path}
		statusColor := tablewriter.FgGreenColor
		if f.Status == utils.StatusModified {
			statusColor = tablewriter.FgBlueColor
		}
		if f.Staging {
			stagedTable.Rich(row, []tablewriter.Colors{{tablewriter.Bold, statusColor}, {}})
		} else {
			notStagedTable.Rich(row, []tablewriter.Colors{{tablewriter.Bold, statusColor}, {}})
		}

	}

	if stagedTable.NumLines() > 0 {
		fmt.Println(stagedMessage)
		stagedTable.Render()
	}

	if notStagedTable.NumLines() > 0 {
		fmt.Println(notStagedMessage)
		notStagedTable.Render()
	}
}

func getStagingInfo(stagingFile string) (utils.FileInfoArr, error) {
	metadata, err := utils.ReadMetadata(stagingFile)
	if err != nil {
		return nil, err
	}

	var filesInfo utils.FileInfoArr
	for _, m := range metadata {
		filesInfo = append(filesInfo, utils.FromFileMetadataToFileInfo(m))
	}

	sort.Sort(filesInfo)
	return filesInfo, nil
}

func createTable(writer io.Writer) *tablewriter.Table {
	table := tablewriter.NewWriter(writer)
	table.SetBorder(false)
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetTablePadding("\t")

	return table
}

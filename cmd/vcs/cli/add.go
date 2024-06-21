package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "This allows you to track all file status (created, modified)",
	Example: "vcs add main.go images/",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAddCommand(indexPath, args)
	},
}

func runAddCommand(stagingPath string, filePaths []string) error {
	//metadataMap, err := filedata.ReadMetadataMap(stagingPath)
	//if err != nil {
	//	return err
	//}
	//
	//determineStagingAreaFiles(filePaths, metadataMap)
	//
	//statusFilePtr, _ := openFile(statusFilePath)
	//defer statusFilePtr.Close()
	//
	//stagingFilePtr, _ := openFileAppendMode(stagingAreaFilePath)
	//defer stagingFilePtr.Close()
	//
	//// we need to truncate status folder in order to catch overwrite file status
	//if err := clearFileContent(statusFilePtr); err != nil {
	//	return err
	//}
	//
	//for _, metadata := range fileNameToMetadata {
	//	lineStr := metadata.toFormatFileMetadataForFile()
	//
	//	_, _ = statusFilePtr.WriteString(lineStr)
	//	if metadata.GoToStaging {
	//		_, _ = stagingFilePtr.WriteString(lineStr)
	//	}
	//}

	//files, err := utils.ReadFilesFromWorkingDir(".")
	//if err != nil {
	//	return err
	//}
	//
	//metadata := utils.UpdateMetadata(files)
	//if err = utils.WriteMetadata(stagingPath, metadata); err != nil {
	//	return err
	//}

	return nil
}

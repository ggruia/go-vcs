package cli

import (
	"flag"
	"fmt"
	"go-vcs/cmd/vcs/internal"
	"log"
)

func HandleAdd(addCmd *flag.FlagSet, args []string) {
	err := addCmd.Parse(args)
	if err != nil {
		return
	}
	
	for _, filePath := range addCmd.Args() {
		err := internal.AddFile(filePath)
		if err != nil {
			log.Fatalf("Failed to add file %s: %v", filePath, err)
		}
		fmt.Printf("File %s added.\n", filePath)
	}
}

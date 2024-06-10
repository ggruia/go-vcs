package cli

import (
	"flag"
	"fmt"
	"go-vcs/cmd/vcs/internal"
	"log"
	"os"
)

func HandleCommit(commitCmd *flag.FlagSet, commitMessage *string) {
	err := commitCmd.Parse(os.Args[2:])
	if err != nil {
		return
	}

	if *commitMessage == "" {
		log.Fatal("commit message is required")
	}
	err = internal.CreateCommit(*commitMessage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Changes committed.")
}

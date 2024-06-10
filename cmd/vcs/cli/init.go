package cli

import (
	"flag"
	"fmt"
	"go-vcs/cmd/vcs/internal"
	"log"
	"os"
)

func HandleInit(initCmd *flag.FlagSet) {
	err := initCmd.Parse(os.Args[2:])
	if err != nil {
		return
	}

	err = internal.InitRepo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Repository initialized.")
}

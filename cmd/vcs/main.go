package main

import (
	"flag"
	"go-vcs/cmd/vcs/cli"
	"log"
	"os"
)

func main() {
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	commitCmd := flag.NewFlagSet("commit", flag.ExitOnError)
	checkoutCmd := flag.NewFlagSet("checkout", flag.ExitOnError)

	// Commit message flag
	commitMessage := commitCmd.String("m", "", "Commit message")

	// Checkout commit ID flag
	checkoutCommitID := checkoutCmd.String("id", "", "Commit ID to checkout")

	if len(os.Args) < 2 {
		log.Fatal("expected 'init', 'add', 'commit', or 'checkout' subcommands")
	}

	switch os.Args[1] {
	case "init":
		cli.HandleInit(initCmd)
	case "add":
		cli.HandleAdd(addCmd, os.Args[2:])
	case "commit":
		cli.HandleCommit(commitCmd, commitMessage)
	case "checkout":
		cli.HandleCheckout(checkoutCmd, checkoutCommitID)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}

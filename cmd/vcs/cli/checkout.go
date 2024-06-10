package cli

import (
	"flag"
	"fmt"
	"go-vcs/cmd/vcs/internal"
	"log"
	"os"
)

func HandleCheckout(checkoutCmd *flag.FlagSet, checkoutCommitID *string) {
	err := checkoutCmd.Parse(os.Args[2:])
	if err != nil {
		return
	}

	if *checkoutCommitID == "" {
		log.Fatal("commit ID is required")
	}
	err = internal.Checkout(*checkoutCommitID)
	if err != nil {
		log.Fatalf("Failed to checkout commit %s: %v", *checkoutCommitID, err)
	}
	fmt.Printf("Checked out commit %s.\n", *checkoutCommitID)
}

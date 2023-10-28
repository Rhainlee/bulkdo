package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/rhainlee/bulkdo"
)

const BulkDoHome = ".bulkdo"
const DefaultItemFile = ".bulkdoitems"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a command.")
		os.Exit(1)
	}
	cmdName := os.Args[1]
	fmt.Printf("Execute command '%v'", cmdName)

	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cmdTplPath := filepath.Join(user.HomeDir, BulkDoHome, cmdName+".tpl")
	if _, err := os.Stat(cmdTplPath); os.IsExist(err) {
		fmt.Printf("Command Template path '%v' does not exist.", cmdTplPath)
		os.Exit(1)
	}

	if _, err := os.Stat(DefaultItemFile); os.IsExist(err) {
		fmt.Printf("DefaultItemFile path '%v' does not exist.", DefaultItemFile)
		os.Exit(1)
	}
	  
	cmdTplFile, cmdTplFileErr := os.Open(cmdTplPath)
	if cmdTplFileErr != nil {
		fmt.Println(cmdTplFileErr)
		os.Exit(1)
	}

	itemFile, itemFileErr := os.Open(DefaultItemFile)
	if itemFileErr != nil {
		fmt.Println(itemFileErr)
		os.Exit(1)
	}

	outs, outErr := bulkdo.BulkDo(cmdTplFile, itemFile)
	if outErr != nil {
		fmt.Println(outErr)
		os.Exit(1)
	}
	for _, out := range outs {
		fmt.Println(out)
	}
}

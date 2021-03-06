package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
)

var ImportCMD = cli.Command{
	Name:  "import",
	Usage: "import tenets from another lingo file",
	Description: `

  Import all tenet's from a hosted .lingo file
	"lingo import github.com/waigani/juju.lingo"

  This command expects the import path to end in ".lingo".

`[1:],
	Action: func(c *cli.Context) {
		fmt.Println("import not implemented")
	},
}

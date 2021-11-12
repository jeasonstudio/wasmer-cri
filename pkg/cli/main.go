package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Webassembly Open Container Interface CLI",
		Usage: "Pull and push webassembly images with OCI registry.",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

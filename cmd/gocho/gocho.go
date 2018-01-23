package main

import (
	"github.com/donkeysharp/gocho/pkg/cmds"
	"os"
)

func main() {
	cmds.New().Run(os.Args)
}

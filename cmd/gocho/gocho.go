package main

//go:generate go-bindata -o ../../assets/assets_gen.go -pkg assets ../../ui/build/...

import (
	"github.com/donkeysharp/gocho/pkg/cmds"
	"os"
)

func main() {
	cmds.New().Run(os.Args)
}

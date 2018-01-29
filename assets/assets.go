package assets

import (
	"github.com/elazarl/go-bindata-assetfs"
)

func AssetFS() *assetfs.AssetFS {
	return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "../../ui/build"}
}

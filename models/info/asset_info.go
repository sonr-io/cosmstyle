package info

import "github.com/sonr-io/cosmstyle/models/html"

type AssetInfo struct {
	Ticker    string
	Name      string
	IsDefault bool
	Icon      html.Icon
}

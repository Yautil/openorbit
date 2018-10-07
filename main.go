package main

import (
	"engo.io/engo"
	"github.com/Yautil/openorbit/scenes"
)

var (
	GlobalOpts = engo.RunOptions{
		NoRun:               false,
		Title:               "MyOpenOrbit",
		HeadlessMode:        false,
		Fullscreen:          false,
		Width:               1000,
		Height:              600,
		VSync:               false,
		NotResizable:        false,
		ScaleOnResize:       false,
		FPSLimit:            0,
		OverrideCloseAction: false,
		StandardInputs:      false,
		MSAA:                1,
		AssetsRoot:          "",
		MobileWidth:         0,
		MobileHeight:        0,
	}
)

func main() {
	engo.Run(GlobalOpts, &scenes.TestScene{})
}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/airtonix/rpg/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

var (
	name    = "App Name"
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	fmt.Printf("%s@%s, commit %s, built at %s", name, version, commit, date)

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowSizeLimits(300, 200, -1, -1)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	rand.Seed(time.Now().UTC().UnixNano())

	g := engine.NewGame(&game.Game{})

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}

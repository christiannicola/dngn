package main

import (
	"flag"
	internal "github.com/christiannicola/dngn/cmd/dngn/ebiten"
	"github.com/christiannicola/dngn/pkg/debug"
	"os"
	"runtime/pprof"
)

var (
	cpuProfile    = flag.Bool("cpu", false, "write cpu profile to file")
	memoryProfile = flag.Bool("memory", false, "write memory profile to file")
)

func main() {
	const width, height = 100, 30

	flag.Parse()

	if *cpuProfile == true {
		f, err := os.Create("cpu.prof")
		if err != nil {
			panic(err)
		}

		if err = pprof.StartCPUProfile(f); err != nil {
			panic(err)
		}

		defer pprof.StopCPUProfile()
	}

	l := debug.NewLogger(nil)
	l.SetPrefix("MAIN")
	l.SetLevel(debug.LogLevelDebug)

	l.Debug("constructing game")

	game, err := internal.NewGame(width, height, "dngn")
	if err != nil {
		panic(err)
	}

	if err = game.Renderer.Run(game.Draw, game.Update, 100*16, 30*32, "dngn"); err != nil {
		panic(err)
	}

	if *memoryProfile == true {
		f, err := os.Create("mem.prof")
		if err != nil {
			panic(err)
		}

		if err = pprof.WriteHeapProfile(f); err != nil {
			panic(err)
		}

		defer f.Close()
	}
}

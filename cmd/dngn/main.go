package main

import (
	"flag"
	internal "github.com/christiannicola/dngn/cmd/dngn/sdl"
	"os"
	"runtime/pprof"
)

var (
	cpuProfile    = flag.Bool("cpu", false, "write cpu profile to file")
	memoryProfile = flag.Bool("memory", false, "write memory profile to file")
)

func main() {
	const width, height, fps = 100, 30, 60
	const frameDelay = 1000 / fps

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

	game, err := internal.NewGame(width, height, fps, "dngn")
	if err != nil {
		panic(err)
	}

	for game.IsRunning {
		if err = game.HandleEvents(); err != nil {
			panic(err)
		}

		if err = game.Update(); err != nil {
			panic(err)
		}

		if err = game.Draw(); err != nil {
			panic(err)
		}
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

	game.Shutdown()
}

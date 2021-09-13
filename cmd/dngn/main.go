package main

import (
	"flag"
	internal "github.com/christiannicola/dngn/cmd/dngn/sdl"
	"github.com/veandco/go-sdl2/sdl"

	"os"
	"runtime/pprof"
)

var (
	cpuProfile    = flag.Bool("cpu", false, "write cpu profile to file")
	memoryProfile = flag.Bool("memory", false, "write memory profile to file")
)

func main() {
	const width, height, fps = 6 * 16, 6 * 9, 60
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

	var frameStart, frameTime int32 = 0, 0

	for game.IsRunning {
		frameStart = int32(sdl.GetTicks())

		if err = game.HandleEvents(); err != nil {
			panic(err)
		}

		if err = game.Update(); err != nil {
			panic(err)
		}

		if err = game.Draw(); err != nil {
			panic(err)
		}

		frameTime = int32(sdl.GetTicks()) - frameStart

		if frameDelay > frameTime {
			sdl.Delay(uint32(frameDelay - frameTime))
		}
		// println(1000 / (int32(sdl.GetTicks()) - frameStart))
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

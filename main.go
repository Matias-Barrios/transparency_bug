package main

import (
	"log"
	"os"

	"github.com/Matias-Barrios/transparency_bug/SDL"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	window, renderer, err := SDL.InitSDL()
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		panic(err)
		return
	}

	if err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Println(err)
		return
	}

	defer sdl.Quit()
	defer window.Destroy()
	defer renderer.Destroy()

	background := SDL.GetTexture(window, renderer, "backgrounds/sky.png", 0)
	block := SDL.GetTexture(window, renderer, "assets/blue.png", 50)
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				os.Exit(0)
				break
			}
			SDL.DrawStuff(renderer, background, 0, 0, 800, 600)
			SDL.DrawStuff(renderer, block, 200, 200, 100, 100)

			renderer.Present()
			renderer.Clear()
		}
	}
}

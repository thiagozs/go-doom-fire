package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/thiagozs/go-doom-fire/engine"
	"github.com/thiagozs/go-doom-fire/renderer"
)

func main() {

	cfg := engine.Config{
		Width:  120,
		Height: 40,
	}

	e := engine.New(cfg)
	r := renderer.NewTerminalRenderer(e)
	screen := r.GetScreen()

	defer screen.Fini()

	// tamanho inicial real
	w, h := screen.Size()
	e.Resize(w, h)

	running := true

	// captura Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		for running {

			select {
			case <-sigChan:
				running = false
				return
			default:
				ev := screen.PollEvent()

				switch ev := ev.(type) {

				case *tcell.EventKey:
					switch ev.Key() {
					case tcell.KeyEscape, tcell.KeyCtrlC:
						running = false
						return
					}

				case *tcell.EventResize:
					screen.Sync()

					w, h := screen.Size()
					e.Resize(w, h)
				}
			}
		}
	}()

	for running {
		e.Update()
		r.Render()
		time.Sleep(30 * time.Millisecond)
	}
}

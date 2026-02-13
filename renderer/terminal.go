package renderer

import (
	"github.com/gdamore/tcell/v2"
	"github.com/thiagozs/go-doom-fire/engine"
)

type TerminalRenderer struct {
	screen tcell.Screen
	engine *engine.Engine
}

func NewTerminalRenderer(e *engine.Engine) *TerminalRenderer {
	screen, _ := tcell.NewScreen()
	screen.Init()

	return &TerminalRenderer{
		screen: screen,
		engine: e,
	}
}

func (r *TerminalRenderer) Render() {
	r.screen.Clear()

	for y := 0; y < r.engine.Config.Height; y++ {
		for x := 0; x < r.engine.Config.Width; x++ {

			val := r.engine.Fire[y][x]

			if val <= 0 {
				continue
			}

			style := tcell.StyleDefault.Foreground(fireColor(val))
			r.screen.SetContent(x, y, '█', nil, style)
		}
	}

	r.screen.Show()
}

func (r *TerminalRenderer) GetScreen() tcell.Screen {
	return r.screen
}

func fireColor(v int) tcell.Color {
	switch {
	case v > 45:
		return tcell.Color228 // amarelo quente (não branco)
	case v > 38:
		return tcell.Color220
	case v > 30:
		return tcell.Color214
	case v > 22:
		return tcell.Color208
	case v > 16:
		return tcell.Color202
	case v > 10:
		return tcell.Color196
	case v > 5:
		return tcell.Color160
	case v > 2:
		return tcell.Color124
	default:
		return tcell.Color52
	}
}

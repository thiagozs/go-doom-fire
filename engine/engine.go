package engine

import (
	"math/rand"
)

type Engine struct {
	Config    Config
	Fire      [][]int
	Wind      int
	Intensity int
}

func New(cfg Config) *Engine {
	e := &Engine{
		Config:    cfg,
		Wind:      0,
		Intensity: 50,
	}

	e.allocate()
	return e
}

func (e *Engine) allocate() {
	e.Fire = make([][]int, e.Config.Height)
	for i := range e.Fire {
		e.Fire[i] = make([]int, e.Config.Width)
	}
}

func (e *Engine) Resize(width, height int) {
	e.Config.Width = width
	e.Config.Height = height

	e.Fire = make([][]int, height)
	for i := range e.Fire {
		e.Fire[i] = make([]int, width)
	}
}

func (e *Engine) Update() {

	if e.Config.Height <= 1 {
		return
	}

	// Base do fogo (linha inferior)
	for x := 0; x < e.Config.Width; x++ {
		// varia levemente a intensidade na base para quebrar faixas
		jitter := rand.Intn(5) - 2
		val := e.Intensity + jitter
		if val < 0 {
			val = 0
		}
		e.Fire[e.Config.Height-1][x] = val
	}

	// Propagação
	for y := 1; y < e.Config.Height; y++ {
		for x := 0; x < e.Config.Width; x++ {

			// fonte vem da linha abaixo com leve deslocamento lateral
			srcX := x + e.Wind + rand.Intn(3) - 1
			if srcX < 0 {
				srcX = 0
			} else if srcX >= e.Config.Width {
				srcX = e.Config.Width - 1
			}

			decay := rand.Intn(4)
			val := e.Fire[y][srcX] - decay
			if val < 0 {
				val = 0
			}

			// escreve na linha acima com drift lateral
			dstX := max(x-rand.Intn(2), 0)

			e.Fire[y-1][dstX] = val
		}
	}
}

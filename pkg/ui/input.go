package ui

var (
	ok     = Input{"ok"}
	cancel = Input{"cancel"}
	quit   = Input{"quit"}

	// NOTE (c.nicola): Directional input for the game & ui screens
	n  = Input{"n"}
	ne = Input{"ne"}
	e  = Input{"e"}
	se = Input{"se"}
	s  = Input{"s"}
	sw = Input{"sw"}
	w  = Input{"w"}
	nw = Input{"nw"}
)

type Input struct {
	name string
}

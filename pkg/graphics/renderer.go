package graphics

type (
	RenderCallback = func(Surface) error

	UpdateCallback = func() error

	Renderer interface {
		GetRendererName() string
		SetWindowIcon(path string)
		Run(r RenderCallback, u UpdateCallback, width, height int, title string) error
		IsDrawingSkipped() bool
		CreateSurface(surface Surface) (Surface, error)
		NewSurface(width, height int) Surface
		IsFullScreen() bool
		SetFullScreen(fullScreen bool)
		SetVSyncEnabled(vsync bool)
		IsVSyncEnabled() bool
		GetCursorPos() (int, int)
		CurrentFPS() float64
	}
)

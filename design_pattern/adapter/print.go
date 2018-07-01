package adapter

type (
	// Print print interface
	Print interface {
		PrintWeak()
		PrintStrong()
	}
)

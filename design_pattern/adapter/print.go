package adapter

// Role: Client

type (
	// Print print interface
	Print interface {
		PrintWeak()
		PrintStrong()
	}
)

package builder

type (
	// Builder builder interface
	Builder interface {
		MakeTitle(title string) string
		MakeString(str string) string
		MakeItems(items []string) string
		Close() string
	}
)

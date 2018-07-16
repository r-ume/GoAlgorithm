package builder

type (
	// TextBuilderImpl text builder implementation
	TextBuilderImpl struct {
	}
)

// NewTextBuilder returns new text builder implementation
func NewTextBuilder() Builder {
	return &TextBuilderImpl{}
}

// MakeTitle makes title
func (t *TextBuilderImpl) MakeTitle(title string) string {
	return "# " + title + "\n"
}

// MakeString makes string
func (t *TextBuilderImpl) MakeString(str string) string {
	return "## " + str + "\n"
}

// MakeItems make items
func (t *TextBuilderImpl) MakeItems(items []string) string {
	var result string
	for _, item := range items {
		result += "- " + item + "\n"
	}
	return result
}

// Close close
func (t *TextBuilderImpl) Close() string {
	return "\n"
}

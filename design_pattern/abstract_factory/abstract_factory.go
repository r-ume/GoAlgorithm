// AbstractFactory パターンとは、
// インスタンスの生成を専門に行うクラスを用意することで、
// 整合性を必要とされる一連のオブジェクト群を間違いなく生成するためのパターンです。

package abstractfactory

// Interfaces

// Factory interface
type Factory interface {
	CreateLink(caption, url string) Link
	CreateTray(caption string) Tray
	CreatePage(title, author string) Page
}

// Page << Tray << Item

// Item interface
type Item interface {
	ToString() string
}

// Link interface
type Link interface {
	Item
}

// Tray interface
type Tray interface {
	Item
	AddToTray(item Item)
}

// BaseTray struct
type BaseTray struct {
	Tray []Item
}

// AddToTray add to tray(list)
func (b *BaseTray) AddToTray(item Item) {
	b.Tray = append(b.tray, item)
}

// BasePage struct
// ItemインタフェースにあるToString関数を持つ構造体のスライスがコンテントに入る。
type BasePage struct {
	Content []Item
}

// AddToContent add to content
func (bp *BasePage) AddToContent(item Item) {
	bp.Content = append(bp.Content, item)
}

// MdLink << MdTray << MdPage

// MdLink 1. MdLink struct
type MdLink struct {
	Caption string
	URL     string
}

// ToString to string
func (m *Mdlink) ToString() string {
	return "[" + m.Caption + "](" + m.URL + ")"
}

// MdTray 2. MdTray struct
type MdTray struct {
	BaseTray
	Caption string
}

// ToString to string
func (md *MdTray) ToString() string {
	tray := "- " + md.Caption + "\n"
	for _, item := range md.Tray {
		tray += item.ToString() + "\n"
	}
	return tray
}

// MdPage 3. MdPage struct
type MdPage struct {
	BasePage
	Title  string
	Author string
}

// Output output
func (m *MdPage) Output() string {
	content := "title: " + m.Title + "\n"
	content += "author: " + m.Author

	for _, item := range m.Content {
		content += item.ToString() + "\n"
	}

	return content
}

// MdFactory struct
type MdFactory struct {
}

// CreateLink create link
func (mf *MdFactory) CreateLink(caption, url string) Link {
	return &MdLink{Caption, url}
}

// CreateTray create tray
func (mf *MdFactory) CreateTray(caption string) Tray {
	return &MdTray{Caption: caption}
}

// CreatePage create page
func (mf *MdFactory) CreatePage(title, author string) Page {
	return &MdPage{Title: title, Author: author}
}

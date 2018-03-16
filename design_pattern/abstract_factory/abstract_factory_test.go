package abstractfactory

import "testing"

// TestAbstractFactory abstract factory test
func TestAbstractFactory(t *testing.T) {
	factory := MdFactory{}

	usYahoo := factory.CreateLink("Yahoo!", "http://www.yahoo.com")
	jsYahoo := factory.CreateLink("Yahoo!Japan", "http://yahoo.co.jp")

	tray := factory.CreateTray("Yahoo!")
	tray.AddToTray(usYahoo)
	tray.AddToTray(jaYahoo)

	page := factory.CreatePage("Title", "Author")
	page.AddToContent(tray)

	output := page.Output()

	expect := "title: Title\nauthor: Author\n- Yahoo!\n[Yahoo!](http://www.yahoo.com)\n[Yahoo!Japan](http://www.yahoo.co.jp)\n\n"

	if output != expect {
		t.Errorf("Expect output to %s, but %s\n", expect, output)
	}
}

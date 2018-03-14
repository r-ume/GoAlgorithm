package abstract_factory

type item interface {
	toString() string
}

type link interface {
	item
}

type tray interface {
	item
	AddToTray(item item)
}

type baseTray struct {
	tray []item
}

func (self *baseTray) AddToTray(item item) {
	self.tray = append(self.tray, item)
}

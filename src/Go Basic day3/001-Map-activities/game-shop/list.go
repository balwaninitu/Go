package main

type printer interface {
	print()
}

type list []printer

func (l list) print() {
	for _, item := range l {
		item.print()
	}
}

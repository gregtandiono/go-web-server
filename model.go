package main

// Model struct serves as base model with default behaviours
type Model interface {
	fetch()
	save()
	update()
	destroy()
}

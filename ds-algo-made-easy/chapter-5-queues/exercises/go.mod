module example.com/exercises.go

go 1.20

replace example.com/stacks => ../../chapter-4-stacks/stacks

replace example.com/queue => ../queue

require (
	example.com/queue v0.0.0-00010101000000-000000000000
	example.com/stacks v0.0.0-00010101000000-000000000000
)

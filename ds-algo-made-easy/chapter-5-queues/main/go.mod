module example.com/main

go 1.20

replace example.com/queue => ../queue

require (
	example.com/exercises v0.0.0-00010101000000-000000000000
	example.com/queue v0.0.0-00010101000000-000000000000
)

require example.com/stacks v0.0.0-00010101000000-000000000000 // indirect

replace example.com/exercises => ../exercises

replace example.com/stacks => ../../chapter-4-stacks/stacks

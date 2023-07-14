module example.com/main

go 1.20

replace example.com/stacks => ../stacks

require (
	example.com/exercises v0.0.0-00010101000000-000000000000
	example.com/stacks v0.0.0-00010101000000-000000000000
)

replace example.com/exercises => ../exercises

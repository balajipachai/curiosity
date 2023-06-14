module example.com/main

go 1.20

replace example.com/lists => ../lists

require (
	example.com/doublylists v0.0.0-00010101000000-000000000000
	example.com/lists v0.0.0-00010101000000-000000000000
	example.com/queue v0.0.0-00010101000000-000000000000
	example.com/sets v0.0.0-00010101000000-000000000000
	example.com/stack v0.0.0-00010101000000-000000000000
	example.com/tuples v0.0.0-00010101000000-000000000000
)

replace example.com/doublylists => ../doublylists

replace example.com/sets => ../sets

replace example.com/tuples => ../tuples

replace example.com/queue => ../queue

replace example.com/stack => ../stack

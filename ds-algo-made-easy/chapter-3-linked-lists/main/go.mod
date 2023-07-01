module example.com/main

go 1.20

replace example.com/singlylist => ../singlylist

require (
	example.com/doublylist v0.0.0-00010101000000-000000000000
	example.com/singlylist v0.0.0-00010101000000-000000000000
)

replace example.com/doublylist => ../doublylist

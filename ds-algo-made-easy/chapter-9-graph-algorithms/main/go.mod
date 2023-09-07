module example.com/main

go 1.20

replace example.com/graphs => ../graphs

require example.com/graphs v0.0.0-00010101000000-000000000000

require example.com/stacks v0.0.0-00010101000000-000000000000 // indirect

replace example.com/stacks => ../../chapter-4-stacks/stacks

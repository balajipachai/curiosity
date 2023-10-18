module example.com/main

go 1.20

replace example.com/sorting => ../sorting

require example.com/sorting v0.0.0-00010101000000-000000000000

require example.com/trees v0.0.0-00010101000000-000000000000 // indirect

replace example.com/trees => ../../chapter-6-trees/trees

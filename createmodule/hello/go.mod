module createmodule/hello

go 1.22.3

replace createmodule/greet => ../greet

require createmodule/greet v0.0.0-00010101000000-000000000000

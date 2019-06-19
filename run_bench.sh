GOMAXPROCS=1 taskset 0xFF000000 go test -bench . -benchtime 300000000x -benchmem -timeout 720m

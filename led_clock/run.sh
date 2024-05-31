go test -bench=. -benchmem  -benchtime 10s -cpuprofile='cpu.prof' -memprofile='mem.prof'

pprof mem.prof

pprof cpu.prof



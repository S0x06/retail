
#性能测试
go test -test.bench=".*"

#性能分析
go test -bench=".*" -cpuprofile=cpu.profile ./


#覆盖率测试
go test -coverprofile=cover.out

#覆盖率分析
go tool cover -func=cover.out
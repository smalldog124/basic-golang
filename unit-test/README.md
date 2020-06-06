# unit test
unit test คือการทดสอบการทำงานของฟังก์ชันและมีอิสระจากฟังก์ชันอื่น การที่จะทำให้ฟังก์ชันอิสละจากกันอาจใช้ mock, stub, dummies ช่วยด้วย
<img src="https://miro.medium.com/max/1400/1*h0pVZsx7-rBLpthLlsQUnw.png">

[A Visual Tutorial on Every Type of Test You Can Write](https://medium.com/better-programming/a-visual-tutorial-on-every-type-of-test-you-can-write-ec9b83edcf35)

## command
รัน test 
```sh
$ go test ./unit-test/example1
```
test coverage
```sh
$ go test ./unit-test/example1 --cover
```
test coverage report
```sh
$ go test ./unit-test/example1 -coverprofile=coverage.out
$ go tool cover -html=coverage.out
```

รัน benchmarks
```sh
$ go test ./unit-test/example1 -bench=. -benchmem
```
เช็คการใช้งาน memory และ cpu
```sh
$ go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile cpuprofile.out
$ go tool pprof cpuprofile.out
```
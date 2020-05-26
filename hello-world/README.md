# hello world
เรามาเริ่มต้นด้วยสิ่งที่เราทำก่อนเริ่มเหมือนทุกๆ ภาษานั้น แสดงข้อความ "hello world"

## package ที่นำมาใช้
- [fmt](https://pkg.go.dev/fmt?tab=doc) เป็นเครื่องมือจัดการกับ formatted I/O ที่คล้ายกับ printf และ scanf ของภาษา C

## run program
```sh
$ go run hello-world.go
```

## build เป็น binary file 
```sh
$ go build hello-world.go
or
$ go build -o ./bin/hello-world hello-world.go
```

## run binary
```sh
$ ./hello-world
```
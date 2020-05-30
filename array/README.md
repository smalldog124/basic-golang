# array
array เป็น data structure ที่เป็นการจองพื้นที่ของ memory ที่กำหนดขนาดไว้เฉพาะเจาะจง

## ประกาศตัวแปร array
```go
// Long declaration
// zero values
var numbers [4]int

// Short declaration
// กำหนดค่าเริ่ม
friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
```

สิ่งที่โปรแกรมจอง memory ไว้ 4 ช่องที่เป็นการข้อมูลเป็น int เมื่อประกาศ [4]int
<img src="https://blog.golang.org/slices-intro/slice-array.png">
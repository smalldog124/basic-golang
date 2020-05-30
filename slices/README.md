# slices
  เป็นอีก 1 data structure คล้ายๆกับ array แต่ไม่กำหนดขนาดไว้ทำให้การจัดการข้อมูลได้ดีขึ้น จะเห็นบ่อยๆใน code ของคนอื่น

## ประกาศตัวแปร slices
```go
// Long declaration
// zero values
var name = make([]string,5,5)

// Short declaration
// กำหนดค่าเริ่ม
friends := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
```
ใน slices มีการเก็บความยาวและความจุ <br>
ptr คือ pointer เก็บตำแหน่งที่จะเข้าถึงข้อมูล <br>
len คือ <br> 
cap คือ <br>
<img src="https://blog.golang.org/slices-intro/slice-struct.png">

สิ่งที่โปรแกรมจอง memory ไว้ 5 ช่องที่เป็นการข้อมูลเป็น byte เมื่อประกาศ make([]byte,5,5)
<img src="https://blog.golang.org/slices-intro/slice-1.png">
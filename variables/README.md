# variables
## วิธีประกาศ Variable
1. <b>Long declaration</b> ใช้คำว่า var นำหน้าตามด้วย ชื่อ ตามด้วย Type ตัวภาษาจะเซตค่าให้เป็น default หรือเรียกว่า zero value
```go
var productName string
var price float64
var amount int
```

2. <b>Short declaration</b> ใช้ := เป็น key โดยมีรูปแบบ ชื่อ := ค่าของตัวแปร ตัวของภาษาจะใส่ Type ให้เราเอง
```go
productName := "huawei p20" // type string
price := 20000.00           // type float64
amount := 10                // typr int

```

3. <b>Multiple declaration</b> ใช้ var แล้วก็มีวงเล็บครอบ
```go
var (
      productName string
      price,amount = 20000.00,10
)
```
productName เป็น Type string มีค่าเป็น “”<br>
price เป็น Type float64 มีค่าเป็น 20000.00<br>
amount เป็น Type int มีค่าเป็น 10

4. <b>Short multiple declaration</b>
```go
prdoctName,prince,amount := "huawei p20",20000.00,10
```
productName เป็น Type string มีค่าเป็น “huawei p20”<br>
price เป็น Type float64 มีค่าเป็น 20000.00<br>
amount เป็น Type int มี่ค่าเป็น 10

### Zero Values
	Type Initialized Value
	Boolean false
	Integer 0
	Floating Point 0
	Complex 0i
	String "" (empty string)
	Pointer nil
# basis-golang
<img width="100" height="300" src="https://cdn.worldvectorlogo.com/logos/gopher.svg" >

## ความเป็นมาของภาษา
ภาษา Golang ถูกพัฒนาขึ้นมาในปี 2003 โดย McCabe และ Clark ทาง google ได้นำมาเขียนโปรแกรมที่แข่งเล่น [AlphaGo](https://en.wikipedia.org/wiki/AlphaGo)

ต่อมาในปี 2007 google ได้มาพัฒนาต่อในเรื่อง munticore,networked machines โดยมีการคุยกันว่าต้องการความสามารถของภาษาอื่นที่ใช้งานอยู่ใน google มารวมกัน
- เป็นภาษาที่ต้องกำหนดชนิดของข้อมูล(static type) และ มีประสิทธิภาพในการทำงาน(run-time) เหมือน C++
- ตัวภาษาอ่านง่าย และ ใช้งานง่าย เหมือน Python หรือ JavaScript
- การทำงานที่เกี่ยวกับ networking และ multiprocessing ที่ประสิธิภาพสูง

ในเดือน พฤศจิกายน 2006 ทาง google ได้ปล่อยตัวภาษาใน developer ใช้งาน และในเดือน มีนาคม 2012 ได้ปล่อยเวอร์ชั่น 1.0 ออกมา ซึ่งปัจจุบันเป็นเวอร์ชั่น 1.14 (ณ วันที่ 23 พฤษาคม 2020)

อ้างอิง: [wiki Golang](https://en.wikipedia.org/wiki/Go_(programming_language))
## ทำไมต้องภาษา Golang + แนวคิดของภาษา
<b>แนวคิดของภาษา Go </b>

<i> "Build fast, reliable, and efficient software at scale" </i>

ตัวภาษาเกิดขึ้นใน google ที่นำมาจัดการเรื่อง network และ cloud ซึ่งจะโดดเด่นเรื่องการจัดการ network และ memory อีกทั้งยังตอบโจทย์เรื่องของ scale อีกด้วย ซึ่งสามารถทำไปใช้งาน [Cloud & Network Services](https://go.dev/solutions/cloud/), [Web Development](https://go.dev/solutions/webdev/), [Command-line Interfaces (CLIs)](https://go.dev/solutions/clis/) และ [Development Operations & Site Reliability Engineering (SRE)](https://go.dev/solutions/devops/)

<b> ทำไมต้องเลือกภาษา Go </b>
- ง่ายต่อการเรียนรู้


## หัวข้อที่จะได้เรียน
วันแรก
- life cycle
- Hello world
- go build
- variables
- for, for reage
- array,slices,maps
- structs

วันที่สอง
- function,multiple return values
- unit test
- methods
- goroutines
- channels
- demo web service by golang

# basis-golang
<img width="100" src="https://cdn.worldvectorlogo.com/logos/gopher.svg" />

## ความเป็นมาของภาษา
ภาษา Golang ถูกพัฒนาขึ้นมาในปี 2003 โดย McCabe และ Clark ทาง google ได้นำมาเขียนโปรแกรมที่แข่งเล่น [AlphaGo](https://en.wikipedia.org/wiki/AlphaGo)

ต่อมาในปี 2007 google ได้มาพัฒนาต่อในเรื่อง munticore,networked machines โดยมีการคุยกันว่าต้องการความสามารถของภาษาอื่นที่ใช้งานอยู่ใน google มารวมกัน
- เป็นภาษาที่ต้องกำหนดชนิดของข้อมูล(static type) และ มีประสิทธิภาพในการทำงาน(run-time) เหมือน C++
- ตัวภาษาอ่านง่าย และ ใช้งานง่าย เหมือน Python หรือ JavaScript
- การทำงานที่เกี่ยวกับ networking และ multiprocessing ที่ประสิธิภาพสูง

ในเดือน พฤศจิกายน 2006 ทาง google ได้ปล่อยตัวภาษาใน developer ใช้งาน และในเดือน มีนาคม 2012 ได้ปล่อยเวอร์ชั่น 1.0 ออกมา ซึ่งปัจจุบันเป็นเวอร์ชั่น 1.14 (ณ วันที่ 23 พฤษาคม 2020)

อ้างอิง: [wiki Golang](https://en.wikipedia.org/wiki/Go_(programming_language))
## ทำไมต้องภาษา Golang + แนวคิดของภาษา
### แนวคิดของภาษา Go

<i> "Build fast, reliable, and efficient software at scale" </i>

ตัวภาษาเกิดขึ้นใน google ที่นำมาจัดการเรื่อง network และ cloud ซึ่งจะโดดเด่นเรื่องการจัดการ network และ memory อีกทั้งยังตอบโจทย์เรื่องของ scale อีกด้วย ซึ่งสามารถทำไปใช้งาน [Cloud & Network Services](https://go.dev/solutions/cloud/), [Web Development](https://go.dev/solutions/webdev/), [Command-line Interfaces (CLIs)](https://go.dev/solutions/clis/) และ [Development Operations & Site Reliability Engineering (SRE)](https://go.dev/solutions/devops/)

### ทำไมต้องเลือกภาษา Go

<b> ง่ายต่อการเรียนรู้(Easy to learn) </b>
>ในตอนนั้น 1 ในสมาชิกของทีมคนที่รู้จัก Go แต่ผ่านไป 1 เดือน ทุกคนในทีมสามารถเขียน Go ได้และยังสร้าง enpoints ขึ้นมาได้ มันเป็นมีความยืดหยุ่นของการใช้งานและ แนวคิดของมันโคตรเจ๋ง(วิธีจัดการกับ native concurrency, การจัดการ... และ ลดต้นทุน+ความเร็ว) นั้นช่วยให้นักพัฒนาระหว่างทำงาน นอกจากนี้ตัวสัญลักษณ์น่ารักอีกด้วย 
>> [Jaime Enrique Garcia Lopez, Senior Software Development Manager at CapitalOne](https://medium.com/capital-one-tech/a-serverless-and-go-journey-credit-offers-api-74ef1f9fde7f)

<b> มีประสิทธิภาพ(Efficient) </b>
>เป็นภาษาที่มีขนาดเล็กและยัง compiles เร็วทำให้เหล่านักพัฒนามีความสุข Go เป็นภาษาที่มีขนาดเล็กและ complines เร็วทำให้เราไปสนใจปัญหาจริงๆ และไม่ต้องใช้เครื่องมืออื่นมาช่วยแก้ปัญหา Code, test, debug ระยะเวลาที่รู้ว่ามันทำงานถูกหรือไม่เร็วมากจนลืมว่าเรากำลังเขียนโปรแกรม ทำให้ code ดูสะอาดขึ้นมีแต่ business logic
>>[Clayton Coleman, lead engineer, Open Shift at RedHat](https://blog.gopheracademy.com/birthday-bash-2014/openshift-3-old-dogs-new-tricks/)

<b> ทรงพลัง(Powerful) </b>
> Go ทำได้ดีสำหรับเรื่องการ scel และ service ทีเขียนด้วย Go จะมีการใช้งาน memory ที่น้อย(footprints) เพราะว่า code ถูกทำให้เป็น binary ทำให้ง่ายต่อการย้ายและยัง build,deploy ง่ายด้วย ความสามารถเหลาะนั้น Go จึงยิบไปสร้างสำหรับ microservices คุณสามารถนำไป deploy อุปกรณ์ที่ประสิทธิภาพสุงได้ง่ายและที่สามารถจัดการกับ scale ได้ดีเช่น Kubernetes
>>[Matt Boyle, lead software engineer at Curve](https://www.computerweekly.com/blog/Open-Source-Insider/Golang-or-go-home-how-Curve-is-taking-Golang-to-new-heights)
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

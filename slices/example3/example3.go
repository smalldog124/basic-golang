package main

import "fmt"

func main() {
	var data []string
	lastCap := cap(data)

	// Append ~100k strings เข้าใน slice.
	for record := 1; record <= 1e5; record++ {

		value := fmt.Sprintf("Rec: %d", record)
		data = append(data, value)

		// ถ้า cap ของ slice เปลี่ยนให้แสดงผลลัพธ์ของการเปลี่ยนแปลง
		if lastCap != cap(data) {

			// คำนวน % ของการเปลี่ยน
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100

			lastCap = cap(data)

			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n", &data[0], record, cap(data), capChg)
		}
	}
}

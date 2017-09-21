package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := App{}
	a.Initialize("sysinfo", "sysinfo", "sysinfo", "127.0.0.1")

	d, _ := getMachineIDs(a.DB)
	dd, _ := json.Marshal(d)
	fmt.Println(string(dd))

	p, _ := getPackagesByMI(a.DB, `19e5190061f94c9498a9951f8df592a3`)
	pp, _ := json.Marshal(p)
	fmt.Println(string(pp))

	s, _ := getSysInfoByMI(a.DB, `19e5190061f94c9498a9951f8df592a3`)
	ss, _ := json.Marshal(s)
	fmt.Println(string(ss))

	a.Run(":8080")
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

type Infrastructure struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func main() {
	old_infra := flag.String("old", "", "infrastuktur lama tenant")
	new_infra := flag.String("new", "", "infrastuktur baru tenant")

	flag.Parse()

	if *old_infra == "" || *new_infra == "" {
		fmt.Println("argument yang berisi infrastruktur tidak boleh " +
			"kosong agar dapat melakukan migrasi data")
		return
	}

	var oldInfra, newInfra Infrastructure
	json.Unmarshal([]byte(*old_infra), &oldInfra)
	json.Unmarshal([]byte(*new_infra), &newInfra)

	// lakukan migrasi menggunakan
	// data koneksi DB infrastruktur lama ke baru...
	fmt.Println(oldInfra, newInfra)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

type Product struct {
	Name, Category string
	Price          float64
}

var Products = []Product{
	{"Kayak", "Watersports", 279},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}

type ConfigData struct {
	Username           string
	AdditionalProducts []Product
}

var Config ConfigData

func main() {
	raw_config, err := os.Open("./config.json")
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	defer raw_config.Close()
	decoder := json.NewDecoder(raw_config)
	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	for _, p := range Config.AdditionalProducts {
		Printfln("%v from cat %v costs %v", p.Name, p.Category, p.Price)
	}
	out_file, err := os.OpenFile("./config2.json", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	defer out_file.Close()
	encoder := json.NewEncoder(out_file)
	err = encoder.Encode(Products)
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	filepath.WalkDir(pwd, func(path string, d fs.DirEntry, err error) error {
		name := d.Name()
		info, _ := d.Info()
		mode := info.Mode().Perm()
		fmt.Println("Name:", name, "Mode:", mode)
		return err
	})

}

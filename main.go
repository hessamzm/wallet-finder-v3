package main

import (
	"bufio"
	"fmt"
	"github.com/foxnut/go-hdwallet"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"regexp"
	"time"
)

// address to find in txt
var path = "file.txt"

type BTC struct {
	Address string
	Mem     string
}

// / i use en word and just see 1xBTC address
func main() {

	// use such more case for ex : i try to finde 12 word mnemonic and 1001 normal index
	for {
		mnemonic, err := hdwallet.NewMnemonic(12, hdwallet.English)
		if err != nil {
			panic(err)
		}
		master, err := hdwallet.NewKey(
			hdwallet.Mnemonic(mnemonic),
		)
		if err != nil {
			panic(err)
		}

		wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.BTC), hdwallet.AddressIndex(1001))
		address, _ := wallet.GetAddress()
		// for test i use this address from txt file to test runned
		//	var address string = "1DxJW2b8nU5d8Z6bLDDfg8ueditTeN2H99"
		addresses := scanner()

		if contains(addresses, address) {

			// if i find anythink i save ite in db and print to screen
			fmt.Println(mnemonic)
			db := dbclient()
			db.AutoMigrate(&BTC{})
			db.Create(&BTC{Address: address, Mem: mnemonic})
			break
		}
	}
	for {
		mnemonic, err := hdwallet.NewMnemonic(16, hdwallet.English)
		if err != nil {
			panic(err)
		}
		master, err := hdwallet.NewKey(
			hdwallet.Mnemonic(mnemonic),
		)
		if err != nil {
			panic(err)
		}

		wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.BTC), hdwallet.AddressIndex(1001))
		address, _ := wallet.GetAddress()
		//	var address string = "1DxJW2b8nU5d8Z6bLDDfg8ueditTeN2H99"
		addresses := scanner()
		//addresses
		//fmt.Println(address)
		//fmt.Println(addresses)
		if contains(addresses, address) {
			fmt.Println(mnemonic)
			db := dbclient()
			db.AutoMigrate(&BTC{})
			db.Create(&BTC{Address: address, Mem: mnemonic})
			break
		}
	}
	for {
		mnemonic, err := hdwallet.NewMnemonic(24, hdwallet.English)
		if err != nil {
			panic(err)
		}
		master, err := hdwallet.NewKey(
			hdwallet.Mnemonic(mnemonic),
		)
		if err != nil {
			panic(err)
		}

		wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.BTC), hdwallet.AddressIndex(1001))
		address, _ := wallet.GetAddress()
		//	var address string = "1DxJW2b8nU5d8Z6bLDDfg8ueditTeN2H99"
		addresses := scanner()
		//addresses
		//fmt.Println(address)
		//fmt.Println(addresses)
		if contains(addresses, address) {
			fmt.Println(mnemonic)
			db := dbclient()
			db.AutoMigrate(&BTC{})
			db.Create(&BTC{Address: address, Mem: mnemonic})
			break
		}
	}
	// base in 12 en word and index 1
	for {
		mnemonic, err := hdwallet.NewMnemonic(12, hdwallet.English)
		if err != nil {
			panic(err)
		}
		master, err := hdwallet.NewKey(
			hdwallet.Mnemonic(mnemonic),
		)
		if err != nil {
			panic(err)
		}

		wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.BTC), hdwallet.AddressIndex(1))
		address, _ := wallet.GetAddress()
		//	var address string = "1DxJW2b8nU5d8Z6bLDDfg8ueditTeN2H99"
		addresses := scanner()
		//addresses
		//fmt.Println(address)
		//fmt.Println(addresses)
		if contains(addresses, address) {
			fmt.Println(mnemonic)
			db := dbclient()
			db.AutoMigrate(&BTC{})
			db.Create(&BTC{Address: address, Mem: mnemonic})
			break
		}
	}
	for {
		mnemonic, err := hdwallet.NewMnemonic(24, hdwallet.English)
		if err != nil {
			panic(err)
		}
		master, err := hdwallet.NewKey(
			hdwallet.Mnemonic(mnemonic),
		)
		if err != nil {
			panic(err)
		}

		wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.BTC), hdwallet.AddressIndex(1))
		address, _ := wallet.GetAddress()
		//var address string = "12ib7dApVFvg82TXKycWBNpN8kFyiAN1dr"
		addresses := scanner()
		//addresses
		//fmt.Println(address)
		//fmt.Println(addresses)
		if contains(addresses, address) {
			fmt.Println(mnemonic)
			db := dbclient()
			db.AutoMigrate(&BTC{})
			db.Create(&BTC{Address: address, Mem: mnemonic})
			break
		}
	}
	for {
		mnemonic, err := hdwallet.NewMnemonic(18, hdwallet.English)
		if err != nil {
			panic(err)
		}
		master, err := hdwallet.NewKey(
			hdwallet.Mnemonic(mnemonic),
		)
		if err != nil {
			panic(err)
		}

		wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.BTC), hdwallet.AddressIndex(1))
		address, _ := wallet.GetAddress()
		//var address string = "12ib7dApVFvg82TXKycWBNpN8kFyiAN1dr"
		addresses := scanner()
		//addresses
		//fmt.Println(address)
		//fmt.Println(addresses)
		if contains(addresses, address) {
			fmt.Println(mnemonic)
			db := dbclient()
			db.AutoMigrate(&BTC{})
			db.Create(&BTC{Address: address, Mem: mnemonic})
			break
		}
	}
}

// i use my sql to save mnemonic its so sample and easy to use
func dbclient() *gorm.DB {
	time.Sleep(50 * time.Millisecond) // تاخیر قبل از تلاش مجدد

	dsn := "root:root@tcp(127.0.0.1:8889)/finderv3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

// its example file scanner to fine my fv address from file
func scanner() []string {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// Create a new scanner to read the file
	scanner := bufio.NewScanner(file)

	// Compile the regular expression pattern to match Bitcoin addresses
	// in this case i just want address who start by 1
	pattern := `[1][a-zA-Z0-9]{25,34}`
	re := regexp.MustCompile(pattern)

	// Iterate through the file and extract the addresses
	addresses := make(map[string]bool)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			addresses[match] = true
		}
	}

	// Convert the map to a slice
	result := make([]string, 0, len(addresses))
	for addr := range addresses {
		result = append(result, addr)
	}
	return result

}

// very sample contains tools for find string from arry
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
			fmt.Println(v)
		}
	}
	return false
}

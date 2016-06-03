package main

//import (
//	"log"
//
//	"github.com/BurntSushi/toml"
//)
//
//type tomlConfig struct {
//	UrlKibana  string
//	Addr       string
//	Country    string
//	DecryptKey string
//	Signature  string
//	Roles      map[string]roles
//}
//
//type roles struct {
//	Id       int
//	Username string
//	Password string
//}
//
//func readConfig() {
//	if _, err := toml.DecodeFile("src/cmd/reverse_proxy/conf.ini", &config); err != nil {
//		log.Println(err)
//		return
//	}
//
//	log.Println("===============Config file===============")
//	log.Println("         country ==>", config.Country)
//	log.Println("       urlKibana ==>", config.UrlKibana)
//	log.Println("            addr ==>", config.Addr)
//	log.Println("      decryptKey ==>", config.DecryptKey)
//	log.Println("       signature ==>", config.Signature)
//	log.Println("  Admin username ==>", config.Roles["Admin"].Username)
//	log.Println("        Admin pw ==>", config.Roles["Admin"].Password)
//	log.Println("Partner username ==>", config.Roles["Partner"].Username)
//	log.Println("      Partner pw ==>", config.Roles["Partner"].Password)
//	log.Println("=========================================\n")
//}

package main

//
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
//	Admin      admin
//	Partner    partner
//}
//
//type admin struct {
//	Username string
//	Password string
//}
//
//type partner struct {
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
//	log.Println("  Admin username ==>", config.Admin.Username)
//	log.Println("        Admin pw ==>", config.Admin.Password)
//	log.Println("Partner username ==>", config.Partner.Username)
//	log.Println("      Partner pw ==>", config.Partner.Password)
//	log.Println("=========================================\n")
//
//}

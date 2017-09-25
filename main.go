package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	netPort, proto, webPort            string
	dbuser, dbpassword, dbname, dbhost string
	configFile                         = flag.String("configFile", "config.json", "The sysinfo-server confgiguration file.")
	a                                  = App{}
)

func init() {
	flag.Parse()
	viper.SetConfigFile(*configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
	netPort = viper.GetString("net.port")
	proto = viper.GetString("net.proto")
	webPort = viper.GetString("web.port")
	dbuser := viper.GetString("db.user")
	dbpassword := viper.GetString("db.password")
	dbname := viper.GetString("db.dbname")
	dbhost := viper.GetString("db.host")
	a.Initialize(dbuser, dbpassword, dbname, dbhost)
	ensureTableExists()
	log.Println(netPort, proto, webPort, " is used")
}
func main() {
	go startNetServer()
	a.Run(":" + webPort)
}

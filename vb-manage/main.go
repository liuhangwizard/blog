package main

import (
	"fmt"
	_ "vb-server/init"
	VBServer "vb-server/server"
)

func main(){
	fmt.Println("Volute Blog Manage Server Starting...");
	VBServer.RunServer()
}

package go_util

import (
	"fmt"
	"io/ioutil"
)

type banner struct {
	version string
	serverName string
	mode string
}

func NewBanner(serverName , version ,mode string)*banner{
	return &banner{
		version:version,
		serverName:serverName,
		mode:mode,
	}
}

func(b *banner)Print(){
	fmt.Println("")
	data, err := ioutil.ReadFile("./banner.txt")
	if err == nil {
		fmt.Printf("%s\n",string(data))
	}
	fmt.Printf(":: %v ::	(v%v %v)\n", b.serverName, b.version, b.mode)
	fmt.Println("")
}
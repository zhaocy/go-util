package go_util

import (
	"fmt"
	"testing"
)

func TestGenSonyflake(t *testing.T)  {
	GenSonyflake()
}

func TestUniqueId(t *testing.T) {
	fmt.Println(UniqueId())
}

func TestCreateCaptcha(t *testing.T) {
	for i:=0;i<30;i++{
		captcha := Gen8Digit()
		fmt.Println(captcha)
	}
}

func TestGenUUIDv4(t *testing.T){
	genUUIDv4()
}

func TestGenSid(t *testing.T) {
	GenSid()
}

func TestGen8Digit(t *testing.T) {
	fmt.Println(Gen8Digit())
}

func TestGenBetterGUID(t *testing.T){
	genBetterGUID()
}

func TestGenKsuid(t *testing.T){
	GenKsUid()
}

func TestGenXid(t *testing.T){
	fmt.Println(GenXid())
}
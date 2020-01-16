package go_util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/chilts/sid"
	"github.com/kjk/betterguid"
	"github.com/rs/xid"
	"github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
	"io"
	"log"
	rmath "math/rand"
	"time"
)

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

func GenXid() string{
	return xid.New().String()
}

func GenKsUid() string{
	return ksuid.New().String()
}

func genBetterGUID() {
	id := betterguid.New()
	fmt.Printf("github.com/kjk/betterguid:   %s\n", id)
}

/*func genUlid() {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	fmt.Printf("github.com/oklog/ulid:       %s\n", id.String())
}*/

func GenSonyflake() uint64{
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	// Note: this is base16, could shorten by encoding as base62 string
	fmt.Printf("github.com/sony/sonyflake:   %x\n", id)
	return id
}

func GenSid() {
	id := sid.Id()
	fmt.Printf("github.com/chilts/sid:       %s\n", id)
}

func genUUIDv4() {
	id := uuid.NewV4()
	fmt.Printf("github.com/satori/go.uuid:   %s\n", id)
}

func Gen8Digit() string {
	return fmt.Sprintf("%08v", rmath.New(rmath.NewSource(time.Now().UnixNano())).Int31n(100000000))
}


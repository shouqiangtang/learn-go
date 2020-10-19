package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func extractAddress2(userExt string) string {
	var res interface{}
	if err := json.Unmarshal([]byte(userExt), &res); err != nil {
		log.Println(err)
		return ""
	}

	ue := res.(map[string]interface{})
	if address, ok := ue["address"]; ok {
		addr, err := u2s(address.(string))
		if err == nil {
			return addr
		}
	}
	return ""
}

func u2s(form string) (to string, err error) {
	bs, err := hex.DecodeString(strings.Replace(form, `\u`, ``, -1))
	if err != nil {
		return
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		binary.Read(br, binary.BigEndian, &r)
		to += string(r)
	}
	return
}

func main() {

	f, err := os.Open("/Users/tangshouqiang/quanzhoushi.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		strs := strings.Split(line, "\t")

		extractAddress := extractAddress2(strs[1])

		pt := ""
		if strs[2] == "12" {
			pt = "实物奖品"
		} else if strs[2] == "1" {
			pt = "电影票"
		}

		// if pt != "" && strs[0] != "" {
		fmt.Printf("%s,%s,%s\n", pt, strs[0], extractAddress)
		// }
	}
}

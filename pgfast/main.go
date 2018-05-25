package main

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/coreos/etcd/client"
	"github.com/koding/multiconfig"
)

type SsLocalConfig struct {
	Local_port      int        `json:"local_port"`
	Server_password [][]string `json:"server_password"`
}

type PgConfig struct {
	Subscribe_url string `required:"true"`
}

func main() {
	m := multiconfig.NewWithPath("./pgconfig.toml")

	config := new(PgConfig)

	m.Load(config)
	m.MustLoad(config)

	ss_local_config := get_ss_local_config(config.Subscribe_url)
	set_etcd(ss_local_config)

	ticker := time.NewTicker(time.Hour)
	for _ = range ticker.C {
		ss_local_config := get_ss_local_config(config.Subscribe_url)
		set_etcd(ss_local_config)
	}
}

func get_ss_local_config(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error in response")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	sDec, _ := b64.StdEncoding.DecodeString(string(body))
	strSlice := strings.Split(string(sDec), "\n")

	ss_config := SsLocalConfig{}

	ss_config.Local_port = 1080

	for _, v := range strSlice {
		if strings.HasPrefix(v, "ssr://") {
			ssrDecode, _ := b64.StdEncoding.DecodeString(v[6:])
			reg := regexp.MustCompile(`(?P<ip>.*?):(?P<port>\d+):(\w+):(?P<method>.*?):(.*?):(?P<pass>.*?)\/`)
			arr := reg.FindStringSubmatch(string(ssrDecode))
			if len(arr) == 7 {
				// arr := strings.Split(string(ssrDecode), ":")
				// fmt.Println(arr[0], arr[1], arr[5])
				//非常奇怪，为什么后面加个"="就获得正确结果呢？
				//有问题的字符串:M0JoQURKOUo
				pwd, _ := b64.StdEncoding.DecodeString(arr[6] + "=")
				server := []string{arr[1] + ":" + arr[2], string(pwd), arr[4]}
				ss_config.Server_password = append(ss_config.Server_password, server)
			}

		}
	}
	x, _ := json.MarshalIndent(ss_config, "", "    ")
	return string(x)
}

func set_etcd(value string) {
	cfg := client.Config{
		Endpoints: []string{"http://localhost:2379"},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)

	resp, err := kapi.Set(context.Background(), "/ss-local", value, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Set is done. Metadata is %q\n", resp)
	}
}

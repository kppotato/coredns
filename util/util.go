package util

import (
	"encoding/json"
	"github.com/kppotato/coredns/g"
	"github.com/kppotato/coredns/model"
	"regexp"
	"strings"
)

var (
	reg, _ = regexp.Compile(`/x\d{1,}$`)
)

func Reverse(l []string) {
	for i := 0; i < int(len(l)/2); i++ {
		li := len(l) - i - 1
		l[i], l[li] = l[li], l[i]
	}
}

func Etcdkey2Host(key, value string) *model.Dns {
	temp := reg.ReplaceAllString(key, "")
	temp = strings.Replace(temp, g.Etcd_path, "", 1)
	list := strings.Split(temp, "/")
	Reverse(list)
	aaa := model.A{}
	json.Unmarshal([]byte(value), &aaa)
	return &model.Dns{Origin: strings.Join(list, "."), NameServer: aaa.Host, TTL: aaa.TTL, Key: key, Value: value}
}

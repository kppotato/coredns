package g

import (
	"time"
	"github.com/patrickmn/go-cache"
)

var(
	Etcd_url = make([]string,0)
	Etcd_path ="/skydns"
	Exit = make(chan struct{},1)
	Mycache = cache.New(60*time.Minute,60*time.Minute)
	Cache_Name ="dns"

)

package etcd

import (
	"github.com/coreos/etcd/client"
	"mtime.com/corednsUI/g"
	"mtime.com/corednsUI/util"
	"mtime.com/corednsUI/model"
	//"mtime.com/corednsUI/controller"
	"fmt"
	"os"
	"context"
	"strings"
	"github.com/patrickmn/go-cache"
)

var(
	kapi client.KeysAPI
	EtcdDao = &etcddao{}
	NewMessage = make(chan []*model.Dns,1)
)

func OninitCheck()  {
	c, err := client.New(client.Config{
		Endpoints:               g.Etcd_url,
		Transport:               client.DefaultTransport,
	})
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	kapi = client.NewKeysAPI(c)
	_,err = kapi.Get(context.Background(),g.Etcd_path,&client.GetOptions{})
	//fmt.Println(rep.Node.Value)
	//_, err = kapi.Set(context.Background(),ippath, "0",&client.SetOptions{PrevExist:client.PrevExist,PrevValue:"0",Dir:false})
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
type etcddao struct {

}

func (this *etcddao) DnsList() []*model.Dns{
	result,found :=g.Mycache.Get(g.Cache_Name)
	if found{
		return result.([]*model.Dns)
	}else{
		fmt.Println(33300000)
		mymap ,err :=etcdList()
		if err !=nil{
			return nil
		}
		result :=make([]*model.Dns,0,len(mymap))
		for k,v :=range mymap{
			result=append(result, util.Etcdkey2Host(k,v))
		}
		return result
	}
}
func etcdALL()[]*model.Dns{
	mymap ,err :=etcdList()
	if err !=nil{
		return nil
	}
	result :=make([]*model.Dns,0,len(mymap))
	for k,v :=range mymap{
		result=append(result, util.Etcdkey2Host(k,v))
	}
	g.Mycache.Set(g.Cache_Name,result,cache.DefaultExpiration)
	return result
}
func (this *etcddao) DnsAdd(key,value string) (bool,error){
	return etcdAdd(key,value)
}
func (this *etcddao) DnsDel(key string) error{
	return etcdDel(key)
}
func (this *etcddao) DnsEdit(key,value string) error{
	return etcdEdit(key,value)
}
func (this *etcddao) DnsGet(key string) (*model.Dns,error) {
	node,err := etcdGet(key)
	if err != nil{
		return nil,err
	}
	return util.Etcdkey2Host(node.Key,node.Value),nil
}

func etcdGetmap(node *client.Node,mymap map[string]string){
	if node.Dir{
		for _,x :=range node.Nodes{
			if x.Dir {
				resp,err := kapi.Get(context.Background(),x.Key,&client.GetOptions{})
				if err != nil {
					continue
				}
				etcdGetmap(resp.Node,mymap)
			}else{
				mymap[x.Key]=x.Value
			}
		}
	} else{
		mymap[node.Key]=node.Value
	}
}
func etcdList()(map[string]string,error){
	node,err := etcdGet(g.Etcd_path)
	if err != nil{
		return nil,err
	}
	result :=make(map[string]string,10000)
	etcdGetmap(node,result)
	return result,nil
}
func etcdGet(key string) (*client.Node,error){
	resp,err := kapi.Get(context.Background(),key,&client.GetOptions{})
	if err != nil {
		return nil,err
	}
	return resp.Node,nil
}
func etcdAdd(key,value string) (bool,error){
	keylist := strings.Split(key,".")
	util.Reverse(keylist)
	prekey :=strings.Join(keylist,"/")
	if !strings.HasPrefix(prekey,"/"){
		prekey="/"+prekey
	}
	key=g.Etcd_path+prekey
	_, err := kapi.Set(context.Background(),key,value,&client.SetOptions{PrevExist:client.PrevNoExist,Dir:false})
	if err != nil {
		if e,ok := err.(client.Error);ok{
			if e.Code == client.ErrorCodeNodeExist{
				return false,fmt.Errorf("数据已经存在！")
			}
		}
		return false,err
	}
	return true,nil
}
func etcdDel(key string) error{
	fmt.Println(key)
	_,err := kapi.Delete(context.Background(),key,&client.DeleteOptions{Recursive:true})
	if err != nil{
		fmt.Println(key,err)
		return err
	}
	return nil
}
func etcdEdit(key,value string) error{
	_, err := kapi.Set(context.Background(),key, value,&client.SetOptions{PrevExist:client.PrevExist,Dir:false})
	if err != nil{
		return err
	}
	return nil
}

func WatchEtcd(){
	watcher := kapi.Watcher(g.Etcd_path,&client.WatcherOptions{Recursive: true})
	fmt.Println(122222)
	for{
		select {
		case <- g.Exit:
			break
		default:

		}
		res,err := watcher.Next(context.Background())
		if err != nil{
			continue
		}
		if res.Action == "expire"{
			continue
		}else if res.Action =="set" || res.Action == "update" || res.Action == "create" || res.Action == "delete"{
			fmt.Println(res.Action)
			result:=etcdALL()
			if result !=nil{
				NewMessage <- result
			}
		}

	}
}









# coredns
coredns web管理界面。

* 页面使用adminlite
* 前后端通讯使用了 websocket，无需刷新页面。
* watch后端etcd的变化，通过websocket通知页面变更。


# 列表页
![展示](/12.PNG)
## 编辑页
![展示](/13.PNG)

- 运行golang主程序 主程序需要的参数，
 <table>
 <thead>
 <tr>
  <td>参数名</td>
  <td>描述</td>
  <td>例子</td>
  <td>必须字段</td>
 </tr>
  </thead>
 <tbody>
  <tr>
  <td>etcdurl</td>
  <td>etcd集群地址</td>
  <td>./corednsui --etcdurl http://xxxxxx:port,http://xxxxxx:port</td>
   <td>必须字段</td>
 </tr>
    <tr>
  <td>etcdpath</td>
  <td>etcd数据存储地址</td>
  <td>./corednsui --etcdpath skydns</td>
     <td>必须字段</td>
 </tr>
   </tbody>
 </table>

/***************************
@File        : distribute.go
@Time        : 2021/04/15 17:13:36
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : go简易分布式
****************************/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

//用于json和结构体对象的互转
type NodeInfo struct {
	NodeId     int    `json:"nodeId"`     //节点ID，通过随机数生成
	NodeIpAddr string `json:"nodeIpAddr"` //节点ip地址
	Port       string `json:"port"`       //节点端口号
}

//添加一个节点到集群的一个请求或者响应的标准格式
type AddToClusterMessage struct {
	Source  NodeInfo `json:"source"`
	Dest    NodeInfo `json:"dest"`
	Message string   `json:"message"`
}

//将节点信息格式化输出
func (node *NodeInfo) String() string {
	return "NodeInfo {nodeId:" + strconv.Itoa(node.NodeId) + ", nodeIpAddr:" + node.NodeIpAddr + ", port:" + node.Port + "}"
}

//将添加节点信息格式化
func (req AddToClusterMessage) String() string {
	return "AddToClusterMessage:{\n  source:" + req.Source.String() + ",\n  dest: " + req.Dest.String() + ",\n  message:" + req.Message + " }"
}

var dest NodeInfo

func main() {
	makeMasterOnError := flag.Bool("makeMasterOnError", false, "make this node master if unable to connect to the cluster ip provided.")
	clusterip := flag.String("clusterip", "192.168.0.15:8001", "ip address of any node to connnect")
	myport := flag.String("myport", "8001", "ip address to run this node on. default is 8001.")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano()) //种子
	myid := rand.Intn(9999999)

	//获取ip地址
	myIp := getLocalIp()
	//创建nodeInfo结构体
	me := NodeInfo{NodeId: myid, NodeIpAddr: myIp, Port: *myport}
	dest.NodeIpAddr = strings.Split(*clusterip, ":")[0]
	dest.Port = strings.Split(*clusterip, ":")[1]
	fmt.Println("我的节点信息：", me.String(), dest.String())
	//尝试连接到集群，在已连接的情况下向集群发送请求
	ableToConnect := connectToCluster(me, dest)
	fmt.Println(ableToConnect)
	//如果dest节点不存在，则me节点为主节点启动，否则直接退出系统
	if ableToConnect || (!ableToConnect && *makeMasterOnError) {
		if *makeMasterOnError {
			fmt.Println("将启动me节点为主节点")
		}
		dest.NodeId = me.NodeId
		log.Println(dest.NodeId, "输出")
		listenOnPort(me)
	} else {
		fmt.Println("正在退出系统，请设置me节点为主节点")
	}
}

//发送请求时格式化json包有用的工具
func getAddToClusterMessage(source NodeInfo, dest NodeInfo, message string) AddToClusterMessage {
	return AddToClusterMessage{
		Source: NodeInfo{
			NodeId:     source.NodeId,
			NodeIpAddr: source.NodeIpAddr,
			Port:       source.Port},
		Dest: NodeInfo{
			NodeId:     dest.NodeId,
			NodeIpAddr: dest.NodeIpAddr,
			Port:       dest.Port},
		Message: message,
	}
}
func connectToCluster(me NodeInfo, dest NodeInfo) bool {
	//连接到socket的相关细节信息
	log.Println(dest.String())
	if dest.NodeId == -1 {
		log.Println("hhhh")
		return false
	}
	connOut, err := net.DialTimeout("tcp", dest.NodeIpAddr+":"+dest.Port, time.Duration(1)*time.Second)
	log.Println(connOut, err)
	if err != nil {
		if _, ok := err.(net.Error); ok {
			fmt.Println("不能连接到集群", me.NodeId)
			return false
		}
	} else {
		fmt.Println("连接到集群")
		text := "Hi nody.. 请添加我到集群"
		requestMessage := getAddToClusterMessage(me, dest, text)
		json.NewEncoder(connOut).Encode(&requestMessage)
		decoder := json.NewDecoder(connOut)
		var responseMessage AddToClusterMessage
		decoder.Decode(&responseMessage)
		fmt.Println("得到数据响应:\n" + responseMessage.String())
		if responseMessage.Source.NodeId == 0 {
			return false
		}
		return true
	}
	return false
}

//me节点连接其它节点成功或者自身成为主节点之后开始监听别的节点在未来可能对它自身的连接
func listenOnPort(me NodeInfo) {
	fmt.Println("启动主节点监听")
	//监听即将到来的信息
	ln, _ := net.Listen("tcp", fmt.Sprint(me.NodeIpAddr+":"+me.Port))
	//接受连接
	for {
		fmt.Print("listen")
		connIn, err := ln.Accept()
		if err != nil {
			fmt.Println("连接失败")
			if _, ok := err.(net.Error); ok {
				fmt.Println("Error received while listening.", me.NodeId)
			}
		} else {
			fmt.Println("等待接收信息")
			var requestMessage AddToClusterMessage
			json.NewDecoder(connIn).Decode(&requestMessage)
			fmt.Println("Got request:\n" + requestMessage.String())

			text := "已添加你到集群"
			responseMessage := getAddToClusterMessage(me, requestMessage.Source, text)
			json.NewEncoder(connIn).Encode(&responseMessage)
			connIn.Close()
			log.Println("hhhh")
		}
	}
}

// 获取本地ip地址
func getLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "null"
}

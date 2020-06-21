package zookeeper

import (
	"encoding/json"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// ServiceNode : 存储在节点的data中，表示服务发现的地址信息
type ServiceNode struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

// SdClient : 服务发现的客户端结构体SdClient
type SdClient struct {
	zkServers []string // 多个zookeeper主机地址
	zkRoot    string   // 服务根结点
	conn      *zk.Conn // zk的客户端连接
}

// NewClient : 编写构造器，创建根结点
func NewClient(zkServers []string, zkRoot string,
	timeout int) (*SdClient, error) {
	client := new(SdClient)
	client.zkServers = zkServers
	client.zkRoot = zkRoot
	// 连接服务器
	conn, _, err := zk.Connect(zkServers, time.Duration(timeout)*time.Second)
	if err != nil {
		return nil, err
	}
	client.conn = conn
	// 创建服务根结点
	if err := client.ensureRoot(); err != nil {
		client.Close()
		return nil, err
	}
	return client, nil
}

// Close : 关闭连接，释放临时节点
func (s *SdClient) Close() {
	s.conn.Close()
}

// ensureRoot : 创建根结点
func (s *SdClient) ensureRoot() error {
	exists, _, err := s.conn.Exists(s.zkRoot)
	if err != nil {
		return err
	}
	if !exists {
		if _, err := s.conn.Create(s.zkRoot, []byte(""), 0,
			zk.WorldACL(zk.PermAll)); err != nil && err != zk.ErrNodeExists {
			return err
		}
	}
	return nil
}

// Register : 服务注册
// 先要创建/api/user节点作为服务列表的父结点。然后创建一个保护顺序临时(
// ProtectedEphemeralSequential)子节点，同时将地址信息存储在节点中。
// 什么叫保护顺序临时节点，首先它是一个临时节点，会话关闭后自动消失。其次它是个顺序
// 节点，zookeeper自动在名称后面增加自增后缀，确保节点名称的唯一性。同时还是个保护性节点，
// 节点前缀新增加了GUID字段，确保断开重连后临时可以和客户端状态对接上。
func (s *SdClient) Register(node *ServiceNode) error {
	if err := s.ensureName(node.Name); err != nil {
		return err
	}
	path := s.zkRoot + "/" + node.Name + "/n"
	data, err := json.Marshal(node)
	if err != nil {
		return err
	}
	if _, err := s.conn.CreateProtectedEphemeralSequential(
		path, data, zk.WorldACL(zk.PermAll)); err != nil {
		return err
	}
	return nil
}

func (s *SdClient) ensureName(name string) error {
	path := s.zkRoot + "/" + name
	exists, _, err := s.conn.Exists(path)
	if err != nil {
		return err
	}
	if !exists {
		if _, err := s.conn.Create(path, []byte(""), 0,
			zk.WorldACL(zk.PermAll)); err != nil && err != zk.ErrNodeExists {
			return err
		}
	}
	return nil
}

// GetNodes : 消费者获取服务列表
func (s *SdClient) GetNodes(name string) ([]*ServiceNode, error) {
	path := s.zkRoot + "/" + name
	// 获取子节点名称
	childs, _, err := s.conn.Children(path)
	if err != nil {
		if err == zk.ErrNoNode {
			return []*ServiceNode{}, nil
		}
		return nil, err
	}
	nodes := []*ServiceNode{}
	for _, child := range childs {
		fullPath := path + "/" + child
		data, _, err := s.conn.Get(fullPath)
		if err != nil {
			if err == zk.ErrNoNode {
				continue
			}
			return nil, err
		}
		node := new(ServiceNode)
		if err = json.Unmarshal(data, node); err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

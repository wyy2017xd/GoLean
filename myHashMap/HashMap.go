package myHashMap

import "fmt"

//存储map的kv链表的node
type KV struct {
	 key string
	 value string
}

//存储键值对的链表类

type ListNode struct {
	//节点数据
	data KV
	//写一个数据的节点
	nextLink *ListNode
}

//创建链表

func createLink() *ListNode {
	var node = &ListNode{KV{"",""}, nil}
	return node
}

//哈希碰撞时，键值对会存储在新的链表节点上，添加节点

func (link *ListNode) addNode(data KV) int {
	var count = 0;
	tail := link
	//循环找到节点的尾结点
	for {
		count ++
		if tail.nextLink == nil {
			break
	} else {
		tail = tail.nextLink
		}

	}
	var newNode = &ListNode{data, nil}
	tail.nextLink = newNode
	return count +1
}

//创建hash木桶（数组）的个数，初始化每个数组对应的链表
const BucketCount  = 16
type HashMap struct {
	Buckets [BucketCount]*ListNode//数组
}

//创建hashMap
func CreateHashMap() *HashMap {
	myMap := &HashMap{}
	for i := 0; i < BucketCount; i++ {
		myMap.Buckets[i] = createLink()
	}
	return myMap
}

//自定义hash算法
func HashCode(key string) int {
	var sum = 0
	for i := 0; i < len(key); i++ {
		sum += int(key[i])
	}
	return sum%BucketCount
}

//添加KV

func (myMap *HashMap) AddKeyValue(key, value string) {
	//求hashcode
	var index = HashCode(key)
	//获取头节点
	var link = myMap.Buckets[index]
	//插入节点
	if link.nextLink == nil {
		link.data.key = key
		link.data.value = value
	} else {
		index := link.addNode(KV{key, value})
		fmt.Println("node count is %d", index)

	}
}

//从map中取得对应keyd的value

func (myMap *HashMap) GetValueByKey(key string) string {
	//求hashcode
	index := HashCode(key)
	//获取头结点
	link := myMap.Buckets[index]
	head := link
	value := ""
	for {
		if head.data.key == key {
			value = head.data.value
		} else {
			head = head.nextLink
		}
	}
	return value
}
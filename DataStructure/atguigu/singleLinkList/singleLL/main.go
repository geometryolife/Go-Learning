package main

import "fmt"

// 定义 HeroNode 结构
type HeroNode struct {
	no       int
	name     string
	nickName string
	next     *HeroNode // 表示指向下一个节点
}

// 给链表插入一个节点
// 方法1: 在单链表的最后插入(后插法)
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	// 思路: 先找到该链表的最后一个节点
	// 1. 创建一个辅助节点
	temp := head
	for {
		// 当辅助节点找到了链表的最后一个节点，指针域为空
		if temp.next == nil {
			break
		}
		// 让temp 不断指向下一个节点
		temp = temp.next
	}

	// 2. 将 newnewHeroNode 加入到链表的最后
	temp.next = newHeroNode
}

// 给链表插入一个节点
// 方法2: 插入时，根据 no 的编号从小到大插入
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	// 思路: 找到适当的节点，进行插入
	// 1. 创建一个辅助节点
	temp := head
	flag := true
	// 让插入的节点的 no 与 temp 的下一个节点的 no 进行比较
	for {
		if temp.next == nil {
			// 当辅助节点找到了链表的最后一个节点，指针域为空
			break
		} else if temp.next.no > newHeroNode.no { // 改变此处可以实现次序排序或相同编号插入
			// 此时符合条件，newHeroNode 插入到 temp 后面
			break
		} else if temp.next.no == newHeroNode.no {
			// 此时链表已经存在 no 相同的节点，不允许插入，保证唯一性
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("无法插入，已经存在 no =", newHeroNode.no)
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}

	// 2. 将 newnewHeroNode 加入到链表的最后
}

// 删除一个节点
func DelHeroNode(head *HeroNode, id int) {

	// 1. 创建一个辅助节点
	temp := head
	flag := false
	// 找到要删除的节点的 no 与 temp 的下一个节点的 no 进行比较
	for {
		if temp.next == nil {
			// 当辅助节点找到了链表的最后一个节点，指针域为空
			break
		} else if temp.next.no == id {
			// 此时找到了符合的元素
			flag = true
			break
		}
		temp = temp.next
	}
	// 删除找到的元素
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("Sorry，要删除的 id 不存在!")
	}
}

// 显示链表的所有节点信息
func ListHeroNode(head *HeroNode) {
	// 1. 创建一个辅助节点
	temp := head

	// 2. 判断链表是否为空链表
	if temp.next == nil {
		fmt.Println("空空入也!")
		return
	}

	// 2. 遍历链表
	for {
		fmt.Printf("[%d, %s, %s]==>", temp.next.no, temp.next.name,
			temp.next.nickName)
		// 判断辅助节点是否到了链表的末尾
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func main() {
	// 1. 创建头节点
	head := &HeroNode{}

	// 2. 创建一个新的 HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickName: "及时雨",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickName: "玉麒麟",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickName: "豹子头",
	}
	hero4 := &HeroNode{
		no:       3,
		name:     "吴用",
		nickName: "智多星",
	}

	// 3. 加入
	// Test1
	// InsertHeroNode(head, hero1)
	// InsertHeroNode(head, hero2)
	// InsertHeroNode(head, hero3)
	// InsertHeroNode(head, hero4)

	// Test2
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero2)
	InsertHeroNode2(head, hero4)

	// 4. 显示
	ListHeroNode(head)
	fmt.Println()

	// 5. 删除
	DelHeroNode(head, 1)
	DelHeroNode(head, 3)
	ListHeroNode(head)
	fmt.Println()
}

/*
>>> Execution Result:
Test1:
[1, 宋江, 及时雨]==>[2, 卢俊义, 玉麒麟]==>[3, 林冲, 豹子头]==>[3, 吴用, 智多星]==>

Test2:
顺序插入:
无法插入，已经存在 no = 3
[1, 宋江, 及时雨]==>[2, 卢俊义, 玉麒麟]==>[3, 林冲, 豹子头]==>

倒序插入:
无法插入，已经存在 no = 3
[3, 林冲, 豹子头]==>[2, 卢俊义, 玉麒麟]==>[1, 宋江, 及时雨]==>

允许插入 no 相同的人物:
[1, 宋江, 及时雨]==>[2, 卢俊义, 玉麒麟]==>[3, 吴用, 智多星]==>[3, 林冲, 豹子头]==>
[3, 吴用, 智多星]==>[3, 林冲, 豹子头]==>[2, 卢俊义, 玉麒麟]==>[1, 宋江, 及时雨]==>

删除节点:
无法插入，已经存在 no = 3
[1, 宋江, 及时雨]==>[2, 卢俊义, 玉麒麟]==>[3, 林冲, 豹子头]==>
[2, 卢俊义, 玉麒麟]==>
*/

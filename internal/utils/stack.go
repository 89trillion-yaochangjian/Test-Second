package utils

import (
	"sync"
)

type Item interface{}

type ItemStack struct {
	items []string
	lock  sync.RWMutex
}

// 初始化栈
func NewStack() *ItemStack {
	s := &ItemStack{}
	s.items = []string{}
	return s
}

// 添加元素
func (s *ItemStack) Push(str string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, str)
}

// 获取栈顶元素
func (s *ItemStack) Top() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.items) == 0 {
		return ""
	}
	//获取栈顶元素
	str := s.items[len(s.items)-1]
	return str
}

// 弹出元素
func (s *ItemStack) Pop() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.items) == 0 {
		return ""
	}
	//获取栈顶元素
	str := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return str
}

// IsEmpty 判断栈是否为空
func (s *ItemStack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

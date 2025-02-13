package main

import "fmt"

type DLink struct {
	m_Data int
	m_Next *DLink
	m_Prev *DLink
}

type Queue struct {
	m_Capacity int
	m_Size     int
	m_Queue    *DLink
}

func NewQueue(capacity int) *Queue {
	q := &Queue{capacity, 0, nil}
	q.m_Queue = &DLink{-1, nil, nil}
	q.m_Queue.m_Next = q.m_Queue
	q.m_Queue.m_Prev = q.m_Queue

	return q
}

func (q *Queue) IsEmpty() bool {
	return q.m_Size == 0
}

func (q *Queue) Size() int {
	return q.m_Size
}

func (q *Queue) Front() int {
	return q.m_Queue.m_Prev.m_Data
}

func (q *Queue) Back() int {
	return q.m_Queue.m_Next.m_Data
}

func (q *Queue) Push(element int) {
	if q.m_Size == q.m_Capacity {
		return
	}

	pE := &DLink{element, nil, nil}
	pE.m_Next = q.m_Queue.m_Next
	pE.m_Prev = q.m_Queue
	q.m_Queue.m_Next.m_Prev = pE
	q.m_Queue.m_Next = pE
	q.m_Size++
}

func (q *Queue) Pop() {
	if q.IsEmpty() {
		return
	}

	q.m_Queue.m_Prev.m_Prev.m_Next = q.m_Queue
	q.m_Queue.m_Prev = q.m_Queue.m_Prev.m_Prev
	q.m_Size--
}

type Stack struct {
	m_Queue *Queue
}

func NewStack(capacity int) *Stack {
	return &Stack{NewQueue(capacity)}
}

func (s *Stack) IsEmpty() bool {
	return s.m_Queue.IsEmpty()
}

func (s *Stack) Size() int {
	return s.m_Queue.Size()
}

func (s *Stack) Top() int {
	return s.m_Queue.Back()
}

func (s *Stack) Push(element int) {
	s.m_Queue.Push(element)
}

func (s *Stack) Pop() {
	tempQueue := NewQueue(s.m_Queue.Size())
	for s.m_Queue.Size() > 1 {
		tempQueue.Push(s.m_Queue.Front())
		s.m_Queue.Pop()
	}

	s.m_Queue.Pop()

	for !tempQueue.IsEmpty() {
		s.m_Queue.Push(tempQueue.Front())
		tempQueue.Pop()
	}
}

func PrintTest(s *Stack) {
	// Commented is the normal implementation. The current code is as per the questions' requirement
	n := s.Top()
	fmt.Print(n, " ")
	fmt.Println(s.IsEmpty(), s.Size())
}

func main() {
	s := NewStack(10)
	PrintTest(s)

	s.Push(5)
	PrintTest(s)

	s.Push(6)
	PrintTest(s)

	s.Push(7)
	PrintTest(s)

	s.Pop()
	PrintTest(s)

	s.Pop()
	PrintTest(s)

	s.Pop()
	PrintTest(s)
}

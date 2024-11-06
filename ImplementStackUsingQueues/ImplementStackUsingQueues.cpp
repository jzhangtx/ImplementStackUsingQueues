// ImplementStackUsingQueues.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>
#include <queue>

class Queue
{
public:
	Queue(int capacity)
		: m_Capacity(capacity)
	{
	}

	int size()
	{
		return static_cast<int>(m_Queue.size());
	}

	bool isEmpty()
	{
		return m_Queue.empty();
	}

	int front()
	{
		if (m_Queue.size() == 0)
			return -1;

		return m_Queue.front();
	}

	int back()
	{
		if (m_Queue.size() == 0)
			return -1;

		return m_Queue.back();
	}

	void push(int element)
	{
		if (m_Queue.size()== static_cast<size_t>(m_Capacity))
			return;

		m_Queue.push(element);
	}

	void pop()
	{
		//
		// commented because the question does not require. it will
		//  throw an exception if the queue is empty
		// 
		//if (m_Queue.empty())
		//	return;

		m_Queue.pop();
	}

private:
	std::queue<int> m_Queue;
	int m_Capacity;
};


// Implement the Stack class
class Stack {
public:

	Stack(int capacity)
		:m_Queue(capacity)
	{}

	bool isEmpty()
	{
		return m_Queue.isEmpty();
	}

	int size()
	{
		return m_Queue.size();
	}

	int top()
	{
		return m_Queue.back();
	}

	void push(int element)
	{
		m_Queue.push(element);
	}

	void pop()
	{
		Queue tempQueue(m_Queue.size());
		while (m_Queue.size() != 1)
		{
			tempQueue.push(m_Queue.front());
			m_Queue.pop();
		}
		m_Queue.pop();
		while (!tempQueue.isEmpty())
		{
			m_Queue.push(tempQueue.front());
			tempQueue.pop();
		}
	}

private:
	Queue m_Queue;
};

int main()
{
	int i = 0;
	bool finished = false;
	Stack stack(10);

	while (!finished)
	{
		std::cout << stack.top() << " ";;
		std::cout << (stack.isEmpty() ? "true" : "false") << " " << stack.size() << std::endl;
		switch (i++)
		{
		case 0:
			stack.push(5);
			break;

		case 1:
			stack.push(6);
			break;

		case 2:
			stack.push(7);
			break;

		case 3:
		case 4:
		case 5:
			stack.pop();
			break;

		default:
			finished = true;
		}
	}	
}

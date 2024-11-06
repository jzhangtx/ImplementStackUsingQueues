Implement Stack using Queues


Implement a stack using one or more queues.

The Stack class should support the following methods:

int size()
boolean isEmpty()
int top()
void push(int element)
void pop()
You can assume that you've access to a Queue class which provides the following methods:

int size()
boolean isEmpty()
int front()
int back()
void push(int element)
void pop()
Main method:
Stack created of size 10
Printed: stack.top() stack.isEmpty() stack.size()
Pushed: 5
Printed: stack.top() stack.isEmpty() stack.size()
Pushed: 6
Printed: stack.top() stack.isEmpty() stack.size()
Pushed: 7
Printed: stack.top() stack.isEmpty() stack.size()
Popped
Printed: stack.top() stack.isEmpty() stack.size()
Popped
Printed: stack.top() stack.isEmpty() stack.size()
Popped
Printed: stack.top() stack.isEmpty() stack.size()
Expected Output
-1 true 0
5 false 1
6 false 2
7 false 3
6 false 2
5 false 1
-1 true 0
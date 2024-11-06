ImplementStackUsingQueues: ImplementStackUsingQueues.o
	g++ -g -o ImplementStackUsingQueues.exe ImplementStackUsingQueues.o -pthread    

ImplementStackUsingQueues.o: ImplementStackUsingQueues/ImplementStackUsingQueues.cpp
	g++ -g  -c -pthread ImplementStackUsingQueues/ImplementStackUsingQueues.cpp

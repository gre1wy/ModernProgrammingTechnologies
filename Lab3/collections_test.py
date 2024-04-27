from collections import deque

stack = deque(maxlen=3)

stack.append(1)
stack.append(2)
stack.append(3)
stack.append(4)

print(stack.pop()) 
print(stack.pop())  
print(stack.pop())  

queue = deque()

queue.append(1)
queue.append(2)
queue.append(3)
queue.insert(2, 0)
print(queue.popleft())  
print(queue.popleft())  
print(queue.popleft())  
print(queue.popleft())  
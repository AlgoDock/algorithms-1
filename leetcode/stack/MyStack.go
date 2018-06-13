/* https://leetcode.com/problems/implement-stack-using-queues/#/description
Implement the following operations of a stack using queues.

    push(x) -- Push element x onto stack.
    pop() -- Removes the element on top of the stack.
    top() -- Get the top element.
    empty() -- Return whether the stack is empty.

Notes:

You must use only standard operations of a queue -- which means only push to back, peek/pop from front, size, and is empty operations are valid.
Depending on your language, queue may not be supported natively. You may simulate a queue by using a list or deque (double-ended queue),
as long as you use only standard operations of a queue.

You may assume that all operations are valid (for example, no pop or top operations will be called on an empty stack).

*/

package lstack

type MyStack struct {
	q []int
}

/** Initialize your data structure here. */
func MyStackConstructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.q = append(this.q, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	r := this.q[len(this.q)-1]
	this.q = this.q[:len(this.q)-1]
	return r
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.q[len(this.q)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.q) == 0
}

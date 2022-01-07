package challenge

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

/* Problem:
Note: Try to solve this task in O(n) time using O(1) additional space, where n is the number of elements in the list, since this is what you'll be asked to do during an interview.

Given a singly linked list of integers l and an integer k, remove all elements from list l that have a value equal to k.

*/
func RemoveKFromList(l *ListNode, k int) *ListNode {
	c := l
	if c.Value == k && c.Next != nil {
		c = c.Next
	} else {
		return nil
	}
	for c.Next != nil {
		if c.Value == k {
			c.Next = c.Next.Next
		} else {
			c = c.Next
		}
	}
	return c
}

/* Problem:
Note: Try to solve this task in O(n) time using O(1) additional space, where n is the number of elements in l, since this is what you'll be asked to do during an interview.

Given a singly linked list of integers, determine whether or not it's a palindrome.

Note: in examples below and tests preview linked lists are presented as arrays just for simplicity of visualization: in real data you will be given a head node l of the linked list
*/
func IsListPalindrome(l *ListNode) bool {
	var list []interface{}
	c := l

	for c != nil {
		list = append(list, c.Value)
		c = c.Next
	}
	for i := 0; i < (len(list) / 2); i++ {

		if list[i] != list[(len(list)-1)-i] {
			return false
		}
	}
	return true
}

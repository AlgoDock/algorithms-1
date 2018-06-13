/* https://leetcode.com/problems/partition-list/description/
Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

For example,
Given 1->4->3->2->5->2 and x = 3,
return 1->2->2->4->3->5.
*/

package lll

func partition(head *ListNode, x int) *ListNode {
	ltHead, gteHead := &ListNode{Val: 0}, &ListNode{Val: 0}
	ltCur, gteCur := ltHead, gteHead

	for cur := head; cur != nil; cur = cur.Next {
		tmp := &ListNode{Val: cur.Val}
		if cur.Val < x {
			ltCur.Next, ltCur = tmp, tmp
		} else {
			gteCur.Next, gteCur = tmp, tmp
		}
	}
	ltCur.Next = gteHead.Next
	return ltHead.Next
}

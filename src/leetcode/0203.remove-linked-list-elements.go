/**
 * Remove Linked List Elements - LeetCode
 * https://leetcode.com/problems/remove-linked-list-elements/
 */

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
    headPre := ListNode{Next:head}
    tmp := &headPre

    for tmp.Next != nil {
        if tmp.Next.Val == val {
            tmp.Next = tmp.Next.Next
        } else {
            tmp = tmp.Next
        }
    }
    
    return headPre.Next
}


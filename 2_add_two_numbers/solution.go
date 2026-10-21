type ListNode struct {
    Val  int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := &ListNode{Val: 0}
    current := dummyHead
    carry := 0
    
    for l1 != nil || l2 != nil || carry != 0 {
        val1, val2 := 0, 0
        if l1 != nil {
            val1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            val2 = l2.Val
            l2 = l2.Next
        }
        
        totalSum := val1 + val2 + carry
        carry = totalSum / 10
        newDigit := totalSum % 10
        
        current.Next = &ListNode{Val: newDigit}
        current = current.Next
    }
    
    return dummyHead.Next
}
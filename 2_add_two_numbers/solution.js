class ListNode {
    constructor(val = 0, next = null) {
        this.val = val;
        this.next = next;
    }
}

var addTwoNumbers = function(l1, l2) {
    let dummyHead = new ListNode(0);
    let current = dummyHead;
    let carry = 0;
    
    while (l1 !== null || l2 !== null || carry !== 0) {
        let val1 = l1 !== null ? l1.val : 0;
        let val2 = l2 !== null ? l2.val : 0;
        
        let totalSum = val1 + val2 + carry;
        carry = Math.floor(totalSum / 10);
        let newDigit = totalSum % 10;
        
        current.next = new ListNode(newDigit);
        current = current.next;
        
        if (l1 !== null) l1 = l1.next;
        if (l2 !== null) l2 = l2.next;
    }
    
    return dummyHead.next;
};
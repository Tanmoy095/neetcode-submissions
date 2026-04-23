/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {

if head == nil {
	return false
}
exist := make(map[*ListNode]bool)
curr:=head 
for curr!= nil{
	if exist[curr]{
		return true
	}
	exist[curr]=true
	curr=curr.Next

}
return false
    
}
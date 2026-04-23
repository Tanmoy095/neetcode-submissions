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

sPtr:= head //slow pointer 
fPtr:= head  //fast pointer 
for fPtr!= nil && fPtr.Next!= nil{
	sPtr = sPtr.Next
	fPtr = fPtr.Next.Next
if sPtr == fPtr{
	return true
}
}
return false 
}
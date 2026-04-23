/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */




func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
dummyNode:= &ListNode{
}
current:=  dummyNode
ptr1 := list1
ptr2 := list2 

for ptr1!=nil && ptr2 !=nil {
		if ptr1.Val <= ptr2.Val {
			current.Next = ptr1
			ptr1 = ptr1.Next

		}else if ptr2.Val <= ptr1.Val {
			
			current.Next = ptr2
			ptr2 = ptr2.Next
		}
		current = current.Next 
	}
	if ptr1!=nil{
			current.Next = ptr1

	} 
	if ptr2!=nil{
			current.Next = ptr2

	} 
		return dummyNode.Next

	
}

    


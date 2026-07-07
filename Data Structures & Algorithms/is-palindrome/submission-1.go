func isPalindrome(s string) bool {

//  Filter out non-alphanumeric characters and normalize to lowercase
    var clean []rune
    for _, r := range s {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            clean = append(clean, unicode.ToLower(r))
        }
    }

    // Check if the slice is a palindrome by comparing from both ends
    n := len(clean)
    for i := 0; i < n/2; i++ {
        // Compare elements from the start and end moving inward
        if clean[i] != clean[n-1-i] {
            return false
        }
    }

    return true
}

/*
To validate the palindrome, I use a two-step process that prioritizes order preservation. First, 
I filter out all non-alphanumeric characters and normalize the remaining runes to lowercase, storing them sequentially in a slice.
Next, I use a symmetric mirror check by iterating up to n/2 of the slice's length. 
Using the formula n - 1 - i, I pair the element at the index i from the front with its exact structural counterpart from the back.
If any paired elements mismatch, I can immediately return false. By stopping exactly at the halfway mark,
I avoid redundant checks on the second half of the string, resulting in an optimal O(n) time complexity solution
*/


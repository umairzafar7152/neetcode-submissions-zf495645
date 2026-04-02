func isPalindrome(s string) bool {
    l, r := 0, len(s)-1

    for l < r {
        for l < r && !isAlphanumeric(s[l]) {
            l++
        }
        for l < r && !isAlphanumeric(s[r]) {
            r--
        }
        if !equalFold(s[l], s[r]) {
            return false
        }
        l++
        r--
    }
    return true
}

func isAlphanumeric(b byte) bool {
    return (b >= 'a' && b <= 'z') ||
           (b >= 'A' && b <= 'Z') ||
           (b >= '0' && b <= '9')
}

func equalFold(b1, b2 byte) bool {
    if b1 >= 'A' && b1 <= 'Z' {
        b1 += 32
    }
    if b2 >= 'A' && b2 <= 'Z' {
        b2 += 32
    }
    return b1 == b2
}
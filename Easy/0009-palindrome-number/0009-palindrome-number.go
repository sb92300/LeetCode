func isPalindrome(x int) bool {
    if (x < 0) {
        return false
    }
    y := x;
    z := 0;

    // var result bool;
    for 0 < x {
        if y > 0 {
            z = z * 10 + y % 10;
            y = y / 10;
        } else {
            break;
        }
    }

return z == x;
}
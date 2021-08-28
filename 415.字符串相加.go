/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
    rune_num1 := []rune(num1)
    rune_num2 := []rune(num2)
    i,j := len(rune_num1)-1,len(rune_num2)-1
    carry := 0
    res := make([]rune,0)
    for i>=0 || j>=0 {
        var (
            n1 int
            n2 int
        )
        if i<0 {
            n1 = 0
        }else {
            n1 = int(rune_num1[i])-int('0')
        }
        if j<0 {
            n2 = 0
        }else {
            n2 = int(rune_num2[j])-int('0')
        }

        sum := n1+n2+carry
        carry = sum/10
        sum = sum%10
        
        res = append([]rune{rune(sum)+'0'},res...)
        i--
        j--
    }
    if carry != 0 {
        res = append([]rune{'1'},res...)
    }
    return string(res)
}
// @lc code=end


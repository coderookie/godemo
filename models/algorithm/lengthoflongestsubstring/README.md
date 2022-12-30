# 无重复字符的最长子串
leetcode: https://leetcode.cn/problems/longest-substring-without-repeating-characters/?favorite=2cktkvj <br />
思路
    1、声明一个结果字符串 <br />
    2、第一次循环第i个字符的时候, 设置第一次循环结果字符串 <br />
    3、第二个嵌套循环从i + 1个字符 <br />
    4、如果第二个嵌套循环中, 有字符存在在第一次循环结果字符串中, 则看第一次循环结果字符串的长度是否大于结果字符串 <br />
    5、如果大于结果字符串, 将结果字符串设置成第一次循环结果字符串 <br />
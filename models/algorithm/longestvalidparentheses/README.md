# 最长的有效括号
leetcode: https://leetcode.cn/problems/longest-valid-parentheses/?favorite=2cktkvj <br />
思路 <br />
    1、声明字符串ret <br />
    2、写一个函数, 用来确认是否是有效的括号func1 <br />
    3、第一层循环字符串, 当前key, i <br />
    4、从j = i + 1开始, 循环字符串, 将s[i:j + 1]的字符串, 判断该字符串是否是有效括号 <br />
    5、如果是有效括号, 判断长度跟ret对比, 如果比ret长, 则ret = 该字符串 <br />
    6、返回ret的长度 <br />
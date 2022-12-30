# 最长回文子串
leetcode: https://leetcode.cn/problems/longest-palindromic-substring/ <br />
思路
    1、写个函数判断字符串是否是回文字符串 <br />
    2、从字符串的第一个字符开启第一次循环 <br />
    3、第一次循环的当前字符的下一个字符, 开启第二次循环, 判断当前字符到第二次循环末尾的所有字符串是否是回文字符串 <br />
    4、如果有字符串是, 则查看长度大小 <br />
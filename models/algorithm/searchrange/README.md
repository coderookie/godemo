# 
leetcode: https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/?favorite=2cktkvj <br />
思路 <br />
    1、设置start和end都等于-1 <br />
    2、循环数组, 当start = -1并且当前值等于target的时候, start和end都赋值成当前key <br />
    3、当当前值等于target的时候, end等于当前key <br />
    4、返回[]int{start, end} <br />
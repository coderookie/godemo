# 寻找两个正序数组的中位数
leetcode: https://leetcode.cn/problems/median-of-two-sorted-arrays/?favorite=2cktkvj <br />
思路
    1、循环数组1, 同时记录数组2的当前下标k2, 初始化0
    2、初始化数组3
    3、查看数组1的当前值和数组2的k2下标的值, 如果v1 > v2, v1推入数组3, 反之, 将v2推入数组3, 同时循环数组2, 以此类推, 直到数组1和2合并成一个有序数组
    4、二分查找数组3的中位数
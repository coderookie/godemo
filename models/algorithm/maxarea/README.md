# 盛最多水的容器
leetcode: https://leetcode.cn/problems/container-with-most-water/?favorite=2cktkvj <br />
思路 <br />
    1、声明area <br />
    2、循环第一遍数组, 记录当前的键值k1、v1 <br />
    3、在当前的值后面, 开始第二次循环, area1 = (k2 - k1) * min(v1, v2) <br />
    4、对比area1和area大小, 返回 <br />
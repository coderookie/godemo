# 电话号码的字母组合
leetcode: https://leetcode.cn/problems/letter-combinations-of-a-phone-number/?favorite=2cktkvj <br />
思路 <br />
    1、将数组映射成字符的数组. eg: 2 => ["a", "b", "c"], 3 => ["d", "e", "f"]等 <br />
    2、将目标数字, 转换成二维的字符数组 <br />
    3、声明返回值 <br />
    4、声明函数func1, 参数1: 组合1, 参数2: 组合2, 返回组合1和组合2的排列组合 <br />
    5、循环二维数组, 将里面的一维数组作为参数2, 函数的返回值(默认空)作为参数1, 循环到二维数组结束为止, 返回 <br />
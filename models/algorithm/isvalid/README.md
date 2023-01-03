# 有效的括号
leetcode: https://leetcode.cn/problems/valid-parentheses/?favorite=2cktkvj <br />
思路 <br />
    1、申请一个slice <br />
    2、设置一个map, 括号与括号的对应关系 ")" => "(", "]" => "["等 <br />
    3、循环字符串, 取出当前字符的map中的对应括号, 如果括号等于slice最后一个元素, 则删除slice最后一个元素, 否则将当前括号放入slice <br />
    4、如果最后slice是空的, 则是有效的括号, 否则无效 <br />
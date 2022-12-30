# 删除链表的倒数第N个结点
leetcode: https://leetcode.cn/problems/remove-nth-node-from-end-of-list/?favorite=2cktkvj <br />
思路 <br />
    1、计算链表长度, 倒数N个节点, 换算成正数M个节点 <br />
    2、设置新的头结点, 头结点的next等于要删除的原链表 <br />
    3、循环链表, 当前节点等于M的时候, 当前节点的前一个节点等于当前节点的下一个节点, 退出循环 <br />
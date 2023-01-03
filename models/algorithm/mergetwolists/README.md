# 合并两个有序链表
leetcode: https://leetcode.cn/problems/merge-two-sorted-lists/?favorite=2cktkvj <br />
思路 <br />
    1、递归思想 <br />
    2、返回条件设定 l1等于nil, 返回l2, l2等于nil, 返回l1, l1 && l2都等于nil, 返回nil <br />
    3、当l1的节点小于等于l2的节点, l1的下一个节点则等于l1的下一个节点和l2的合并结果 <br />
    4、当l1的节点大于l2的节点, l2的下一个节点等于l1的节点和l2的下一个节点的合并结果 <br />
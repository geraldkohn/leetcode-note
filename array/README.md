# 数组

## 数组理论基础

## 二分查找
注意边界条件

## 移除元素
1. 暴力解法: 两遍for循环, 第一个for循环遍历数组, 第二个for循环更新数组
2. 双指针法: 通过一个快指针和一个慢指针在一个for循环下完成两个for循环的工作
> 快指针: 遍历原来数组的下标
慢指针: 遍历新数组的下标

这道题不允许用两个数组来解决, 只能在一个数组上操作. 假设能在两个数组上解决, 快指针就是代表原来数组的下标, 慢指针就是代表新数组的下标. 将快指针指向的元素覆盖慢指针指向的元素, 就是这道题的解题精髓.

## 长度最小的子数组
1. 暴力解法
2. 滑动窗口解法

滑动窗口的精妙之处在于根据当前子序列和大小的情况，不断调节子序列的起始位置。从而将O(n^2)的暴力解法降为O(n).

## 螺旋矩阵
模拟行为, 注意找到行为的逻辑点. 然后再进行设计算法.
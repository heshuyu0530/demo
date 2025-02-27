package leetcode

/*
给定一个正整数数组 nums 和一个整数 target 。

向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：

例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

 

示例 1：

输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
示例 2：

输入：nums = [1], target = 1
输出：1
 

提示：

1 <= nums.length <= 20
0 <= nums[i] <= 1000
0 <= sum(nums[i]) <= 1000
-1000 <= target <= 1000

*/

func FindTargetSumWays(nums []int, target int) int {
	var count int
	var backtrack func(int, int)
    backtrack = func(index, sum int) {
        if index == len(nums) { // 遍历到数组末尾
            if sum == target { // 当前表达式的和为目标值
                count++ // 计数加一
            }
            return // 返回上一层
        }
        // 考虑添加正号
        backtrack(index+1, sum+nums[index])
        // 考虑添加负号
        backtrack(index+1, sum-nums[index])
    }
    backtrack(0, 0) // 从索引 0 和和为 0 开始回溯
    return count
}

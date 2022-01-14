/**
 * @param {number[]} nums
 * @param {number} k
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var rotate = function (nums, k) {
    for (let i = 0; i < k; i++) {
        let replaced
        let nextValue
        let nextIndex

        for (let j = 0; j < nums.length; j++) {
            if (j == nums.length - 1) {
                nextIndex = 0
            } else {
                nextIndex = j + 1
            }
            
            nextValue = nums[nextIndex]
            
            console.log(`i ${i} j ${j} nextIndex ${nextIndex} nextValue ${nextValue}`)

            nums[nextIndex] = replaced === undefined ? nums[j] : replaced
            replaced = nextValue
        }
    }

    console.log('nums', JSON.stringify(nums, null, 4))
}

rotate([2147483647,-2147483648,33,219,0], 4)
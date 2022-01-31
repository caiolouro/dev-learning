/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
	const complementsIndexes = new Map()

	for (let i = 0; i < nums.length; i++) {
		const element = nums[i]
		if (complementsIndexes.has(element)) {
			console.log(`${complementsIndexes.get(element)} ${i}`)
			return
            // return [complementsIndexes.get(element), i]
		} else {
			complementsIndexes.set(target - element, i)
		}
	}
}

// twoSum([2,7,11,15], 18)
twoSum([3, 2, 4], 7)

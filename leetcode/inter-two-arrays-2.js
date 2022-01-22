/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var intersect = function (nums1, nums2) {
	const nums1count = new Map()
	const result = []

	for (let i = 0; i < nums1.length; i++) {
		const num = nums1[i]
		let count = nums1count.get(num)
		if (count > 0) nums1count.set(num, count + 1)
		else nums1count.set(num, 1)
	}

	for (let i = 0; i < nums2.length; i++) {
		const num = nums2[i]
		let count = nums1count.get(num)
		if (count > 0) {
			result.push(num)
			nums1count.set(num, --count)
		}
	}

	console.log("result", JSON.stringify(result, null, 4))
}

// For this implementation, the nums1 array should be the shorter one, so in some cases, the for loop will end quickly
var forcedInputSortingIntersect = function (nums1, nums2) {
	const ascSortFn = function (a, b) {
		return a - b
	}
	nums1.sort(ascSortFn)
	nums2.sort(ascSortFn)

	console.log("nums1", JSON.stringify(nums1, null, 4))
	console.log("nums2", JSON.stringify(nums2, null, 4))

	let j = 0
	const result = []

	for (let i = 0; i < nums1.length; i++) {
		const num1 = nums1[i]

		if (num1 > nums2[j]) {
			while (num1 > nums2[j]) j++
		}

		if (num1 == nums2[j]) {
			result.push(num1)
			j++
			// i++
		}
	}

	console.log("result", JSON.stringify(result, null, 4))
}

// intersect([4, 9, 5], [9, 4, 9, 8, 4])
forcedInputSortingIntersect([4, 9, 5, 1, 10, 1], [1, 9, 4, 9, 8, 4, 1, 1])

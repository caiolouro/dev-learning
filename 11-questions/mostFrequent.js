function most_frequent(input) {
    if (!input || input.length === 0) return null;

    const map = new Map();
    const max = { key: undefined, count: 0 };

    input.forEach(key => {
        let count = map.get(key);
        if (count) count++;
        else count = 1;
        map.set(key, count);

        if (count > max.count) {
            max.key = key;
            max.count = count;
        }
    })

    return max.key;
}

const shouldBeOne = most_frequent([1, 3, 1, 3, 2, 1]);
console.log(shouldBeOne);

const shouldBeThree = most_frequent([3, 3, 1, 3, 2, 1]);
console.log(shouldBeThree);

const shouldBeNull = most_frequent([]);
console.log(shouldBeNull);

const shouldBeZero = most_frequent([0]);
console.log(shouldBeZero);

const shouldBeNegativeOne = most_frequent([0, -1, 10, 10, -1, 10, -1, -1, -1, 1]);
console.log(shouldBeNegativeOne);
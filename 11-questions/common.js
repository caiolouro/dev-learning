function common_elements_v1(inputA, inputB) {
    const common = [];

    const smaller = inputA;
    const bigger = inputB;
    if (inputA.length > inputB.length) {
        smaller = inputB;
        bigger = inputA;
    }

    let foundInnerIndex = -1;
    for (let i = 0; i < smaller.length; i++) {
        const outerKey = smaller[i];
        const firstInnerIndex = (foundInnerIndex > -1) ? (foundInnerIndex + 1) : 0;
        
        for (let j = firstInnerIndex; j < bigger.length; j++) {
            const innerKey = bigger[j];
            if (outerKey === innerKey) {
                common.push(outerKey);
                foundInnerIndex = j;
            }
        }
    }

    return common;
};

function common_elements(inputA, inputB) {
    const common = [];

    let i = 0, j = 0;

    while (i < inputA.length && j < inputB.length) {
        const a = inputA[i];
        const b = inputB[j];

        if (a === b) {
            common.push(a);
            i++;
            j++;
        } else if (a < b) {
            i++;
        } else if (b < a) {
            j++;
        }
    }

    return common;
};

const result1 = common_elements([1, 3, 4, 6, 7, 9], [1, 2, 4, 5, 9, 10]);
console.log('result1', result1);

const result2 = common_elements([1, 2, 9, 10, 11, 12], [0, 1, 2, 3, 4, 5, 8, 9, 10, 12, 14, 15]);
console.log('result2', result2);

const result3 = common_elements([0, 1, 2, 3, 4, 5], [6, 7, 8, 9, 10, 11]);
console.log('result3', result3);
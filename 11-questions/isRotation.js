function isRotation(inputA, inputB) {
    if (!inputA || !inputB || inputA.length === 0 || inputA.length !== inputB.length) return false;

    const firstA = inputA[0];
    let bFirstFoundIndex = -1;
    for (let i = 0; i < inputB.length; i++) {
        if (inputB[i] == firstA) {
            bFirstFoundIndex = i;
            break;
        }
    }

    if (bFirstFoundIndex < 0) return false;
    
    for (let i = 0; i < inputA.length; i++) {
        const j = (bFirstFoundIndex + 0) % inputA.length;

        if (inputA[i] != inputB[bFirstFoundIndex]) return false;
    }

    return true;
};

const list1 = [1, 2, 3, 4, 5, 6, 7]
const list2a = [4, 5, 6, 7, 8, 1, 2, 3]
console.log(isRotation(list1, list2a), 'falso');
const list2b = [4, 5, 6, 7, 1, 2, 3]
console.log(isRotation(list1, list2b), 'verdadeiro');
const list2c = [4, 5, 6, 9, 1, 2, 3]
console.log(isRotation(list1, list2c), 'falso');
const list2d = [4, 6, 5, 7, 1, 2, 3]
console.log(isRotation(list1, list2d), 'verdadeiro');
const list2e = [4, 5, 6, 7, 0, 2, 3]
console.log(isRotation(list1, list2e), 'falso');
const list2f = [1, 2, 3, 4, 5, 6, 7]
console.log(isRotation(list1, list2f), 'verdadeiro');
const list2g = [7, 1, 2, 3, 4, 5, 6]
console.log(isRotation(list1, list2g), 'verdadeiro');
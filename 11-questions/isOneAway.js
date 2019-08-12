function is_one_away_two_whiles(inputA, inputB) {    
    let indexA = 0, indexB = 0, diffs = 0;
    while (indexA < inputA.length) {
        const a = inputA[indexA];
        const b = inputB[indexB];
        if (a === b) {
            indexA++;
            indexB++;
            continue;
        }
        
        if (++diffs > 1) return false;

        if (a === inputB[indexB + 1]) {
            indexB++;
            continue;
        }

        if (inputA[indexA + 1] === inputB[indexB]) {
            indexA++;
            continue;
        }

        if (inputA[indexA + 1] === inputB[indexB + 1]) {
            indexA++;
            indexB++;
            continue;
        }
    }

    while (indexB < inputB.length) {
        if (++diffs > 1) return false;
    }

    return true;
};

function is_one_away_strategies(inputA, inputB) {
    const lengthDiff = Math.abs(inputA.length - inputB.length);
    if (lengthDiff > 1) {
        return false;
    } else if (lengthDiff === 1) {
        let bigger = inputA;
        let smaller = inputB;

        if (inputA.length < inputB.length) {
            bigger = inputB;
            smaller = inputA;
        }

        let indexSmaller = 0, diffs = 0;
        for (let i = 0; i < bigger.length; i++) {
            if (bigger[i] !== smaller[indexSmaller]) {
                if (++diffs > 1) return false;
            } else {
                indexSmaller++;
            }
        }

        return true;
    } else {
        let diffs = 0;
        for (let i = 0; i < inputA.length; i++) {
            if (inputA[i] !== inputB[i]) {
                if (++diffs > 1) return false;
            }
        }
        return true;
    }
};

function is_one_away(inputA, inputB) {
    const lengthDiff = Math.abs(inputA.length - inputB.length);
    if (lengthDiff > 1) {
        return false;
    } else if (lengthDiff === 1) {
        let bigger = inputA;
        let smaller = inputB;

        if (inputA.length < inputB.length) {
            bigger = inputB;
            smaller = inputA;
        }

        let i = 0, diffs = 0;
        while (i < smaller.length) {
            if (bigger[i + diffs] !== smaller[i]) {
                if (++diffs > 1) return false;
                continue;
            }

            i++;
        }

        return true;
    } else {
        let diffs = 0;
        for (let i = 0; i < inputA.length; i++) {
            if (inputA[i] !== inputB[i]) {
                if (++diffs > 1) return false;
            }
        }
        return true;
    }
};

console.log(is_one_away("abcde", "abcd"), 'True');
console.log(is_one_away("abde", "abcde"), 'True');
console.log(is_one_away("a", "a"), 'True');
console.log(is_one_away("abcdef", "abqdef"), 'True');
console.log(is_one_away("abcdef", "abccef"), 'True');
console.log(is_one_away("abcdef", "abcde"), 'True');
console.log(is_one_away("aaa", "abc"), 'False');
console.log(is_one_away("abcde", "abc"), 'False');
console.log(is_one_away("abc", "abcde"), 'False');
console.log(is_one_away("abc", "bcc"), 'False');
console.log(is_one_away("abcdefighj", "abcdefghjk"), 'False');
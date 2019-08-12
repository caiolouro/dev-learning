function non_repeating_two_maps(input) {
    if (!input || input.length === 0) return null;

    const uniqueChars = new Map();
    const nonUniqueChars = new Map();
    for (let i = 0; i < input.length; i++) {
        const char = input[i];
        if (uniqueChars.get(char)) {
            uniqueChars.delete(char);
            nonUniqueChars.set(char, true);
        } else if (!nonUniqueChars.get(char)) {
            uniqueChars.set(char, true);
        }
    }

    return uniqueChars.size > 0 ? uniqueChars.keys().next().value : null;
};

function non_repeating(input) {
    if (!input || input.length === 0) return null;

    const chars = new Map();

    for (let i = 0; i < input.length; i++) {
        const char = input[i];
        const count = chars.get(char);
        if (count) chars.set(char, count + 1);
        else chars.set(char, 1);
    }

    for (let entry of chars.entries()) {
        if (entry[1] === 1) return entry[0];
    }

    return null;
};

console.log('c', non_repeating("abcab"));
console.log('null', non_repeating("abab"));
console.log('c', non_repeating("aabbbc"));
console.log('d', non_repeating("aabbdbc"));
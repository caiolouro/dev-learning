function printField(field) {
    field.forEach(r => console.log(r));
};

function revealSafeSurroundings(field, numRows, numColumns, rowNumber, columnNumber) {
    field[rowNumber][columnNumber] = -2;

    for (let i = rowNumber - 1; i <= rowNumber + 1; i++) {
        for (let j = columnNumber - 1; j <= columnNumber + 1; j++) {
            if (i < 0 || j < 0 || i >= numRows || j >= numColumns) continue;
            
            if (field[i][j] === 0) revealSafeSurroundings(field, numRows, numColumns, i, j);
        }
    }
};

function click(field, numRows, numColumns, clickedRow, clickedColumn) {
    if (field[clickedRow][clickedColumn] !== 0) return field;

    revealSafeSurroundings(field, numRows, numColumns, clickedRow, clickedColumn);

    return field;
};

const field1 = [
    [0, 0, 0, 0, 0],
    [0, 1, 1, 1, 0],
    [0, 1, -1, 1, 0]
];

console.log(printField(click(field1, 3, 5, 0, 0)));

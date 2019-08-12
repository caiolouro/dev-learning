function mine_sweeper_r_c_b(bombs, num_rows, num_cols) {
    const map = [];

    for (let i = 0; i < num_rows; i++) {
        const r = [];
        for (let j = 0; j < num_cols; j++) {
            if (bombs.some(bomb => bomb[0] === i && bomb[1] === j)) {
                r.push(-1);
                continue;
            }

            const lowR = i === 0 ? 0 : i - 1;
            const highR = i === num_rows - 1 ? i : i + 1;
            const lowC = j === 0 ? 0 : j - 1;
            const highC = j === num_cols - 1 ? j : j + 1;

            let bombsCount = 0;
            bombs.forEach(bomb => {
                if (bomb[0] >= lowR && bomb[0] <= highR && bomb[1] >= lowC && bomb[1] <= highC) bombsCount++;
            });

            r.push(bombsCount);
        }
        map.push(r);
    }
    
    return map;
};

function mine_sweeper(bombs, num_rows, num_cols) {
    const map = [];

    for (let i = 0; i < num_rows; i++) {
        const r = [];
        for (let j = 0; j < num_cols; j++) {
            r.push(0);
        }
        map.push(r);
    }

    bombs.forEach(bomb => {
        const bombR = bomb[0];
        const bombC = bomb[1];

        for (let i = bombR - 1; i <= bombR + 1; i++) {
            for (let j = bombC - 1; j <= bombC + 1; j++) {
                if (i < 0 || i > num_rows - 1 || j < 0 || j > num_cols - 1) continue;

                if (i === bombR && j === bombC) {
                    map[i][j] = -1;
                    continue;
                }

                if (map[i][j] > -1) map[i][j]++;
            }
        }
    });
    
    return map;
};

print_map(mine_sweeper([[0, 2], [2, 0]], 3, 3));
console.log();
print_map(mine_sweeper([[0, 0], [0, 1], [1, 2]], 3, 4));
console.log();
print_map(mine_sweeper([[1, 1], [1, 2], [2, 2], [4, 3]], 5, 5));

function print_map(map) {
    map.forEach(r => console.log(r));
};
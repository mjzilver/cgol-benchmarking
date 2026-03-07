const fs = require('fs');
const readline = require('readline');

const OUTPUT_FILE = 'output.txt';
let board = [];
let size = -1;
let silent = false;
let serverMode = false;

const neighbors = [
    [-1, -1], [-1, 0], [-1, 1],
    [0, -1], [0, 1],
    [1, -1], [1, 0], [1, 1]
];

function parseArgs() {
    const args = process.argv.slice(2);
    let filePath = null;
    let amount = 10;

    for (let i = 0; i < args.length; i++) {
        if (args[i] === '--file' && i + 1 < args.length) {
            filePath = args[i + 1];
            i++;
        } else if (args[i] === '--amount' && i + 1 < args.length) {
            amount = parseInt(args[i + 1]);
            i++;
        } else if (args[i] === '--silent') {
            silent = true;
        } else if (args[i] === '--server') {
            serverMode = true;
        }
    }

    return { filePath, amount };
}

function parseBoard(filePath) {
    const data = fs.readFileSync(filePath, 'utf8');
    const lines = data.trim().split('\n');
    size = lines.length;
    board = [];

    for (let y = 0; y < size; y++) {
        board[y] = [];
        for (let x = 0; x < size; x++) {
            board[y][x] = lines[y][x] === '1' ? 1 : 0;
        }
    }
}

function nextState() {
    const newBoard = Array(size).fill(null).map(() => Array(size).fill(0));

    for (let y = 0; y < size; y++) {
        for (let x = 0; x < size; x++) {
            let count = 0;

            for (let i = 0; i < 8; i++) {
                const ny = (y + neighbors[i][0] + size) % size;
                const nx = (x + neighbors[i][1] + size) % size;
                count += board[ny][nx];
            }

            if (board[y][x] === 1) {
                if (count < 2 || count > 3) {
                    newBoard[y][x] = 0;
                } else {
                    newBoard[y][x] = 1;
                }
            } else {
                if (count === 3) {
                    newBoard[y][x] = 1;
                } else {
                    newBoard[y][x] = 0;
                }
            }
        }
    }

    board = newBoard;
}

function printBoard() {
    let output = '';
    for (let y = 0; y < size; y++) {
        for (let x = 0; x < size; x++) {
            output += board[y][x];
        }
        output += '\n';
    }
    fs.writeFileSync(OUTPUT_FILE, output);
}

function runSimulation(filePath, amount) {
    parseBoard(filePath);
    
    for (let i = 0; i < amount; i++) {
        nextState();
    }
    
    if (!silent) {
        printBoard();
    }
}

function serverLoop() {
    console.log('READY');
    
    const rl = readline.createInterface({
        input: process.stdin,
        output: process.stdout,
        terminal: false
    });

    rl.on('line', (line) => {
        const trimmed = line.trim();
        
        if (trimmed === '') return;
        
        if (trimmed === 'SHUTDOWN') {
            rl.close();
            process.exit(0);
        }
        
        if (trimmed.startsWith('RUN ')) {
            const parts = trimmed.substring(4).trim().split(/\s+/);
            if (parts.length >= 1) {
                const filePath = parts[0];
                const reqAmount = parts.length >= 2 ? parseInt(parts[1]) : 1;
                
                parseBoard(filePath);
                for (let i = 0; i < reqAmount; i++) {
                    nextState();
                }
                
                if (!silent) {
                    printBoard();
                }
                
                console.log('DONE');
            } else {
                console.log('ERROR bad request');
            }
        } else {
            console.log('ERROR unknown command');
        }
    });
}

function main() {
    const { filePath, amount } = parseArgs();

    if (serverMode) {
        serverLoop();
    } else {
        if (!filePath) {
            console.error('Error: --file is required');
            process.exit(1);
        }
        runSimulation(filePath, amount);
    }
}

main();

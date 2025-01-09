#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

#define OUTPUT_FILE "output.txt"

int **board;
int size = -1;
int amount = -1;
int silent = 0;
char *inputFile = "invalid";

int neighbors[8][2] = {
    {-1, -1}, {-1, 0}, {-1, 1},
    {0, -1}, {0, 1},
    {1, -1}, {1, 0}, {1, 1}
};

void initBoard() {
    board = (int **)malloc(size * sizeof(int *));
    for (int i = 0; i < size; i++) {
        board[i] = (int *)malloc(size * sizeof(int));
    }
}

void freeBoard() {
    for (int i = 0; i < size; i++) {
        free(board[i]);
    }
    free(board);
}

int **nextState() {
    int **newBoard = (int **)malloc(size * sizeof(int *));
    for (int i = 0; i < size; i++) {
        newBoard[i] = (int *)malloc(size * sizeof(int));
    }

    for (int y = 0; y < size; y++) {
        for (int x = 0; x < size; x++) {
            int count = 0;

            for (int i = 0; i < 8; i++) {
                int ny = (y + neighbors[i][0] + size) % size;
                int nx = (x + neighbors[i][1] + size) % size;
                count += board[ny][nx];
            }

            if (board[y][x] == 1) {
                if (count < 2 || count > 3) {
                    newBoard[y][x] = 0;
                } else {
                    newBoard[y][x] = 1;
                }
            } else {
                if (count == 3) {
                    newBoard[y][x] = 1;
                } else {
                    newBoard[y][x] = 0;
                }
            }
        }
    }

    freeBoard();
    return newBoard;
}

void printBoard() {
    FILE *file = fopen(OUTPUT_FILE, "w");
    for (int y = 0; y < size; y++) {
        for (int x = 0; x < size; x++) {
            fprintf(file, "%d", board[y][x]);
        }
        fprintf(file, "\n");
    }
    fclose(file);
}

void parseBoard() {
    FILE *file = fopen(inputFile, "r");
    if (!file) {
        fprintf(stderr, "File %s not found!\n", inputFile);
        exit(1);
    }

    char line[1024];
    int row = 0;

    while (fgets(line, sizeof(line), file)) {
        char *ptr = line;

        while (isspace(*ptr)) {
            ptr++;
        }

        if (size == -1) {
            size = strlen(ptr);
            while (size > 0 && isspace(ptr[size - 1])) {
                size--;
            }
            initBoard();
        }

        for (int i = 0; i < size; i++) {
            if (isdigit(ptr[i])) {
                board[row][i] = ptr[i] - '0';
            } else {
                fprintf(stderr, "Unexpected character '%c' in input\n", ptr[i]);
                exit(1);
            }
        }
        row++;
    }

    fclose(file);
}

int main(int argc, char *argv[]) {
    for (int i = 1; i < argc; i++) {
        if (strcmp(argv[i], "--file") == 0) {
            inputFile = argv[++i];
        } else if (strcmp(argv[i], "--amount") == 0) {
            amount = atoi(argv[++i]);
        } else if (strcmp(argv[i], "--silent") == 0) {
            silent = 1;
        }
    }

    parseBoard();

    for (int i = 0; i < amount; i++) {
        board = nextState();
    }

    if (!silent) {
        printBoard();
    }

    freeBoard();
    return 0;
}

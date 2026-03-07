#include <iostream>
#include <fstream>
#include <vector>
#include <string>
#include <cstring>

const char* OUTPUT_FILE = "output.txt";

std::vector<std::vector<int>> board;
int size = -1;
int amount = -1;
bool silent = false;
const char* inputFile = "invalid";

const int neighbors[8][2] = {
    {-1, -1}, {-1, 0}, {-1, 1},
    {0, -1}, {0, 1},
    {1, -1}, {1, 0}, {1, 1}
};

void parseBoard() {
    std::ifstream file(inputFile);
    if (!file.is_open()) {
        std::cerr << "File " << inputFile << " not found!\n";
        exit(1);
    }

    std::string line;
    size = -1;
    board.clear();

    while (std::getline(file, line)) {
        if (line.empty()) continue;

        if (size == -1) {
            size = line.length();
        }

        std::vector<int> row;
        row.reserve(size);
        for (int i = 0; i < size; i++) {
            row.push_back(line[i] - '0');
        }
        board.push_back(row);
    }

    file.close();
}

std::vector<std::vector<int>> nextState() {
    std::vector<std::vector<int>> newBoard(size, std::vector<int>(size, 0));

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

    return newBoard;
}

void printBoard() {
    std::ofstream file(OUTPUT_FILE);
    for (int y = 0; y < size; y++) {
        for (int x = 0; x < size; x++) {
            file << board[y][x];
        }
        file << '\n';
    }
    file.close();
}

int main(int argc, char* argv[]) {
    bool serverMode = false;

    for (int i = 1; i < argc; i++) {
        if (strcmp(argv[i], "--file") == 0 && i + 1 < argc) {
            inputFile = argv[++i];
        } else if (strcmp(argv[i], "--amount") == 0 && i + 1 < argc) {
            amount = std::atoi(argv[++i]);
        } else if (strcmp(argv[i], "--silent") == 0) {
            silent = true;
        } else if (strcmp(argv[i], "--server") == 0) {
            serverMode = true;
        }
    }

    if (!serverMode) {
        parseBoard();

        for (int i = 0; i < amount; i++) {
            board = nextState();
        }

        if (!silent) {
            printBoard();
        }

        return 0;
    }

    std::cout << "READY" << std::endl;

    std::string line;
    while (std::getline(std::cin, line)) {
        if (line.empty()) continue;

        if (line == "SHUTDOWN") {
            break;
        }

        if (line.substr(0, 4) == "RUN ") {
            std::string params = line.substr(4);
            size_t space = params.find(' ');
            
            std::string filePath = params.substr(0, space);
            int reqAmount = 1;
            
            if (space != std::string::npos) {
                reqAmount = std::atoi(params.substr(space + 1).c_str());
            }

            inputFile = filePath.c_str();
            parseBoard();

            for (int i = 0; i < reqAmount; i++) {
                board = nextState();
            }

            if (!silent) {
                printBoard();
            }

            std::cout << "DONE" << std::endl;
        } else {
            std::cout << "ERROR unknown command" << std::endl;
        }
    }

    return 0;
}

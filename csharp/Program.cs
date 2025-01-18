using System;
using System.IO;

class Program
{
    const string OUTPUT_FILE = "output.txt";
    static int[,] board;
    static int size = -1;
    static int amount = -1;
    static bool silent = false;
    static string inputFile = "invalid";

    static int[,] neighbors = new int[,]
    {
        {-1, -1}, {-1, 0}, {-1, 1},
        {0, -1}, {0, 1},
        {1, -1}, {1, 0}, {1, 1}
    };

    static void Main(string[] args)
    {
        for (int i = 0; i < args.Length; i++)
        {
            if (args[i] == "--file")
            {
                inputFile = args[++i];
            }
            else if (args[i] == "--amount")
            {
                amount = int.Parse(args[++i]);
            }
            else if (args[i] == "--silent")
            {
                silent = true;
            }
        }

        ParseBoard();

        for (int i = 0; i < amount; i++)
        {
            board = NextState();
        }

        if (!silent)
        {
            PrintBoard();
        }
    }

    static void ParseBoard()
    {
        var lines = File.ReadAllLines(inputFile);
        size = lines.Length;
        board = new int[size, size];

        for (int y = 0; y < size; y++)
        {
            for (int x = 0; x < size; x++)
            {
                board[y, x] = lines[y][x] - '0';
            }
        }
    }

    static int[,] NextState()
    {
        var newBoard = new int[size, size];

        for (int y = 0; y < size; y++)
        {
            for (int x = 0; x < size; x++)
            {
                int count = 0;

                for (int i = 0; i < 8; i++)
                {
                    int ny = (y + neighbors[i, 0] + size) % size;
                    int nx = (x + neighbors[i, 1] + size) % size;
                    count += board[ny, nx];
                }

                if (board[y, x] == 1)
                {
                    if (count < 2 || count > 3)
                    {
                        newBoard[y, x] = 0;
                    }
                    else
                    {
                        newBoard[y, x] = 1;
                    }
                }
                else
                {
                    if (count == 3)
                    {
                        newBoard[y, x] = 1;
                    }
                    else
                    {
                        newBoard[y, x] = 0;
                    }
                }
            }
        }

        return newBoard;
    }

    static void PrintBoard()
    {
        using (var writer = new StreamWriter(OUTPUT_FILE))
        {
            for (int y = 0; y < size; y++)
            {
                for (int x = 0; x < size; x++)
                {
                    writer.Write(board[y, x]);
                }
                writer.Write("\n");
            }
        }
    }
}
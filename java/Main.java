import java.io.*;
import java.util.*;

public class Main {
    private static final String OUTPUT_FILE = "output.txt";
    private static int[][] board;
    private static int size = -1;
    private static int amount = -1;
    private static boolean silent = false;
    private static String inputFile = "invalid";

    private static final int[][] neighbors = {
        {-1, -1}, {-1, 0}, {-1, 1},
        {0, -1}, {0, 1},
        {1, -1}, {1, 0}, {1, 1}
    };

    public static void main(String[] args) {
        boolean serverMode = false;
        for (int i = 0; i < args.length; i++) {
            if (args[i].equals("--file")) {
                inputFile = args[++i];
            } else if (args[i].equals("--amount")) {
                amount = Integer.parseInt(args[++i]);
            } else if (args[i].equals("--silent")) {
                silent = true;
            } else if (args[i].equals("--server")) {
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

            return;
        }

        System.out.println("READY");
        System.out.flush();

        try (BufferedReader reader = new BufferedReader(new InputStreamReader(System.in))) {
            String line;
            while ((line = reader.readLine()) != null) {
                line = line.trim();
                if (line.isEmpty()) continue;
                if (line.equals("SHUTDOWN")) break;

                if (line.startsWith("RUN ")) {
                    String[] parts = line.substring(4).trim().split("\\s+");
                    if (parts.length >= 1) {
                        inputFile = parts[0];
                        int reqAmount = 1;
                        if (parts.length >= 2) {
                            try {
                                reqAmount = Integer.parseInt(parts[1]);
                            } catch (NumberFormatException e) {
                                reqAmount = 1;
                            }
                        }

                        parseBoard();
                        for (int i = 0; i < reqAmount; i++) {
                            board = nextState();
                        }

                        if (!silent) printBoard();

                        System.out.println("DONE");
                        System.out.flush();
                    } else {
                        System.out.println("ERROR bad request");
                        System.out.flush();
                    }
                } else {
                    System.out.println("ERROR unknown command");
                    System.out.flush();
                }
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private static void parseBoard() {
        try {
            List<String> lines = new ArrayList<>();
            try (BufferedReader reader = new BufferedReader(new FileReader(inputFile))) {
                String line;
                while ((line = reader.readLine()) != null) {
                    lines.add(line);
                }
            }

            size = lines.size();
            board = new int[size][size];

            for (int y = 0; y < size; y++) {
                String line = lines.get(y);
                for (int x = 0; x < size; x++) {
                    board[y][x] = line.charAt(x) - '0';
                }
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private static int[][] nextState() {
        int[][] newBoard = new int[size][size];

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

    private static void printBoard() {
        try (PrintWriter writer = new PrintWriter(new FileWriter(OUTPUT_FILE))) {
            for (int y = 0; y < size; y++) {
                for (int x = 0; x < size; x++) {
                    writer.print(board[y][x]);
                }
                writer.print("\n");
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}

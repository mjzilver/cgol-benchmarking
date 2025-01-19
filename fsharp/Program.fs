open System
open System.IO

type clArgs =
    { silent: bool
      file: string option
      amount: int option }

type board = int array array

let rec parseArgs (args: string list) (parsedArgs: clArgs) =
    match args with
    | [] -> parsedArgs
    | "--silent" :: rest -> parseArgs rest { parsedArgs with silent = true }
    | "--file" :: fileName :: rest -> parseArgs rest { parsedArgs with file = Some fileName }
    | "--amount" :: amountStr :: rest ->
        match Int32.TryParse(amountStr) with
        | (true, value) -> parseArgs rest { parsedArgs with amount = Some value }
        | (false, _) ->
            printfn "Invalid value for --amount: %s" amountStr
            parsedArgs
    | unknown :: rest ->
        printfn "Unknown argument: %s" unknown
        parseArgs rest parsedArgs

let parseLine (line: string) : int array =
    line
    |> Seq.map (fun c ->
        match c with
        | '0' -> 0
        | '1' -> 1
        | _ -> failwithf "Invalid character in file: '%c'." c)
    |> Array.ofSeq

let parseFile (fileName: string) : int * board =
    let board: board = File.ReadAllLines fileName |> Array.map parseLine
    let size = Array.length board
    size, board

let directions: (int * int) list = [
  (-1, -1); (-1, 0); (-1, 1);
  ( 0, -1);          ( 0, 1);
  ( 1, -1); ( 1, 0); ( 1, 1)
]

let countNeighbors (board: board) (size: int) (y, x) : int =
    List.fold
        (fun acc (dy, dx) ->
            let ny = (y + dy + size) % size
            let nx = (x + dx + size) % size
            acc + board[ny][nx])
        0
        directions

let nextState (board: board) (size: int) : board =
    Array.init size (fun y ->
        Array.init size (fun x ->
            let nCount = countNeighbors board size (y, x)

            match board[y][x] with
            | 1 -> if nCount = 2 || nCount = 3 then 1 else 0
            | 0 -> if nCount = 3 then 1 else 0
            | _ -> failwith "Invalid cell"))

let boardToString (board: board) : string = 
    board
    |> Array.map (fun row ->
        row 
        |> Array.map string 
        |> String.concat ""
    )
    |> String.concat "\n"
    |> fun res -> res + "\n"

let printBoard (board: board) (fileName: string) =
    File.WriteAllText(fileName, boardToString board)

[<EntryPoint>]
let main argv =
    let defaultArgs =
        { silent = false
          file = None
          amount = None }

    let parsedArgs = parseArgs (argv |> Array.toList) defaultArgs

    let size, board = parseFile parsedArgs.file.Value

    let rec simulateSteps (board: board) (n: int) : board =
        if n = 0 then board
        else simulateSteps (nextState board size) (n - 1)

    let finalBoard = simulateSteps board parsedArgs.amount.Value

    if not parsedArgs.silent then
        printBoard finalBoard "./output.txt"

    0

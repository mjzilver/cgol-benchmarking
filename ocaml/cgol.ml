type board = int array array

type args = {
  silent : bool;
  amount : int;
  file : string;
}

let parse_args () : args =
  let file = ref "" in
  let amount = ref 0 in
  let silent = ref false in
  let flags = [
    ("--file", Arg.Set_string file, "Input file to parse board");
    ("--amount", Arg.Set_int amount, "Amount to simulate");
    ("--silent", Arg.Set silent, "Run without printing the result")
  ] in
  Arg.parse flags (fun _ -> ()) "Usage: program --file <filename> --amount <number>";
  { file = !file; amount = !amount; silent = !silent }

let parse_line (line : string) : int array =
  String.to_seq line 
    |> Seq.map (fun c -> 
      match c with
      | '0' -> 0
      | '1' -> 1
      | _ -> failwith "Unknown char"
    ) 
    |> Array.of_seq

let parse_file (filename : string) : int * board =
  let input = open_in filename in
  let rec parse_lines acc =
    match input_line input with
    | line -> parse_lines (parse_line line :: acc)
    | exception End_of_file -> close_in_noerr input; 
    List.rev acc
  in
  let rows = parse_lines [] in
  let size = List.length rows in
  size, Array.of_list rows

let directions = [
  (-1, -1); (-1, 0); (-1, 1);
  ( 0, -1);          ( 0, 1);
  ( 1, -1); ( 1, 0); ( 1, 1)
]

let count_neighbors (board : board) (size : int) (y, x) =
  List.fold_left (fun acc (dy, dx) ->
    let ny = (y + dy + size) mod size in
    let nx = (x + dx + size) mod size in
    acc + board.(ny).(nx)
  ) 0 directions

let next_state (board : board) (size : int) : board =
  Array.init size (fun y ->
    Array.init size (fun x ->
      let neighbors = count_neighbors board size (y, x) in
      if board.(y).(x) = 1 then
        if neighbors = 2 || neighbors = 3 then 1 else 0
      else
        if neighbors = 3 then 1 else 0
    )
  )

let print_board (board : board) (filename : string) =
  let board_string =
    Array.map (fun row ->
      Array.map string_of_int row
      |> Array.to_list
      |> String.concat ""
    ) board
    |> Array.to_list
    |> String.concat "\n"
  in
  let output = open_out filename in
  output_string output (board_string ^ "\n");
  close_out output  

let simulate (args : args) (output : string) =
  let size, board = parse_file args.file in
  let rec simulate_steps board n =
    if n = 0 then board
    else simulate_steps (next_state board size) (n - 1)
  in
  let final_board = simulate_steps board args.amount in
  if not args.silent then print_board final_board output

let () =
  let args = parse_args () in
  let output = "./output.txt" in
  simulate args output

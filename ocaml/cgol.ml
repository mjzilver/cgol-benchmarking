let parse_args () =
  let file = ref "" in
  let amount = ref 0 in
  let silent = ref false in
  let flags = [
    ("--file", Arg.Set_string file, "Input file to parse board");
    ("--amount", Arg.Set_int amount, "Amount to simulate");
    ("--silent", Arg.Set silent, "Run without printing the result")
  ] in
  Arg.parse flags (fun _ -> ()) "Usage: program --file <filename> --amount <number>";
  (!file, !amount, !silent)

let parse_line line =
  String.to_seq line 
    |> Seq.map (fun c -> int_of_char c - int_of_char '0') 
    |> Array.of_seq

let parse_file filename =
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

let count_neighbors board size (y, x) =
  List.fold_left (fun acc (dy, dx) ->
    let ny = (y + dy + size) mod size in
    let nx = (x + dx + size) mod size in
    acc + board.(ny).(nx)
  ) 0 directions

let next_state board size =
  Array.init size (fun y ->
    Array.init size (fun x ->
      let neighbors = count_neighbors board size (y, x) in
      if board.(y).(x) = 1 then
        if neighbors = 2 || neighbors = 3 then 1 else 0
      else
        if neighbors = 3 then 1 else 0
    ))

let print_board board filename =
  let output = open_out filename in
  Array.iter (fun row ->
    Array.iter (fun value -> output_string output (string_of_int value)) row;
    output_string output "\n";
  ) board;
  close_out output

let simulate file amount silent output =
  let size, board = parse_file file in
  let rec simulate_steps board n =
    if n = 0 then board
    else simulate_steps (next_state board size) (n - 1)
  in
  let final_board = simulate_steps board amount in
  if not silent then print_board final_board output

let () =
  let file, amount, silent = parse_args () in
  let output = "./output.txt" in
  simulate file amount silent output

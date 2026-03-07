type board = int array array

type args = {
  silent : bool;
  amount : int;
  file : string;
  server : bool;
}

let parse_args () : args =
  let file = ref "" in
  let amount = ref 0 in
  let silent = ref false in
  let server = ref false in
  let flags = [
    ("--file", Arg.Set_string file, "Input file to parse board");
    ("--amount", Arg.Set_int amount, "Amount to simulate");
    ("--silent", Arg.Set silent, "Run without printing the result")
    ; ("--server", Arg.Set server, "Run in server mode (accept RUN/SHUTDOWN on stdin)")
  ] in
  Arg.parse flags (fun _ -> ()) "Usage: program --file <filename> --amount <number>";
  { file = !file; amount = !amount; silent = !silent; server = !server }

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
  let base_args = parse_args () in
  if base_args.server then (
    print_endline "READY";
    try
      while true do
        let line = read_line () in
        if String.trim line = "SHUTDOWN" then raise Exit;
        if String.length line >= 4 && String.sub line 0 4 = "RUN " then (
          let rest = String.sub line 4 (String.length line - 4) in
          let parts = String.split_on_char ' ' rest |> List.filter (fun s -> String.length s > 0) in
          match parts with
          | file :: amt :: _ ->
              let args = { base_args with file = file; amount = int_of_string amt } in
              simulate args "./output.txt";
              print_endline "DONE"
          | file :: [] ->
              let args = { base_args with file = file; amount = 1 } in
              simulate args "./output.txt";
              print_endline "DONE"
          | _ -> print_endline "ERROR bad request"
        ) else (
          print_endline "ERROR unknown command"
        )
      done
    with
    | End_of_file -> ()
    | Exit -> ()
  ) else (
    let args = base_args in
    let output = "./output.txt" in
    simulate args output
  )

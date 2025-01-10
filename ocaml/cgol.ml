let file = ref ""
let amount = ref (-1)
let board = ref [||]
let size = ref (-1)
let output = "./output.txt"
let silent = ref false

let flags = [
  ("--file", Arg.Set_string file, "Input file to parse board");
  ("--amount", Arg.Set_int amount, "Amount to simulate");
  ("--silent", Arg.Set silent, "Run without priting the result");
]

let parse_line (y : int) (line : string) =
  let cols = String.to_seq line |> List.of_seq |> List.map (fun c -> int_of_string (String.make 1 c)) in
  if !size = -1 then (
    size := List.length cols;
    board := Array.make_matrix !size !size 0
  );
  List.iteri (fun x value -> !board.(y).(x) <- value) cols

let parse_file () =
  let y = ref 0 in
  let input = open_in !file in
  try
    while true do
      let line = input_line input in
      parse_line !y line;
      incr y
    done
  with End_of_file ->
    close_in_noerr input

let print_board () =
  let output_file = open_out output in 
  Array.iter (fun row ->
    Array.iter (fun value ->
      Printf.fprintf output_file "%d" value 
    ) row;
    Printf.fprintf output_file "\n"
  ) !board;
  close_out output_file

let count_neighbors y x =
  let directions = [ (-1, -1); (-1, 0); (-1, 1);
                     (0, -1);         (0, 1);
                     (1, -1); (1, 0); (1, 1) ] in
  List.fold_left (fun acc (dy, dx) ->
    let ny = (y + dy + !size) mod !size in
    let nx = (x + dx + !size) mod !size in
    if !board.(ny).(nx) = 1 then
      acc + 1
    else
      acc
  ) 0 directions

let next_state () =
  let new_board = Array.make_matrix !size !size 0 in
  for y = 0 to !size - 1 do
    for x = 0 to !size - 1 do
      let neighbors = count_neighbors y x in
      if !board.(y).(x) = 1 then (
        if neighbors = 2 || neighbors = 3 then
          new_board.(y).(x) <- 1
      ) else (
        if neighbors = 3 then
          new_board.(y).(x) <- 1
      )
    done
  done;
  board := new_board

let () =
  Arg.parse flags
    (fun _ -> ()) 
    "Usage: program --file <filename> --amount <number>";
    
  parse_file();
  
  for _ = 1 to !amount do 
    next_state ();
  done;

  if not !silent then (
    print_board ();
  )

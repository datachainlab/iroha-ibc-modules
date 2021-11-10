type arg_type =
  | ArgString
  | ArgBytes
  | ArgBytesM of int (* bytesM *)
  | ArgUintM of int (* uintM *)
  | ArgIntM of int (* intM *)

let type_to_solidity = function
  | ArgString -> "string memory"
  | ArgBytes -> "bytes memory"
  | ArgBytesM m -> Printf.sprintf "bytes%d" m
  | ArgUintM m -> Printf.sprintf "uint%d" m
  | ArgIntM m -> Printf.sprintf "int%d" m

let type_to_selector = function
  | ArgString -> "string"
  | ArgBytes -> "bytes"
  | ArgBytesM m -> Printf.sprintf "bytes%d" m
  | ArgUintM m -> Printf.sprintf "uint%d" m
  | ArgIntM m -> Printf.sprintf "int%d" m

type arg_spec = {
  name: string;
  typ: arg_type;
}

let arg_to_solidity {name; typ} =
  Printf.sprintf "%s _%s" (type_to_solidity typ) name

type api_spec = {
  name: string;
  args: arg_spec list;
}

exception Invalid_line of string
exception Invalid_first_part of string
exception Invalid_arg_type of string

let parse_arg_type word : arg_type =
  let open Core in
  match word with
    | "string" -> ArgString
    | "bytes" -> ArgBytes
    | w ->
        let suffix s ~pos =
          let len = String.length s - pos in
          String.sub s ~pos ~len
        in
        if String.is_prefix w ~prefix:"int" then
          ArgIntM (suffix w ~pos:3 |> Int.of_string)
        else if String.is_prefix w ~prefix:"uint" then
          ArgUintM (suffix w ~pos:4 |> Int.of_string)
        else if String.is_prefix w ~prefix:"bytes" then
          ArgBytesM (suffix w ~pos:5 |> Int.of_string)
        else
          raise (Invalid_arg_type w)

let parse_first part : string * arg_type list =
  let open Core in
  let words = Str.split (Str.regexp "[,()]+") part in
  match words with
    | api_name :: arg_type_names ->
        let arg_types = List.map arg_type_names ~f:parse_arg_type in
        api_name, arg_types
    | _ -> raise (Invalid_first_part part)

let parse_second part : string list =
  let open Core in
  let words = String.split part ~on:',' in
  List.map words ~f:String.strip

let parse_line line : api_spec =
  let open Core in
  let parts = String.split_on_chars ~on:['\t'] line in
  match parts with
    | first :: rest ->
        let (api_name, arg_types) = parse_first first in
        let arg_names =
          match rest with
          | second :: [] -> parse_second second
          | [] -> []
          | _ -> raise (Invalid_line line)
        in
        let args : arg_spec list =
          List.zip_exn arg_names arg_types |>
          List.map ~f:(fun (name, typ) -> {name; typ})
        in
        {name = api_name; args}
    | _ -> raise (Invalid_line line)

let print_templated_function api =
  let open Core in
  let tmpl = "
    event %s(bytes result);
    function %s(%s) internal returns (bytes memory) {
        bytes memory payload = abi.encodeWithSignature(%s);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, \"DELEGATECALL to %s failed\");
        emit %s(result);
        return result;
    }
" ^^ "" in
  let evname = Printf.sprintf "%sCalled" (String.capitalize api.name) in
  let solargs =
    List.map api.args ~f:arg_to_solidity |>
    String.concat ~sep:", "
  in
  let selector =
    let args =
      List.map api.args ~f:(fun arg -> type_to_selector arg.typ) |>
      String.concat ~sep:","
    in
    Printf.sprintf "%s(%s)" api.name args
  in
  let selargs =
    List.map api.args ~f:(fun arg -> Printf.sprintf ",_%s" arg.name) |>
    String.concat
  in
  let sel_with_args = Printf.sprintf "\"%s\"%s" selector selargs in
  Printf.printf tmpl evname api.name solargs sel_with_args api.name evname

let () =
  let open Core in
  In_channel.read_lines "api.tsv" |>
  List.map ~f:parse_line |>
  List.iter ~f:print_templated_function

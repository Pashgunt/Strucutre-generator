# Json to Golang Structure

This Go program is a tool for processing JSON files and converting their contents into data structures corresponding to
Go types.

The program starts by checking the presence of the path to the JSON file in the command line arguments. If the path to
the file is not specified, the program will display a usage message and exit.

After successful conversion, the data is displayed on the screen to check the correctness of the parsing. The user can
also add additional data operations if necessary for a specific application.

It is important to note that the program is not adapted to the structure of a specific JSON file that it needs to
process, but is ready to process any file.

**In order for the program to work, you need to run the following command**

```Go
go run main.go -file ../../testdata/example.json -structure Test -mode file
```

at startup, you must pass 3 parameters
> 1. `-file ../path_to_file.json` - _path ti json file for generating Go structure_
>2. `-structure Default` - _name of golang generated structure_
>3. `-mode cli || file` - _mode response for get output format result, for example **cli** - return result in console
and **file** - return result in file -structure.go (`/cmd/structures/Y-m-d/unix/-structure.go`)_

---

### Examples
This example shows almost all the possibilities of converting JSON to Golang structure
Input json
```json
[
{
"_____id___": 1,
"_test_": "2",
"tags": [
"dolor",
"reprehenderit",
"aute",
"occaecat",
"reprehenderit",
"dolor",
"sint"
],
"_friends_": [
{
"id": 0,
"name": "Robbie Whitaker"
},
{
"id": 1,
"name": "Guerra Dillon"
},
{
"id": 2,
"name": "Noble Burris"
}
],
"global": "iam",
"user": {
"name": "Pavel"
}
}
]
```
Output Golang structure
```Go
type Test []struct {
User struct {
Name string `json:"name"`
} `json:"user"`
ID      float64       `json:"_____id___"`
Test    string        `json:"_test_"`
Tags    []interface{} `json:"tags"`
Friends []struct {
ID   float64 `json:"id"`
Name string  `json:"name"`
} `json:"_friends_"`
Global string `json:"global"`
}
```
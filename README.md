# ft-task

# Unit tests
choose which unit test you want to execute.
If you want to run the calculator repository tests you need to execute `make calculator-repo-tests`. </br>
If you want to run the facade tests you need to execute `make calculator-facade-tests`. </br>
If you want to run the calculator api tests you need to execute `make calculator-api-tests`. </br>
If you want to run all tests you need to execute `make all-tests`. </br>

You can get the coverage of the tests with `make get-coverage` - this will make a file that contains the rows that are covered by the tests. </br>
!!!IMPORTANT - In order to run `make get-coverage` there should not be any focused tests.

# calculator-cli
There are 2 available commands that can be run from the terminal.</br>
`evaluate` - It will give you the result of the expression</br>
`validate` - It will validate the expression</br>
Both commands need a flag -e/--expression which is the expression that will be evaluated or validated.
In order to check if it is running properly you should navigate to `cd ./cmd/calcularot-cli/` then run
`go run main.go [COMMAND] [FLAG]`. You can run `go run main.go -h/--help` for a help.

<img width="1019" alt="image" src="https://github.com/valioromanov/ft-task/assets/15982194/ac19c105-f9cb-4086-a89f-42aa4cd74681">


# calculator-api
This provides 3 endpoints:
 - \evaluate - POST
 - \validate - POST
 - \errors - GET
 The \evaluate endpoint receives an object
`{
    "expression": "What is <number> <operator> <number> [..<operator> <number>]"
}`
and returns HTTP 200OK with the result of the expression when everything is valid. You can use only these four operators:
  - plus
  - minus
  - multiplied by
  - divided by </br>
If the body is not valid HTTP 403 Bad Request will be returned with the proper message.</br>

The \validate endpoint receives the same object as above `{
    "expression": "What is <number> <operator> <number> [..<operator> <number>]"
}` </br> 
This endpoint returns HTTP 200 OK with </br>
`{
    "valid": true
}` when the body is valid. </br>
And HTTP 403 Bad Request with </br>
`{
    "valid": false,
    "reason": <reason>
}` when the expression is not valid</br>

The \errors endpoint returns all errors that occurred while the server is up. </br>
The response is HTTP 200 OK with </br>
`[
    {
        "expression": <expression>,
        "endpoint": <endpoint>,
        "frequency": <frequency>,
        "type": <type>
    }
    ...
]`
</br>
!!!IMPORTANT - If you want to start the server from localhost you should navigate to `cd ./cmd/calcurator-api/` and then `go run main.go`
It will start a server which listen on localhost:8080

<img width="1019" alt="image" src="https://github.com/valioromanov/ft-task/assets/15982194/3c045cb9-92b4-4c28-aad7-10cdd9f1790b">

<img width="1019" alt="image" src="https://github.com/valioromanov/ft-task/assets/15982194/c33de81d-b01a-48f0-bafe-d70bb11b4cff">

<img width="1019" alt="image" src="https://github.com/valioromanov/ft-task/assets/15982194/706ff9b9-839a-4400-8a7e-72f09c5fcfc5">

<img width="1007" alt="image" src="https://github.com/valioromanov/ft-task/assets/15982194/70859883-bb3c-4945-acd3-8b0155ab18bc">







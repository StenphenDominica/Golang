1. Tutorial: Getting started


2. Tutorial: Create a module

   1. Call your code from another module
       1. create a module
        go mod init example.com/hello
       2. user the go mode edit command to edit the example.com/hello module to redirect Go tools from its module path (where the module isn`t) to the local directory (where it is)
        go mod edit -replace example.com/greetings=../greetings
       3. run "go mod tidy" command to synchronize the example.com/hello module's dependencies
        go mod tidy
    2. Return and handle an error
    3. Add a test
       1. Ending a file's name with _test.go tells the go test command that this file contains test functions.
       2. Test function
          1. Implement test functions in the same package as the code you're testing.
          2. Create two test functions to test the greetings.Hello function. Test function names have the form TestName, where Name says something about the specific test. Also, test functions take a pointer to the testing package's testing.T type as a parameter. You use this parameter's methods for reporting and logging from your test.
          3. Implement two tests:
             1. TestHelloName calls the Hello function, passing a name value with which the function should be able to return a valid response message. If the call returns an error or an unexpected response message (one that doesn't include the name you passed in), you use the t parameter's Fatal method to print a message to the console and end execution.
             2. TestHelloEmpty calls the Hello function with an empty string. This test is designed to confirm that your error handling works. If the call returns a non-empty string or no error, you use the t parameter's Fatal method to print a message to the console and end execution.
       3. "go test" command


3. 
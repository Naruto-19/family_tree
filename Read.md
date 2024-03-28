To execute the code and set up locally, follow these steps:

1. **Download the code**: Copy the provided Go code into a file named `main.go`.

2. **Install Go**: If you haven't already, install Go from the official [Go website](https://golang.org/doc/install).

3. **Create a directory**: Create a directory for your project and place the `main.go` file inside it.

4. **Run the program**:
   - Open a terminal or command prompt.
   - Navigate to the directory containing the `main.go` file.
   - Run the command `go run main.go` to compile and execute the program.

5. **Interacting with the program**:
   - Once the program is running, it will prompt you with `family-tree>`.
   - You can enter commands like `add`, `connect`, `count`, `father`, or `exit`, followed by appropriate arguments.
   - Here's how you can use each command:
     - `add`: Add a person to the family tree. Provide first name, last name, and gender as arguments.
     - `connect`: Connect relationships between people. Provide first name, last name, relationship, and partner's first name, last name as arguments.
     - `count`: Count the number of sons, daughters, or wives of a person. Provide the count type (`sons`, `daughters`, or `wives`) and the person's first name, last name as arguments.
     - `father`: Get the father's name of a person. Provide the person's first name, last name as arguments.
     - `exit`: Terminate the program.

6. **Executing sample commands**:
   - Here are some sample commands you can try:
     - `add John Doe male` - This will add new male in family tree you have to write complete name and gender.
     - `add Jenny Doe female` -This will add new female in family tree you have to write complete name and gender.
     - `connect Jake Doe as son of Jane Doe` - This will connect Jake Doe as son of Jane doe
     - `connect Jenny Doe as wife of John Doe` - Connects Jenny Doe as wife of John Doe 
     - `connect Tiny Doe as daughter of John Doe` - Connects Tiny Doe as daughter of John Doe 
     - `count sons of John Doe` - This will provide count of sons but remember first add into family tree than connect as son and than you can execute this command 
     - `count wives of John Doe` - This will provide count of wives John Doe
     - `count daughters of John Doe` - This will provide count of daughters of John Doe  
     - `father of Jake Doe` - for father's name
     - `exit` - to exit the command line interface


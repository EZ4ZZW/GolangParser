# GolangParser
A project to parse golang source file and generate functions details.

This Project based on golang AST package.

Get the AST from the source file and visit the FuncDecl type to get the information of the function.

# To Run
modify the path to the file you want to parse in the main.go and the run the following command in terminal
```
go run main.go > output.txt
```
the output.txt contains the Details of the functions in the sourse file you added
# To do
- [x] get the function names from source file
- [x] add param parse
- [ ] design a more clear output format 
- [ ] can parse all function called in some function's body
- [x] can parse all source file in some directory
- [ ] analysis whether these functions are api
- [ ] check this api is called from outside or inside

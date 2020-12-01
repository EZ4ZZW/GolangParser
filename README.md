# GolangParser
A project to parse golang source file and generate functions details.
This Project based on golang AST package
Get the AST from the source file and visit the FuncDecl type to get the information of the function

# To Run
```
go run main.go > output.txt
```

# To do
- [x] get the function names from source file
- [ ] add param parse
- [ ] design a more clear output format 
- [ ] can parse all source file in some directory
- [ ] analysis whether these functions are api
- [ ] check this api is called from outside or inside

# golang-template-example

[![golang-template-example release (latest SemVer)](https://img.shields.io/github/v/release/FrancoPersonal/golang-template-example?sort=semver)](https://github.com/FrancoPersonal/golang-template-example/releases)
[![Go Reference](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://github.com/FrancoPersonal/golang-template-exampl)
[![Test Coverage](images/badge.svg)](cover.html)

# Directory
``` 
📦golang-template-example
 ┣ 📂output
 ┣ 📂src
 ┃ ┣ 📂filegenerator
 ┃ ┃ ┗ 📜generator.goadd
 ┃ ┗ 📂templates
 ┃ ┃ ┗ 📂initialproyect
 ┃ ┃ ┃ ┣ 📂files
 ┃ ┃ ┃ ┃ ┣ 📜makefile
 ┃ ┃ ┃ ┃ ┗ 📜README.md
 ┃ ┃ ┃ ┣ 📜main.go
 ┃ ┃ ┃ ┣ 📜prompt.txt
 ┃ ┃ ┃ ┗ 📜templatejson.json
 ┣ 📦templateBase
 ┃ ┣ 📂files
 ┣ ┣ 📜main.go
 ┣ ┣ 📜prompt.txt
 ┣ ┗ 📜templatejson.json
 ┣ 📜go.mod
 ┣ 📜makefile
 ┗ 📜README.md
```
# Templates

## initialproyect:
**Description :** create a basic solution for golang
**Usage :** 
``` batch
go run src/templates/initialproyect/main.go
```
![alt text](images/image.png)

the output path you can find the files generated

![alt text](images/image-1.png)


# Create a new template
``` batch
go run src/templates/templatebase/main.go TEMPLATENAME
```

## 
``` 
📦{{.TemplateName}}
 ┣ 📂files   # Template files folder
 ┣ 📜main.go # principal aplication
 ┗ 📜templatejson.json # variables
 ```

# MakeFile commands

- **clean :** clean files created
- **init :** initialize repository
- **test :** run unit test
- **coverage :** run unit test and make the HTML with the coverage
- **coverage :** run unit test and make the HTML with the coverage
- **showcoverage :** run unit test and show the HTML with the coverage
- **newtemplate :** example command to create a new template

example:
``` batch
makefile init
```
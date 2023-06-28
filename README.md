# GO - CEM304

Repository for studying and practicing GOLANG applied in a school management system project

## DOCKER - Containers
 - [X] Create dockerfile for Golang
 - [X] Create docker-compose.yaml
 - [X] Config DevContainer on VSCODE
 - [X] Config P10k and zsh history

## App
 - [X] Create repositório on Github
 - [X] Create entity accounts
 - [X] Test entity users
 - [X] Create service users and test
 - [X] Create persistence db and test
 - [X] Create cli with cobra-cli

## WebServer
 - [X] Create and setting command line for webserver
 - [X] Create adapter http
 - [X] Install and setting [chi](https://go-chi.io/#/) 
 - [X] Create jsonError to manipulate messages errors
 - [X] Create User handler server
   - [X] Create user
   - [X] Get user
   - [X] List users
   - [X] Change role
   - [X] Change password
   - [X] Change avatar
 - [X] Auth handler
   - [X] Login
   - [X] Refreshtoken
   - [X] Authentication and authorization
 - [X] Create Classroom handler server
   - [X] Create Classroom
   - [X] Get by id Classroom
   - [X] Get by name Classrooms
   - [X] List Classrooms
     - [X] Include Students
   - [X] Update ANNE Classroom
   - [X] Activate Classroom
   - [X] Desactivate Classroom
   - [X] Import file csv and registry Classrooms
 - [ ] Create Students handler server
   - [X] Create Students
   - [X] Get by id Students
   - [X] Get by name Students
   - [X] List Students
   - [ ] Update ANNE Students
   - [ ] Activate Students
   - [ ] Desactivate Students
   - [X] Import file csv and registry Classrooms
   - [X] Change Classroom
   - [X] Check Students on Classrooms in mass
 - [ ] Create Teacher handler server
   - [X] Register Teacher
   - [X] Attach Classroom and Subject
   - [X] Get by id teacher
     - [X] List subjects e classrooms
   - [X] Get by name teacher
   - [X] Get by cpf teacher
     - [ ] List subjects e classrooms
   - [ ] Update Teacher
 - [X] Create Subjects handler server
   - [X] Register Subjects
   - [X] Get by id Subjects
   - [X] List by Licenses
   - [X] Get by name Subjects

## Domains - Users
 - [X] Create entity
 - [X] Test entity classroom
 - [X] Create repository classroom
 - [X] Create persistence db and test
 - [X] Create table on postgresql
 - [X] Create cli with cobra-cli
  ## RefreshTokens
  - [X] Create entity
  - [X] Test entity classroom
  - [X] Create repository classroom
  - [X] Create persistence db and test
  - [X] Create table on postgresql

## Domains - Classrooms
 - [X] Create entity
 - [X] Test entity classroom
 - [X] Create repository classroom
 - [X] Create persistence db and test
 - [X] Create table on postgresql
 - [ ] Create cli with cobra-cli

## Domains - Students
 - [X] Create entity
 - [X] Test entity classroom
 - [X] Create repository classroom
 - [X] Create persistence db and test
 - [X] Create table on postgresql

## Domains - Parent
 - [X] Create entity
 - [X] Test entity classroom
 - [X] Create repository Parent
 - [X] Create persistence db and test
 - [X] Create table on postgresql

## Domains - Teacher
 - [X] Create entity
   - [X] Classroom relationship
   - [X] Subject relationship
 - [X] Test entity classroom
 - [X] Create repository Teacher
 - [ ] Create persistence db and test
   - [ ] Classroom relationship
   - [ ] Subject relationship
 - [X] Create table on postgresql
   - [X] Classroom relationship
   - [X] Subject relationship
## Domains - Subjects
 - [X] Create entity
   - [X] Classroom relationship
   - [X] Teacher realationship
 - [X] Test entity classroom
 - [X] Create repository Subject
 - [ ] Create persistence db and test
   - [ ] Classroom relationship
   - [ ] Teacher realationship
 - [X] Create table on postgresql
   - [X] Classroom relationship
   - [X] Teacher realationship

## PDF report generator
 - [X] Create a function for generator PDF
 - [X] Create student's list in pdf
 - [X] Diary Class in pdf

## Utils tools
 - [X] Read and save data of reports
   - [X] Create a function read reports students
   - [X] Create a function on repository to save data
   - [X] Create handler for import reports
   - [X] Update in mass students list
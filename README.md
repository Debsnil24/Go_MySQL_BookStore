# Go_MySQL_BookStore

This is a Golang project that uses a MySQL backend server & CRUD API to manage a BookStore


## Roadmap

### Project Structure
![ProjectStructure](https://github.com/user-attachments/assets/81ec7f7d-e41f-4df0-a16a-9dc3e87b5f51)

### Routes
![Routes](https://github.com/user-attachments/assets/074a8df7-a2bd-448a-bfe1-cf204297fa16)


## Testing Demo

#### GET ALL Books
![GET_ALL-postman](https://github.com/user-attachments/assets/79e117c2-1fd9-4658-beb0-5453cd90826d)

#### GET Book by ID
![GET_BY_ID-postman](https://github.com/user-attachments/assets/145ea1c2-09f1-43c6-b7d8-e74fc00e19dc)

#### Create new Book
![CREATE-postman](https://github.com/user-attachments/assets/b8cb0960-361e-4189-a2d9-0859409a5f56)

#### Delete a Book
![DELETE-postman](https://github.com/user-attachments/assets/e7c3b841-29cd-4add-af71-6a0e0d83c608)

#### Update a Book
![UPDATE-postman](https://github.com/user-attachments/assets/a1c6233f-04d2-447b-b4d6-cef54dbc1ef3)


## Run Locally

Clone the project

```bash
  git clone https://github.com/Debsnil24/Go_MySQL_BookStore.git
```

Go to the project directory

```bash
  cd {Directory}/Go_MySQL_BookStore
```

Install dependencies

```go
  go get "github.com/gorilla/mux"
```
```go
  go get "github.com/jinzhu/gorm"
```
```go
  go get "github.com/jinzhu/gorm/dialects/mysql"
```
Start the server

Browse to the directory of main.go file
```bash
  cd {Directory}/Go_MySQL_BookStore/cmd/main
```
Build the project
```go
  go build
```
Run the main.go file
```go
  go run main.go
```

## Acknowledgements

#### Check Out Akhil Sharma's Youtube Video for Project Tutorial
[GO And MYSQL - 2024 Project 🚀 💣 🔥 - Connect Go with Mysql / Build a Book Management System](https://youtu.be/1E_YycpCsXw?si=oLUV_plDs8y4i70s)

# REST API w/ GO
The aim of this project is to learn GoLang by building a simple REST API.

## 📚 Resources
This project used the videos below as base foundation:  
- [Building a GO API - No External Packages! by Alex Mux](https://www.youtube.com/watch?v=9bRMLKBbFMQ)
- [Complete Backend Engineering Course in Go](https://www.youtube.com/watch?v=h3fqD6IprIA)

This project also uses a Flat Project Structure to minimise complexity for now. Read more on GO Project Structures here:  
[Go - The Ultimate Folder Structure](https://dev.to/ayoubzulfiqar/go-the-ultimate-folder-structure-6gj)

## ⚙ Prerequisites
- **Go 1.22+**  
[Install Go Here!](https://golang.org/dl/)

- **GCC (C Compiler) for SQLite Driver**  
[Download GCC Here!](https://www.mingw-w64.org/downloads/#mingw-w64-builds)

  - Extract and add the `bin/` folder (e.g. `C:\mingw64\bin`) to your System Path.

## 💻 Setup & Usage
Follow these steps to get your development environment set up and operational:  

1. **Clone the Repository**  
    ```bash
    git clone https://github.com/dotping-me/learning-go-with-rest-api.git

    cd learning-go-with-rest-api
    ```

2. **Install GO dependencies**
   ```bash
   go mod tidy
   ```

3. **Run GO API**
   ```bash
   go run .
   ```

4. **Test Routes** (Use `curl`, Postman or any alternatives)
   1. Register a user through a `POST` request on `/user`
   2. Login using credentials on `/login`
      1. Copy generated token
   3. Test other routes
      1. Insert token as `Bearer Token` in `Auth` Headers 
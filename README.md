# REST API w/ GO
The aim of this project is to learn GoLang by building a simple REST API.

## 📚 Resources
This project used the video below as base foundation:  
[Building a GO API - No External Packages! by Alex Mux](https://www.youtube.com/watch?v=9bRMLKBbFMQ)

This project also uses a Flat Project Structure to minimise complexity for now. Read more on GO Project Structures here:  
[Go - The Ultimate Folder Structure](https://dev.to/ayoubzulfiqar/go-the-ultimate-folder-structure-6gj)

## ⚙ Prerequisites
- Go 1.22+   
[Install Go Here!](https://golang.org/dl/)

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

4. **Test Routes**
    ```bash
    curl -H "Authorization: Bearer 123" "http://localhost:8000/user/profile?id=USER0"
    ```
# 🚀 Tranzify API
A simple Morse translator API built on **Golang** and the **Gin** framework.
## 📬 API Endpoints
|Method|Endpoint|Description|
|------|--------|-----------|
| POST |`/api/textToMorse`|Convert text → Morse code|
| POST |`/api/morseToText`|Convert Morse code → text|
## ⚙️ Installation & Run
### 1. Clone the repository 
```
git clone https://github.com/anvarmx/tranzify-api.git
```
### 2. Go to the repository
```
cd tranzify-api
```
### 3. Install dependencies
```
go mod tidy
```
### 4. Run the server
```
go run main.go
```
Server will start by default on: 👉 http://localhost:3030
## 📨 Example request
```
curl -X POST http://localhost:3030/api/textToMorse \
  -H "Content-Type: application/json" \
  -d '{"input": "Hello World"}'
```
## 📩 Response
```
{
    "status": "success",
    "result": {
        "input": "Hello World",
        "output": ".... . .-.. .-.. --- / .-- --- .-. .-.. -..",
        "length": 11
    }
}
```

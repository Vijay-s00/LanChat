# 🗨️ LanChat App

# Table of Contents
1. [Desciption](#description)
2. [Features](#features)
3. [Tech Stack](#techstack)
4. [Architecture](#archi)
5. [Running the App](#run)


## 📌 Description <a name="description"></a>

This is a real-time group chat application built in Go that enables multiple users to communicate via a shared message channel. 
Messages are broadcast to all connected users using a fanout pattern through RabbitMQ, ensuring scalable and asynchronous communication. 
Each message is also persistently stored in a PostgreSQL database for historical access and analysis. 

## ✅ Features <a name="features"></a>
- Real-time messaging using RabbitMQ (fanout exchange)
- Persistent message storage in PostgreSQL
- Easy inspection and management with pgAdmin
- Docker Compose setup for rapid development
- Built with performance and scalability in mind using Go

## 🛠️ Tech Stack <a name="techstack"></a>
|   Component    | Technology |
|----------------|------------|
|    Backend     |     Go     |
| Message Broker |	RabbitMQ  |
|    Database	   | PostgreSQL |
|     DB GUI	   |   pgAdmin  |
|Containerization|    Docker  |

## 📐 Architecture <a name="archi"></a>
![lanchat](https://github.com/user-attachments/assets/23984d7b-1495-4703-9479-94e6d69e330c)

## 🚀 Running the Application <a name="run"></a>
The requirements for running this application are *Go* and *Docker*

1. Clone the repo and enter the folder

   `git clone https://github.com/JuanMartinCoder/LanChat/`
   
   `cd LanChat/`

3. You will need the dependencies used

   `go mod tidy`

4. Start the services with the script

   `./servicesUp start`

5. Connect to the chat with the main app

   `go run ./cmd/main.go`

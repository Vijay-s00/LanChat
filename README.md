# 🗨️ LanChat App
## 📌 Description

This is a real-time group chat application built in Go that enables multiple users to communicate via a shared message channel. 
Messages are broadcast to all connected users using a fanout pattern through RabbitMQ, ensuring scalable and asynchronous communication. 
Each message is also persistently stored in a PostgreSQL database for historical access and analysis. 

## ✅ Features
- Real-time messaging using RabbitMQ (fanout exchange)
- Persistent message storage in PostgreSQL
- Easy inspection and management with pgAdmin
- Docker Compose setup for rapid development
- Built with performance and scalability in mind using Go

## 🛠️ Tech Stack
|   Component    | Technology |
|----------------|------------|
|    Backend     |     Go     |
| Message Broker |	RabbitMQ  |
|    Database	   | PostgreSQL |
|     DB GUI	   |   pgAdmin  |
|Containerization|    Docker  |

## 📐 Architecture
![lanchat](https://github.com/user-attachments/assets/23984d7b-1495-4703-9479-94e6d69e330c)

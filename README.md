Tugas Akhir
Project ini adalah untuk memenuhi salah satu syarat untuk lulus kuliah, project ini juga merupakan microservices untuk logbook kegiatan kampus merdeka di sekitar program studi universitas pasundan

Logo

Roadmap
Additional browser support

Add more integrations

API Reference
Register
  POST /login
Parameter	Type	Description
email	string	Required.
password	string	Required.
Login
  POST /register
Parameter	Type	Description
nrp	string	Required.
name	string	Required.
email	string	Required.
password	string	Required.
Get All Users
  GET /allusers
Parameter	Type	Description
id	int	Required.
add(num1, num2)
Takes two numbers and returns the sum.

Installation
Install my-project with dockercompose

  docker-compose up -b
Migration Database

  make migrateup
  make migratedown
Tech Stack
Client: Swift, Vue, TailwindCSS

Server: Go

Run Locally
Clone the project

  git clone https://github.com/nuriansyah/backend-microservices.git
Go to the project directory

  cd backend-microservices
Start the server

  go run main.go
Authors
@nuriansyah
Badges
Add badges from somewhere like: shields.io

MIT License GPLv3 License AGPL License Go License


# Tugas Akhir 

Project ini adalah untuk memenuhi salah satu syarat untuk lulus kuliah,
project ini juga merupakan microservices untuk logbook kegiatan kampus merdeka
di sekitar program studi universitas pasundan


![Logo](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/th5xamgrr6se0x5ro4g6.png)


## Roadmap

- Additional browser support

- Add more integrations


## API Reference

#### Register

```http
  POST /login
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email` | `string` | **Required**. |
| `password` | `string` | **Required**. | 


#### Login

```http
  POST /register
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `nrp`      | `string` | **Required**. |
| `name` | `string` | **Required**. |
| `email` | `string` | **Required**. | 
| `password` | `string` | **Required**.| 

#### Get All Users

```http
  GET /allusers
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. |



#### add(num1, num2)

Takes two numbers and returns the sum.


## Installation

Install my-project with dockercompose

```bash
  docker-compose up -b
```
Migration Database
```bash
  make migrateup
```
```bash
  make migratedown
```
    
## Tech Stack

**Client:** Swift, Vue, TailwindCSS

**Server:** Go 


## Run Locally

Clone the project

```bash
  git clone https://github.com/nuriansyah/backend-microservices.git
```

Go to the project directory

```bash
  cd backend-microservices
```


Start the server

```bash
  go run main.go
```


## Authors

- [@nuriansyah](https://www.github.com/nuriansyah)

# Commit types
* `build`: Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)
* `ci`: Changes to CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)
* **`chore`: Changes which doesn't change source code or tests e.g. changes to the build process, auxiliary tools, libraries**
* `docs`: Documentation only changes
* **`feat`: A new feature**
* **`fix`: A bug fix**
* `perf`: A code change that improves performance
* `refactor`:  A code change that neither fixes a bug nor adds a feature
* `revert`: Revert something
* `style`: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
* `test`: Adding missing tests or correcting existing tests

## Badges

Add badges from somewhere like: [shields.io](https://shields.io/)

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![GPLv3 License](https://img.shields.io/badge/License-GPL%20v3-yellow.svg)](https://opensource.org/licenses/)
[![AGPL License](https://img.shields.io/badge/license-AGPL-blue.svg)](http://www.gnu.org/licenses/agpl-3.0)
[![Go License](https://img.shields.io/github/go-mod/go-version/nuriansyah/backend-microservices)](https://go.dev/LICENSE)


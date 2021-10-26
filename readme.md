# Backend Service

## Instalation

Set up variable environment on your device.
export PATH_CONF = "../config"
export FILE_CONF = "local.conf"

When you want to start this service, set up this aliases.
alias backend-run = "gin --port 8080 --appPort 8081 --path . --build ./app run main.go"
and running with alias backend-run

## Introduction

Merupakan service core **(Security)** yang bisa di duplikasi ke beberapa aplikasi.
Mulai dari register user, login dll. Service ini menggunakan bahasa Google Language (golang)
yang di support Framework [Gin Gonic](https://github.com/gin-gonic/gin).

### Guide Styling Code

Kita harus memiliki aturan dalam menerapkan code.
Agar code yang dihasilkan mudah dibaca, dipahami & di maintenance.
Beberapa hal yang harus diperhatikan adalah :

- **Design Pattern**
  - Service ini mengadopsi Structural Pattern (Decorator Pattern)
  - Terdiri dari beberapa layer
    - Delivery -> Layer untuk menerima request dari client / service dan juga bertugas melakukan validasi dari setiap request yang diterima.
    - Usecase -> Business layer, hanya untuk management proses atau flowchart dari api tersebut.
    - Entity -> Entitas layer, untuk mendapatkan data pada database atau data dari request api.
- **Penamaan function**
  - Pada endpoint handler harus sesuai dengan usecase yang digunakan (Get, Create, Edit & Remove) sebagai awalan penamaan function pada Usecase & Delivery.
  - Pada repository harus menggunakan awalan (Read, Insert, Update & Delete).
- **Penamaan variable** tidak boleh di singkat dan harus jelas merepresentasikan isi dari variable tersebut.
- **Handle error** pada Usecase layer harus menggunakan logs.
- **Clean code** menerapkan comment atau instruction pada setiap baris line.
- **Bahasa** untuk penamaan Function, Variable & Comment wajib menggunakan Bahasa Inggris.
- **Struct / Object** dibuat terpisah antara object table & object request response. Untuk object request response harus menggunakan akhiran (NamaFunctionRequest / NamaFunctionResponse).
- **Endpoint** harus menerapkan sebagai berikut
  - /api/v1/example -> Digunakan untuk method (POST & PUT)
  - /api/v1/examples -> Digunakan untuk method (GET) listing, request payload on header
  - /api/v1/example/:id -> Digunakan untuk method (GET & DELETE)
  - /api/v1/example?car=toyota&type=yaris -> Digunakan untuk method (GET), ketentuan berlaku

### Pemrograman Modular

Service ini masih menggunakan metode **monolith**, tetapi menggunakan teknik modular.
Agar setiap modul dapat di pisahkan dan kedepannya dapat di pecah menjadi beberapa **microservices**.
Setiap modul hanya dapat mengkonsumsi tablenya sendiri. Dilarang keras untuk melakukan join cross modul.
Pemrograman Modular di pisahkan menjadi beberapa modul, yaitu :

- **Modul Security**
  Security memiliki beberapa table antara lain :
  1. mst_user
  2. mst_user_status
  3. mst_group
  4. trx_register_user
  5. mst_register_user_status
  6. mst_service
  7. mst_api
  8. mst_group_access_api
  9. mst_device
  10. mst_module
  11. mst_group_access_module
  12. mst_forgot_password
  13. mst_forgot_password_status

### Unit Test

Setiap proses yang ada di entity wajib dibuatkan **Unit Test** (Integration Test)
agar memudahkan dalam proses pengembangan.

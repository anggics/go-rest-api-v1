# go-rest-api-v1
create, update, read, delete menggunakan rest api. digunakan untuk integrasi system to system


GET    /                         --> rest-api/interface/server.setupRouter.func1 (3 handlers)
GET    /customer/v1/getAllData   --> rest-api/interface/server.(*customerHandler).getAllDataCustomer.func1 (3 handlers)
POST   /customer/v1/add          --> rest-api/interface/server.(*customerHandler).customerSaveData.func1 (3 handlers)
POST   /customer/v1/getById      --> rest-api/interface/server.(*customerHandler).getDataCustomerById.func1 (3 handlers)
PUT    /customer/v1/update       --> rest-api/interface/server.(*customerHandler).customerUpdate.func1 (3 handlers)
DELETE /customer/v1/delete       --> rest-api/interface/server.(*customerHandler).customerDelete.func1 (3 handlers)
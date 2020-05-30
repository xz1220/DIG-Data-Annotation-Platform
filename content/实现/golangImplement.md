# Impplement Based on Golang
---
In this section, We will introduce labelProject implemented by golang ,GIn (Http router framework), Gorm (ORM framework). The content was orginized as Dao layer(define the data structs according to dataTable of database labelproject in MYSQL),Repository layer (expose and implement the interface of CURD , supporting to edit the database with the structs defined), Service layer, Controller layer.

## Dao layer
----
We created four files ,contained four categories of structs, Image(such as Image informations and label datas), Task, User, Video. Structs exectly brought into correspondence with data Table in Mysql.Also, we add notes for almost every struct member to ensure GORM (ORM framework we used in the project) works.
 
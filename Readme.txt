REST_CRUD_API_JWT_SWAGGER

/Basically this proj is build using Go, here you can perform crud operations(using ur Mysql database), user authentication/generating tokens and you can use Swagger for a UI/

Create A CrudRestApi Folder.
Then create a models folder inside the CrudRestApi folder.
Inside the models folder paste the client.go, projects.go, logintokens.go and users.go
Now run the Vscode and open the CrudRestApi folder in it.
Now create an executable file by running go mod eg."filename.exe" and build the project using "go build" and then ".\restapi.exe" //in my case i have created a restapi.exe file using go mod//.
You can now use Insomnia or the Postman server to pass json inputs to do the Crud operations.
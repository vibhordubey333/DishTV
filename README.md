# DishTV
CLI based implementation of , user using/subscribing services provided by D2H . Written in Golang. 

 #### Description:
 
 1. To check if go is installed type : go version
 
 2. Starting point of program is in "main" directory.
 
 3. To run program , fire command in terminal inside main dir : ``` go run main.go ```    
 
 4. "menuservice" is responsible for displaying menu.
 
 5. "recharge" is self-explanatory , allows user to recharge balance.
 
 6. "subscriptioncompute" , computes subsriptiion charge , discount to be given. It also provides       listing of channels and services which user can subscribe to.
 
 7. Testcases can be found in directory "testcases". To generate use command : ``` gotests -all menu_service.go > menu_service_test.go ```
 
 8. Mockery is used to generate the dummy data : ``` mockery -name= Interface_Name ```
 
 9. To test case coverage : ```go test -cover ```
 
 

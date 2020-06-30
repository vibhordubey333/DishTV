# DishTV
CLI based implementation of , user using/subscribing services provided by D2H . Written in Golang. 

 #### Description:
 
 1. To check if go is installed type : go version
 
 2. Starting point of program is in "main" directory.
 
 3. "menuservice" is responsible for displaying menu.
 
 4. "recharge" is self-explanatory , allows user to recharge balance.
 
 5. "subscriptioncompute" , computes subsriptiion charge , discount to be given. It also provides       listing of channels and services which user can subscribe to.
 
 6. Testcases can be found in directory "testcases". To generate use command : ``` gotests -all menu_service.go > menu_service_test.go ```
 
 7. Mockery is used to generate the dummy data : ``` mockery -name= Interface_Name ```
 
 8. To test case coverage : ```go test -cover ```
 
 

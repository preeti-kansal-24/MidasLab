# MidasLab


STEPS to run the project

Initial Setup
1. Go to terminal -> run make file using `make up` command.
   You should see the the following output -  
 
        ✔ Network midaslab_default        Created                                                                                                                                                              0.0s
        ✔ Container midaslab-zookeeper-1  Started                                                                                                                                                                 0.7s
        ✔ Container midaslab-postgres-1   Started                                                                                                                                                                 0.8s
        ✔ Container midaslab-kafka-1      Started
2. Next step is to create configuration to run main.go. 
We need to provide two mandatory fields in flags to pass as command line
arguments namely ->
   1.      --twilio-sid= and --twilio-auth= and twilio-from-phone-number
   
3. Now run the configuration which runs main.go, which will finally setup 
database, kafka producer and consumer, repository layer, service layer, 
grpc handler and run the service at :50001 port. 

4. To test the grpc's you can simply use postman -> grpc feature to uplaod the 
proto file and use the 0.0.0.0:50001 host and port to hit the grpc. Select
the grpc you wish to run and generate an example msg from postman and run it.

5. Following grpc's are implemented
      
         1. SignUpWithPhoneNumber : this takes profile as input like name, email and phonenumber 
         and saves it to db. The profile then generated is pushed to a topic where the consumer sits and
         generates a 6 digit otp. It then tries to send the otp to the given user's phone number 
         using twilio api. The auth id and sid and from ph no of twilio needs to be 
         provided as input as flags. The generated otp is also saved to db so that
         when user uses "VerifyNumber" api, the otp can be matched.
   
         2. VerifyNumber - Takes phone number and otp as input and returns success/failure
         based on the otp matched
   
         3. Generate otp - Takes phone number as input and generates a 6 digit otp
         for the user, sends an sms on the phone number using twilio and saves otp to db
   
         4. Login - This takes phone number and otp as input from the above step and  
         sends back success or failure based on otp provided
   
         5. GetProfile - Once the login is successful, user can get their profile using this api
         where phone number is given as input and o/p is the profile
   


      
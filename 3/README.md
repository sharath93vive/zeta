# README FOR GO-REDIS-API IMPLEMENTATION

implementation : 
      http://oldblog.antirez.com/post/autocomplete-with-redis.html

      Achieved using the methods suggested in the above link

1. Steps to reproduce:

   -> Docker image  [In case facing trouble building image locally]
           ImageName : sharath94z/goredisapi:alpine 

   -> Testing API locally without docker
          
          If have added two binary files for different platform [built in MacOS]
             1. goredis_darwin - macOS
             2. goredis_linux_amd64_link - Linux/ container
 
          If running in Linux with provided binary
             $./goredis_linux_amd64_link

          To build locally and run the generated binary [Requied to import required lib]
          $ go build goredis.go
          $./goredis
          
          Command to test API in postman or curl:
          1. curl http://localhost:32776/add_word/word\=foo   
             Success
          2. curl http://localhost:32776/add_word/word\=foob
             Success 
          3. curl http://localhost:32776/add_word/word\=foobr
             Success
          4. curl http://localhost:32776/autocomplete/query\=foo
             [foo foob foobar foobr]

          Note: Program will fail if the auto-search is done on the string which is not present in redis
          ERROR: 2020/04/25 13:05:17 Unable to query ZRANK redis : redigo: nil returned
    
    -> Testing API using Docker image

       Required images:
          1. sharath94z/goredisapi:alpine
          2. redis
       
       Pull Docker image : $docker pull sharath94z/goredisapi:apline
       url : https://hub.docker.com/r/sharath94z/goredisapi
       
       port exposed: 8080

       Steps : 
           
           1. docker run --name redis -p 6379:6379 -d redis
           2. docker run --name goredis --link redis:redis -P -d sharath94z/goredisapi:alpine
           3. Test with the curl command with host port number

       Note: Dockerfile is also provided to build the image locally [Binary files need to be present in the current directory ]


# filestore

#To start Server,
    go run main.go

#Server docker image 
    sreethecool2/filestore:latest

#To start client
#calling client.GetClient with url of the server will return client instance
#calling client.Start will start the filestore cli
#please find the sample client.


Known design Issues:
#uploaded file is stored in upload/ which is inside docker containers. 
#so on restart of docker upload data will be lost.
#Better way will be to store in file storage server like s3 or mounting volume in server and will be shared by image


#!/bin/bash
echo "gen swagger and run api , option -help"

if [ $1 == "-help" ];
then
echo "********************** Option ***********************"
echo "*    ./start.sh [arguments]                         *"
echo "*The arguments are :                                *"
echo "*    -air     : run hot reload                      *"
echo "*    -api     : run only api                        *"
echo "*    -swagger : gen lnly api document               *"
echo "*    -all     : gen swagger and run api             *"
echo "*    -emtry   : gen swagger and run api             *"
echo "*    -test    : test code Quality                   *"
echo "*****************************************************"

elif [ $1 == "-api" ];
then
echo "run only api"
go run cmd/api/main.go

elif [ $1 == "-swagger" ];
then
echo "run only api"
swag init -g cmd/api/main.go

elif [ $1 == "-all" ];
then
echo "run swagger and api"
swag init -g cmd/api/main.go
go run cmd/api/main.go

elif [ $1 == "-air" ];
then
echo "run hot reload api"
swag init -g cmd/api/main.go
air -c scripts/air/air.toml

elif [ $1 == "-test" ];
then
echo "run test code Quality"
go test ./... -coverprofile=coverage.out && sonar-scanner

else
swag init -g cmd/api/main.go
go run cmd/api/main.go
fi




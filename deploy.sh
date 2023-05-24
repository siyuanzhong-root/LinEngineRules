#!/bin/sh

#NECCESSARY DEFINE ROOT_ENVIROMENT!!!!!!!
CODE_PATH="/userdata/engine_build"
BUILD_ENV_IMAGE_NAME="build_env:1.0"
BUILD_ENV_CONTAINER_NAME="build_env"
MAIN_FUNCTION_PATH="server/main.go"
GOLANG_BUILD_NAME="engineserver1.0"
APP_IMAGE_NAME="engine_server:1.0"
APP_CONTAINER_NAME="engine1.0"


cd $CODE_PATH

docker rm -f $(docker ps  |  grep "engine*"  | awk '{print $1}')
docker rmi -f $(docker images  |  grep "engine*"  | awk '{print $3}')

echo "BUILD_CONTAINER IS $(docker ps -aqf "name=$BUILD_ENV_CONTAINER_NAME")"


if  test -z "$(docker ps -aqf "name=$BUILD_ENV_CONTAINER_NAME")"
then
echo "-------STEP1:$BUILD_ENV_IMAGE_NAME not exist,begin build golang_env--------"
docker build -t $BUILD_ENV_IMAGE_NAME .
docker run  -it -d --name $BUILD_ENV_CONTAINER_NAME -v $CODE_PATH:/siyuanzhong/gobuild $BUILD_ENV_IMAGE_NAME
else
echo "---------STEP1:$BUILD_ENV_IMAGE_NAME EXIST---------"
fi 
echo "---------STEP1:$BUILD_ENV_IMAGE_NAME FINISH---------"


echo "---------STEP2:BUILD GOALANG BINARY PROCESS---------"
docker start $BUILD_ENV_CONTAINER_NAME
docker exec -it $BUILD_ENV_CONTAINER_NAME /bin/sh -c "go build -o $GOLANG_BUILD_NAME $MAIN_FUNCTION_PATH &&exit"
echo "---------STEP2:BUILD GOALANG BINARY FINISH---------"

echo "---------STEP3:BUILD $APP_IMAGE_NAME PROCESS---------"
docker build -f Dockerfile_Engine -t $APP_IMAGE_NAME .
echo "---------STEP3:BUILD $APP_IMAGE_NAME FINISH----------"


echo "---------STEP4:RUN $APP_CONTAINER_NAME PROCESS---------"
docker run -d --name $APP_CONTAINER_NAME -v $CODE_PATH/database:/database -v /etc/localtime:/etc/localtime --restart=always -p 20080:10080 $APP_IMAGE_NAME
echo "---------STEP4:RUN $APP_CONTAINER_NAME FINISH---------"
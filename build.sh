docker build -t mirrors2/liketriple:latest .

docker rmi $(docker images -q -f dangling=true)

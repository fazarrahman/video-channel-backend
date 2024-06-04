# video-channel-backend

# Install docker engine
Please follow the instructions from this link : https://docs.docker.com/engine/install/

# Install postgredb on docker
Run this command
sudo docker run --name postgres -e POSTGRES_USER=admin -e POSTGRES_PASSWORD='Pass123$$' -e POSTGRES_DB=video_db -p 5432:5432 -d postgres:13
# video-channel-backend

# Install docker engine
Please follow the instructions from this link : https://docs.docker.com/engine/install/

# Run these command on the terminal
sudo docker network create video
sudo docker compose -f docker-compose.db.yml up -d
sudo docker compose -f docker-compose.backend.yml up --build -d

Backend will run at localhost:4000 and postgre db at port 5432
The postgresdb credentials :
Host: localhost
Port: 5432
Username: admin
Password: Pass123
Database: video_db

# Endpoints
## User auth endpoints
1. Create user
Curl :
curl --location 'http://localhost:4000/api/users/register' \
--header 'Content-Type: application/json' \
--data-raw ' { 
    "username": "example", 
    "email": "example@example.com", 
    "password": "password" 
}'

2. Sign In 
Curl :
curl --location 'http://localhost:4000/api/users/signin' \
--header 'Content-Type: application/json' \
--data-raw '{ 
    "email": "example@example.com", 
    "password": "password" 
}'

## Film Endpoints
1. Get All Films :
Curl :
curl --location 'http://localhost:4000/api/films'

2. Create Films :
Curl :
curl --location 'http://localhost:4000/api/films'
--header 'Content-Type: application/json'
--data '{
"title": "Example Title",
"description": "Example Description",
"image_thumbnail": "http://example.com/image.jpg"
}'

3. Get Films By Id :
Curl :
curl --location 'http://localhost:4000/api/films/5'

4. Update Films :
Curl :
curl --location --request PUT 'http://localhost:4000/api/films/3'
--header 'Content-Type: application/json'
--data '{
"title": "Updated Title",
"description": "Updated Description",
"image_thumbnail": "http://example.com/updated_image.jpg"
}'

5. Delete Films :
Curl :
curl --location --request DELETE 'http://localhost:4000/api/films/5'
--header 'Content-Type: application/json'
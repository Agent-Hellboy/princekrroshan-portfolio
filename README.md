# my-portfolio

sudo docker-compose down --remove-orphans  && sudo docker-compose up -d --force-recreate


# testing 
## client 

docker build -t myreactapp .
docker run -d -p 3000:3000 myreactapp

same for server 

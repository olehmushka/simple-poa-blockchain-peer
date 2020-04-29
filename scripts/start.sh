docker-compose build
docker-compose up -d --remove-orphans --force-recreate
docker-compose logs --tail=70 -f

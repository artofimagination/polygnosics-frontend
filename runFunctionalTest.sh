# pip3 install -r tests/requirements.txt
docker-compose --file docker-compose-test.yml down
docker stop $(docker ps -aq)
docker rm $(docker ps -aq)
docker system prune -f
docker-compose --file docker-compose-test.yml up --build --force-recreate -d frontend
python3 -m pytest -v tests/functional
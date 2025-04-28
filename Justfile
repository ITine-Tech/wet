start-redis:
    redis-server

start-redis-cli:
    redis-cli

get-redis:
    redis-cli keys '*'

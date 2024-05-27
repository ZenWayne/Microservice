### cluster call
`redis-cli -a 123 --cluster call 172.20.0.1:6379 zrange LB_BlackrockCaverns 0 2 withscores`

### zset clear 
`zremrangebyrank LB_BlackrockCaverns 0 2`
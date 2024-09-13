# Rate Limiter

Rate Limiter is a system which limits the number of requests a user can make in a given time window. It is used to protect the system from abuse and to make sure that the system is not overwhelmed by the requests.

## Requirements
### Functional Requirements
1. **Rate Limiting**: Limit the number of requests a user can make in a given time window.
2. **Handle Multiple Rate Limit Logics**: The system should be able to handle multiple rate limit logics.

### Non Functional Requirements
1. **Scalable**: The system should be able to handle a large number of requests.
2. **Low Latency**: The system should have low latency.
3. **Fault Tolerant**: The system should be fault-tolerant.
4. **Consistent**: The system should be consistent.

## Capacity Estimation
- 1 Billion Users
- 20 Services to be rate-limited
- User_id = 8 bytes & integer to store the count of requests = 4 bytes
- The rate limit logic can be implemented on user_id or better ip_address

## API Design
```
boolean isRequestAllowed(String user_id, String service_id, Date request_time);
```

## High Level Design
// TODO

## Additional Articles:

#### Wobble Bug in Github Rate Limiter [link](https://www.youtube.com/watch?v=LAdUhenKDJ8&t=288s&ab_channel=ArpitBhayani)
Github Used memcache Initally and used a single cluster Memcache server. This cache stores cache for all there use case and also for rate limiting purpose. 
**The problem came when the data increase and the cache evicting the rate limiting keys. This caused the rate limiter to fail.**
</br>
**Solution:**
1. They moved to a redis as the distributed cache system. 
2. It has a Simple Sharding and replication Setup. It can also have multiple shards to increase the reads (It is important as Rate Limiter is read heavy system).
3. Redis Supports Lua scripts for rate limiting logics and can handle it atomically and efficiently.

**Additional Points:**
After Migrating to Redis, they faced wobble bug. In which the rate limiter is sending the TTL of the key and it used to wobble.
The wobble is due to precision error in the time. The lua script is sending the TTL to the rate limiter and then it add time.Now to get the exipry time. But due to network latency, this used to wobble.
</br>
This was fixed by using some extra space in key to store reset_at time in the key , which is the absolute time the key should expire. This stops the wobble.



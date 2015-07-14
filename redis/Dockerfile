# Redis is an open-source, networked, in-memory, key-value data store 
# with optional durability. It is written in ANSI C. The development of 
# Redis has been sponsored by Pivotal since May 2013; before that, it was 
# sponsored by VMware. According to the monthly ranking by DB-Engines.com, 
# Redis is the most popular key-value store. The name Redis means 
# REmote DIctionary Server.
FROM redis:3.0.2

COPY redis.conf /usr/local/etc/redis/redis.conf
CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]

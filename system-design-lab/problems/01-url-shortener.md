Problem Statement
Given a url "https://long-url-to-shorten-code.com/do-something" return a shortened url "https://magix-short.com/do".

Design a url shortening service that allows users to create short urls. Redirect users to the long url when given the short-code.

1. ***Requirements***

    *Functional Requirements*
    - create short url
    - redirect short url
    - Url expiry
    - Custom alias
    
    *Non-functional Requirements*
    - High Availability
    - Low latency
    - Scalibility
    - Durability

    **Out of Scope**
    - Analytics
    - user management

2. Traffic 
    - 10:1 read:write ratio
    - 100M reads/ day which is approx 1157 average qps. 11570 peak qps.
    - 10M writes/day whic is 115 average qps, 1150 peak qps

3. Data
    - short-code type string (max len 10)
    - long url varchar 50-100
    - create_at time
    - expiry_at time

4. Flows
 
   Phase 1: 
    read: client ---(short code)--> query db for the short code ---shortcode--->database [SQL]
    write: client ---(long url)--> server (generates the hash code, shorten. Encode to base62) ---data model--->database [SQL]
   
   Phase 2: With increase in request may be peak reading data from database would increase the request latency.
            We can implement a cache aside strategy with redis (database is the hotspot).
    read: client ---(short-code)--> queries cache for (short-code)-->read from db(if cache miss) --->database--> add to cache -----> redirect to long url[SQL]
    write: client ---(long url)--> server (generates the hash code. Encode to base62) ---data model--->database-->save to   redis cache [SQL], 

    Phase 3
        With the flow in Phase 2 we might have solved the read latency to some extent. We have 10M/ day which 117 writes qps. With hashing is quite possible using base62 we have short code which can be generated for two urls.
        Approach 1:
           Generate the short code as (hash code + unique string based on timestamp).
           Check if the string is already mapped to any other url.
        
        cons checking always whether collision happening and keep generating.


6. How to scale?
    - We can have read/write replicas for database for read requests. 
    - Add multiple servers with load balancing.
        

5. API Design

    - POST /shorten {url:"long-url","short-code":"short-url"}
    - GET /{short-code}

What I learnt?
- Base64 has character =, /, ?,+ which have special meaning in urls. so use base64url or base62
- Snowflake string generation
- base62 encoding (using character length) starting 000000, 000001, ..., up to zzzzzzz,

Just prepopulate your database with all possible short-codes like 000000, 000001, ..., up to zzzzzzz.
When someone shortens a URL, you simply pick the next available code. No hashing, no collisions, no magic numbers.

character(n) | Formula | Result | Human Scale 
4 | 62 power 4 | 14,776,336 | ~14.7 million
5 | 62 power 5 | 916,132,832 | ~916 million
6 | 62 power 6 | 56,800,225,584 | ~56 billion
7 | 62 power 7 | 3,521,614,606,208 | ~3.5 trillion
8 | 62 power 8 | 218,340,105,584,896 | ~218 trillion



Twitter’s Snowflake algorithm the famous 64-bit ID generator. It produces unique, time-sortable IDs using a mix of timestamp, machine ID, and a local sequence number.
It’s brilliant. It’s scalable. It’s collision-free.
And when you Base62-encode those 64 bits, you get something around 11 characters long. Not bad, but still not the elegant 6–8-character short URLs we see in the wild.


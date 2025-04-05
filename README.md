Nice little proxy server that caches responses in a map, Fun project to teach me more about go. Run the server on some port and add the origin.  

-Lets say your port is 8080 and origin is http://origin.com/. Now if you send requests to this server at some path /path it will forward that request to http://origin.com/path and cache it if hasn't been done already. If a request has been made again here then it returns the cached response instead of hitting the origin again. 

To run the proxy, enter the following command:
```bash
caching-proxy --port <port> --origin https://your-origin.com/
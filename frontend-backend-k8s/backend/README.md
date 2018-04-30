# Mini Hello from Kubernetes App and API

A small app showing various info about the server. 

Currently, there are two endpoints in the API:

### "/health_check" endpoint:

Gives info about the server and the redis connection:

```
$> curl $(SERVICE_IP)/health_check
{"alive": true, "redis_conn": "good"}
```

The return value is a JSON dictionary with `alive` and `redis_conn` keys.

### "/" endpoint:

Prints the hostname of the machine, version of the program, Total # of requests on
this pod, Total # of requests in the whole system, App Uptime and Log Time.
```
$> curl localhost:8080
Hello from Kubernetes! Running on Mac-mini.local | version: 0.3
Total number of requests to this pod:1
Total number of requests in all system:6
App Uptime: 2.163356829s
Log Time: 2018-04-30 06:06:21.856840565 +0200 CEST m=+2.165491052
```
Number of requests are counted by hits on only this (`/`) endpoint.

#### TODO List
- [ ] Write Tests
- [ ] Add more endpoints.
- [ ] "/" endpoint doesn't return JSON, it breaks the API structure. Fix it.
- [ ] Code Cleaning and Restructure

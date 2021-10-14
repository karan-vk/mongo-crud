# GO + MONGODB

## Environment

```
go-version = 1.17.2
```
```env
MONGO_URI=mongodb://admin:OuOp66RvQpGma6u0@SG-sud-47257.servers.mongodirector.com:27017/admin
DB=sud
PORT=3000
APP_ENV=development
JWT_SECRET=dfsodhfjksh35453d45jk4534f5hs45d3l45fhadkhfadhu345234
```

### Deployed APP
```https://rocky-island-43692.herokuapp.com```

may not be available at all times 


## Create User
```[POST] /api/user ```
### Body
```json
{
	"name": "Test",
	"dob": "25-10-2001",
	"address":"abc street",
	"description":"summa",
	"createdAt":"2021-10-14T11:28:15.997Z"
	
}
```
### Response
```json
{
  "data": {
    "InsertedID": "61681484a8c956597ac3936d"
  },
  "message": "User inserted successfully",
  "success": true
}
```
<hr>

## Get token 

``` /api/auth ```
### Requst Body
Usually I would send a username and password and use bcrypt to verify the password
```json
{
	"id":"61681484a8c956597ac3936d"
}
```

### Response

```json
{
  "success": true,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiVGVzdCJdLCJleHAiOjE2MzQyOTc0MjgsImlhdCI6MTYzNDIxMTAyOCwianRpIjoiNjE2ODE0ODRhOGM5NTY1OTdhYzM5MzZkIn0.6bCMPGAiAdJS5oHwH8H3t5YPAppE_Z9flJ0GJtrFJgE",
  "user": {
    "_id": "61681484a8c956597ac3936d",
    "name": "Test",
    "dob": "25-10-2001",
    "address": "abc street",
    "description": "summa",
    "createdAt": "2021-10-14T11:29:07.831Z"
  }
}
```
<hr>

## Get All users

```[GET] /api/user ``` 
<hr>

## Get One user

```[GET] /api/user/:id ``` 
<hr>

## Update User
``` [PUT] /api/user/:id ```

### Body
```json
{
	"name": "Test",
	"dob": "25-10-2001",
	"address":"abc street",
	"description":"summa",
	"createdAt":"2021-10-14T11:28:15.997Z"
	
}
```
### Response
```json
{
  "message": "User updated successfully",
  "success": true
}
```
<hr>

## Delete User
``` [DELETE] /api/user/:id ```



## Geo Friend Find 
The Implementation would go like the following: -
Build an API so that the Frontend sends the current location of that device 
- Step 1 <br>
I set it up in Redis Using the GEOADD functionality
Now we have a system for querying users’ location 
- Step 2 <br>
There are 2 approaches <br>
Approach 1 :  Calculate the distance between my location and my friends location and if the distance is below a certain radius like 4km I return it back <br>
**GEODIST x N friends || complexity: O( N log(M))** <br>
Approach 2 : Find everyone within radius 4km(assumption) and filter out only my friends <br>
**GEORADIUS + N friends || complexity:  O(N(M+log(P)))** <br>
M is the number of elements inside the bounding box of the circular area <br>
We can also use neo4j to do the same kind of queries but the execution speed will be slower when compared to Redis
## Logging 
I had a logging system built with Rust + Redis where all the data will be pushed to Redis Streams and then asynchronously pushed to Timescale DB and then visualized in Grafana.
Mostly logging can be like an trackable event like user.created, user.updated and many more 
For my case I’ve done it as an HTTP API but if I were to reimplement it would be using GRCP/WebSocket/Direct function calls.
Friend and Follow implementation that I have done for my personal project                                             




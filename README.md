# Project Setup Guide

## 1. Clone the Repository
To run the application, you need to download the source code from the GitHub repository:

## 2. Start the Application with Docker Compose
Run the following command to build and start the application:
```sh
docker-compose up --build
```

## 3. Send a Test Request
You can use `curl` to send a test request to the API:
```sh
curl --location 'http://localhost:8080/api/v1/card/notification' \
--header 'Content-Type: application/json' \
--data '{
  "orderType": "SendOtp",
  "sessionId": "500cf308-e666-4639-aa9f-f6376015d1b4",
  "card": "4433**1409",
  "eventDate": "2023-04-06 22:52:34.930150 +00:00",
  "websiteUrl": "https://somon.tj"
}'
```

## 4. Check Worker Logs
To monitor the worker logs, run the following command:
```sh
docker-compose logs --follow
```

This will display real-time logs of the running services.




Backend developer test
=================

## Summary
Write notification service to notify user for card activities on his card. You need to write HTTP server which accepts events in JSON format via POST method and stores them in some sort of storage
(could store either in db or some basic memory storage). You also need to write worker/job which will notify those events to client. Notification can be mocked just by printing them to terminal.

- Try to use only standard library
- Provide instructions on how to build and run the code
- Keep your solution simple
- Add your solution to github and share the link with us
- We would like to see some tests too

### Sample Events
```javascript
{
  "orderType": "Purchase",
  "sessionId": "29827525-06c9-4b1e-9d9b-7c4584e82f56",
  "card": "4433**1409",
  "eventDate": "2023-01-04 13:44:52.835626 +00:00",
  "websiteUrl": "https://amazon.com"
},

{
  "orderType": "CardVerify",
  "sessionId": "500cf308-e666-4639-aa9f-f6376015d1b4",
  "card": "4433**1409",
  "eventDate": "2023-04-07 05:29:54.362216 +00:00",
  "websiteUrl": "https://adidas.com"
},

{
  "orderType": "SendOtp",
  "sessionId": "500cf308-e666-4639-aa9f-f6376015d1b4", 
  "card": "4433**1409", 
  "eventDate": "2023-04-06 22:52:34.930150 +00:00",
  "websiteUrl": "https://somon.tj"
}
```
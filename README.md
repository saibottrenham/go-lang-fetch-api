### Project Assignment
You have been tasked with building a Go app that retrieves input in structured JSON (the format of which is to your choosing), validates the input, logs the input to STDOUT, and sends a request (the format of which is to your choosing) to a HTTP based API, with the input being part of the request message.

### Guidance
* Provide a detailed walkthrough of your thought process to understand any architectural & other considerations, that help you arrive at a proposed solution
* Provide a descriptive overview of the the solution
* Document challenges, considerations, requirements

### Notes
* We don’t expect you to provide a compiled app, however, the codebase must be compilable
* Feel free to provide the codebase as a file, or as a Git repository
* You can also provide a Go Playground link, if you prefer
* Feel free to provide diagrams, and other material, as relevant to support your response


This will execute the sender endpoint which will log the payload to STDOUT
and send a subsequent request to the receiver endpoint which will return the 
payload `t.From.Name`

* Provide a detailed walkthrough of your thought process to understand any architectural & other considerations, that help you arrive at a proposed solution

I'm not familiar with Go Lang (come from a python backgound) so I did not really know what I was doing. My goal was to create an api that is able to call itself on a different endpoint which seemed sufficient for this task.

* Provide a descriptive overview of the the solution

### How to Run
```
go run app.go
```

then execute
```
curl -X POST -H "Content-Type: application/json" -d '{
    "to" : [
 	    {"email": "tobiasmahnert@gmail.com"},
 	    {"email": "tobiasmahnert@web.de"}
    ],
    "from" : {
        "email_address_id": 16426,
        "name": "tobias"
    },
    "schedule": 1624166184,
    "subject": "test subject",
    "body": "test body"
}' http://localhost:8082/sender
```
  
* Document challenges, considerations, requirements

My main challenge was that I did not ceate anything in Go Lang yet therefore I might not do everything in a Go lang fashion. I tried to incoporate universal programing paradigms by seperating the sender and receiver functionality into respective functions. 

Both sender and receiver have been tested with unittests inside app_test.go. For the tests I mocked responses and simulated http request interactions. 
My apologies in advance when this whole Repo is a big dumpster fire, happy to learn how to do it in a better way... 
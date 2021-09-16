# gin-test
A dummy application built using the Gin framework in Go

- It's designed to respond to customers of different brands based on the reciever's phone number. If they ask for the nearest store, NearestLoc() responds to that. // supported, but no actual implementation.

- Recieves a JSON API request body and accordingly returns a JSON text response and/or location response.

- main.go       - main file, server runs here
                - currently only has a GET endpoint

- controller.go - interface level, handles services
                - decides whether to call a text or location response based on JSON binding
              
- service.go    - implementation level, real work happens here
                - lookup sender's number on database and accordingly send a text decided by Desision() based on input text (RespondText())
                - send nearest store location if needed (RespondLoc())

- data.go       - contains structures converted to JSON formats for the API.

- logger.go     - just a logger.


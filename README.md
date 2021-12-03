# What is this
- Simple & Primitive multi client communication system.
   - e.g. chat system
- for learning   

https://user-images.githubusercontent.com/836774/144528442-199c889b-32cc-4185-a5d4-90eceaf2c7bb.mp4

# Supported
- Broadcast message
- Unicast message

# Not Supported
- Multi Room
- Mutex Control
- Error Handlling
- Ping/Pong based Alive monitoring 

# Protocol
## Structure
```
|----Header(1byte)-----|-----Body(max255byte)-----|
|     BodyLength       |          Body            |
|______________________|__________________________|
```

# Usage

## 1. Setup

```shell
git clone git@github.com:kappa-lab/very-simple-chat.git
```

## 2. Run Server

```shell
cd very-simple-chat
go run .
```

## 3. Join client
- Open new terminal (Establish Client Window)

```shell
cd very-simple-chat/client
go run .
```

## 4. Send Message
Input command into Client Window.

### Broadcast
tartget:255, send all clients.

```shell
{"target":255, "message":"hello evrybody"}
```

### Unicast
tartget:n(<255), send single clinet.
 
```shell
{"target":1, "message":"hello 1"}
```

## 5. Leave client
`Ctrl+C`

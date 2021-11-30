# What is this
- Simple & Primitive multi client communication system.
   - e.g. chat system
- for larning   

# Supported
- Broadcast message
- Unicast message

# Not Supported
- Multi Room
- Mutex Control
- Error Handlling
- Ping/Pong based Alive monitoring 

# Protocol
WIP

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
Input command into Client Window

### Broadcast
- tartget 255 as Broadcast

```shell
{"target":255, "message":"hello evrybody"}
```

### Unicast
- tartget 255 as Broadcast
 
```shell
{"target":1, "message":"hello 1"}
```

## 5. Leave client
`Ctrl+C`

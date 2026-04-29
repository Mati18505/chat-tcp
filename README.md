# What it is doing?
One client reads terminal input and sends messages to server.
Second client receives them from server an prints into terminal. 
Server broadcasts all received messages.

# Why?
This project was created to learn how to wrap the TCP protocol into usable message protocol.
TCP connection is wrapped in:
- framing based on length information from the headers
- message marshalling

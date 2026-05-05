# What is it?
One client reads terminal input and sends messages to server.
Second client receives them from server an prints into terminal. 
Server broadcasts all received messages.
All built on pure tcp.

# Why?
This project was created to learn how to wrap the TCP protocol into basic usable message protocol.
TCP connection is wrapped in:
- message framing based on length information from the headers
- message marshalling

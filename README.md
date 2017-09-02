golang im server chat with heatbeat check
===========

This is a demo tcp chat client and server with BigEndian binary protocol.

include heatbeat check

##The Protocol

A 4 or 8 byte int in big endian format is sent first. This int represents the length of the message to be sent
and is considered the header. The message is then sent.

This allows one to avoid the messiness of having to choose a delimiter and having arbitrary buffer sizes.
By using such a protocol we can determine the size of the buffer to be used for the message by reading
the 4 byte header and allocating a buffer large enough to accomodate the message.

If a message of more than 4GB needs to be sent, the header can be increased to 8 bytes. This has to be determined
carefully when creating the protocol since once defined the protocol cannot be extended without breaking backwards
compatibility.


## Useage

configFile.FlAppInfo.CONN_HeatBeatEnable = 1 enable heatcheck,server will send a pack, client should send a response; if CONN_HeatBeatEnable = 0,close server heat check


=================================
- [x] support chat msg
- [x] support broadcast communication
- [ ] support group communication
- [ ] support message encryption
- [ ] support file transfer
- [ ] support audio
- [ ] support video
- [ ] support RCP Protobuf
- [ ] support Microservice 

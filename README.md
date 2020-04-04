# Matrix Send

This is a little utility that allows you to send a message into a given
channel.

```
export MATRIX_USER=someuser
export MATRIX_PASSWORD=somepassword
export MATRIX_HOMESERVER_URL=https://matrix.org

$ echo "hello" | matrix-send --room ...
```

# Backend for Realtime Chat Webapp
<https://github.com/prince6635/realtime-chat-webapp>

* Packages
    * ~~(depreciated) RethinkDB: $ go get -u github.com/dancannon/gorethink~~
    * RethinkDB: $ go get -u gopkg.in/gorethink/gorethink.v3
    * WebSocket: $ go get -u github.com/gorilla/websocket
    * Decode: $ go get -u github.com/mitchellh/mapstructure
    
* Run
    ```
    under project folder: 
    $ go build
    $ ./realtime-chat-webapp-backend
    $ go clean
    
    or 
    
    $ go run *.go
    ```

* Test
    * Test client connects to the server with "channel subscribe" event, then send a "channel add" event
    ```
    http://jsbin.com/casilicapo/edit?js,console,output
    
    let msg = {
      name: 'channel add',
      data: {
        name: 'Hardware Support'
      }
    }
    let subMsg = {
      name: 'channel subscribe'
    }
    
    let ws = new WebSocket('ws://localhost:8080');
    
    ws.onopen = () => {
     ws.send(JSON.stringify(subMsg));
     ws.send(JSON.stringify(msg));
    }
    
    ws.onmessage = (e) => {
     console.log(e.data);
    }
    ```
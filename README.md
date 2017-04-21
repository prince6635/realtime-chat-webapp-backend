# Backend for Realtime Chat Webapp
<https://github.com/prince6635/realtime-chat-webapp>

* Packages
    * RethinkDB: $ go get -u github.com/dancannon/gorethink
    * WebSocket: $ go get -u github.com/gorilla/websocket
    * Decode: $ go get -u github.com/mitchellh/mapstructure
    
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
<template>
  <div id="chat">
    <div class="messages">
      <div class="item" v-for="(message) in messages" :key="message.id" v-bind:class="{self: message.self}">
        <div class="avatar"></div>
        <div class="name">{{ message.name }}</div>
        <div class="message">{{ message.message }}</div>
      </div>
    </div>
    <div class="message-input">
      <input class="text" v-on:keydown.enter.prevent="send()" v-model="messageText" type="text">
      <button class="send" @click="send()">Send</button>
    </div>
  </div>
</template>

<script>
  export default {
    name: "Chat",
    data: function() {
      return {
        messages: [],
        messageText: null,
        socket: null,
        token: null,
      };
    } ,
    methods: {
      send() {
        let message = JSON.stringify({
          name: this.name,
          message: this.messageText,
          token: this.token,
        })
        this.socket.send(message)
        this.messageText = null

        return false
      },
      authorize() {
        this.name = sessionStorage.getItem('name')
        if (this.name == null) {
          this.name = prompt("Enter your name", "")
          sessionStorage.setItem('name', this.name)
          if (this.name == null) {
            location.reload()
          }
        } else {
          this.getAuthenticationToken()
        }
      },
      getAuthenticationToken() {
        fetch('/auth', {
          method: 'post',
          body: JSON.stringify({
            name: this.name
          })
        }).then(response => {
          return response.json()
        }).then(response => {
          if (response.token.length > 0) {
            this.token = response.token
          }
        })
      },
    },
    created() {
      this.authorize()
      this.socket = new WebSocket('ws://' + window.location.host + '/ws')
      this.socket.addEventListener('message', (e) => {
        let msg = JSON.parse(e.data)
        if (msg.token === this.token) {
          msg.self = true
        }
        this.messages.push(msg)
      })
    }
  }
</script>

<style>
#chat {
  display: block;
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  width: 100%;
  height: 100vh;
  background: #2c3e50;
}

.messages {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 60px;
  overflow-x: auto;
}

.messages .item {
  width: 60%;
  height: auto;
  padding: 10px;
  display: block;
  background: bisque;
  border-radius: 10px;
  margin: 10px;
  float: left;
}

.messages .item.self {
  position: relative;
  float: right;
  background: #00ddff;
}

.messages .name {
  font-weight: bold;
  padding-bottom: 10px;
}

.message-input {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  background: silver;
  height: 60px;
}

.message-input .text {
  position: absolute;
  top: 10px;
  left: 10px;
  right: 10px;
  bottom: 10px;
  display: block;
  width: calc(100% - 135px);
  border: 1px solid #999999;
  border-radius: 5px;
}

.message-input .send {
  width: 100px;
  height: 40px;
  display: block;
  position: absolute;
  top: 10px;
  right: 10px;
  bottom: 10px;
  background: #0060FF;
  font-size: 18px;
  font-weight: bold;
  color: #FFFFFF;
  border: 1px solid #999999;
  border-radius: 5px;
}
</style>
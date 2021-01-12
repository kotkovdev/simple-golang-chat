<template>
  <div id="app">
    <ChatComponent></ChatComponent>
  </div>
</template>

<script>
import ChatComponent from "@/components/ChatComponent";

export default {
  name: 'App',
  components: {
    ChatComponent
  }
}
</script>

<style>
body {
  background: rgba(0,0,0,0.2);
  margin: 0;
  padding: 0;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: left;
  color: #2c3e50;
  max-width: 1000px;
  height: 100vh;
  display: block;
  position: relative;
  margin: 0 auto;
}
</style>

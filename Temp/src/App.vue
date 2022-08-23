<template>
  <img alt="List logo" src="./assets/todo-icon.png" align = "center">
  <br><h1>Welcome to your list!</h1>
  <p>Select the item to show details!</p>
  <MyListData/>
</template>

<script>
import MyListData from './components/MyListData.vue'
export default {
  name: 'App',
  components: { MyListData }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
h1{
  color: #151B8D;
  border-bottom: 1px solid #ddd;
  display: inline-block;
  padding-bottom: 10px;
}
</style>

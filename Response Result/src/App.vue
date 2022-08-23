<template>
  <img alt="List logo" src="./assets/todo-icon.png" align = "center">
  <br><h1>Welcome to your list!</h1>
  <p>Select the item to show details!</p>
  <div v-for="element in myList" v-bind:key = "element.id">
    <button @click = "toggleShowItem(element)">task id : {{element.id}}</button>
      <div v-if = "element.show_item">
        <div class = "backdrop" @click = "closeModal">
            <div class = "modal" :class = "{blue : theme ==='blue'}">
                <h1>{{header}}</h1>
                <p align ="left">
                  <br><b>id : </b>{{element.id}}
                  <br><b>task name : </b>{{element.item.Name}}
                </p>
                <ul align ="left" v-for="sub in element.item.SubItem" v-bind:key = "sub">
                  - {{sub}}
                </ul>
                <p align ="left">
                  <br><b>status : </b>{{element.status}}
                  <br><b>duration time : </b>{{element.duration_time}}
                </p>
            </div>
        </div>
      </div>
  </div>
</template>

<script>
import TodoService from "./TodoService";
export default {
  name: 'App',
  data() {
    return {
      header: 'My item',
      theme: 'blue',
      myList: []
    }
  },
  methods : {
    toggleShowItem(element){
      element.show_item = !element.show_item
    },
    retrievelist(){
      TodoService.getAll().then(response => {
        this.myList = response.data;
        console.log(response.data);
      })
      .catch(e => {
        console.log(e);
      });
    },
    closeModal(){
      this.$emit('close')
    }
  },
  created(){
    this.retrievelist();
  }
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
    border: none;
    padding: 0px;
  }

  .modal{
    width: 400px;
    padding: 20px;
    margin: 100px auto;
    background: rgb(255, 255, 255);
    border-radius: 10px;
  }

  .backdrop{
    top: 0;
    position: fixed;
    background: rgba(0,0,0,0.5);
    width: 100%;
    height: 100%;
  }

  /* override p in global.css */
  .modal p{
    font-style: normal;
  }

  .modal.blue{
    background: #43BFC7;
    color: #151B54;
  }
</style>

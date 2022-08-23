<template>
    <div v-for="element in myList" v-bind:key = "element.id">
    <button @click = "toggleShowItem(element)">task id : {{element.id}}</button>
      <div v-if = "element.show_item">
        <MyListData :header = "header" :element = "element" :theme = "theme" @close = "toggleShowItem(element)" />
        <div class = "backdrop" @click = "closeModal">
            <div class = "modal" :class = "{blue : theme ==='blue'}">
                <h1>{{header}}</h1>
                <p align ="left">
                    <b>id : </b>{{element.id}}
                    <br><b>task : </b>{{element.item}}
                    <br><b>status : </b>{{element.status}}
                    <br><b>durationTime : </b>{{element.duration_time}}
                </p>
            </div>
        </div>
      </div>
  </div>
</template>

<script>
import TodoService from "../TodoService";
export default({
    name: 'MyTodoList',
    data() {
        return {
          showList: false,
          header: 'My item',
          theme: 'blue',
          myList: []
        }
    },
    methods : {
        toggleShowList(){
            this.showList = !this.showList
        },
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
    },
})
</script>

<style scoped >
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
    h1{
        color: black;
        border: none;
        padding: 0px;
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
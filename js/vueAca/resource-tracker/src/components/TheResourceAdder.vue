<template>
  <!-- <div class="error"> -->
  <!--   <h1 class="text-3xl text-red-500">Error</h1> -->
  <!-- </div> -->
  <form class="ml-2" action="">
    <h2 class="label">Title:</h2>
    <input 
    required
    type="text" 
    placeholder="Enter Title" 
    v-model="enteredVal.title"
    class="input">

    <h2 class="label">Description:</h2>
    <input 
    type="text" 
    required
    v-model="enteredVal.description"
    placeholder="Enter Description" 
    class="input">

    <h2 class="label">Link:</h2>
    <input 
    v-model="enteredVal.link"
    required
    type="text" 
    placeholder="Enter link"  
    class="input">
    <p></p>
    <button @click.prevent="emitToParent" class="bg-green-600 text-white mt-4 p-2 px-4 md:px-8 lg:px-12 rounded">Add</button>
  </form>
  <!-- <error-handle></error-handle> -->
</template>

<script>
import errorHandle from './TheErrorHandle.vue'
export default {
  emits: ['val-above'],
  components: {
      errorHandle
    },
  data() {
      return {
          enteredVal:{
               title:'', 
               description:'', 
               link:''
            }
        }
    },
  methods: {
      emitToParent() {
          /* if(this.enteredVal.title !== '' && this.enteredVal.description !== '' && this.enteredVal.link !== ''){
              return this.$emit('val-above', this.enteredVal)
            }
          else {
            return this.$emit('val-above', false)
          } */
          return this.$emit('val-above', this.enteredVal)
        },
    }
}
</script>

<style scoped>
  .label{
      @apply text-black text-2xl mt-4 font-bold
    }
  .input{
      @apply bg-gray-100 border border-black w-3/4 p-2 my-2
    }
/*
  .error{
      position: relative;
      color: red;
      z-index: 1;
    }
  .error::before {
      content: "x";
      color: black;
      position: absolute;
      background: black;
      height: 90vh;
      inset: 0;
      opacity: 0.8;
      z-index: 0;
    } */
</style>

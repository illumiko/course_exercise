<template>
  <div class="assContainer">
    <h1 class="assHeader">Assigmnet 4</h1>
    <div class="assUnderline"></div>
    <div>
      <span class="text-2xl text-white">enter name <input v-model='userInput1' type="text" class="inputText w-full sm:w-1/3"></span>
      <div class="my-4 grid sm:grid-cols-2 sm:gap-x-4"> 
        <p class="textOutput text-center my-2" :class="[{notVisible: show}, userInput1]" >Style Me {{userInput1}}</p>
        <p class="btn my-2 my-0 sm:mx-4 grid place-items-center" @click='toggleVis()'>Toggle Style Paragraph</p>
      </div>
    </div>
    <!-- <div class="taskSeparator"></div> -->
    <div class="">
      <input type="text" v-model='userInput2'  class="inputText w-full sm:w-2/4" >
      <p class="textOutput my-8 text-center" :style="setBgC">Change</p> 
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      userInput1: '',
      userInput2: '',
      show: false
    }
  },
  computed: {
   setBgC () {
     const val = `background:${this.userInput2}`
     return val
   } 
  },
  methods: {
    toggleVis() {
      this.show = !this.show
      console.log(this.show)
    }
  }
}
</script>

<style scoped>
  .usr1{
    @apply bg-blue-500;
  }
  .usr2{
    @apply bg-indigo-500;
  }
  .visible{
    opacity: 1;
  }
  .notVisible{
    opacity: 0;
  }
  
  
</style>

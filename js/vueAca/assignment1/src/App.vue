<template>
  <h1 class="assHeader text-white pt-4">Assignments</h1>
  <Ass1 />
  <Ass2 />
  <Ass3 />
  <Ass4 />
  <Ass5 />
  <Ass6a :age="ass6Obj.age" :name="ass6Obj.name" class="my-0 z-20"/>
  <Ass6b @changeInfo="changeInitiated" class="-my-14 z-10"/>
</template>

<script>
  import Ass1 from './components/ass1.vue'
  import Ass2 from './components/ass2.vue'
  import Ass3 from './components/ass3.vue'
  import Ass4 from './components/ass4.vue'
  import Ass5 from './components/ass5.vue'
  import Ass6a from './components/ass6a.vue'
  import Ass6b from './components/ass6b.vue'

export default {
  name: 'App',
  components: {
    Ass1,
    Ass2,
    Ass3,
    Ass4,
    Ass5,
    Ass6a,
    Ass6b,
   },
   data() {
     return {
      ass6Obj: {
        name: 'N/A',
        age: 'N/A',
      }
     };
   },
   methods: {
     changeInitiated (n,a) {
       this.ass6Obj.name = n
       this.ass6Obj.age = a
     },
   }
}
</script>

<style>
  body {
    background:#555;
   }
   .btn {
     @apply cursor-pointer  select-none text-xl bg-red-500 hover:bg-green-500 block p-1 rounded-md text-center;
    }
    .taskSeparator {
      @apply border-2 border-solid border-green-600 rounded-2xl;
    }
  .textOutput{
    @apply text-xl sm:text-2xl bg-green-300 p-2 mb-4 rounded-2xl text-center;
   }
  .text{
    @apply text-xl sm:text-2xl bg-green-500 p-2 mb-4 rounded-2xl text-center;
   }
  .inputLabel{
    @apply text-xl sm:text-2xl text-white;
  }
  .inputText{
    @apply w-3/4 sm:w-2/4 text-black bg-gray-300;
  }
  .assHeader {
    @apply text-white text-3xl sm:text-4xl text-center mb-2;
  }
  .assContainer {
    @apply bg-gray-800 p-10 rounded-2xl m-12 sm:m-16 mt-8 lg:mx-60;
  }
  .assUnderline {
    @apply border-2 border-black border-solid rounded-2xl mb-10;
  }
</style>

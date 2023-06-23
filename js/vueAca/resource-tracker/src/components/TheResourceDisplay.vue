<template>

  <h1 class="text-black mb-4 text-3xl ml-2 pt-3 font-bold">Resources:</h1>
  <div 
  class="bor my-4 p-2 md:my-8 md:p-4 flex flex-row justify-between" 
  :key="res.title" 
  v-for="res in resources
  ">
    <div class="flex flex-col items-baseline justify-between">
      <h1 class="text-black ml-2 text-2xl md:text-4xl font-semibold mt-1 mb-2">{{res.title}}</h1>
      <p class="ml-5 text-black text-xl md:text-2xl">{{res.description}}</p>
      <a class="ml-5 text-purple-600 font-medium text-xl md:text-2xl" target="_blank" :href="res.link">Go to website</a>
    </div>
      <button @click="$emit('removeItem', res.title)" class="text-black p-2 bg-red-500 text-xs rounded">X</button>
  </div>
</template>

<script>
import baseCard from './baseCard.vue'
export default {
  props: ['resources'],
  emits: ['removeItem'],
  components:[
    baseCard,
  ]
}
</script>

<style scoped>
  .bor{
      border-bottom: 3px solid black;
    }
</style>

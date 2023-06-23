<template>
    <div 
    :class="[true ? 'block' : 'hidden']"
    class="understandable text-black bg-white grid items-center rounded p-4 text-center" >
      <div>
        <h1 class="text-red-500 text-2xl font-bold font-mono">Please Fill out the form</h1>
        <button @click="$emit('errorOk')" class="px-4 py-2 rounded my-4 bg-purple-600">OK</button>
      </div>
    </div>
</template>

<script>
export default {
  emits: ['errorOk'],
  props: ['manage'],
  data() {
      return{
          displayModal: true,
        }
    },
  methods: {
    }
  
}
</script>

<style scoped>
  .understandable{
      background: black;
      border: red 2px dashed;
      position: absolute;
      top: 10vh;
      bottom: 20vh;
      left: 20vw;
      height: clamp(20vw, 40vw, 80vw); 
      width: 60vw;
    }
</style>

<template>
<div class="assContainer">
  <h1 class="assHeader">Assignment 2</h1>
  <div  class="assUnderline"></div>
    <p class="text-2xl text-white inline">Click to Alert:</p>
    <button v-on:click='alerting' class="btn ml-2 inline mb-8">alert</button>
  <div class="taskSeparator"></div>
    <input type="text" v-on:input='liveUpdateInput' class="my-4 inputText">
    <p class="textOutput">Output: {{inputText}}</p>
  <div class="taskSeparator"></div>
    <input type="text" v-on:keydown.enter='liveUpInE' class="my-4 inputText">
    <p class="textOutput">Output: {{updatedText}}</p>
</div>

</template>

<script>
export default {
  data() {
    return {
      inputText: 'sd',
      updatedText: ''
      }
    },
  methods: {
      alerting() {
        return alert('hello world')
        },
      liveUpdateInput(e) {
        this.inputText = e.target.value
        },
      liveUpInE(e) {
          this.updatedText = e.target.value
       }
    }
}
</script>

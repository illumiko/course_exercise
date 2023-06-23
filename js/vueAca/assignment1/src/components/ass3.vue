<template>
  <div class="assContainer">
    <h1 class="assHeader">Assigment 3</h1>
    <div class="assUnderline"></div>
    <div class="flex sm:flex-row justify-center">
      <p class="btn my-4 w-full mx-2 w-max" @click="this.count+=5">Add 5</p>
      <p class="btn my-4 w-full mx-2 w-max" @click='this.count++'>Add 1</p>
    </div>
    <p class="textOutput">Result: {{conditionOfCount}}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      count:0,
    }
  },
  watch: {
    conditionOfCount() {
      setTimeout(() => {
         this.count = 0
       },5000)
    }
  },
  computed: {
    conditionOfCount() {
      if (this.count > 0){
        if(this.count < 37) {
            return 'not there yet'
          } else if (this.count === 37) {
            return '37'
         } else {
           return 'more '
         }
       } else {
          return 'click button to start'
       }
    }

  },
}
</script>


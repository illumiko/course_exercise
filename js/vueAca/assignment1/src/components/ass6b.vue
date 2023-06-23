<template>
  <div class="assContainer">
    <form action="">
      <input type="text" v-model="name" placeholder="Enter Name" class="inputText" required>
      <input type="text" v-model="age" placeholder="Enter Age" class="inputText mt-4" required>
      <button @click.prevent="submit" type="submit" class="btn mt-8 text-lg">Submit</button>
    </form>

  </div>
  
</template>
<script>
export default {
  data() {
    return{
      name: '',
      age: '',
    }
  },
  methods: {
    submit() {
      if(this.name !== '' && this.age !== '') {
        this.$emit('changeInfo', this.name, this.age) 
      }
    }
  }
  
}
</script>

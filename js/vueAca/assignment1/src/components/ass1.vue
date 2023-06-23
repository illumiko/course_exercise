<template>
  <div class="assContainer">
    <h1 class="assHeader">Assignment One</h1>
    <div class="assUnderline"></div>
    <p class="text">Name: {{name}}</p>
    <p class="text">Age(+5): {{age + 5 }}</p>
    <p class="text" id="randomNumber">Random Number: {{rando(1,0)}}</p>
    <p class="text" id="randomNumber">Random Number(without): {{Math.random() * (1-0) + 0}}</p>
    <img class="w-48 mx-auto sm:mx-0 h-36 rounded-2xl mb-4 bg-center bg-no-repeat bg-cover" v-bind:src="imgUrl"  alt="">
    <span class="inputLabel">Enter Text: <input class="inputText" :value=this.val ></span>
  </div>
    
</template>

<script>
export default {
  data() {
    return{
        name: 'illumiko',
        age: 12,
        imgUrl: 'https://c.wallhere.com/photos/94/13/anime_anime_girls_Sword_Art_Online_Yuuki_Asuna_long_hair_brown_eyes-19126.jpg!d',
        val: 'entered'
      }
    },
  methods: {
    rando(max, min) {
      return Math.random() * (max-min) + min
    }
  },
}
</script>

<style scoped>
</style>

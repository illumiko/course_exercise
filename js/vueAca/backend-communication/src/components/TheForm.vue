<template>
  <form class="" @submit.prevent="submitForm">
    <!-- radio btns -->
    <h2 class="formHeader">Radio Buttons</h2>
    <div class="radio">
      <div class="">
        <input type="radio" name="rad" v-model="formInputs.radBox" value="Nvim" id="rad1">
        <label for="rad1">Nvim</label>
      </div>
      <div>
        <input type="radio" name="rad" v-model="formInputs.radBox" value="Vim" id="rad2">
        <label for="rad2">Vim</label>
      </div>
      <div>
        <input type="radio" name="rad" v-model="formInputs.radBox" value="VsCode" id="rad3">
        <label for="rad3">Vscode</label>
      </div>
    </div>
    <!-- Cheakbox -->
    <h2 class="formHeader">Checkbox Buttons</h2>
    <div class="checkbox">
      <div>
        <input type="checkbox" name="check" v-model="formInputs.checkBox" value="Nvim" id="check1">
        <label for="check1">Nvim</label>
      </div>
      <div>
        <input type="checkbox" name="check" v-model="formInputs.checkBox" value="Vim" id="check2">
        <label for="check2">Vim</label>
      </div>
      <div>
        <input type="checkbox" name="check" v-model="formInputs.checkBox" value="Vscode" id="check3">
        <label for="check3">Vscode</label>
      </div>
    </div>
    <!-- Input -->
    <div class="input">
      <h2 class="formHeader">Text Input</h2>
      <div class="flex flex-col items-center">
        <span class="mr-4">Enter something</span>
        <input type="text" v-model="formInputs.text" class="py-1 px-1 w-3/4 md:w-2/4 lg:w-1/4 text-black">
      </div>
    </div>
    <!-- Custom form field -->
    <h2 class="formHeader">Recommneded?</h2>
    <would-recommend v-model:rating="this.formInputs.rating"></would-recommend>
    <!-- <div>{{this.formInputs}}</div> -->
    <!-- <would-recommend></would-recommend> -->
    <button class="m-auto block bg-green-500 p-2 px-4 cursor-pointer text-black">Submit</button>
  </form>
  <div> {{this.fetchedData}} </div>
  <button 
    class="px-4 py-2 border border-green-400 m-auto block" 
    type="button" @click="this.loadData()">Click to see what other People use
  </button>
  <!-- Load submitted data -->
  <p class="text-red-400 text-2xl py-8 text-center" v-if="(this.fetchedData.length == 0 && this.isLoading == false)">No data</p>
  <div
    v-for="i in this.fetchedData"
    :key="i.id"
    v-else-if="this.fetchedData.lenth !== 0"
  >
    <h1 class="underline text-center py-8">User:{{i.id}}</h1>
    <p class="text-green-200">Using: <span class="text-white"> {{i.editor}}</span></p>
    <p class="text-green-200">Prefers: <span class="text-white"> {{i.preferance}}</span></p>
    <p class="text-green-200">Rating: <span class="text-white"> {{i.rating}}</span></p>
    <p class="text-green-200">Comment: <span class="text-white"> {{i.comments}}</span></p>
  </div>
  <p class="py-8 text-center text-blue-400" v-else-if="this.isLoading === true">Loading...</p>
  

</template>

<script>
import wouldRecommend from './TheWouldRecommend.vue'
export default {
  components:{
    wouldRecommend
  },
  data() {
    return {
      formInputs:{
        rating:"",
        text:"",
        radBox:"",
        checkBox:[],
      },
      fetchedData:[],
      isLoading: null,
    }
  },
  methods: {
    async submitForm (){
      fetch("https://idk-manh-default-rtdb.asia-southeast1.firebasedatabase.app/submitServey.json", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          name: "Eyanat",
          rating: this.formInputs.rating,
          editor: this.formInputs.radBox,
          prefer: this.formInputs.checkBox,
          comments: this.formInputs.text
        }),
      })
      this.resetForm()
    },
    resetForm(){
      // resets the form
      for (let x in this.formInputs){
        this.formInputs[x] = ""
      }
      // because the checkbox field requires []
      this.formInputs['checkBox'] = []
      window.location.reload()
    },
    async loadData () {
      // fetch("https://idk-manh-default-rtdb.asia-southeast1.firebasedatabase.app/submitServey.json")
      // .then((fetch) => console.log(fetch))
      try{
        this.isLoading = true
        let response = await fetch("https://idk-manh-default-rtdb.asia-southeast1.firebasedatabase.app/submitServey.json",{method:"GET"})
        // this.isLoading = false
        let data = await response.json()
        for (let id in data){
          let parsed = {
            id:id,
            name:data[id].name,
            editor:data[id].editor,
            preferance:data[id].prefer,
            comments:data[id].comments,
            rating:data[id].rating,
          }
          this.fetchedData.push(parsed)
        }
          this.isLoading = false
      }catch(error) {
        console.error(error)
      }
    }, 
  }, //method ends 
}
</script>

<style scoped>
.formHeader{
  @apply text-2xl md:text-3xl text-green-400 py-4 font-bold text-center;
}
.radio div{
  @apply block my-4 ;
}
.radio > div,
.checkbox > div{
  @apply mx-2 md:mx-3
}
.radio, .checkbox{
  @apply flex justify-center text-2xl;
}
.radio input[type=radio],
.checkbox input[type=checkbox]{
  @apply mr-1 md:mr-2;
}
label{
  @apply text-xl;
}
</style>

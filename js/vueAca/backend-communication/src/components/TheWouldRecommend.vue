<template>
  <!-- <h1>hello</h1> -->
  <div class="flex justify-center">
    <button 
      :value="rating"
      class="border my-4 md:m-8 md:mt-0 ml-2 mr-2 p-2 border-green-400 hover:bg-green-400 hover:text-black transition-all"
      :class="[item.selected ? 'bg-green-400 text-black ' : '']"
      @click.prevent="toggle"
      v-for="item in feedback" :key="item.id"
      type="Submit"
     >{{item.text}}
    </button>
  </div>
</template>

<script>
export default {
  emits:['update:rating'],
  props:['rating'],
  data() {
    return {
      feedback:[
        {
          id: "yes",
          text: "Yes, I would",
          selected: false,
        },
        {
          id: "no",
          text: "No, I wouldnt",
          selected: false,
        },
      ]
    }
  },
  methods: {
    toggle(e){
      this.feedback.forEach(i => {
        i.selected = false
        if(i.text == e.target.innerHTML){
          i.selected = true
        }
      })
      let usrSelected = this.feedback.filter(i => i.selected === true)
      this.$emit("update:rating",usrSelected[0].text)
    },
  },
}
</script>

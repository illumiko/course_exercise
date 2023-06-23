<template>
  <div class="assContainer">
    <h1 class="assHeader">Assigment 5
      <div class="assUnderline"></div>
    </h1>
    <div class="listContainer">
      <p class="btn transition-all duration-100" :class="{active: listVisibility}" @click="this.listVisibility = !this.listVisibility">{{dynamicText}}</p>
      <h1 class="assHeader my-4 mt-10">List Adder</h1>
      <input type="text" class="inputText w-full py-2" placeholder="Enter Things To List" v-model="listName" @keypress.enter="this.lists.push(listName)">
      <div v-if="this.listVisibility" class="list">
        <p class="listItems transition-all " v-for="item in lists" :key='item' @click="removeLists">{{item}}</p>
        <p v-if='lists.length == 0' class="text my-8">no list items</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      listName: '',
      listVisibility: true,
      lists: []
    }
  },
  computed: {
    dynamicText() {
      if(this.listVisibility) {
          return 'visible'
      } else {
          return 'hidden'
      }
    }
  },
  methods: {
    removeLists(e) {
      let x = e.target.innerText
      this.lists.splice(x,1)
    }
  }

}
</script>

<style scoped>
  .listItems{
    @apply cursor-pointer relative bg-pink-500 my-12 p-2 rounded-xl text-xl sm:text-2xl pl-6;
  }

  .listItems:before{
      content: 'Click to remove';
      transition: all 75ms ease;
      position: absolute;
      inset: 0;
      display: flex;
      align-items: center;
      background-color: #000;
      color: white;
      opacity: 0;
      padding-left: 1rem;
      @apply rounded-xl;
  }
  .listItems:hover:before{
    opacity: 1;
  }

  .active{
    @apply bg-yellow-500;
  }

</style>

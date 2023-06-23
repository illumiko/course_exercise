<template>
  <div class="bg-gray-800 min-h-screen pb-20 pt-5 font-sans">
    <h1 class="text-4xl hidden text-center py-6 bg-gray-50 text-black rounded font-semibold">Resources</h1>
    <base-card>
      <div class="flex flex-row ml-2 md:ml-4">
        <p v-for="x in tabs" 
        v-bind:key="x.id"
        @click="changeCmp"
        class="bg-purple-600 text-white mr-4 md:mr-8 p-2 rounded"
        :class="[x.vis ? 'opacity-100': 'opacity-60']"
        >{{x.text}}</p>
      </div>
    </base-card>

    <!-- SHOWING ADDED RESOURCES -->
    <base-card  v-if="this.tabs[0].vis">
      <resource-display 
      :resources="this.resourcesBase"
      @removeItem="filterArray"
      ></resource-display>
      <h1 class="ml-3 text-red-500 font-bold text-3xl" v-if="resourcesBase.length == 0">No resources</h1>
    </base-card>

    <!-- FOR ADDING NEW RESOURCES -->
    <base-card v-if="this.tabs[1].vis">
      <resource-adder 
      @val-above="add"
      ></resource-adder>
    </base-card>

      <!-- HANDLING THE ERROR SHOWN WHEN INPUT EMPTY -->
      <error-handle 
      @errorOk="removeModal"
      v-if="manageModal && tabs[1].vis"
      ></error-handle>
  </div>
</template>

<script>
import baseCard from './components/baseCard.vue'
import resourceDisplay from './components/TheResourceDisplay.vue'
import resourceAdder from './components/TheResourceAdder.vue'
import errorHandle from './components/TheErrorHandle.vue'

export default {
  name: 'App',
  computed:{
    viewTab() {
        return this.tabs.filter(i => i.vis == true)
      }

    },
  data() {
    return{
      manageModal: false,
      // MANGES TABS
      tabs:[
          {
            id:'idk',
            vis: true,
            text: 'Resources'
          },
          {
            id:'two',
            vis: false,
            text: 'Add More'
          },
        ],
      // MANAGES RESOURCES
      resourcesBase: [
          {
            title: 'Javascript Info',
            link: 'https://javascript.info/',
            description: 'Javascript documentation'
          },
          {
            title: 'Google',
            link: 'https://www.google.com/',
            description: 'Defacto Search Engine'
          },
        ],
      }
    },
  components: {
    errorHandle,
    baseCard,
    resourceAdder,
    resourceDisplay,
    },

  methods: {
      filterArray(x) { // for removing an item from resources
          console.log(x)
          const find = this.resourcesBase.filter(i => i.title == x)
          const findIndex = this.resourcesBase.indexOf(find[0])
          this.resourcesBase.splice(findIndex, 1)
        },

      changeCmp(e) { // changing tab view
          this.tabs.forEach(i => i.vis = false)
          const got = e.target.innerText
          const found = this.tabs.filter(i => i.text == got)
          found[0].vis = true
        },

      idGen() { // to generate unique ids
          const id = Math.random()*(100 - 1) + 1
          return Math.floor(id)
        },

      add(x){ // HANDLING EMIT FROM RESOURCE ADDER
          const {title, description, link} = x
          console.log(title, description, link)
          if(title && description && link){ // checks if all the input are empty or not
            const newRes = {
                title: title,
                description: description,
                link: link,
              }
            this.resourcesBase.push(newRes)
            }
          else{
              this.manageModal = true
            }
        },
      removeModal() {
          this.manageModal = false // removes displayeed modal
        }
    }
}

</script>

<style>
  div{
      color:white;
    }
</style>

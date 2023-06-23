<template>
  <h1 class="header">Monster Slayer</h1>
  <div class="container">
    <div class="colF">
      <div class="healthInfo">
        <div class="monster">
        <p class="name font-semibold">Monster</p>
          <div class="healthBar p-2">
            <div class="actulaHealthMonster" :style="{width: monsterHp + '%'}">{{monsterHp}}</div>
          </div>
        </div> 

        <div class="me my-10">
        <p class="name font-bold">Me</p>
          <div class="healthBar">
            <div class="actulaHealthMe" :style='{width: playerHp + "%"}'>{{playerHp}}</div>
          </div>
        </div> 
      </div>

      <div class="userOption">
        <div class="containerBtns">
          <p @click='playerAtk()' class="userOptionBtn">Attack</p>
          <p @click='healPlayer()' class="userOptionBtn">Heal</p>
          <p @click='playerSpecialAtk()' class="userOptionBtn" :class="{'bg-gray-600 pointer-events-none': !disableBtn()}">Special Attack</p>
          <p @click='monsterHp = 1; playerHp = 0; charge=0' class="userOptionBtn">Surrender</p>
        </div>
      </div>
    </div>

    <div v-if="Blog.length !== 0" class="battleLog">
      <h1 class="text-center text-3xl my-4 mb-10 text-white">Battle log</h1>
      <div class="logs">
          <p class="text-2xl text-white p-4 text-yellow-500" v-for="(log, index) in Blog" :class="[{'text-green-500': colorCodedText(`${log}`)},]"  :key='index'> {{log}}</p>
      </div>
    </div>
  </div>

  <div v-if="playerHp <= 0" class="loseScreen">
    <div class="bg-black grid place-items-center m-4 absolute top-1/3 bottom-1/3 left-0 right-0 opacity-90">
      <p class="text-4xl text-white pt-4 text-center font-extrabold">YOU LOSE</p>
      <p @click="monsterHp = 100; playerHp = 100; charge=0" class="userOptionBtn bg-yellow-500 hover:bg-yellow-700">Restart</p>
    </div>
  </div>

  <div v-if="monsterHp <= 0" class="winScreen">
    <div class="bg-black grid place-items-center m-4 absolute top-1/3 bottom-1/3 left-0 right-0 opacity-90">
      <p class="text-4xl text-white pt-4 text-center font-extrabold">YOU WIN</p>
      <p @click='monsterHp = 100; playerHp = 100; charge=0' class="userOptionBtn bg-yellow-500 hover:bg-yellow-700">Restart</p>
    </div>
  </div>

</template>

<script>

export default {
  name: 'App',
  data() {
    return{
      playerHp: 100,
      monsterHp: 100,
      Blog: [],
      charge: 0,
    }
  },
  methods: {
    randoNum(min,max) {
      return  Math.floor(Math.random() * (max - min)) + min
    },
    monsterAtk() {
      let dmg = this.randoNum(5,18)
      let logMsg = `Monster attacks player for ${dmg}`
      this.Blog.push(logMsg)
      this.playerHp -= dmg
    },
    playerAtk() {
      let dmg = this.randoNum(5,12)
      this.monsterHp -= dmg
      let logMsg = `Player attacks monster for ${dmg}`
      this.Blog.push(logMsg)
      this.monsterAtk()
      this.charge++
    },
    colorCodedText(text) {
      const regex = /Player/g
      return regex.test(text)
    },
    healPlayer() {
      let healAmount = this.randoNum(3,20)
      this.playerHp += healAmount
      let logMsg = `Player heals by ${healAmount}`
      this.Blog.push(logMsg)
      this.monsterAtk()
      this.charge++
    },
    disableBtn() {
      if(this.charge >= 2) {
        return true
      }else {
        return false
      }
    },
    playerSpecialAtk() {
      let dmg = this.randoNum(10, 20)
      this.monsterHp -= dmg
      let logMsg = `Player uses special attack on monster dealing ${dmg}`
      this.Blog.push(logMsg)
      this.monsterAtk()
      this.charge = 0
    },
  },
  computed: {
    
  }

}
</script>


<style>
/* layout */
body{
  @apply bg-gray-900
}
.container{
  @apply m-auto lg:grid lg:grid-cols-4 lg:grid-rows-1
}
.colF{
  @apply col-span-3
}
.name{
  @apply text-xl my-4 sm:text-2xl text-white;
}
.header{
  @apply text-3xl sm:text-4xl my-10 text-center text-white
}
/* user actions */
.userOptionBtn{
  @apply bg-gray-400 w-36 text-center my-4 py-4 font-semibold sm:py-8 sm:w-40 lg:w-52 sm:text-lg rounded-xl cursor-pointer select-none hover:bg-gray-500
}
.userOption{
  @apply my-8
}
.containerBtns{
  @apply grid justify-center sm:grid-cols-2 sm:grid-rows-2 sm:place-items-center
}
/* battle log */
.battleLog{
  @apply my-8 border-4 border-gray-600 border-solid mx-10 lg:mx-0 lg:my-0 h-96 lg:h-96 overflow-y-scroll
}
.logs{
  @apply m-auto
}
.logText{
  @apply text-center text-2xl my-2 text-blue-300
}
/* health bar */
.healthInfo{
  @apply mx-8 sm:mx-20 bg-gray-700 p-8 rounded-2xl
}
.healthBar{
  @apply bg-gray-300;
  height: 2.5rem;
  width: 100%;
  position: relative;
}
.actulaHealthMonster{
  position: absolute;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  inset: 0;
  height: 2.5rem;
  width: calc(100%);
  @apply bg-green-500;
}
.actulaHealthMe{
  @apply bg-green-500;
  position: absolute;
  display: flex;
  align-items: center;
  justify-content: end;
  inset: 0;
  height: 2.5rem;
  width: calc(100%)
}
</style>

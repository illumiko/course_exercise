<template>
  <div class="mainContainer bg-gray-900 text-xl text-white min-h-screen">
    <h1 class="text-4xl py-8 font-semibold accent text-center">Form</h1>
    <the-card>
      <the-form></the-form>
    </the-card>
  </div>
</template>

<script>
import theForm from './components/TheForm.vue'
import theCard from './components/TheCard.vue'
export default {
  name: 'App',
  components: {
    theForm,
    theCard
  }
}
</script>

<style>
:root{
  --bg: #313131;
  --fg: #525252;
  --accent: #ca3e47;
}
.mainContainer{
  padding:clamp(1rem, 1vh, 50rem);
  padding-top: 0;
}
@media (min-width:580px) {
  .mainContainer{
    padding:clamp(5rem, 15vh, 50rem);
    padding-top: 0;
  }
}
@media (min-width:1000px) {
  .mainContainer{
    padding:clamp(5rem, 40vh, 100rem);
    padding-top: 0;
  }
}
@media (min-width:1600px) {
  .mainContainer{
    padding:clamp(5rem, 60vh, 100rem);
    padding-top: 0;
  }
}
</style>

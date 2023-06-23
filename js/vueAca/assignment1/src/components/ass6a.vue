<template>
  <div class="assContainer">
    <h1 class="assHeader">Communication between Parent and components</h1>
    <div class="assUnderline"></div>
    <div>
      <p class="text">Name: {{name}}</p>
      <p class="text">Age: {{age}}</p>
    </div>

  </div>
</template>

<script>
export default {
  props: ['name', 'age'],
}
</script>

<template>
  <main>
    <div class="left">
      <p class="tips">Ctrl + Q Hide</p>
    </div>
    <div class="right">
      <p class="quit" @click="quit"> X </p>
    </div>
  </main>
</template>

<script setup>
import {Quit} from "../../wailsjs/runtime/runtime.js";

function quit() {
  Quit();
}
</script>

<style scoped>
  main {
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 100vh;
    padding: 0 20px;
    box-sizing: border-box;
    overflow: hidden;
  }

  .left {
    display: flex;
    align-items: center;
  }

  .quit {
    cursor: pointer;
    padding: 10px;
    box-sizing: border-box;
  }

  p {
    color: rgba(255, 255, 255, 0.3);
  }
</style>
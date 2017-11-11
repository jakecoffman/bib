<template>
  <div class="history">
    <div class="flex align-center">
      <h2 class="flex-1">History</h2>
      <a class="align-end back pager" @click="$router.go(-1)">🔙</a>
    </div>
    <div v-if="history.length === 0">
      No history yet!
    </div>
    <div v-else>
      <ul class="right">
        <li v-for="item of history">
          <router-link :to="{name: 'Bible', params: {book: item.book, chapter: item.chapter}}">
            {{item.book}} {{item.chapter}}
          </router-link>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  import {mapState} from 'vuex'

  export default {
    name: 'History',
    data() {
      return {
      }
    },
    computed: {
      ...mapState({
        history: 'history'
      })
    },
    methods: {
    }
  }
</script>

<style>
  a {
    cursor: pointer;
  }

  .history {
    font-size: 36pt;
  }

  .right {
    text-align: right;
  }
</style>

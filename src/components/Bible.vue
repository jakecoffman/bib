<template>
  <div class="bible">
    <aside v-if="error" class="flex">
      <span class="flex-1">Error: {{error}}</span>
      <button @click="dismiss()" class="align-end">X</button>
    </aside>
    <nav class="flex">
      <h2 v-if="book" class="flex-1 pointer" @click="picker()">{{book.name}} {{chapter}} 📖</h2>
      <p class="align-end">
        <router-link :to="{name: 'History'}" class="page">📜</router-link>
        <a class="page" v-show="chapter > 1" @click="back()">⬅️</a>
        <a class="page" @click="next()">➡️</a>
      </p>
    </nav>
    <spinner v-if="isLoadingVerses"></spinner>
    <span v-if="!isLoadingVerses" v-for="verse in verses">
      <span v-html="verse.Text"></span>
    </span>
    <p class="right">
      <a class="page" v-show="chapter > 1" @click="back()">⬅️</a>
      <a class="page" @click="next()">➡️</a>
    </p>
  </div>
</template>

<script>
  import bible from '@/services/bible'
  import Spinner from './Spinner'
  import {mapState} from 'vuex'

  export default {
    name: 'Bible',
    data () {
      return {
        verses: [],
        chapters: [],
        error: '',
        isLoadingVerses: true,
        isLoadingChapters: true
      }
    },
    computed: {
      ...mapState({
        book: state => {
          return bible.books.find(b => b.abbr === state.book)
        },
        chapter: 'chapter'
      })
    },
    watch: {
      async '$route.params.book' (newBook) {
        this.$store.commit('setBook', newBook)
        this.isLoadingVerses = true
        this.verses = await bible.fetchVerses(this.book.abbr, this.chapter)
        this.isLoadingVerses = false
      },
      async '$route.params.chapter' (newChapter) {
        this.$store.commit('setChapter', newChapter)
        this.isLoadingVerses = true
        this.verses = await bible.fetchVerses(this.book.abbr, this.chapter)
        this.isLoadingVerses = false
      }
    },
    async mounted () {
      this.isLoadingVerses = true
      this.isLoadingChapters = true

      this.$store.commit('setBook', this.$route.params.book)
      if (this.book === undefined) {
        this.error = 'Book not found'
        return
      }
      this.$store.commit('setChapter', this.$route.params.chapter)

      this.verses = await bible.fetchVerses(this.book.abbr, this.chapter)
      this.isLoadingVerses = false
    },
    methods: {
      dismiss() {
        this.error = ''
      },
      back() {
        if (this.chapter === 1) {
          return
        }
        this.$router.push({name: 'Bible', params: {book: this.book.abbr, chapter: this.chapter - 1}})
      },
      next() {
        this.$router.push({name: 'Bible', params: {book: this.book.abbr, chapter: this.chapter + 1}})
      },
      picker() {
        this.$router.push({name: 'Picker'})
      }
    },
    components: {
      Spinner
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .bible {
    font-size: 24pt;
  }
  .bible nav {
    display: flex;
  }
  .page {
    margin-left: 20px;
    font-size: 42pt;
  }
  .bible aside {
    background: red;
    padding: 10px;
    border-radius: 5px;
    font-weight: bolder;
  }
  .pointer {
    cursor: pointer;
  }
  .bible a {
    cursor: pointer;
  }
  .right {
    float: right;
  }
</style>

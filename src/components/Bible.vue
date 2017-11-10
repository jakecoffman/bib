<template>
  <div class="hello">
    <aside v-if="error" class="flex">
      <span class="flex-1">Error: {{error}}</span>
      <button @click="dismiss()" class="align-end">X</button>
    </aside>
    <nav class="flex">
      <h2 v-if="book" class="flex-1 pointer" @click="picker()">{{book.name}} {{chapter}} 📖</h2>
      <p class="align-end">
        <a class="page" :class="{disabled: chapter === 1}" @click="back()">⬅️</a>
        <a class="page" @click="next()">➡️</a>
      </p>
    </nav>
    <spinner v-if="isLoadingVerses"></spinner>
    <span v-else v-for="verse in verses">
      <span v-html="verse.Text"></span>
    </span>
    <p class="right">
      <a class="page" :class="{disabled: chapter === 1}" @click="back()">⬅️</a>
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
        this.verses = await bible.fetchVerses(this.book, this.chapter)
        this.isLoadingVerses = false
        this.chapters = await bible.fetchChapters(this.book.abbr)
        this.isLoadingChapters = false
      },
      async '$route.params.chapter' (newChapter) {
        this.$store.commit('setChapter', newChapter)
        this.verses = await bible.fetchVerses(this.book, this.chapter)
        this.isLoadingVerses = false
      }
    },
    async mounted () {
      if (this.$route.params.book) {
        this.$store.commit('setBook', this.$route.params.book)
        if (this.book === undefined) {
          this.error = 'Book not found'
          return
        }
      } else {
        this.$router.push({name: 'Chapter', params: {book: 'Gen', chapter: 1}})
      }
      if (this.$route.params.chapter) {
        this.$store.commit('setChapter', this.$route.params.chapter)
      } else {
        this.$router.push({name: 'Chapter', params: {book: this.book.abbr, chapter: 1}})
      }

      this.verses = await bible.fetchVerses(this.book, this.chapter)
      this.isLoadingVerses = false
      this.chapters = await bible.fetchChapters(this.book.abbr)
      this.isLoadingChapters = false
    },
    methods: {
      dismiss() {
        this.error = ''
      },
      back() {
        if (this.chapter === 1) {
          return
        }
        this.isLoadingVerses = true
        this.$router.push({name: 'Chapter', params: {book: this.book.abbr, chapter: this.chapter - 1}})
      },
      next() {
        this.isLoadingVerses = true
        this.$router.push({name: 'Chapter', params: {book: this.book.abbr, chapter: this.chapter + 1}})
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
  nav {
    display: flex;
  }
  aside {
    background: red;
    padding: 10px;
    border-radius: 5px;
    font-weight: bolder;
  }
  .pointer {
    cursor: pointer;
  }
  a {
    cursor: pointer;
  }
  .page {
    font-size: 90pt;
  }
  .right {
    float: right;
  }
</style>

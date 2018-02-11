<template>
  <div class="bible">
    <transition name="slide">
      <header v-show="isHeaderVisible">
        <nav class="flex align-center">
          <h2 v-if="book" class="flex-1 pointer" @click="picker()">
            {{book.name}} {{chapter}}
          </h2>
          <div class="align-end">
            <a class="pager" @click="picker()">📖</a>
            <router-link :to="{name: 'History'}" class="pager">📜</router-link>
            <a class="pager" v-show="chapter > 1" @click="back()">⬅️</a>
            <a class="pager" @click="next()">➡️</a>
          </div>
        </nav>
      </header>
    </transition>
    <aside v-if="error" class="flex">
      <span class="flex-1">Error: {{error}}</span>
      <button @click="dismiss()" class="align-end">X</button>
    </aside>
    <main>
      <spinner v-if="isLoadingVerses"/>
      <span v-html="verses"></span>
      <span id="main"></span>
    </main>
    <footer>
      <p class="right">
        <a class="pager" v-show="chapter > 1" @click="back()">⬅️</a>
        <a class="pager" @click="next()">➡️</a>
      </p>
    </footer>
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
        verses: "",
        chapters: [],
        error: '',
        isLoadingVerses: true,
        isLoadingChapters: true,

        // scrolly stuff
        isHeaderVisible: true,
        scroller: null,
        didScroll: false,
        previousScrollPos: 0
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
      },
      handleScroll() {
        this.didScroll = true
      },
      hasScrolled() {
        const scrollPos = document.scrollingElement.scrollTop
        if (scrollPos > this.previousScrollPos) {
          this.isHeaderVisible = false
        } else if (scrollPos < this.previousScrollPos) {
          this.isHeaderVisible = true
        }
        this.previousScrollPos = scrollPos
      }
    },
    components: {
      Spinner
    },
    created: function () {
      window.addEventListener('scroll', this.handleScroll)
      this.scroller = setInterval(() => {
        if (this.didScroll) {
          this.hasScrolled()
          this.didScroll = false
        }
      }, 100)
    },
    destroyed: function () {
      window.removeEventListener('scroll', this.handleScroll)
      clearInterval(this.scroller)
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
  $header: 115px;
  $hfont: 5rem;
  .bible {
    font-size: 24pt;
    margin: 0;
  }
  .bible header {
    background: #3c3c3c;
    height: $header;
    position: fixed;
    top: 0;
    width: 100%;
    font-size: $hfont - 3rem;
    border-bottom: 1px #646464 solid;
  }
  .bible .pager {
    font-size: $hfont;
  }
  .bible nav {
    display: flex;
    padding: 10px;
  }
  .bible aside {
    background: red;
    padding: 10px;
    border-radius: 5px;
    font-weight: bolder;
  }
  .bible main {
    padding: $header 0.5rem 0.5rem;
  }
  .pointer {
    cursor: pointer;
  }
  .bible a {
    cursor: pointer;
  }
  .bible h2 {
    margin: 0;
  }
  .right {
    float: right;
  }
  .slide-enter-active {
    transition: all 0.3s ease-in-out
  }
  .slide-leave-active {
    transition: all 0.3s ease-in-out
  }
  .slide-enter, .slide-leave-to {
    transform: translateY(-$header);
  }
</style>

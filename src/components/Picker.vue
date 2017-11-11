<template>
  <div class="picker">
    <div class="flex">
      <div id="ot" v-if="pickingBook">
        <h2>Old</h2>
        <ul>
          <li v-for="(b, index) in books.slice(0, 39)" class="choosable" :class="{chosen: b === book}">
            <a @click="pick(b.abbr)">{{index+1}}. {{b.name}}</a>
          </li>
        </ul>
      </div>
      <div id="nt" v-if="pickingBook">
        <h2>New</h2>
        <ul>
          <li v-for="(b, index) in books.slice(39)" class="choosable" :class="{chosen: b === book}">
            <a @click="pick(b.abbr)">{{index+40}}. {{b.name}}</a>
          </li>
        </ul>
      </div>
      <div id="chaps" v-if="!pickingBook">
        <div class="flex align-center">
          <h2 class="flex-1">{{book.name}}</h2>
          <a class="align-end pager back" @click="pickBookAgain()">🔙</a>
        </div>
        <spinner v-if="fetchingChapters"></spinner>
        <ul class="flex wrap">
          <li v-for="c in chapters" class="chapter-choose">
            <a @click="goto(c)">{{c}}</a>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
  import bible from '@/services/bible'
  import Spinner from '@/components/Spinner'
  import {mapState} from 'vuex'

  export default {
    name: 'Picker',
    data() {
      return {
        books: bible.books,
        chapters: [],
        pickingBook: true,
        fetchingChapters: false
      }
    },
    components: {
      Spinner
    },
    computed: {
      ...mapState({
        book: state => {
          return bible.books.find(b => b.abbr === state.book)
        },
        chapter: 'chapter'
      })
    },
    mounted() {
      this.pickingBook = true
    },
    methods: {
      cancel() {
        this.$router.go(-1)
      },
      async pick(bookAbbr) {
        this.$store.commit('setBook', bookAbbr)
        this.fetchingChapters = true
        this.pickingBook = false
        this.chapters = []
        this.chapters = await bible.fetchChapters(bookAbbr)
        this.fetchingChapters = false
      },
      goto(chapter) {
        this.$store.commit('setChapter', chapter)
        this.$router.replace({name: 'Bible', params: {book: this.book.abbr, chapter: chapter}})
      },
      pickBookAgain() {
        this.pickingBook = true
      }
    }
  }
</script>

<style>
  .chosen {
    background: red;
  }

  .choosable {
    padding-top: 10px;
  }

  ul {
    list-style: none;
  }

  a {
    cursor: pointer;
  }

  #ot {
    width: 50%;
  }

  #nt {
    width: 50%;
  }

  #chaps {
    width: 100%;
  }

  .picker {
    font-size: 36pt;
  }

  .chapter-choose {
    padding: 10px 10px 10px 15px;
  }
</style>

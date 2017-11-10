<template>
  <div>
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
        <h2><a @click="pickBookAgain()"><</a> {{book.name}}</h2>
        <spinner v-if="fetchingChapters"></spinner>
        <ul class="flex wrap">
          <li v-for="c in chapters" class="choosable">
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
    async mounted() {
      this.chapters = await bible.fetchChapters(this.book.abbr)
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
        this.$router.replace({name: 'Chapter', params: {book: this.book.abbr, chapter: chapter}})
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
    padding: 10px;
  }

  ul {
    list-style: none;
  }

  a {
    cursor: pointer;
  }

  .between {
    justify-content: space-between;
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

  body {
    font-size: 36pt;
  }
</style>

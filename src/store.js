import Vuex from 'vuex'
import Vue from 'vue'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    book: 'Gen',
    chapter: 1
  },
  mutations: {
    setBook (state, book) {
      state.book = book
    },
    setChapter (state, chapter) {
      state.chapter = parseInt(chapter)
    }
  }
})

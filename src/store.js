import Vuex from 'vuex'
import Vue from 'vue'
import createPersistedState from 'vuex-persistedstate'

Vue.use(Vuex)

export default new Vuex.Store({
  plugins: [createPersistedState()],
  state: {
    book: 'Gen',
    chapter: 1,
    history: []
  },
  mutations: {
    setBook (state, book) {
      state.book = book
    },
    setChapter (state, chapter) {
      state.chapter = parseInt(chapter)

      // enter history entry
      const entry = {book: state.book, chapter: state.chapter}
      let found = false
      for (let i=0; i<state.history.length; i++) {
        let item = state.history[i]
        if (item.book === entry.book && item.chapter === entry.chapter) {
          found = true
          state.history.splice(i, 1)
          state.history.unshift(entry)
        }
      }
      if (!found) {
        state.history.unshift(entry)
      }
    }
  }
})

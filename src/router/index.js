import Vue from 'vue'
import Router from 'vue-router'
import Bible from '@/components/Bible'
import Picker from '@/components/Picker'
import History from '@/components/History'

import store from '@/store'

Vue.use(Router)

export default new Router({
  scrollBehavior (to, from, savedPosition) {
    return { x: 0, y: 0 }
  },
  routes: [
    {
      path: '/',
      redirect: to => {
        if (store.state.book) {
          return {name: 'Bible', params: {book: store.state.book, chapter: store.state.chapter || 1}}
        }
        return {name: 'Bible', params: {book: 'Gen', chapter: '1'}}
      }
    },
    {
      path: '/bible/:book',
      redirect: to => to.path + '/1'
    },
    {
      path: '/bible/:book/:chapter',
      name: 'Bible',
      component: Bible
    },
    {
      path: '/picker',
      name: 'Picker',
      component: Picker
    },
    {
      path: '/history',
      name: 'History',
      component: History
    }
  ]
})

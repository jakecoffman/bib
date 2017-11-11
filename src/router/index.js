import Vue from 'vue'
import Router from 'vue-router'
import Bible from '@/components/Bible'
import Picker from '@/components/Picker'

Vue.use(Router)

export default new Router({
  scrollBehavior (to, from, savedPosition) {
    return { x: 0, y: 0 }
  },
  routes: [
    {
      path: '/',
      redirect: {name: 'Bible', params: {book: 'Gen', chapter: '1'}}
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
    }
  ]
})

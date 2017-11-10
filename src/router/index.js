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
      name: 'Bible',
      component: Bible
    },
    {
      path: '/:book',
      name: 'Book',
      component: Bible
    },
    {
      path: '/:book/:chapter',
      name: 'Chapter',
      component: Bible
    },
    {
      path: '/picker',
      name: 'Picker',
      component: Picker
    }
  ]
})

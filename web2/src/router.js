import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
    },
    {
      path: '/signin',
      name: 'signin',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/Signin.vue')
    },
    {
      path: '/vote',
      name: 'vote',
      component: () => import('./views/Gvote.vue')
    },
    {
      path: '/writer',
      name: 'writer',
      component: () => import('./views/Writer.vue')
    },
    {
      path: '/content/:cat/:id',
      name: 'content',
      component: () => import('./views/Garticle.vue'),
      props: {
        cat: {
          type: String,
          required: true
        },
        id: {
          type: Number,
          required: true
        }
      }
    },
    {
      path: '/profile/:id',
      name: 'profile',
      component: () => import('./views/Gprofile.vue'),
      props: {
        id: {
          type: Number,
          required: false
        }
      }
    },
  ]
})

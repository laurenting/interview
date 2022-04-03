import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Interviewer from '~/views/InterviewerHome/index.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  },
  {
    path: '/interviewer',
    component: Interviewer,
    children: [
      {
        path: 'create_questionnaire',
        name: 'CreateQuestionnaire',
        component: () => import('~/views/InterviewerHome/CreateQuestionnaire/index.vue')
      },
      {
        path: 'view_questionnaire',
        name: 'view_questionnaire',
        component: () => import('~/views/InterviewerHome/ViewQuestionnaire/index.vue')
      }
    ]
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router

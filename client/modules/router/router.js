import { createRouter, createWebHistory } from 'vue-router';

import SideBar from 'components/SideBar.vue';
import CardCollection from 'views/CardCollection.vue';
import CardEditor from 'views/CardEditor.vue';
import Home from 'views/Home.vue';


export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      components: {
        content: Home,
        sidebar: SideBar,
      },
    },
    {
      path: '/editor',
      name: 'editor',
      components: {
        content: CardEditor,
        sidebar: SideBar,
      },
    },
    {
      path: '/collection',
      name: 'collection',
      components: {
        content: CardCollection,
        sidebar: SideBar,
      },
    },
  ],
});

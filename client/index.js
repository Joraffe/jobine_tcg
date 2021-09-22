import { createApp } from 'vue';
import App from 'App.vue';

import { testLog } from 'helpers/logging';
import { router } from 'router/router';

testLog('Webpack JS loaded!');

const app = createApp(App);

app.use(router);

app.mount('#app');

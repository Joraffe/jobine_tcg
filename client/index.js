import { createApp } from 'vue';
import App from 'App.vue';

import { testLog } from 'helpers/logging';

testLog('Webpack JS loaded!');

createApp(App).mount('#app');

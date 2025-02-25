import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import marked from 'marked';
const markedMixin = {
    methods: {
         md: function (input) {
            return marked (input);
        },
    },
};

const app = createApp(App).mixin(markedMixin)

app.use(store)
// Initialize auth state
store.dispatch('auth/initialize')
app.use(router)
app.use(ElementPlus)

app.mount('#app')

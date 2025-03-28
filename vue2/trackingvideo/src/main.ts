import Vue from 'vue'
import App from './App.vue'
import VueYouTube from 'vue-youtube';




Vue.config.productionTip = false;

// Đăng ký plugin cho Vue 2
Vue.use(VueYouTube);

new Vue({
  render: h => h(App),
})
.$mount('#app')

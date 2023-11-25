import { createApp } from "vue";
import App from "./App.vue";

import { createPinia } from "pinia";
import router from "./router";

import "bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";

const pinia = createPinia();
createApp(App).use(pinia).use(router).mount("#app");

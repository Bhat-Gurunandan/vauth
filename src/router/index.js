import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import AboutView from "@/views/AboutView.vue";
import LoginView from "@/views/LoginView.vue";
import { useAuthStore } from "@/store/auth.js";

const routes = [
    {
        path: "/",
        name: "home",
        component: HomeView,
        meta: { requireAuth: true },
    },
    {
        path: "/about",
        name: "about",
        component: AboutView,
        meta: { requireAuth: true },
    },
    {
        path: "/login",
        name: "login",
        component: LoginView,
        meta: { requireAuth: false },
    },
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
});

router.beforeEach(async (to, from) => {
    console.log("To: ", to, "\nFrom: ", from);
    let authStore = useAuthStore();
    console.log(authStore.currentUser);
    if (to.meta.requireAuth) {
        if (!authStore.currentUser && to.name !== "login") {
            return { name: "login" };
        }
    }
});

export default router;

import { defineStore } from "pinia";

export const useAuthStore = defineStore("Auth", {
    state: () => ({
        apiHost: process.env.VUE_APP_DBAPI_ROOT,
        currentUser: null,
    }),
    actions: {},
});

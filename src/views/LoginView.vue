<template>
    <div class="row">
        <div class="col-md-4 offset-md-4 text-start">
            <form @submit.prevent="doLogin()">
                <div class="form-group mb-3">
                    <label for="email"><strong>Email</strong></label>
                    <input
                        type="email"
                        class="form-control"
                        id="email"
                        v-model="user.email"
                        aria-describedby="emailHelp"
                        placeholder="Enter email"
                    />
                    <small id="emailHelp" class="form-text text-muted"
                        >We'll never share your email with anyone else.</small
                    >
                </div>
                <div class="form-group mb-3">
                    <label for="password"><strong>Password</strong></label>
                    <input
                        type="password"
                        class="form-control"
                        id="password"
                        v-model="user.password"
                        placeholder="Password"
                    />
                </div>
                <div class="form-group form-check mb-3">
                    <input
                        type="checkbox"
                        class="form-check-input"
                        id="rememberMe"
                        v-model="user.rememberMe"
                    />
                    <label class="form-check-label" for="rememberme"
                        ><strong>Remember Me</strong></label
                    >
                </div>
                <div class="text-center">
                    <button type="submit" class="btn btn-primary">
                        Submit
                    </button>
                </div>
            </form>
        </div>
    </div>
</template>

<script>
export default {
    name: "LoginView",
    data() {
        return {
            user: {},
        };
    },
    methods: {
        doLogin() {
            console.log(this.user);
        },
    },
};
</script>

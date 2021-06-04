<template>
    <v-app>
        <Header app />
        <v-main>
            <router-view />
        </v-main>
        <Footer app />
    </v-app>
</template>

<script>
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";

export default {
    name: "App",
    components: {
        Header,
        Footer,
    },
};
</script>
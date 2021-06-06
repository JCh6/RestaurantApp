<template>
    <v-footer app dark padless>
        <v-card class="flex" tile color="white">
            <v-card-title>
                <v-spacer></v-spacer>
                <v-btn
                    v-for="(image, index) of images"
                    :key="index"
                    :href="image.link"
                    :target="image.target"
                    class="ml-3"
                    small
                    icon
                >
                    <v-img
                        width="20"
                        :src="require(`@/assets/${image.name}`)"
                    />
                </v-btn>
                <v-spacer></v-spacer>
            </v-card-title>
        </v-card>
    </v-footer>
</template>

<script>
export default {
    name: "Footer",
    data: () => ({
        images: [
            {
                link: "https://github.com/JCh6/RestaurantApp",
                name: "github.svg",
                target: "_blank",
            },
            {
                link: "https://www.linkedin.com/in/c95/",
                name: "linkedin.svg",
                target: "_blank",
            },
        ],
    }),
};
</script>

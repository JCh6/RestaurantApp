<template>
    <v-data-table
        hide-default-footer
        class="elevation-2"
        :headers="header"
        :items="data"
        @click:row="handleClick"
    ></v-data-table>
</template>

<script>
export default {
    name: "DataTable",
    props: {
        header: Array,
        data: Array,
        handleClick: Function,
    },
};
</script>

<style>
.v-data-table__wrapper:hover {
    cursor: pointer;
}
</style>
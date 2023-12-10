<template>
    <a-layout class="min-h-screen">
        <Header></Header>
        <a-layout-content class="flex justify-between max-w-7xl w-full min-h-max text-slate-950  bg-white" style="margin: 25px  auto;">
            <a-menu v-model:openKeys="openKeys" v-model:selectedKeys="selectedKeys" class="" style="width: 216px;"
                mode="inline" :items="items" @click="handleClick" />
                <div class="p-6 w-full">
                    <slot />
                </div>
        </a-layout-content>
        <a-layout-footer class="flex justify-center">
            footer
        </a-layout-footer>
    </a-layout>
</template>
<script lang="ts" setup>
import { h, ref } from 'vue';
import {
    AppstoreAddOutlined,
    SettingOutlined,
} from '@ant-design/icons-vue';
import type { MenuProps } from 'ant-design-vue';

const selectedKeys = ref([]);
const openKeys = ref([]);
const items = ref([
    {
        key: 'setting',
        icon: () => h(SettingOutlined),
        label: 'Setting',
        title: 'Developer Setting',
    },
    {
        key: 'spora',
        icon: () => h(AppstoreAddOutlined),
        label: 'Spora',
        title: 'Create New App',
    },
]);
const handleClick: MenuProps['onClick'] = e => {
    switch (e.key) {
        case "spora":
            navigateTo("/developer/spora")
            break;
        default:
            navigateTo("/developer")
            break;
    }
};
</script>
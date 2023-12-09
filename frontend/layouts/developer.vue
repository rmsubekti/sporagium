<template>
    <a-layout class="min-h-screen">
        <a-layout-header style="background-color:rgb(164, 177, 255)adding: 0;padding: 0 20px;">
            <div class="flex max-w-7xl align-middle justify-between h-full" style="margin: 0 auto;">
                <div></div>
                <a-dropdown-button :trigger="['click']" style="display: inline;">
                    Dropdown
                    <template #overlay>
                        <a-menu @click="">
                            <a-menu-item key="1">
                                <UserOutlined />
                                1st menu item
                            </a-menu-item>
                            <a-menu-item key="2">
                                <UserOutlined />
                                2nd menu item
                            </a-menu-item>
                            <a-menu-item key="3">
                                <UserOutlined />
                                3rd item
                            </a-menu-item>
                        </a-menu>
                    </template>
                    <template #icon>
                        <UserOutlined />
                    </template>
                </a-dropdown-button>
            </div>
        </a-layout-header>
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
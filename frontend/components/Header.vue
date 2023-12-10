<template>
    <a-layout-header style="padding: 0;padding: 0 20px;">
        <div class="flex max-w-7xl align-middle justify-between h-full" style="margin: 0 auto;">
            <div></div>
            <a-dropdown-button :trigger="['click']" style="display: inline;">
                <span v-if="userName">{{ userName }}</span>
                <span v-else> Guest </span>
                <template #overlay>
                    <a-menu @click="handleMenuClick">
                        <a-menu-item v-if="!authenticated" key="login">
                            <UserOutlined />
                            Login
                        </a-menu-item>
                        <a-menu-item v-if="authenticated" key="logout">
                            <UserOutlined />
                            Log Out
                        </a-menu-item>
                    </a-menu>
                </template>
                <template #icon>
                    <UserOutlined />
                </template>
            </a-dropdown-button>
        </div>
    </a-layout-header>
</template>

<script setup lang="ts">
import type { MenuProps } from 'ant-design-vue';
const authenticated = !!localStorage.getItem('nuxt-cred');
const userName = localStorage.getItem('nuxt-uname');

const handleMenuClick: MenuProps['onClick'] = e => {
  switch (e.key) {
    case "login":
        navigateTo("/login")
        break;
    case "logout":
        handleLogout()
        break;
    default:
        break;
  }
};

const handleLogout=() =>{
    localStorage.removeItem('nuxt-cred');
    localStorage.removeItem('nuxt-uname');
    localStorage.removeItem('nuxt-uid');
    location.reload()
}
</script>
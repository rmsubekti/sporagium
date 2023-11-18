<template>
  <a-layout class="min-h-screen">
    <a-layout-sider v-model:collapsed="collapsed" :trigger="null" collapsible>
      <div class="logo" />
      <a-menu v-model:selectedKeys="selectedNavMenu" theme="dark" mode="inline">
        <a-menu-item key="1">
          <user-outlined />
          <span>
            <NuxtLink to="/dashboard/user">User</NuxtLink>
          </span>
        </a-menu-item>
        <a-menu-item key="2">
          <video-camera-outlined />
          <span>
            <NuxtLink to="/dashboard/spora">Spora</NuxtLink>
          </span>
        </a-menu-item>
        <a-menu-item key="3">
          <upload-outlined />
          <span>nav 3</span>
        </a-menu-item>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header style="background: #fff; padding: 0">
        <menu-unfold-outlined v-if="collapsed" class="trigger px-6" @click="() => (collapsed = !collapsed)" />
        <menu-fold-outlined v-else class="trigger px-6" @click="() => (collapsed = !collapsed)" />
      </a-layout-header>
      <a-layout-content class="mx-6 my-4 p-6 bg-white min-h-max">
        <slot />
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>
<script lang="ts" setup>
const collapsed = ref<boolean>(false);
</script>
<style></style>
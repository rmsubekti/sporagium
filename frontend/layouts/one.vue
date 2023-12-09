<template>
    <a-layout class="min-h-screen">
        <a-layout-header style="background-color:cornsilk;padding: 0;padding: 0 20px;">
            <div class="flex max-w-4xl align-middle justify-between h-full" style="margin: 0 auto;">
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
        <a-layout-content class="max-w-4xl w-full min-h-max text-slate-950  bg-white" style="margin: 25px  auto;">
            <slot />
        </a-layout-content>
        <a-layout-footer class="flex justify-center">
            footer
        </a-layout-footer>
    </a-layout>
</template>
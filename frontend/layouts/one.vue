<template>
    <a-layout class="min-h-screen">
        <Header></Header>
        <a-layout-content class="max-w-4xl w-full min-h-max text-slate-950  bg-white" style="margin: 25px  auto;">
            <slot />
        </a-layout-content>
        <a-layout-footer class="flex justify-center">
            footer
        </a-layout-footer>
    </a-layout>
</template>
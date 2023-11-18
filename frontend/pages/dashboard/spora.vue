<template>
    <Head>
        <title>Spora App</title>
    </Head>
    <NuxtLayout name="dashboard">
        Spora Dashboard
        <NewSpora @submitted="() => { refresh() }" />
        <a-table :columns="columns" :data-source="spora?.data" :expand-column-width="100">
            <template #bodyCell="{ column, record }">
                <template v-if="column.key === 'action'">
                    <a-popconfirm title="Are you sure？" @confirm="confirm">
                        <template #icon><question-circle-outlined style="color: red" /></template>
                        <a href="#">Delete</a>
                    </a-popconfirm>
                    <a-popconfirm title="Are you sure？" @confirm="generateSecret(record.id)">
                        <template #icon><question-circle-outlined style="color: red" /></template>
                        <a>Create Secret</a>
                    </a-popconfirm>
                </template>
            </template>
            <template #expandedRowRender="{ column,record }">
                <a-table :columns="innerColumn" :data-source="record.clients">
                </a-table>
            </template>
            <template #expandColumnTitle>
                <span style="color: red">More</span>
            </template>
        </a-table>
    </NuxtLayout>
</template>

<script lang="ts" setup>
const config = useRuntimeConfig();
const columns = [
    { title: 'Name', dataIndex: 'name', key: 'name', fixed: true },
    { title: 'Homepage', dataIndex: 'homepage', key: 'homepage' },
    { title: 'Description', dataIndex: 'description', key: 'description' },
    { title: 'Action', key: 'action' },
];
const innerColumn = [
    { title: 'ID', dataIndex: 'id', key: 'secret' },
    { title: 'Secret', dataIndex: 'secret', key: 'secret' },
    { title: 'Callback', dataIndex: 'domain', key: 'domain' },
];
const jwtToken = localStorage.getItem('nuxt-cred');
const { data: spora, error, refresh } = await useFetch(`${config.public.apiBase}/spora`, {
    headers: { Authorization: `Bearer ${jwtToken}` },
});

const confirm = (e: MouseEvent) => {
    console.log(e);
};
const generateSecret = async (id: string) => {
    const { data, error } = await useFetch(`${config.public.apiBase}/spora/${id}`, {
        headers: { Authorization: `Bearer ${jwtToken}` },
        onRequest({ request, options }) {
            options.method = "PATCH";
        },
        onResponse({ response }) {
            refresh()
        },
        onResponseError({ response }) {
            console.log(response?._data);
        },
    });
};
</script>
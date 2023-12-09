<template>
    <NuxtLayout name="developer">
        <a-form :model="sporaFormState" :label-col="labelCol" :wrapper-col="wrapperCol" @finish="onFinish"
            @finishFailed="onFinishFailed">
            <a-form-item label="Application name">
                <a-input class="w-full" v-model:value="sporaFormState.name" :rules="[{ required: true, message: 'Name is required!' }]" />
            </a-form-item>
            <a-form-item label="Homepage">
                <a-input v-model:value="sporaFormState.homepage" />
            </a-form-item>
            <a-form-item label="Callback Url">
                <a-input v-model:value="sporaFormState.callback_url" />
            </a-form-item>
            <a-form-item label="Description">
                <a-textarea v-model:value="sporaFormState.description" />
            </a-form-item>
            <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
                <a-button html-type="submit" type="primary" class="bg-blue-700 ">Create</a-button>
            </a-form-item>
        </a-form>
    </NuxtLayout>
</template>
<script lang="ts" setup>
import { reactive, toRaw } from 'vue';
const config = useRuntimeConfig();
interface SporaFormState {
    name: string;
    homepage: string;
    callback_url: string;
    description: string;
}
const sporaFormState = reactive<SporaFormState>({
    name: "",
    homepage: "",
    callback_url: "",
    description: "",
});

const onFinish = async (values: any) => {
    const jwtToken = localStorage.getItem('nuxt-cred');
    const { data, error } = await useFetch(`${config.public.apiBase}/spora`, {
        headers: { Authorization: `Bearer ${jwtToken}` },
        onRequest({ request, options }) {
            options.method = "POST";
            options.body = toRaw(sporaFormState);
        },
        onResponse({ response }) {
            if (response.status == 200) {
            }
        },
        onResponseError({ response }) {
            console.log(response?._data);
        },
    });
};
const onFinishFailed = async (errorInfo: any) => {
    console.log("Failed:", errorInfo);
};
const labelCol = { style: { width: '150px' } };
const wrapperCol = { span: 14 };
</script>
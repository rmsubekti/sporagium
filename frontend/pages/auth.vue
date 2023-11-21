<template>
    <Head>
        <title>Login</title>
    </Head>
    <NuxtLayout name="blank">
            <!-- <a-form-item placeholder="Username" name="email"
                :rules="[{ required: true, message: 'Please input your username or email!' }]">
                <a-input v-model:value="formState.email">
                    <template #prefix>
                        <UserOutlined class="site-form-item-icon" />
                    </template>
                </a-input>
            </a-form-item>

            <a-form-item placeholder="Password" name="password"
                :rules="[{ required: true, message: 'Please input your password!' }]">
                <a-input-password v-model:value="formState.password">
                    <template #prefix>
                        <LockOutlined class="site-form-item-icon" />
                    </template>
                </a-input-password>
            </a-form-item>

            <a-form-item>
                <div class="flex justify-between">
                    <a-form-item name="remember" no-style>
                        <a-checkbox v-model:checked="formState.remember">Remember me</a-checkbox>
                    </a-form-item>
                    <a class="login-form-forgot" href="">Forgot password</a>
                </div>
            </a-form-item> -->

            <a href="/o/authorize">Authorize</a>

    </NuxtLayout>
</template>
<script lang="ts" setup>
// interface FormState {
//     email: string;
//     password: string;
//     remember: boolean;
// }
const config = useRuntimeConfig();
// const formState = reactive<FormState>({
//     email: '',
//     password: '',
//     remember: true,
// });
const onFinish = async (values: any) => {
    const { data, error } = await useFetch(`${config.public.apiBase}/login`, {
        onRequest({ request, options }) {
            options.method = "POST";
            options.body = values;
        },
        onResponse({ response }) {
            if (response.status === 200) {
                localStorage.setItem('nuxt-cred', response._data.data.token);
                navigateTo('/dashboard')
            }
        },
        onResponseError({ response }) {
            console.log(response?._data);
        },
    });
};

const auth = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
};
</script>
<style scoped></style>
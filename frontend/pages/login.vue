<template>
    <Head>
        <title>Login</title>
    </Head>
    <NuxtLayout name="blank">
        <a-form :model="formState" name="login" @finish="onFinish" @finishFailed="onFinishFailed">
            <a-form-item placeholder="Username" name="email"
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
            </a-form-item>

            <a-form-item>
                <a-button type="primary" html-type="submit" class="bg-blue-700">
                    Log in
                </a-button>
                Or
                <NuxtLink to="/register">register now!</NuxtLink>
            </a-form-item>
        </a-form>
    </NuxtLayout>
</template>
<script lang="ts" setup>
interface FormState {
    email: string;
    password: string;
    remember: boolean;
}
const config = useRuntimeConfig();
const formState = reactive<FormState>({
    email: '',
    password: '',
    remember: true,
});
const onFinish = async (values: any) => {
    const { data, error } = await useFetch(`${config.public.apiBase}/login`, {
        onRequest({ request, options }) {
            options.method = "POST";
            options.body = values;
        },
        onResponse({ response }) {
            if (response.status === 200) {
                localStorage.setItem('nuxt-cred', response._data.data.token);
                localStorage.setItem('nuxt-uname', response._data.data.name);
                localStorage.setItem('nuxt-uid', response._data.data.id);
                navigateTo('/')
            }else if(response.status === 307 )(
                navigateTo(response._data.data)
            )
        },
        onResponseError({ response }) {
            console.log(response?._data);
        },
    });
};

const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
};
</script>
<style scoped></style>
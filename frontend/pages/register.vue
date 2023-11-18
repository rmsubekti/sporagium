<template>
  <Head>
    <title>Register</title>
  </Head>
  <NuxtLayout name="blank">
    <a-form :model="registerFormState" name="register" style="max-width: 500px" :label-col="labelCol" :wrapper-col="wrapperCol"
      @finish="onFinish" @finishFailed="onFinishFailed">
      <a-form-item label="Full Name" name="name" :rules="[{ required: true, message: 'Full Name is required!' }]">
        <a-input v-model:value="registerFormState.name" />
      </a-form-item>

      <a-form-item label="Username" name="username">
        <a-input v-model:value="registerFormState.username" />
      </a-form-item>

      <a-form-item label="Phone" name="phone">
        <a-input v-model:value="registerFormState.phone" />
      </a-form-item>

      <a-form-item label="Email" name="email" :rules="[{ required: true, message: 'Required!' }]">
        <a-input v-model:value="registerFormState.email" />
      </a-form-item>

      <a-form-item label="Password" name="password" :rules="[{ required: true, message: 'Required!' }]">
        <a-input-password v-model:value="registerFormState.password" />
      </a-form-item>

      <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
        <a-button type="primary" html-type="submit" class="bg-blue-700 ">
          Register
        </a-button>
        Or
        <NuxtLink to="/login">login now!</NuxtLink>
      </a-form-item>
    </a-form>
  </NuxtLayout>
</template>

<script lang="ts" setup>
import { reactive, toRaw } from 'vue';
const config = useRuntimeConfig();
const labelCol = { style: { width: '150px' } };
const wrapperCol = { span: 14 };

interface RegisterFormState {
  name: string;
  username: string;
  email: string;
  phone: string;
  password: string;
}

const registerFormState = reactive<RegisterFormState>({
  name: "",
  username: "",
  email: "",
  phone: "",
  password: "",
});

const onFinish = async (values: any) => {
  const { data, error } = await useFetch(`${config.public.apiBase}/register`, {
    onRequest({ request, options }) {
      options.method = "POST";
      options.body = values;
    },
    onResponse({ response }) {
      setTimeout(() => {
        navigateTo("/login")
      }, 1000 * 7);
    },
    onResponseError({ response }) {
      console.log(response?._data);
    },
  });
};

const onFinishFailed = async (errorInfo: any) => {
  console.log("Failed:", errorInfo);
};
</script>
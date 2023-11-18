// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: {
    enabled: true,

    timeline: {
      enabled: true,
    },
  },
  ssr: false,
  pages: true,
  modules: [
    '@ant-design-vue/nuxt',
    '@nuxt/ui'
  ],
  runtimeConfig: {
    public: {
      apiBase: "http://localhost:3030/v1"
    }
  },
})

export default defineNuxtRouteMiddleware((to, from) => {
  const authenticated = !!localStorage.getItem('nuxt-cred')
    if (to.path.includes("/developer") && !authenticated) {
      return navigateTo('/login')
    }
  })
  
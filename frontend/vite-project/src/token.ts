import { ref } from 'vue';

export const token = ref<string | null>(null);

export const setToken = (newToken: string) => {
  token.value = newToken;
};
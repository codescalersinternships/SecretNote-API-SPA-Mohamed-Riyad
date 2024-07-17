import { ref } from 'vue';

export const userId = ref<number | null>(null);

export const setUserId = (id: number) => {
  userId.value = id;
};
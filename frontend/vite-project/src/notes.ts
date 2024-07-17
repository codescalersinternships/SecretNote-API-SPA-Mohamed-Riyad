import { ref } from 'vue';

export const notes = ref<any[]>([]);

export const setNotes = (newNotes: any[]) => {
  notes.value = newNotes;
};
import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080';




export const getNote = async (noteId: number) => {
  try {
    const response = await axios.get(`${API_BASE_URL}/get-note/${noteId}`);
    return response.data;
  } catch (error) {
    console.error('Error getting note:', error);
    throw error;
  }
};

export const getAllNotes = async (userId: number) => {
  try {
    const response = await axios.get(`${API_BASE_URL}/get-all-notes/${userId}`);
    return response.data;
  } catch (error) {
    console.error('Error getting all notes:', error);
    throw error;
  }
};

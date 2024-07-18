import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080';




export const getNote = async (noteId: number, token: string) => {
  try {
    const response = await axios.get(`${API_BASE_URL}/get-note/${noteId}`,{
      headers: {
        Authorization: `Bearer ${token}`
      }
    });
    return response.data;
  } catch (error) {
    console.error('Error getting note:', error);
    throw error;
  }
};

export const getAllNotes = async (token: string) => {
  try {
    const response = await axios.get(`${API_BASE_URL}/get-all-notes`,{
      headers: {
        Authorization: `Bearer ${token}`
      }
    });
    return response.data;
  } catch (error) {
    console.error('Error getting all notes:', error);
    throw error;
  }
};

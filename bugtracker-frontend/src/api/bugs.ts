import { Bug } from '../types/bug';

const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';
const API_PATH = '/api';

export const getBugs = async (): Promise<Bug[]> => {
    try {
        const response = await fetch(`${API_BASE_URL}${API_PATH}/bugs`);
        if (!response.ok) {
            throw new Error('Failed to fetch bugs');
        }
        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching bugs:', error);
        throw error;
    }
};

export const createBug = async (bugData: Omit<Bug, 'id'>) => {
    try {
        const response = await fetch(`${API_BASE_URL}${API_PATH}/bugs`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(bugData),
        });
        if (!response.ok) {
            throw new Error('Failed to create bug');
        }
        return await response.json();
    } catch (error) {
        console.error('Error creating bug:', error);
        throw error;
    }
};

export const updateBug = async (id: string, bugData: Partial<Bug>) => {
    try {
        const response = await fetch(`${API_BASE_URL}${API_PATH}/bugs/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(bugData),
        });
        if (!response.ok) {
            throw new Error('Failed to update bug');
        }
        return await response.json();
    } catch (error) {
        console.error('Error updating bug:', error);
        throw error;
    }
};

export const deleteBug = async (id: string) => {
    try {
        const response = await fetch(`${API_BASE_URL}${API_PATH}/bugs/${id}`, {
            method: 'DELETE',
        });
        if (!response.ok) {
            throw new Error('Failed to delete bug');
        }
        return;
    } catch (error) {
        console.error('Error deleting bug:', error);
        throw error;
    }
}; 
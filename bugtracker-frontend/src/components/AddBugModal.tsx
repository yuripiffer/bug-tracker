import { useState } from 'react';
import { Bug, Priority } from '@/types/bug';

interface AddBugModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSubmit: (bug: Omit<Bug, 'id' | 'status'>) => void;
}

export default function AddBugModal({ isOpen, onClose, onSubmit }: AddBugModalProps) {
  const [formData, setFormData] = useState<Omit<Bug, 'id' | 'status'>>({
    title: '',
    description: '',
    priority: 'Medium'
  });

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-gray-600 bg-opacity-50 flex items-center justify-center">
      <div className="bg-white rounded-lg p-8 max-w-md w-full">
        <h2 className="text-xl font-bold mb-4">Add New Bug</h2>
        <form onSubmit={(e) => {
          e.preventDefault();
          onSubmit(formData);
          onClose();
        }}>
          <div className="mb-4">
            <label className="block text-gray-700 text-sm font-bold mb-2">
              Title
            </label>
            <input
              type="text"
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700"
              value={formData.title}
              onChange={(e) => setFormData({ ...formData, title: e.target.value })}
              required
            />
          </div>

          <div className="mb-4">
            <label className="block text-gray-700 text-sm font-bold mb-2">
              Description
            </label>
            <textarea
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700"
              value={formData.description}
              onChange={(e) => setFormData({ ...formData, description: e.target.value })}
              required
            />
          </div>

          <div className="mb-4">
            <label className="block text-gray-700 text-sm font-bold mb-2">
              Priority
            </label>
            <select
              className="shadow border rounded w-full py-2 px-3 text-gray-700"
              value={formData.priority}
              onChange={(e) => {
                const value = e.target.value as Priority;
                setFormData({ ...formData, priority: value });
              }}
            >
              <option value="Low">Low</option>
              <option value="Medium">Medium</option>
              <option value="High">High</option>
            </select>
          </div>

          <div className="flex justify-end gap-4">
            <button
              type="button"
              onClick={onClose}
              className="bg-gray-500 hover:bg-gray-600 text-white px-4 py-2 rounded"
            >
              Cancel
            </button>
            <button
              type="submit"
              className="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded"
            >
              Add Bug
            </button>
          </div>
        </form>
      </div>
    </div>
  );
} 
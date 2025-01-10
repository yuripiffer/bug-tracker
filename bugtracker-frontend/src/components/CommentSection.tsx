import { useState } from 'react';
import { Comment } from '@/types/comment';

interface CommentSectionProps {
    bugId: number;
    comments: Comment[];
    onCommentAdded: () => void;
}

export default function CommentSection({ bugId, comments = [], onCommentAdded }: CommentSectionProps) {
    const [author, setAuthor] = useState('');
    const [content, setContent] = useState('');
    const [isSubmitting, setIsSubmitting] = useState(false);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setIsSubmitting(true);

        try {
            const response = await fetch(`http://localhost:8080/api/bugs/${bugId}/comments`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ author, content }),
            });

            if (!response.ok) {
                throw new Error('Failed to add comment');
            }

            setAuthor('');
            setContent('');
            onCommentAdded();
        } catch (error) {
            console.error('Error adding comment:', error);
        } finally {
            setIsSubmitting(false);
        }
    };

    return (
        <div className="mt-8">
            <h2 className="text-xl font-semibold mb-4">Comments</h2>
            
            <div className="mb-6">
                <form onSubmit={handleSubmit} className="space-y-4">
                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Your Name
                        </label>
                        <input
                            type="text"
                            value={author}
                            onChange={(e) => setAuthor(e.target.value)}
                            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                            required
                        />
                    </div>
                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Comment
                        </label>
                        <textarea
                            value={content}
                            onChange={(e) => setContent(e.target.value)}
                            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                            rows={3}
                            required
                        />
                    </div>
                    <button
                        type="submit"
                        disabled={isSubmitting}
                        className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 disabled:opacity-50"
                    >
                        {isSubmitting ? 'Adding...' : 'Add Comment'}
                    </button>
                </form>
            </div>

            <div className="space-y-4">
                {(comments || []).map((comment) => (
                    <div key={comment.id} className="bg-gray-50 p-4 rounded-lg">
                        <div className="flex justify-between items-start">
                            <span className="font-medium">{comment.author}</span>
                            <span className="text-sm text-gray-500">
                                {new Date(comment.createdAt).toLocaleString()}
                            </span>
                        </div>
                        <p className="mt-2 text-gray-700">{comment.content}</p>
                    </div>
                ))}
            </div>
        </div>
    );
} 
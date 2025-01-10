import React, { useEffect, useState } from 'react';
import { getBugs, createBug } from '../api/bugs';
import AddBugModal from './AddBugModal';
import DeleteBugButton from './DeleteBugButton';
import LoadingScreen from './LoadingScreen';
import { Bug } from '../types/bug';
import Link from 'next/link';

const BugList: React.FC = () => {
    const [bugs, setBugs] = useState<Bug[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [showLoadingScreen, setShowLoadingScreen] = useState(true);
    const [fadeIn, setFadeIn] = useState(false);

    useEffect(() => {
        const fetchBugs = async () => {
            try {
                const data = await getBugs();
                setBugs(data);
                setLoading(false);
            } catch (error) {
                console.error('Error fetching bugs:', error);
                setError('Failed to fetch bugs');
                setLoading(false);
            }
        };

        fetchBugs();
    }, []);

    const handleLoadingComplete = () => {
        setShowLoadingScreen(false);
        setFadeIn(true);
    };

    const handleAddBug = async (newBug: Omit<Bug, 'id' | 'status'>) => {
        try {
            await createBug({ ...newBug, status: 'Open' });
            const updatedBugs = await getBugs();
            setBugs(updatedBugs);
            setIsModalOpen(false);
        } catch (error) {
            console.error('Failed to create bug:', error);
        }
    };

    const handleBugDeleted = async () => {
        try {
            const updatedBugs = await getBugs();
            setBugs(updatedBugs);
        } catch (error) {
            console.error('Error refreshing bugs:', error);
            setError('Failed to refresh bugs');
        }
    };

    if (showLoadingScreen) {
        return <LoadingScreen onLoadingComplete={handleLoadingComplete} />;
    }

    if (loading) return <div className="text-center p-4">Loading...</div>;
    if (error) return <div className="text-center p-4 text-red-500">Error: {error}</div>;

    return (
        <div className={`min-h-screen bg-gray-100 transition-opacity duration-2000 ${fadeIn ? 'opacity-100' : 'opacity-0'}`}>
            <nav className="bg-white shadow-lg">
                <div className="max-w-7xl mx-auto px-4 py-4">
                    <h1 className="text-2xl font-bold text-gray-800">Bug Tracker</h1>
                </div>
            </nav>

            <main className="max-w-7xl mx-auto px-4 py-8">
                <div className="flex justify-between items-center mb-6">
                    <h2 className="text-xl font-semibold text-gray-700">All Bugs</h2>
                    <button 
                        onClick={() => setIsModalOpen(true)}
                        className="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-md"
                    >
                        Add New Bug
                    </button>
                </div>

                <div className="bg-white shadow-md rounded-lg">
                    <table className="min-w-full">
                        <thead className="bg-gray-50">
                            <tr>
                                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">ID</th>
                                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Title</th>
                                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Status</th>
                                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Priority</th>
                                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
                            </tr>
                        </thead>
                        <tbody className="bg-white divide-y divide-gray-200">
                            {(bugs || []).map((bug) => (
                                <tr key={bug.id}>
                                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                                        <Link href={`/bugs/${bug.id}`} className="text-blue-600 hover:text-blue-900">
                                            {bug.id}
                                        </Link>
                                    </td>
                                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                                        <Link href={`/bugs/${bug.id}`} className="text-blue-600 hover:text-blue-900">
                                            {bug.title}
                                        </Link>
                                    </td>
                                    <td className="px-6 py-4 whitespace-nowrap">
                                        <span className={`px-2 inline-flex text-xs leading-5 font-semibold rounded-full
                                            ${bug.status === 'Open' ? 'bg-red-100 text-red-800' : 
                                            bug.status === 'In Progress' ? 'bg-yellow-100 text-yellow-800' : 
                                            'bg-green-100 text-green-800'}`}>
                                            {bug.status}
                                        </span>
                                    </td>
                                    <td className="px-6 py-4 whitespace-nowrap">
                                        <span className={`px-2 inline-flex text-xs leading-5 font-semibold rounded-full
                                            ${bug.priority === 'High' ? 'bg-red-100 text-red-800' : 
                                            bug.priority === 'Medium' ? 'bg-yellow-100 text-yellow-800' : 
                                            'bg-green-100 text-green-800'}`}>
                                            {bug.priority}
                                        </span>
                                    </td>
                                    <td className="px-6 py-4 whitespace-nowrap text-sm font-medium">
                                        <button className="text-blue-600 hover:text-blue-900 mr-4">Edit</button>
                                        <DeleteBugButton 
                                            bugId={bug.id}
                                            onDelete={handleBugDeleted}
                                        />
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            </main>

            <AddBugModal 
                isOpen={isModalOpen}
                onClose={() => setIsModalOpen(false)}
                onSubmit={handleAddBug}
            />
        </div>
    );
};

export default BugList; 
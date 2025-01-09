import { useState } from 'react';
import Head from 'next/head';
import AddBugModal from '@/components/AddBugModal';

interface Bug {
  id: string;
  title: string;
  description: string;
  status: 'Open' | 'In Progress' | 'Resolved';
  priority: 'Low' | 'Medium' | 'High';
}

export default function Home() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [bugs, setBugs] = useState<Bug[]>([
    {
      id: '1',
      title: 'Login page not responsive',
      description: 'The login page breaks on mobile devices',
      status: 'Open',
      priority: 'High'
    },
    // Sample data for demonstration
  ]);

  const handleAddBug = (newBug: Omit<Bug, 'id' | 'status'>) => {
    const bug = {
      ...newBug,
      id: (bugs.length + 1).toString(),
      status: 'Open' as const
    };
    setBugs([...bugs, bug]);
  }; 

  return (
    <div className="min-h-screen bg-gray-100">
      <Head>
        <title>Bug Tracker</title>
        <meta name="description" content="Track and manage software bugs" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

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
              {bugs.map((bug) => (
                <tr key={bug.id}>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{bug.id}</td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{bug.title}</td>
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
                    <button className="text-red-600 hover:text-red-900">Delete</button>
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
} 
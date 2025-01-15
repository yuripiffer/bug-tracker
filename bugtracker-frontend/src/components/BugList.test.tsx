import { render, screen, fireEvent, waitFor, act } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { useRouter } from 'next/router'
import BugList from './BugList'
import { getBugs, createBug, updateBug, deleteBug } from '../api/bugs'

// Mock next/router
jest.mock('next/router', () => ({
    useRouter: jest.fn()
}))

// Mock next/image
jest.mock('next/image', () => ({
    __esModule: true,
    default: ({ src, alt, ...props }: { src: string, alt: string }) => {
        // eslint-disable-next-line @next/next/no-img-element
        return <img src={src} alt={alt} {...props} />
    }
}))

// Mock the API calls
jest.mock('../api/bugs')

// Mock next/link
jest.mock('next/link', () => ({
    __esModule: true,
    default: ({ children, href }: { children: React.ReactNode, href: string }) => (
        <a href={href}>{children}</a>
    )
}))

describe('BugList', () => {
    const mockRouter = {
        query: {},
        pathname: '/',
        replace: jest.fn()
    }

    const mockBugs = [
        { id: 1, title: 'Bug 1', status: 'Open', priority: 'High', description: 'Test' },
        { id: 2, title: 'Bug 2', status: 'In Progress', priority: 'Medium', description: 'Test' }
    ]

    beforeEach(() => {
        (useRouter as jest.Mock).mockReturnValue(mockRouter);
        (getBugs as jest.Mock).mockResolvedValue(mockBugs);
        jest.clearAllMocks();
    });

    const waitForLoadingToComplete = async () => {
        await waitFor(() => {
            expect(screen.queryByTestId('loading-screen')).not.toBeInTheDocument()
        }, { timeout: 100 });
    }

    const renderAndWaitForData = async () => {
        const result = render(<BugList testMode={true} />);
        await waitFor(() => {
            expect(screen.queryByText('Loading...')).not.toBeInTheDocument();
        });
        // Wait for data to be loaded
        await waitFor(() => {
            expect(screen.getByText('All Bugs')).toBeInTheDocument();
        });
        return result;
    };

    describe('Initial Rendering', () => {
        it('should show loading screen initially', () => {
            render(<BugList />)
            expect(screen.getByTestId('loading-screen')).toBeInTheDocument()
        })

        it('should render bug list after loading', async () => {
            await renderAndWaitForData();
            expect(screen.getByText('All Bugs')).toBeInTheDocument()
        })

        it('should show error message if bugs fetch fails', async () => {
            (getBugs as jest.Mock).mockRejectedValue(new Error('Failed to fetch'))
            render(<BugList />)
            await waitForLoadingToComplete()
            expect(screen.getByText(/error/i)).toBeInTheDocument()
        })
    })

    describe('Bug Operations', () => {
        beforeEach(async () => {
            await renderAndWaitForData();
        })

        it('should open add bug modal when clicking Add New Bug', async () => {
            const addButton = screen.getByRole('button', { name: /add new bug/i });
            await act(async () => {
                fireEvent.click(addButton);
            });
            expect(screen.getByRole('heading', { name: /add new bug/i })).toBeInTheDocument();
        })

        it('should handle creating a new bug', async () => {
            const newBug = {
                id: 3,
                title: 'New Bug',
                status: 'Open',
                priority: 'High',
                description: 'Test'
            }
            ;(createBug as jest.Mock).mockResolvedValue(newBug)
            ;(getBugs as jest.Mock).mockResolvedValue([...mockBugs, newBug])

            await act(async () => {
                fireEvent.click(screen.getByText('Add New Bug'))
            })

            await act(async () => {
                await userEvent.type(screen.getByLabelText('Title'), 'New Bug')
                await userEvent.type(screen.getByLabelText('Description'), 'Test')
                fireEvent.submit(screen.getByRole('button', { name: /add bug/i }))
            })

            expect(createBug).toHaveBeenCalled()
            expect(mockRouter.replace).toHaveBeenCalledWith({
                pathname: '/',
                query: { createdBugTitle: 'New Bug', showCreateNotification: true }
            })
        })

        it('should handle editing a bug', async () => {
            const editButtons = screen.getAllByRole('button', { name: /edit/i });
            await act(async () => {
                fireEvent.click(editButtons[0]);
            });

            await act(async () => {
                const titleInput = screen.getByDisplayValue('Bug 1');
                await userEvent.clear(titleInput);
                await userEvent.type(titleInput, 'Updated Bug');
                const saveButton = screen.getByRole('button', { name: /save changes/i });
                fireEvent.click(saveButton);
            });

            expect(updateBug).toHaveBeenCalledWith('1', expect.objectContaining({ title: 'Updated Bug' }));
        })

        it('should handle deleting a bug', async () => {
            ;(deleteBug as jest.Mock).mockResolvedValue(undefined)
            ;(getBugs as jest.Mock).mockResolvedValue([mockBugs[1]])

            await act(async () => {
                fireEvent.click(screen.getAllByText('Delete')[0])
            })

            await act(async () => {
                const deleteButtons = screen.getAllByText('Delete')
                fireEvent.click(deleteButtons[deleteButtons.length - 1])
            })

            expect(deleteBug).toHaveBeenCalledWith('1')
            expect(mockRouter.replace).toHaveBeenCalledWith({
                pathname: '/',
                query: { deletedBugTitle: 'Bug 1', showDeleteNotification: true }
            })
        })
    })

    describe('Notifications', () => {
        it('should show success notification after creating bug', async () => {
            mockRouter.query = { showCreateNotification: 'true', createdBugTitle: 'New Bug' };
            await renderAndWaitForData();
            await waitFor(() => {
                expect(screen.getByText(/successfully created bug "New Bug"/i)).toBeInTheDocument();
            });
        })

        it('should show success notification after deleting bug', async () => {
            mockRouter.query = { showDeleteNotification: 'true', deletedBugTitle: 'Bug 1' }
            await renderAndWaitForData();
            expect(screen.getByText(/successfully deleted bug/i)).toBeInTheDocument()
        })
    })

    describe('Status and Priority Display', () => {
        beforeEach(async () => {
            await renderAndWaitForData();
        })

        it('should display correct status indicators', async () => {
            expect(screen.getByText('Open')).toHaveClass('bg-red-100')
            expect(screen.getByText('In Progress')).toHaveClass('bg-yellow-100')
        })

        it('should display correct priority indicators', async () => {
            expect(screen.getByText('High')).toHaveClass('bg-red-100')
            expect(screen.getByText('Medium')).toHaveClass('bg-yellow-100')
        })
    })
}) 
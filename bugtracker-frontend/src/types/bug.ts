export type Priority = 'Low' | 'Medium' | 'High';

export interface Bug {
  id: string;
  title: string;
  description: string;
  status: 'Open' | 'In Progress' | 'Resolved';
  priority: Priority;
}
"use client";

import { useEffect, useState } from 'react';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea'; // Assuming Textarea component
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Toaster, toast } from 'sonner';
import { useAuthStore } from '@/stores/auth-store';
import { cn } from '@/lib/utils'; // For conditional class names

interface Subject {
    id: string;
    name: string;
    code: string;
}

interface Staff {
    id: string;
    name: string;
}

interface Assignment {
    id: string;
    title: string;
    description?: string;
    assignmentType: string;
    dueDate: string; // ISO string
    status: string;
    priority: string;
    subjectId?: string;
    staffId?: string;
    // ... other fields as needed for display
}

export default function AssignmentsPage() {
    const { token } = useAuthStore();
    const [assignments, setAssignments] = useState<Assignment[]>([]);
    const [subjects, setSubjects] = useState<Subject[]>([]);
    const [staffMembers, setStaffMembers] = useState<Staff[]>([]);
    const [newAssignment, setNewAssignment] = useState({
        title: '',
        description: '',
        assignmentType: 'assignment',
        dueDate: '', // YYYY-MM-DDTHH:MM:SSZ
        priority: 'medium',
        subjectId: '',
        staffId: '',
    });
    const [isLoading, setIsLoading] = useState(false);

    useEffect(() => {
        if (token) {
            fetchAssignments();
            fetchDependencies();
        }
    }, [token]);

    const fetchAssignments = async () => {
        try {
            const res = await fetch("/api/assignments", {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });
            if (res.ok) {
                const data = await res.json();
                setAssignments(data);
            } else {
                toast.error("Failed to fetch assignments.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while fetching assignments.");
            console.error("Fetch assignments error:", error);
        }
    };

    const fetchDependencies = async () => {
        try {
            const [subjectsRes, staffRes] = await Promise.all([
                fetch("/api/subjects", { headers: { Authorization: `Bearer ${token}` } }),
                fetch("/api/staff", { headers: { Authorization: `Bearer ${token}` } }),
            ]);

            if (subjectsRes.ok) setSubjects(await subjectsRes.json());
            else toast.error("Failed to fetch subjects.");
            if (staffRes.ok) setStaffMembers(await staffRes.json());
            else toast.error("Failed to fetch staff.");

        } catch (error) {
            toast.error("Error fetching assignment dependencies.");
            console.error("Fetch dependencies error:", error);
        }
    };


    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value } = e.target;
        setNewAssignment(prev => ({ ...prev, [name]: value }));
    };

    const handleSelectChange = (name: string, value: string) => {
        setNewAssignment(prev => ({ ...prev, [name]: value }));
    };

    const handleSubmitNewAssignment = async (e: React.FormEvent) => {
        e.preventDefault();
        setIsLoading(true);

        const payload = {
            ...newAssignment,
            subjectId: newAssignment.subjectId || null,
            staffId: newAssignment.staffId || null,
            description: newAssignment.description || null,
        };

        try {
            const res = await fetch("/api/assignments", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify(payload),
            });

            if (res.ok) {
                toast.success("Assignment created successfully!");
                setNewAssignment({ // Reset form
                    title: '',
                    description: '',
                    assignmentType: 'assignment',
                    dueDate: '',
                    priority: 'medium',
                    subjectId: '',
                    staffId: '',
                });
                fetchAssignments(); // Refresh list
            } else {
                const errorData = await res.json();
                toast.error(errorData.error || "Failed to create assignment.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while creating assignment.");
            console.error("Create assignment error:", error);
        } finally {
            setIsLoading(false);
        }
    };

    const getStatusColor = (status: string) => {
        switch (status) {
            case 'pending': return 'text-yellow-500';
            case 'in_progress': return 'text-blue-500';
            case 'completed': return 'text-green-500';
            case 'submitted': return 'text-purple-500';
            case 'overdue': return 'text-red-500';
            default: return 'text-gray-500';
        }
    };

    return (
        <div className="space-y-6">
            <Toaster position="top-center" richColors />

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Create New Assignment</CardTitle>
                </CardHeader>
                <CardContent>
                    <form onSubmit={handleSubmitNewAssignment} className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        <Input name="title" placeholder="Assignment Title" value={newAssignment.title} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        <Textarea name="description" placeholder="Description (Optional)" value={newAssignment.description} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        
                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Assignment Type</label>
                            <Select name="assignmentType" value={newAssignment.assignmentType} onValueChange={(val) => handleSelectChange('assignmentType', val)} required>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select type" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    <SelectItem value="assignment">Assignment</SelectItem>
                                    <SelectItem value="lab_record">Lab Record</SelectItem>
                                    <SelectItem value="project">Project</SelectItem>
                                    <SelectItem value="presentation">Presentation</SelectItem>
                                    <SelectItem value="viva">Viva</SelectItem>
                                    <SelectItem value="quiz">Quiz</SelectItem>
                                    <SelectItem value="report">Report</SelectItem>
                                    <SelectItem value="other">Other</SelectItem>
                                </SelectContent>
                            </Select>
                        </div>

                        <Input name="dueDate" type="datetime-local" placeholder="Due Date" value={newAssignment.dueDate} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white" />
                        
                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Priority</label>
                            <Select name="priority" value={newAssignment.priority} onValueChange={(val) => handleSelectChange('priority', val)} required>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select priority" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    <SelectItem value="low">Low</SelectItem>
                                    <SelectItem value="medium">Medium</SelectItem>
                                    <SelectItem value="high">High</SelectItem>
                                    <SelectItem value="urgent">Urgent</SelectItem>
                                </SelectContent>
                            </Select>
                        </div>

                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Subject (Optional)</label>
                            <Select name="subjectId" value={newAssignment.subjectId || ''} onValueChange={(val) => handleSelectChange('subjectId', val)}>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select subject" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    {subjects.map(sub => (
                                        <SelectItem key={sub.id} value={sub.id}>{sub.name} ({sub.code})</SelectItem>
                                    ))}
                                </SelectContent>
                            </Select>
                        </div>

                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Staff (Optional)</label>
                            <Select name="staffId" value={newAssignment.staffId || ''} onValueChange={(val) => handleSelectChange('staffId', val)}>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select staff" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    {staffMembers.map(staff => (
                                        <SelectItem key={staff.id} value={staff.id}>{staff.name}</SelectItem>
                                    ))}
                                </SelectContent>
                            </Select>
                        </div>
                        
                        <div className="md:col-span-2">
                            <Button type="submit" disabled={isLoading} className="w-full bg-blue-600 hover:bg-blue-700">
                                {isLoading ? "Creating..." : "Add Assignment"}
                            </Button>
                        </div>
                    </form>
                </CardContent>
            </Card>

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Your Assignments</CardTitle>
                </CardHeader>
                <CardContent>
                    {assignments.length === 0 ? (
                        <p className="text-gray-400">No assignments found. Create one above!</p>
                    ) : (
                        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                            {assignments.map(assignment => (
                                <Card key={assignment.id} className="bg-gray-700 border-gray-600 p-4">
                                    <h3 className="font-bold text-lg text-white">{assignment.title}</h3>
                                    {assignment.description && <p className="text-gray-300 text-sm mt-1">{assignment.description}</p>}
                                    <p className={cn("text-sm font-semibold mt-2", getStatusColor(assignment.status))}>Status: {assignment.status}</p>
                                    <p className="text-gray-300 text-sm">Due: {new Date(assignment.dueDate).toLocaleString()}</p>
                                    <p className="text-gray-300 text-sm">Type: {assignment.assignmentType}</p>
                                    <p className="text-gray-300 text-sm">Priority: {assignment.priority}</p>
                                </Card>
                            ))}
                        </div>
                    )}
                </CardContent>
            </Card>
        </div>
    );
}

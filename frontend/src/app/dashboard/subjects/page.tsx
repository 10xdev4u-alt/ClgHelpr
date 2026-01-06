"use client";

import { useEffect, useState } from 'react';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Toaster, toast } from 'sonner';
import { useAuthStore } from '@/stores/auth-store';

interface Subject {
    id: string;
    code: string;
    name: string;
    shortName?: string;
    type: string;
    credits?: number;
    department?: string;
    semester?: number;
    color?: string;
}

export default function SubjectsPage() {
    const { token } = useAuthStore();
    const [subjects, setSubjects] = useState<Subject[]>([]);
    const [newSubject, setNewSubject] = useState<Omit<Subject, 'id' | 'createdAt'>>({
        code: '',
        name: '',
        type: 'core',
        department: 'CSE',
        semester: 6,
    });
    const [isLoading, setIsLoading] = useState(false);

    useEffect(() => {
        if (token) {
            fetchSubjects();
        }
    }, [token]);

    const fetchSubjects = async () => {
        try {
            const res = await fetch("/api/subjects", {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });
            if (res.ok) {
                const data = await res.json();
                setSubjects(data);
            } else {
                toast.error("Failed to fetch subjects.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while fetching subjects.");
            console.error("Fetch subjects error:", error);
        }
    };

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value, type } = e.target;
        setNewSubject(prev => ({
            ...prev,
            [name]: type === 'number' ? parseInt(value) : value,
        }));
    };

    const handleSubmitNewSubject = async (e: React.FormEvent) => {
        e.preventDefault();
        setIsLoading(true);

        try {
            const res = await fetch("/api/timetable/subjects", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify(newSubject),
            });

            if (res.ok) {
                toast.success("Subject created successfully!");
                setNewSubject({ code: '', name: '', type: 'core', department: 'CSE', semester: 6 }); // Reset form
                fetchSubjects(); // Refresh list
            } else {
                const errorData = await res.json();
                toast.error(errorData.error || "Failed to create subject.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while creating subject.");
            console.error("Create subject error:", error);
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="space-y-6">
            <Toaster position="top-center" richColors />

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Create New Subject</CardTitle>
                </CardHeader>
                <CardContent>
                    <form onSubmit={handleSubmitNewSubject} className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        <Input name="code" placeholder="Subject Code (e.g., CS22601)" value={newSubject.code} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="name" placeholder="Subject Name (e.g., Cryptography)" value={newSubject.name} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="shortName" placeholder="Short Name (Optional)" value={newSubject.shortName || ''} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="type" placeholder="Type (e.g., core, lab)" value={newSubject.type} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="credits" type="number" placeholder="Credits (Optional)" value={newSubject.credits || ''} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="department" placeholder="Department" value={newSubject.department} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="semester" type="number" placeholder="Semester" value={newSubject.semester || ''} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="color" placeholder="Color (Hex, Optional)" value={newSubject.color || ''} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <div className="md:col-span-2">
                            <Button type="submit" disabled={isLoading} className="w-full bg-blue-600 hover:bg-blue-700">
                                {isLoading ? "Creating..." : "Add Subject"}
                            </Button>
                        </div>
                    </form>
                </CardContent>
            </Card>

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Subjects List</CardTitle>
                </CardHeader>
                <CardContent>
                    {subjects.length === 0 ? (
                        <p className="text-gray-400">No subjects found. Create one above!</p>
                    ) : (
                        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                            {subjects.map(subject => (
                                <Card key={subject.id} className="bg-gray-700 border-gray-600 p-4">
                                    <h3 className="font-bold text-lg text-white">{subject.name} ({subject.code})</h3>
                                    <p className="text-gray-300 text-sm">Type: {subject.type}</p>
                                    {subject.shortName && <p className="text-gray-300 text-sm">Short: {subject.shortName}</p>}
                                    {subject.department && <p className="text-gray-300 text-sm">Dept: {subject.department}</p>}
                                    {subject.semester && <p className="text-gray-300 text-sm">Sem: {subject.semester}</p>}
                                </Card>
                            ))}
                        </div>
                    )}
                </CardContent>
            </Card>
        </div>
    );
}
